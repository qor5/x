package statusx

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/qor5/x/v3/i18nx"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
)

func TestTranslateError(t *testing.T) {
	ib, _ := i18nx.New(strings.NewReader(`
key,en,zh-CN
invalid_request,Invalid request,无效请求
email.invalid_format,Invalid email format,邮箱格式错误
pre_localized,Pre-localized error,预置本地化错误 %s
`))

	// Create context with Chinese locale
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(
		"x-selected-language", "zh-CN",
		"accept-language", "zh-CN;q=0.9",
	))

	// Create untranslated original error
	originalErr := New(codes.InvalidArgument, "invalid_request", "original error").
		WithExtraDetail(&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "email",
					Description: "invalid format",
					Reason:      "email.invalid_format",
				},
			},
		}).Err()

	t.Run("translate result and idempotence", func(t *testing.T) {
		err := TranslateError(ctx, ib, originalErr)
		originalStatus := status.Convert(err)

		var localizedMsg *errdetails.LocalizedMessage
		for _, detail := range originalStatus.Details() {
			if v, ok := detail.(*errdetails.LocalizedMessage); ok {
				localizedMsg = v
				break
			}
		}
		if assert.NotNil(t, localizedMsg, "should have localized message") {
			assert.Equal(t, "无效请求", localizedMsg.GetMessage())
		}

		// Test multiple translations
		for i := 0; i < 5; i++ {
			err = TranslateError(ctx, ib, err)
			assert.True(t, proto.Equal(originalStatus.Proto(), status.Convert(err).Proto()),
				"Protobuf mismatch after %d translations", i+1)
		}
	})

	t.Run("with localized", func(t *testing.T) {
		err := New(codes.InvalidArgument, "pre_localized", "original message").
			WithLocalized("pre_localized", "10").Err()
		err = TranslateError(ctx, ib, err)

		var localizedMsg *errdetails.LocalizedMessage
		for _, detail := range status.Convert(err).Details() {
			if v, ok := detail.(*errdetails.LocalizedMessage); ok {
				localizedMsg = v
				break
			}
		}
		if assert.NotNil(t, localizedMsg, "should have localized message") {
			assert.Equal(t, "预置本地化错误 10", localizedMsg.GetMessage(), "should have localized message")
		}
	})

	t.Run("with localized message", func(t *testing.T) {
		// Create error with pre-localized message
		preLocalizedErr := New(codes.InvalidArgument, "pre_localized", "original message").
			WithLocalized("pre_localized").
			WithExtraDetail(&errdetails.LocalizedMessage{
				Locale:  "zh-CN",
				Message: "指定本地化错误",
			}).Err()

		// First translation
		err1 := TranslateError(ctx, ib, preLocalizedErr)
		firstDetails := status.Convert(err1).Details()

		// Second translation
		err2 := TranslateError(ctx, ib, err1)
		secondDetails := status.Convert(err2).Details()

		// Verify idempotence
		assert.Equal(t, len(firstDetails), len(secondDetails), "details length should match")
		assert.True(t, proto.Equal(status.Convert(err1).Proto(), status.Convert(err2).Proto()))

		// Check preserved localized message
		var localizedMsg *errdetails.LocalizedMessage
		for _, detail := range secondDetails {
			if v, ok := detail.(*errdetails.LocalizedMessage); ok {
				localizedMsg = v
				break
			}
		}
		if assert.NotNil(t, localizedMsg, "should have localized message") {
			assert.Equal(t, "指定本地化错误", localizedMsg.GetMessage(), "should preserve pre-localized message")
		}
	})

	t.Run("translate by statusv1.ErrorReason", func(t *testing.T) {
		ib2, _ := i18nx.New(strings.NewReader(`
key,en,zh-CN
INVALID_ARGUMENT,Invalid argument,参数无效
`))

		ctx2 := metadata.NewIncomingContext(context.Background(), metadata.Pairs(
			"x-selected-language", "zh-CN",
			"accept-language", "zh-CN;q=0.9",
		))

		err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "original").Err()
		err = TranslateError(ctx2, ib2, err)

		var localizedMsg *errdetails.LocalizedMessage
		for _, detail := range status.Convert(err).Details() {
			if v, ok := detail.(*errdetails.LocalizedMessage); ok {
				localizedMsg = v
				break
			}
		}
		if assert.NotNil(t, localizedMsg, "should have localized message from enum reason") {
			assert.Equal(t, "参数无效", localizedMsg.GetMessage())
		}

		// Also verify embedded default.csv works without overrides
		errDef := New(codes.PermissionDenied, statusv1.ErrorReason_PERMISSION_DENIED.String(), "orig").Err()
		errDef = TranslateError(ctx2, ib2, errDef)
		var lmDef *errdetails.LocalizedMessage
		for _, d := range status.Convert(errDef).Details() {
			if v, ok := d.(*errdetails.LocalizedMessage); ok {
				lmDef = v
				break
			}
		}
		if assert.NotNil(t, lmDef, "should localize from embedded default.csv") {
			assert.Equal(t, "无权限", lmDef.GetMessage())
		}
	})
}
