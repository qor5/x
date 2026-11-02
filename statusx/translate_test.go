package statusx

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/qor5/x/v3/i18nx"
	"github.com/qor5/x/v3/jsonx"
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
		WithDetails(&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "email",
					Description: "invalid format",
					Reason:      "email.invalid_format",
				},
			},
		}).Err()

	t.Run("translate result and idempotence", func(t *testing.T) {
		err := TranslateError(originalErr, ib, ib.LanguageFromContext(ctx))
		originalStatus := status.Convert(err)

		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](originalStatus.Details())
		if assert.NotNil(t, localizedMsg, "should have localized message") {
			assert.Equal(t, "无效请求", localizedMsg.GetMessage())
		}

		// Test multiple translations
		for i := 0; i < 5; i++ {
			err = TranslateError(err, ib, ib.LanguageFromContext(ctx))
			assert.True(t, proto.Equal(originalStatus.Proto(), status.Convert(err).Proto()),
				"Protobuf mismatch after %d translations", i+1)
		}
	})

	t.Run("with localized", func(t *testing.T) {
		err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "original message").
			WithLocalized("pre_localized", "10").Err()
		err = TranslateError(err, ib, ib.LanguageFromContext(ctx))

		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
		if assert.NotNil(t, localizedMsg, "should have localized message") {
			assert.Equal(t, "预置本地化错误 10", localizedMsg.GetMessage(), "should have localized message")
		}
	})

	t.Run("with localized message", func(t *testing.T) {
		// Create error with pre-localized message
		preLocalizedErr := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "original message").
			WithLocalized("pre_localized").
			WithDetails(&errdetails.LocalizedMessage{
				Locale:  "zh-CN",
				Message: "指定本地化错误",
			}).Err()

		// First translation
		err1 := TranslateError(preLocalizedErr, ib, ib.LanguageFromContext(ctx))
		firstDetails := status.Convert(err1).Details()

		// Second translation
		err2 := TranslateError(err1, ib, ib.LanguageFromContext(ctx))
		secondDetails := status.Convert(err2).Details()

		// Verify idempotence
		assert.Equal(t, len(firstDetails), len(secondDetails), "details length should match")
		assert.True(t, proto.Equal(status.Convert(err1).Proto(), status.Convert(err2).Proto()))

		// Check preserved localized message
		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](secondDetails)
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
		err = TranslateError(err, ib2, ib2.LanguageFromContext(ctx2))

		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
		if assert.NotNil(t, localizedMsg, "should have localized message from enum reason") {
			assert.Equal(t, "参数无效", localizedMsg.GetMessage())
		}

		// Also verify embedded default.csv works without overrides
		errDef := New(codes.PermissionDenied, statusv1.ErrorReason_PERMISSION_DENIED.String(), "orig").Err()
		errDef = TranslateError(errDef, ib2, ib2.LanguageFromContext(ctx2))
		lmDef := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(errDef).Details())
		if assert.NotNil(t, lmDef, "should localize from embedded default.csv") {
			assert.Equal(t, "无权限", lmDef.GetMessage())
		}
	})

	t.Run("gotpl template rendering", func(t *testing.T) {
		ib3, _ := i18nx.New(strings.NewReader(`
key,en,zh-CN
user.validation.failed,"User {{.Name}} validation failed: {{.Field}} {{.Error}}","用户 {{.Name}} 验证失败：{{.Field}} {{.Error}}"
order.processing.error,"Order {{.OrderID}} failed to process. Items: {{.ItemList}} Status: {{.Status}}","订单 {{.OrderID}} 处理失败。商品：{{.ItemList}} 状态：{{.Status}}"
complex.nested.data,"Welcome {{.User.Name}}! You have {{.Stats.UnreadCount}} unread messages and {{.Stats.PendingTasks}} pending tasks.","欢迎 {{.User.Name}}！您有 {{.Stats.UnreadCount}} 条未读消息和 {{.Stats.PendingTasks}} 个待处理任务。"
`))

		ctx3 := metadata.NewIncomingContext(context.Background(), metadata.Pairs(
			"x-selected-language", "zh-CN",
			"accept-language", "zh-CN;q=0.9",
		))

		t.Run("map data with template", func(t *testing.T) {
			data := map[string]any{
				"Name":  "张三",
				"Field": "邮箱",
				"Error": "格式无效",
			}
			err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation error").WithLocalized("user.validation.failed", data).Err()
			err = TranslateError(err, ib3, ib3.LanguageFromContext(ctx3))

			localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
			if assert.NotNil(t, localizedMsg, "should have localized message from gotpl") {
				assert.Equal(t, "用户 张三 验证失败：邮箱 格式无效", localizedMsg.GetMessage())
			}
		})

		t.Run("map data with nested template", func(t *testing.T) {
			data := map[string]any{
				"User": map[string]any{
					"Name": "李四",
				},
				"Stats": map[string]any{
					"UnreadCount":  5,
					"PendingTasks": 3,
				},
			}
			err := New(codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "complex error").WithLocalized("complex.nested.data", data).Err()
			err = TranslateError(err, ib3, ib3.LanguageFromContext(ctx3))

			localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
			if assert.NotNil(t, localizedMsg, "should have localized message from nested map") {
				assert.Equal(t, "欢迎 李四！您有 5 条未读消息和 3 个待处理任务。", localizedMsg.GetMessage())
			}
		})

		t.Run("complex data with string template", func(t *testing.T) {
			data := map[string]any{
				"OrderID":  "ORD-12345",
				"ItemList": "苹果(2) 香蕉(3)",
				"Status":   "已取消",
			}
			err := New(codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "order error").WithLocalized("order.processing.error", data).Err()
			err = TranslateError(err, ib3, ib3.LanguageFromContext(ctx3))

			localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
			if assert.NotNil(t, localizedMsg, "should have localized message with string template") {
				expected := "订单 ORD-12345 处理失败。商品：苹果(2) 香蕉(3) 状态：已取消"
				assert.Equal(t, expected, localizedMsg.GetMessage())
			}
		})

		t.Run("struct data with jsonx.MustToMap preprocessing", func(t *testing.T) {
			// Add struct template for this test
			ib5, _ := i18nx.New(strings.NewReader(`
key,en,zh-CN
product.created,"Product {{.name}} created successfully. Price: {{.price}}, Category: {{.category}}, Available: {{.isAvailable}}","产品 {{.name}} 创建成功。价格：{{.price}}，类别：{{.category}}，可用：{{.isAvailable}}"
`))

			type Product struct {
				Name        string  `json:"name"`
				Price       float64 `json:"price"`
				Category    string  `json:"category"`
				IsAvailable bool    `json:"isAvailable"`
			}

			productData := Product{
				Name:        "智能手表",
				Price:       999.99,
				Category:    "电子产品",
				IsAvailable: true,
			}

			// Convert struct to map using jsonx.MustToMap - this is the key improvement!
			data := jsonx.MustToMap(productData)

			err := New(codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "product created").WithLocalized("product.created", data).Err()
			err = TranslateError(err, ib5, ib5.LanguageFromContext(ctx3))

			localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
			if assert.NotNil(t, localizedMsg, "should have localized message from struct via jsonx.MustToMap") {
				expected := "产品 智能手表 创建成功。价格：999.99，类别：电子产品，可用：true"
				assert.Equal(t, expected, localizedMsg.GetMessage())
			}
		})
	})

	t.Run("field localization with jsonx.MustToMap", func(t *testing.T) {
		type ValidationInfo struct {
			CurrentEmail string `json:"currentEmail"`
			ValidDomains []any  `json:"validDomains"`
			HelpURL      string `json:"helpURL"`
		}

		validationInfo := ValidationInfo{
			CurrentEmail: "user@invalid.co",
			ValidDomains: []any{"gmail.com", "outlook.com", "company.com"},
			HelpURL:      "https://help.company.com/email-format",
		}

		// Template for this test
		ib7, _ := i18nx.New(strings.NewReader(`
key,en,zh-CN
field.email.detailed,"Current email: {{.currentEmail}}. Valid domains: {{range .validDomains}}{{.}} {{end}}. Help: {{.helpURL}}","当前邮箱：{{.currentEmail}}。有效域名：{{range .validDomains}}{{.}} {{end}}。帮助：{{.helpURL}}"
`))

		ctx7 := metadata.NewIncomingContext(context.Background(), metadata.Pairs(
			"x-selected-language", "zh-CN",
			"accept-language", "zh-CN;q=0.9",
		))

		data := jsonx.MustToMap(validationInfo)
		err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "Email validation failed").
			WithFieldViolations(
				NewFieldViolation("email", "field.email.detailed", "Email domain is invalid").
					WithLocalizedArgs(data),
			).Err()

		err = TranslateError(err, ib7, ib7.LanguageFromContext(ctx7))

		badRequest := ExtractDetail[*errdetails.BadRequest](status.Convert(err).Details())
		if assert.NotNil(t, badRequest, "should have BadRequest detail") {
			fvs := badRequest.GetFieldViolations()
			if assert.Len(t, fvs, 1, "should have one field violation") {
				fv := fvs[0]
				localizedMsg := fv.GetLocalizedMessage()
				if assert.NotNil(t, localizedMsg, "should have localized message") {
					expected := "当前邮箱：user@invalid.co。有效域名：gmail.com outlook.com company.com 。帮助：https://help.company.com/email-format"
					assert.Equal(t, expected, localizedMsg.GetMessage())
				}
			}
		}
	})

	t.Run("TranslateError logic consistency", func(t *testing.T) {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(
			"x-selected-language", "zh-CN",
		))

		t.Run("skip translation when LocalizedMessage exists", func(t *testing.T) {
			existingLocalized := &errdetails.LocalizedMessage{
				Locale:  "zh-CN",
				Message: "现有的本地化消息",
			}

			err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithDetails(existingLocalized).
				Err()

			translated := TranslateError(err, ib, ib.LanguageFromContext(ctx))
			st := status.Convert(translated)

			localizedMessages := ExtractDetail[*errdetails.LocalizedMessage](st.Details())
			assert.Equal(t, "现有的本地化消息", localizedMessages.Message, "should preserve existing LocalizedMessage")
		})

		t.Run("skip translation when BadRequest exists", func(t *testing.T) {
			existingBadRequest := &errdetails.BadRequest{
				FieldViolations: []*errdetails.BadRequest_FieldViolation{
					{
						Field:       "email",
						Description: "Email is required",
						Reason:      "FIELD_REQUIRED",
					},
				},
			}

			err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithDetails(existingBadRequest).
				WithFieldViolations(
					NewFieldViolation("password", "field.password.required", "Password is required"),
				).
				Err()

			translated := TranslateError(err, ib, ib.LanguageFromContext(ctx))
			st := status.Convert(translated)

			badRequests := ExtractDetail[*errdetails.BadRequest](st.Details())
			assert.NotNil(t, badRequests, "should preserve existing BadRequest")
			assert.Len(t, badRequests.FieldViolations, 1, "should only have original field violation")
			assert.Equal(t, "email", badRequests.FieldViolations[0].Field)
		})

		t.Run("construct BadRequest when none exists", func(t *testing.T) {
			err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithFieldViolations(
					NewFieldViolation("email", "field.email.required", "Email address is required"),
					NewFieldViolation("password", "field.password.weak", "Password is too weak").
						WithLocalizedArgs("minLength", 8),
				).
				Err()

			translated := TranslateError(err, ib, ib.LanguageFromContext(ctx))
			st := status.Convert(translated)

			badRequests := ExtractDetail[*errdetails.BadRequest](st.Details())
			assert.NotNil(t, badRequests, "should construct new BadRequest")
			assert.Len(t, badRequests.FieldViolations, 2, "should have both field violations")

			// Verify first field
			emailViolation := badRequests.FieldViolations[0]
			assert.Equal(t, "email", emailViolation.Field)
			assert.Equal(t, "Email address is required", emailViolation.Description)
			assert.Equal(t, "field.email.required", emailViolation.Reason)
			assert.NotNil(t, emailViolation.LocalizedMessage, "should have localized message")

			// Verify second field
			passwordViolation := badRequests.FieldViolations[1]
			assert.Equal(t, "password", passwordViolation.Field)
			assert.Equal(t, "Password is too weak", passwordViolation.Description)
			assert.Equal(t, "field.password.weak", passwordViolation.Reason)
			assert.NotNil(t, passwordViolation.LocalizedMessage, "should have localized message")
		})
	})
}

func TestTranslateStatusErrorOnly(t *testing.T) {
	t.Run("handles StatusError correctly", func(t *testing.T) {
		ib, _ := i18nx.New(strings.NewReader(`
key,en,zh
test.message,Test message,测试消息
`))

		originalErr := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "original message").WithLocalized("test.message").Err()

		translatedErr, ok := TranslateStatusErrorOnly(originalErr, ib, language.Chinese)
		assert.True(t, ok, "should successfully handle StatusError")
		assert.NotNil(t, translatedErr, "translated error should not be nil")

		// The function returns the result of TranslateError, which may be the same
		// object if no translation is needed, but it's still processed
		st := Convert(translatedErr)
		assert.Equal(t, codes.InvalidArgument, st.Code())

		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](st.Details())
		assert.Equal(t, "zh", localizedMsg.GetLocale())
		assert.Equal(t, "测试消息", localizedMsg.GetMessage())
	})

	t.Run("handles non-StatusError", func(t *testing.T) {
		ib, _ := i18nx.New(strings.NewReader(""))
		originalErr := errors.New("not a status error")

		translatedErr, ok := TranslateStatusErrorOnly(originalErr, ib, language.Chinese)
		assert.False(t, ok, "should not translate non-StatusError")
		assert.Equal(t, originalErr, translatedErr, "should return original error unchanged")
	})
}
