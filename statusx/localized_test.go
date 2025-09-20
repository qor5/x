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

func TestFlattenFieldViolations(t *testing.T) {
	t.Run("flatten mixed single and slice", func(t *testing.T) {
		// ✅ mixed types in a single call
		single1 := &FieldViolation{Field: "email", Reason: "INVALID"}
		slice1 := []*FieldViolation{
			{Field: "name", Reason: "REQUIRED"},
			{Field: "age", Reason: "TOO_YOUNG"},
		}
		single2 := &FieldViolation{Field: "phone", Reason: "INVALID_FORMAT"}

		// Mixed usage: single, slice, single
		result := FlattenFieldViolations(single1, slice1, single2)

		// Verify results
		assert.Len(t, result, 4)
		assert.Equal(t, "email", result[0].Field)
		assert.Equal(t, "name", result[1].Field)
		assert.Equal(t, "age", result[2].Field)
		assert.Equal(t, "phone", result[3].Field)
	})

	t.Run("flatten handles nil", func(t *testing.T) {
		var nilSingle *FieldViolation
		slice := []*FieldViolation{
			{Field: "valid", Reason: "OK"},
		}

		result := FlattenFieldViolations(nilSingle, slice, nilSingle)
		assert.Len(t, result, 1)
		assert.Equal(t, "valid", result[0].Field)
	})

	t.Run("flatten with ToFieldViolations from error", func(t *testing.T) {
		// Create a status error with field violations
		statusErr := New(codes.InvalidArgument, "validation failed").
			WithFieldViolations(
				&FieldViolation{Field: "nested.field1", Reason: "INVALID", Description: "Field1 is invalid"},
				&FieldViolation{Field: "nested.field2", Reason: "REQUIRED", Description: "Field2 is required"},
			).Err()

		// Convert error to field violations using ToFieldViolations
		errorViolations := ToFieldViolations(statusErr, "user")

		// Mix with other violations
		directViolations := []*FieldViolation{
			{Field: "email", Reason: "INVALID_FORMAT"},
		}

		result := FlattenFieldViolations(errorViolations, directViolations)

		// Verify results: main error + 2 nested violations + 1 direct violation = 4 total
		assert.Len(t, result, 4)
		assert.Equal(t, "user", result[0].Field)               // main error field
		assert.Equal(t, "user.nested.field1", result[1].Field) // prefixed nested field
		assert.Equal(t, "user.nested.field2", result[2].Field) // prefixed nested field
		assert.Equal(t, "email", result[3].Field)              // direct violation
	})

	t.Run("flatten with Status.ToFieldViolations method", func(t *testing.T) {
		// Create a status with field violations
		status := New(codes.InvalidArgument, "user data invalid").WithReason("USER_INVALID").
			WithFieldViolations(
				&FieldViolation{Field: "profile.bio", Reason: "TOO_LONG", Description: "Bio is too long"},
				&FieldViolation{Field: "profile.avatar", Reason: "INVALID_FORMAT", Description: "Avatar format is invalid"},
			)

		// Use Status.ToFieldViolations method
		statusViolations := status.ToFieldViolations("request")

		// Mix with single violations
		singleViolation := &FieldViolation{Field: "timestamp", Reason: "EXPIRED"}

		result := FlattenFieldViolations(statusViolations, singleViolation)

		// Verify results: main status error + 2 nested violations + 1 single violation = 4 total
		assert.Len(t, result, 4)
		assert.Equal(t, "request", result[0].Field)                // main status field
		assert.Equal(t, "request.profile.bio", result[1].Field)    // prefixed nested field
		assert.Equal(t, "request.profile.avatar", result[2].Field) // prefixed nested field
		assert.Equal(t, "timestamp", result[3].Field)              // single violation
		assert.Equal(t, "USER_INVALID", result[0].Reason)          // main status reason
	})

	t.Run("flatten protobuf field violations", func(t *testing.T) {
		// Test direct *errdetails.BadRequest_FieldViolation support
		pbSingle := &errdetails.BadRequest_FieldViolation{
			Field:       "pb_email",
			Reason:      "INVALID_FORMAT",
			Description: "Protobuf email invalid",
		}

		pbSlice := []*errdetails.BadRequest_FieldViolation{
			{Field: "pb_name", Reason: "REQUIRED", Description: "Protobuf name required"},
			{Field: "pb_age", Reason: "TOO_YOUNG", Description: "Protobuf age too young"},
		}

		// Mix protobuf types with internal types
		internal := &FieldViolation{Field: "internal", Reason: "CUSTOM"}

		result := FlattenFieldViolations(pbSingle, pbSlice, internal)

		// Verify results: 1 single pb + 2 slice pb + 1 internal = 4 total
		assert.Len(t, result, 4)
		assert.Equal(t, "pb_email", result[0].Field)
		assert.Equal(t, "INVALID_FORMAT", result[0].Reason)
		assert.Equal(t, "Protobuf email invalid", result[0].Description)
		assert.Equal(t, "pb_name", result[1].Field)
		assert.Equal(t, "pb_age", result[2].Field)
		assert.Equal(t, "internal", result[3].Field)
		assert.Equal(t, "CUSTOM", result[3].Reason)
	})

	t.Run("flatten nil protobuf violations", func(t *testing.T) {
		var nilPbSingle *errdetails.BadRequest_FieldViolation
		var nilPbSlice []*errdetails.BadRequest_FieldViolation
		validInternal := &FieldViolation{Field: "valid", Reason: "OK"}

		result := FlattenFieldViolations(nilPbSingle, nilPbSlice, validInternal)

		// Should handle nil protobuf inputs gracefully
		assert.Len(t, result, 1)
		assert.Equal(t, "valid", result[0].Field)
	})
}

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
	originalErr := New(codes.InvalidArgument, "original error").WithReason("invalid_request").
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
		err := New(codes.InvalidArgument, "original message").
			WithLocalized("pre_localized", "10").Err()
		err = TranslateError(err, ib, ib.LanguageFromContext(ctx))

		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
		if assert.NotNil(t, localizedMsg, "should have localized message") {
			assert.Equal(t, "预置本地化错误 10", localizedMsg.GetMessage(), "should have localized message")
		}
	})

	t.Run("with localized message", func(t *testing.T) {
		// Create error with pre-localized message
		preLocalizedErr := New(codes.InvalidArgument, "original message").
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

		err := New(codes.InvalidArgument, "original").Err()
		err = TranslateError(err, ib2, ib2.LanguageFromContext(ctx2))

		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
		if assert.NotNil(t, localizedMsg, "should have localized message from enum reason") {
			assert.Equal(t, "参数无效", localizedMsg.GetMessage())
		}

		// Also verify embedded default.csv works without overrides
		errDef := New(codes.PermissionDenied, "orig").Err()
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
			err := New(codes.InvalidArgument, "validation error").WithLocalized("user.validation.failed", data).Err()
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
			err := New(codes.Internal, "complex error").WithLocalized("complex.nested.data", data).Err()
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
			err := New(codes.Internal, "order error").WithLocalized("order.processing.error", data).Err()
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

			err := New(codes.Internal, "product created").WithLocalized("product.created", data).Err()
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
		err := New(codes.InvalidArgument, "Email validation failed").
			WithFieldViolations(&FieldViolation{
				Field:  "email",
				Reason: "INVALID_DOMAIN",
				Localized: &Localized{
					Key:  "field.email.detailed",
					Args: []any{data},
				},
			}).Err()

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

			err := New(codes.InvalidArgument, "validation failed").
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

			err := New(codes.InvalidArgument, "validation failed").
				WithDetails(existingBadRequest).
				WithFieldViolations(&FieldViolation{
					Field:       "password",
					Description: "Password is required",
					Reason:      "PASSWORD_REQUIRED",
					Localized:   &Localized{Key: "field.password.required"},
				}).
				Err()

			translated := TranslateError(err, ib, ib.LanguageFromContext(ctx))
			st := status.Convert(translated)

			badRequests := ExtractDetail[*errdetails.BadRequest](st.Details())
			assert.NotNil(t, badRequests, "should preserve existing BadRequest")
			assert.Len(t, badRequests.FieldViolations, 1, "should only have original field violation")
			assert.Equal(t, "email", badRequests.FieldViolations[0].Field)
		})

		t.Run("construct BadRequest when none exists", func(t *testing.T) {
			err := New(codes.InvalidArgument, "validation failed").
				WithFieldViolations(
					&FieldViolation{
						Field:       "email",
						Description: "Email address is required",
						Reason:      "EMAIL_REQUIRED",
						Localized:   &Localized{Key: "field.email.required"},
					},
					&FieldViolation{
						Field:       "password",
						Description: "Password is too weak",
						Reason:      "PASSWORD_WEAK",
						Localized:   &Localized{Key: "field.password.weak", Args: []any{"minLength", 8}},
					},
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
			assert.Equal(t, "EMAIL_REQUIRED", emailViolation.Reason)
			assert.NotNil(t, emailViolation.LocalizedMessage, "should have localized message")

			// Verify second field
			passwordViolation := badRequests.FieldViolations[1]
			assert.Equal(t, "password", passwordViolation.Field)
			assert.Equal(t, "Password is too weak", passwordViolation.Description)
			assert.Equal(t, "PASSWORD_WEAK", passwordViolation.Reason)
			assert.NotNil(t, passwordViolation.LocalizedMessage, "should have localized message")
		})
	})
}

func TestToFieldViolations(t *testing.T) {
	t.Run("basic cases", func(t *testing.T) {
		// Nil input
		violations := ToFieldViolations(nil, "user")
		assert.Nil(t, violations)

		// Simple error
		err := New(codes.InvalidArgument, "Email format is invalid").WithReason("INVALID_EMAIL").Err()
		violations = ToFieldViolations(err, "email")
		assert.Len(t, violations, 1)
		assert.Equal(t, "email", violations[0].Field)
		assert.Equal(t, "INVALID_EMAIL", violations[0].Reason)
	})

	t.Run("localization priority", func(t *testing.T) {
		// LocalizedMessage takes priority over Localized
		localizedMsg := &errdetails.LocalizedMessage{Locale: "zh-CN", Message: "预置本地化消息"}
		err := New(codes.InvalidArgument, "test").
			WithLocalized("template.key", "arg1").
			WithDetails(localizedMsg).Err()
		violations := ToFieldViolations(err, "test")

		assert.Len(t, violations, 1)
		assert.Nil(t, violations[0].Localized)
		assert.Equal(t, localizedMsg, violations[0].LocalizedMessage)
	})

	t.Run("nested violations", func(t *testing.T) {
		// Custom BadRequest
		err := New(codes.InvalidArgument, "Form validation failed").
			WithFieldViolations(
				&FieldViolation{Field: "email", Reason: "REQUIRED", Description: "Email is required"},
			).Err()
		violations := ToFieldViolations(err, "form")

		assert.Len(t, violations, 2)
		assert.Equal(t, "form", violations[0].Field)
		assert.Equal(t, "form.email", violations[1].Field)

		// Standard BadRequest
		standardBadRequest := &errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{Field: "username", Reason: "EXISTS", Description: "Username exists"},
			},
		}
		err = New(codes.InvalidArgument, "User validation failed").
			WithDetails(standardBadRequest).Err()
		violations = ToFieldViolations(err, "user")

		assert.Len(t, violations, 2)
		assert.Equal(t, "user", violations[0].Field)
		assert.Equal(t, "user.username", violations[1].Field)
	})
}

func TestLocalizedConstructors(t *testing.T) {
	t.Run("WithLocalized creates localized error", func(t *testing.T) {
		status := WithLocalized(codes.InvalidArgument, "error.validation", "field", "email")

		assert.Equal(t, codes.InvalidArgument, status.Code())
		assert.Equal(t, statusv1.ErrorReason_INVALID_ARGUMENT.String(), status.Reason())
		assert.Equal(t, "error.validation", status.Message())

		localized := status.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "error.validation", localized.Key)
		assert.Len(t, localized.Args, 2)
	})

	t.Run("WrapLocalized wraps existing error with localization", func(t *testing.T) {
		originalErr := errors.New("original error")
		status := WrapLocalized(originalErr, codes.Internal, "error.internal", "service", "user-service")

		assert.Equal(t, codes.Internal, status.Code())
		assert.Equal(t, statusv1.ErrorReason_INTERNAL.String(), status.Reason())
		assert.Equal(t, "error.internal", status.Message())
		assert.True(t, errors.Is(status.Err(), originalErr))

		localized := status.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "error.internal", localized.Key)
		assert.Len(t, localized.Args, 2)
	})
}

func TestTranslateStatusErrorOnly(t *testing.T) {
	t.Run("handles StatusError correctly", func(t *testing.T) {
		ib, _ := i18nx.New(strings.NewReader(`
key,en,zh
test.message,Test message,测试消息
`))

		originalErr := New(codes.InvalidArgument, "original message").WithLocalized("test.message").Err()

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
