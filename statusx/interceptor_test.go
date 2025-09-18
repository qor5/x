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
		err := TranslateError(ctx, ib, originalErr)
		originalStatus := status.Convert(err)

		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](originalStatus.Details())
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

		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
		if assert.NotNil(t, localizedMsg, "should have localized message") {
			assert.Equal(t, "预置本地化错误 10", localizedMsg.GetMessage(), "should have localized message")
		}
	})

	t.Run("with localized message", func(t *testing.T) {
		// Create error with pre-localized message
		preLocalizedErr := New(codes.InvalidArgument, "pre_localized", "original message").
			WithLocalized("pre_localized").
			WithDetails(&errdetails.LocalizedMessage{
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
		err = TranslateError(ctx2, ib2, err)

		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
		if assert.NotNil(t, localizedMsg, "should have localized message from enum reason") {
			assert.Equal(t, "参数无效", localizedMsg.GetMessage())
		}

		// Also verify embedded default.csv works without overrides
		errDef := New(codes.PermissionDenied, statusv1.ErrorReason_PERMISSION_DENIED.String(), "orig").Err()
		errDef = TranslateError(ctx2, ib2, errDef)
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
			err := New(codes.InvalidArgument, "user.validation.failed", "validation error").WithLocalized("user.validation.failed", data).Err()
			err = TranslateError(ctx3, ib3, err)

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
			err := New(codes.Internal, "complex.nested.data", "complex error").WithLocalized("complex.nested.data", data).Err()
			err = TranslateError(ctx3, ib3, err)

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
			err := New(codes.Internal, "order.processing.error", "order error").WithLocalized("order.processing.error", data).Err()
			err = TranslateError(ctx3, ib3, err)

			localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
			if assert.NotNil(t, localizedMsg, "should have localized message with string template") {
				expected := "订单 ORD-12345 处理失败。商品：苹果(2) 香蕉(3) 状态：已取消"
				assert.Equal(t, expected, localizedMsg.GetMessage())
			}
		})

		t.Run("gotpl fallback to positional when template fails", func(t *testing.T) {
			// Create template with syntax error that will fall back to positional args
			ib4, _ := i18nx.New(strings.NewReader(`
key,en,zh-CN
bad.template,"Error {{.InvalidSyntax}","错误 {{.InvalidSyntax}"
fallback.template,"User %s has %d errors","用户 %s 有 %d 个错误"
`))

			// Test with invalid template - should fallback to positional
			data := map[string]any{"Name": "王五"}
			err := New(codes.Internal, "bad.template", "template error").WithLocalized("bad.template", data).Err()
			err = TranslateError(ctx3, ib4, err)

			localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
			if assert.NotNil(t, localizedMsg, "should fallback when template fails") {
				// Should return the raw template as fallback
				assert.Equal(t, "错误 {{.InvalidSyntax}", localizedMsg.GetMessage())
			}

			// Test normal positional args still work
			err2 := New(codes.Internal, "fallback.template", "fallback error").WithLocalized("fallback.template", "张三", 5).Err()
			err2 = TranslateError(ctx3, ib4, err2)

			localizedMsg2 := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err2).Details())
			if assert.NotNil(t, localizedMsg2, "should work with positional args") {
				assert.Equal(t, "用户 张三 有 5 个错误", localizedMsg2.GetMessage())
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

			err := New(codes.Internal, "product.created", "product created").WithLocalized("product.created", data).Err()
			err = TranslateError(ctx3, ib5, err)

			localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
			if assert.NotNil(t, localizedMsg, "should have localized message from struct via jsonx.MustToMap") {
				expected := "产品 智能手表 创建成功。价格：999.99，类别：电子产品，可用：true"
				assert.Equal(t, expected, localizedMsg.GetMessage())
			}
		})
	})

	t.Run("field-level localization with proto support", func(t *testing.T) {
		ib6, _ := i18nx.New(strings.NewReader(`
key,en,zh-CN
field.email.validation.invalid_format,"Email {{.Email}} has invalid format. Expected pattern: {{.Pattern}}. Suggestions: {{range .Suggestions}}{{.}} {{end}}","邮箱 {{.Email}} 格式无效。期望格式：{{.Pattern}}。建议：{{range .Suggestions}}{{.}} {{end}}"
field.age.validation.out_of_range,"Age {{.Current}} is out of valid range {{.Min}}-{{.Max}}","年龄 {{.Current}} 超出有效范围 {{.Min}}-{{.Max}}"
global.validation_failed,"Multiple validation errors occurred","发生多个验证错误"
`))

		ctx6 := metadata.NewIncomingContext(context.Background(), metadata.Pairs(
			"x-selected-language", "zh-CN",
			"accept-language", "zh-CN;q=0.9",
		))

		t.Run("single field with rich template data", func(t *testing.T) {
			data := map[string]any{
				"Email":       "user@invalid.domain",
				"Pattern":     "RFC5322",
				"Suggestions": []any{"@gmail.com", "@outlook.com", "@qq.com"},
			}

			err := New(codes.InvalidArgument, "validation_failed", "Validation failed").
				WithLocalized("global.validation_failed").
				WithFieldLocalized("email", "field.email.validation.invalid_format", data).
				WithDetails(&errdetails.BadRequest{
					FieldViolations: []*errdetails.BadRequest_FieldViolation{
						{
							Field:       "email",
							Reason:      "INVALID_EMAIL_FORMAT",
							Description: "Invalid email format",
						},
					},
				}).Err()

			err = TranslateError(ctx6, ib6, err)

			// Check top-level localized message
			topLevelMsg := ExtractDetail[*errdetails.LocalizedMessage](status.Convert(err).Details())
			if assert.NotNil(t, topLevelMsg, "should have top-level localized message") {
				assert.Equal(t, "发生多个验证错误", topLevelMsg.GetMessage())
			}

			// Check field-level localized message
			badRequest := ExtractDetail[*errdetails.BadRequest](status.Convert(err).Details())
			if assert.NotNil(t, badRequest, "should have BadRequest detail") {
				fvs := badRequest.GetFieldViolations()
				if assert.Len(t, fvs, 1, "should have one field violation") {
					fv := fvs[0]
					assert.Equal(t, "email", fv.GetField())

					localizedMsg := fv.GetLocalizedMessage()
					if assert.NotNil(t, localizedMsg, "should have field-level localized message") {
						expected := "邮箱 user@invalid.domain 格式无效。期望格式：RFC5322。建议：@gmail.com @outlook.com @qq.com "
						assert.Equal(t, expected, localizedMsg.GetMessage())
					}
				}
			}
		})

		t.Run("multiple fields with different templates", func(t *testing.T) {
			emailData := map[string]any{
				"Email":       "bad@email",
				"Pattern":     "RFC5322",
				"Suggestions": []any{"@gmail.com", "@163.com"},
			}

			ageData := map[string]any{
				"Current": 16,
				"Min":     18,
				"Max":     65,
			}

			err := New(codes.InvalidArgument, "validation_failed", "Multiple validation errors").
				WithLocalized("global.validation_failed").
				WithFieldLocalized("email", "field.email.validation.invalid_format", emailData).
				WithFieldLocalized("age", "field.age.validation.out_of_range", ageData).
				WithDetails(&errdetails.BadRequest{
					FieldViolations: []*errdetails.BadRequest_FieldViolation{
						{Field: "email", Reason: "INVALID_EMAIL_FORMAT"},
						{Field: "age", Reason: "OUT_OF_RANGE"},
					},
				}).Err()

			err = TranslateError(ctx6, ib6, err)

			badRequest := ExtractDetail[*errdetails.BadRequest](status.Convert(err).Details())
			if assert.NotNil(t, badRequest, "should have BadRequest detail") {
				fvs := badRequest.GetFieldViolations()
				if assert.Len(t, fvs, 2, "should have two field violations") {
					// Check email field
					for _, fv := range fvs {
						if fv.GetField() == "email" {
							localizedMsg := fv.GetLocalizedMessage()
							if assert.NotNil(t, localizedMsg, "email field should have localized message") {
								expected := "邮箱 bad@email 格式无效。期望格式：RFC5322。建议：@gmail.com @163.com "
								assert.Equal(t, expected, localizedMsg.GetMessage())
							}
						} else if fv.GetField() == "age" {
							localizedMsg := fv.GetLocalizedMessage()
							if assert.NotNil(t, localizedMsg, "age field should have localized message") {
								expected := "年龄 16 超出有效范围 18-65"
								assert.Equal(t, expected, localizedMsg.GetMessage())
							}
						}
					}
				}
			}
		})

		t.Run("batch field localization with WithFieldsLocalized", func(t *testing.T) {
			emailData := map[string]any{
				"Email":       "user@invalid.domain",
				"Pattern":     "RFC5322",
				"Suggestions": []any{"@gmail.com", "@qq.com"},
			}

			ageData := map[string]any{
				"Current": 17,
				"Min":     18,
				"Max":     65,
			}

			// Create field localization configurations
			fields := []*FieldLocalized{
				{
					Field: "email",
					Key:   "field.email.validation.invalid_format",
					Args:  []any{emailData},
				},
				{
					Field: "age",
					Key:   "field.age.validation.out_of_range",
					Args:  []any{ageData},
				},
			}

			err := New(codes.InvalidArgument, "validation_failed", "Multiple validation errors").
				WithLocalized("global.validation_failed").
				WithFieldsLocalized(fields...). // Batch operation!
				WithDetails(&errdetails.BadRequest{
					FieldViolations: []*errdetails.BadRequest_FieldViolation{
						{Field: "email", Reason: "INVALID_EMAIL_FORMAT"},
						{Field: "age", Reason: "OUT_OF_RANGE"},
					},
				}).Err()

			err = TranslateError(ctx6, ib6, err)

			badRequest := ExtractDetail[*errdetails.BadRequest](status.Convert(err).Details())
			if assert.NotNil(t, badRequest, "should have BadRequest detail") {
				fvs := badRequest.GetFieldViolations()
				if assert.Len(t, fvs, 2, "should have two field violations") {
					// Check email field
					for _, fv := range fvs {
						if fv.GetField() == "email" {
							localizedMsg := fv.GetLocalizedMessage()
							if assert.NotNil(t, localizedMsg, "email field should have localized message") {
								expected := "邮箱 user@invalid.domain 格式无效。期望格式：RFC5322。建议：@gmail.com @qq.com "
								assert.Equal(t, expected, localizedMsg.GetMessage())
							}
						} else if fv.GetField() == "age" {
							localizedMsg := fv.GetLocalizedMessage()
							if assert.NotNil(t, localizedMsg, "age field should have localized message") {
								expected := "年龄 17 超出有效范围 18-65"
								assert.Equal(t, expected, localizedMsg.GetMessage())
							}
						}
					}
				}
			}
		})

		t.Run("field with struct data using jsonx.MustToMap", func(t *testing.T) {
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

			data := jsonx.MustToMap(validationInfo)
			err := New(codes.InvalidArgument, "validation_failed", "Email validation failed").
				WithFieldLocalized("email", "field.email.detailed", data).
				WithDetails(&errdetails.BadRequest{
					FieldViolations: []*errdetails.BadRequest_FieldViolation{
						{Field: "email", Reason: "INVALID_DOMAIN"},
					},
				}).Err()

			err = TranslateError(ctx6, ib7, err)

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

			err := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
				WithDetails(existingLocalized).
				Err()

			translated := TranslateError(ctx, ib, err)
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

			err := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
				WithDetails(existingBadRequest).
				WithFieldsViolations(&FieldViolation{
					Field:       "password",
					Key:         "field.password.required",
					Description: "Password is required",
					Reason:      "PASSWORD_REQUIRED",
				}).
				Err()

			translated := TranslateError(ctx, ib, err)
			st := status.Convert(translated)

			badRequests := ExtractDetail[*errdetails.BadRequest](st.Details())
			assert.NotNil(t, badRequests, "should preserve existing BadRequest")
			assert.Len(t, badRequests.FieldViolations, 1, "should only have original field violation")
			assert.Equal(t, "email", badRequests.FieldViolations[0].Field)
		})

		t.Run("construct BadRequest when none exists", func(t *testing.T) {
			err := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
				WithFieldsViolations(
					&FieldViolation{
						Field:       "email",
						Key:         "field.email.required",
						Description: "Email address is required",
						Reason:      "EMAIL_REQUIRED",
					},
					&FieldViolation{
						Field:       "password",
						Key:         "field.password.weak",
						Args:        []any{"minLength", 8},
						Description: "Password is too weak",
						Reason:      "PASSWORD_WEAK",
					},
				).
				Err()

			translated := TranslateError(ctx, ib, err)
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
