package login

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"net/http"
	"net/url"

	"github.com/pquerna/otp"
	"github.com/qor5/web/v3"
	"github.com/qor5/x/v3/i18n"
	. "github.com/theplant/htmlgo"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

func defaultLoginPage(vh *ViewHelper) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		// i18n start
		msgr := i18n.MustGetModuleMessages(ctx.R, I18nLoginKey, Messages_en_US).(*Messages)
		i18nBuilder := vh.I18n()
		var languagesHTML []HTMLComponent
		languages := i18nBuilder.GetSupportLanguages()
		if len(languages) > 1 {
			qn := i18nBuilder.GetQueryName()
			lang := ctx.R.FormValue(qn)
			if lang == "" {
				lang = i18nBuilder.GetCurrentLangFromCookie(ctx.R)
			}
			accept := ctx.R.Header.Get("Accept-Language")
			_, mi := language.MatchStrings(language.NewMatcher(languages), lang, accept)
			for i, l := range languages {
				u, _ := url.Parse(ctx.R.RequestURI)
				qs := u.Query()
				qs.Set(qn, l.String())
				u.RawQuery = qs.Encode()
				elem := Option(display.Self.Name(l)).
					Value(u.String())
				if i == mi {
					elem.Attr("selected", "selected")
				}
				languagesHTML = append(languagesHTML, elem)
			}
		}
		// i18n end

		var oauthHTML HTMLComponent
		if vh.OAuthEnabled() {
			ul := Div().Class("flex flex-col justify-center mt-8 text-center")
			for _, provider := range vh.OAuthProviders() {
				ul.AppendChildren(
					A().
						Href(fmt.Sprintf("%s?provider=%s", vh.OAuthBeginURL(), provider.Key)).
						Class("px-6 py-3 mt-4 font-semibold text-gray-900 bg-white border-2 border-gray-500 rounded-md shadow outline-none hover:bg-yellow-50 hover:border-yellow-400 focus:outline-none").
						Children(
							provider.Logo,
							Text(provider.Text),
						),
				)
			}

			oauthHTML = Div(
				ul,
			)
		}

		wIn := vh.GetWrongLoginInputFlash(ctx.W, ctx.R)
		isRecaptchaEnabled := vh.RecaptchaEnabled()

		var userPassHTML HTMLComponent
		if vh.UserPassEnabled() {
			userPassHTML = Div(
				Form(
					Div(
						Label(msgr.AccountLabel).Class(DefaultViewCommon.LabelClass).For("account"),
						Input("account").Placeholder(msgr.AccountPlaceholder).Class(DefaultViewCommon.InputClass).
							Value(wIn.Account),
					),
					Div(
						Label(msgr.PasswordLabel).Class(DefaultViewCommon.LabelClass).For("password"),
						DefaultViewCommon.PasswordInputWithRevealFunction("password", msgr.PasswordPlaceholder, "password", wIn.Password),
					).Class("mt-6"),
					If(isRecaptchaEnabled,
						Div(
							// recaptcha response token
							Input("token").Id("token"),
						).Class("hidden"),
					),
					Div(
						Button(msgr.SignInBtn).Class(DefaultViewCommon.ButtonClass).
							ClassIf("g-recaptcha", isRecaptchaEnabled).AttrIf("data-sitekey", vh.RecaptchaSiteKey(), isRecaptchaEnabled).AttrIf("data-callback", "onSubmit", isRecaptchaEnabled),
					).Class("mt-6"),
				).Id("login-form").Method(http.MethodPost).Action(vh.PasswordLoginURL()),
				If(!vh.NoForgetPasswordLink(),
					Div(
						A(Text(msgr.ForgetPasswordLink)).Href(vh.ForgetPasswordPageURL()).
							Class("text-gray-500"),
					).Class("text-right mt-2"),
				),
			)
		}

		r.PageTitle = msgr.LoginPageTitle
		var bodyForm HTMLComponent = Div(
			userPassHTML,
			oauthHTML,
			If(len(languagesHTML) > 0,
				Select(
					languagesHTML...,
				).Class("mt-12 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500").
					Attr("onChange", "window.location.href=this.value"),
			),
		).Class(DefaultViewCommon.WrapperClass)

		r.Body = Div(
			Link(StyleCSSURL).Type("text/css").Rel("stylesheet"),
			If(isRecaptchaEnabled,
				Style(`.grecaptcha-badge { visibility: hidden; }`),
				Script("").Src("https://www.google.com/recaptcha/api.js"),
				Script(`
function onSubmit(token) {
	document.getElementById("token").value = token;
	document.getElementById("login-form").submit();
}
`)),
			DefaultViewCommon.Notice(vh, msgr, ctx.W, ctx.R),
			bodyForm,
		)

		return
	}
}

func defaultForgetPasswordPage(vh *ViewHelper) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		msgr := i18n.MustGetModuleMessages(ctx.R, I18nLoginKey, Messages_en_US).(*Messages)

		wIn := vh.GetWrongForgetPasswordInputFlash(ctx.W, ctx.R)
		secondsToResend := vh.GetSecondsToRedoFlash(ctx.W, ctx.R)
		activeBtnText := msgr.SendResetPasswordEmailBtn
		activeBtnClass := DefaultViewCommon.ButtonClass
		inactiveBtnText := msgr.ResendResetPasswordEmailBtn
		inactiveBtnClass := "w-full px-6 py-3 tracking-wide text-white transition-colors duration-200 transform bg-gray-500 rounded-md"
		inactiveBtnTextWithInitSeconds := fmt.Sprintf("%s (%d)", inactiveBtnText, secondsToResend)

		doTOTP := ctx.R.URL.Query().Get("totp") == "1"
		actionURL := vh.SendResetPasswordLinkURL()
		if doTOTP {
			actionURL = MustSetQuery(actionURL, "totp", "1")
		}

		isRecaptchaEnabled := vh.RecaptchaEnabled()

		r.PageTitle = msgr.ForgetPasswordPageTitle
		r.Body = Div(
			Link(StyleCSSURL).Type("text/css").Rel("stylesheet"),
			If(isRecaptchaEnabled,
				Style(`.grecaptcha-badge { visibility: hidden; }`),
				Script("").Src("https://www.google.com/recaptcha/api.js"),
				Script(`
function onSubmit(token) {
	document.getElementById("token").value = token;
	document.getElementById("forget-form").submit();
}
`)),
			DefaultViewCommon.Notice(vh, msgr, ctx.W, ctx.R),
			If(secondsToResend > 0,
				DefaultViewCommon.WarnNotice(msgr.SendEmailTooFrequentlyNotice),
			),
			Div(
				H1(msgr.ForgotMyPasswordTitle).Class(DefaultViewCommon.TitleClass),
				Form(
					Div(
						Label(msgr.ForgetPasswordEmailLabel).Class(DefaultViewCommon.LabelClass).For("account"),
						Input("account").Placeholder(msgr.ForgetPasswordEmailPlaceholder).Class(DefaultViewCommon.InputClass).Value(wIn.Account),
					),
					If(doTOTP,
						Div(
							Label(msgr.TOTPValidateCodeLabel).Class(DefaultViewCommon.LabelClass).For("otp"),
							Input("otp").Placeholder(msgr.TOTPValidateCodePlaceholder).
								Class(DefaultViewCommon.InputClass).
								Value(wIn.TOTP),
						).Class("mt-6"),
					),
					If(isRecaptchaEnabled,
						Div(
							// recaptcha response token
							Input("token").Id("token"),
						).Class("hidden"),
					),
					Div(
						If(secondsToResend > 0,
							Button(inactiveBtnTextWithInitSeconds).Id("submitBtn").Class(inactiveBtnClass).Disabled(true),
						).Else(
							Button(activeBtnText).Class(activeBtnClass).
								ClassIf("g-recaptcha", isRecaptchaEnabled).AttrIf("data-sitekey", vh.RecaptchaSiteKey(), isRecaptchaEnabled).AttrIf("data-callback", "onSubmit", isRecaptchaEnabled),
						),
					).Class("mt-6"),
				).Id("forget-form").Method(http.MethodPost).Action(actionURL),
			).Class(DefaultViewCommon.WrapperClass),
		)

		if secondsToResend > 0 {
			ctx.Injector.TailHTML(fmt.Sprintf(`
<script>
(function(){
    var secondsToResend = %d;
    var btnText = "%s";
    var submitBtn = document.getElementById("submitBtn");
    var interv = setInterval(function(){
        secondsToResend--;
        if (secondsToResend === 0) {
            clearInterval(interv);
            submitBtn.innerText = btnText;
            submitBtn.className = "%s";
            submitBtn.disabled = false;
            return;
        }
        submitBtn.innerText = btnText + " (" + secondsToResend + ")" ;
    }, 1000);
})();
</script>
        `, secondsToResend, inactiveBtnText, activeBtnClass))
		}
		return
	}
}

func defaultResetPasswordLinkSentPage(vh *ViewHelper) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		msgr := i18n.MustGetModuleMessages(ctx.R, I18nLoginKey, Messages_en_US).(*Messages)

		a := ctx.R.URL.Query().Get("a")

		r.PageTitle = msgr.ResetPasswordLinkSentPageTitle
		r.Body = Div(
			Link(StyleCSSURL).Type("text/css").Rel("stylesheet"),
			DefaultViewCommon.Notice(vh, msgr, ctx.W, ctx.R),
			Div(
				H1(fmt.Sprintf("%s %s.", msgr.ResetPasswordLinkWasSentTo, a)).Class("leading-tight text-2xl mt-0 mb-4"),
				H2(msgr.ResetPasswordLinkSentPrompt).Class("leading-tight text-1xl mt-0"),
			).Class(DefaultViewCommon.WrapperClass),
		)
		return
	}
}

func defaultResetPasswordPage(vh *ViewHelper) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		msgr := i18n.MustGetModuleMessages(ctx.R, I18nLoginKey, Messages_en_US).(*Messages)

		wIn := vh.GetWrongResetPasswordInputFlash(ctx.W, ctx.R)

		doTOTP := ctx.R.URL.Query().Get("totp") == "1"
		actionURL := vh.ResetPasswordURL()
		if doTOTP {
			actionURL = MustSetQuery(actionURL, "totp", "1")
		}

		var user interface{}

		r.PageTitle = msgr.ResetPasswordPageTitle

		query := ctx.R.URL.Query()
		id := query.Get("id")
		if id == "" {
			r.Body = Div(Text("user not found"))
			return r, nil
		} else {
			user, err = vh.FindUserByID(id)
			if err != nil {
				if err == ErrUserNotFound {
					r.Body = Div(Text("user not found"))
					return r, nil
				}
				panic(err)
			}
		}
		token := query.Get("token")
		if token == "" {
			r.Body = Div(Text("invalid token"))
			return r, nil
		} else {
			storedToken, _, expired := user.(UserPasser).GetResetPasswordToken()
			if expired {
				r.Body = Div(Text("token expired"))
				return r, nil
			}
			if token != storedToken {
				r.Body = Div(Text("invalid token"))
				return r, nil
			}
		}

		r.Body = Div(
			Link(StyleCSSURL).Type("text/css").Rel("stylesheet"),
			Script("").Src(ZxcvbnJSURL),
			DefaultViewCommon.Notice(vh, msgr, ctx.W, ctx.R),
			Div(
				H1(msgr.ResetYourPasswordTitle).Class(DefaultViewCommon.TitleClass),
				Form(
					Input("user_id").Type("hidden").Value(id),
					Input("token").Type("hidden").Value(token),
					Div(
						Label(msgr.ResetPasswordLabel).Class(DefaultViewCommon.LabelClass).For("password"),
						DefaultViewCommon.PasswordInputWithRevealFunction("password", msgr.ResetPasswordPlaceholder, "password", wIn.Password),
						DefaultViewCommon.PasswordStrengthMeter("password"),
					),
					Div(
						Label(msgr.ResetPasswordConfirmLabel).Class(DefaultViewCommon.LabelClass).For("confirm_password"),
						DefaultViewCommon.PasswordInputWithRevealFunction("confirm_password", msgr.ResetPasswordConfirmPlaceholder, "confirm_password", wIn.ConfirmPassword),
					).Class("mt-6"),
					If(doTOTP,
						Div(
							Label(msgr.TOTPValidateCodeLabel).Class(DefaultViewCommon.LabelClass).For("otp"),
							Input("otp").Placeholder(msgr.TOTPValidateCodePlaceholder).
								Class(DefaultViewCommon.InputClass).
								Value(wIn.TOTP),
						).Class("mt-6"),
					),
					Div(
						Button(msgr.Confirm).Class(DefaultViewCommon.ButtonClass),
					).Class("mt-6"),
				).Method(http.MethodPost).Action(actionURL),
			).Class(DefaultViewCommon.WrapperClass),
		)
		return
	}
}

func defaultChangePasswordPage(vh *ViewHelper) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		msgr := i18n.MustGetModuleMessages(ctx.R, I18nLoginKey, Messages_en_US).(*Messages)

		wIn := vh.GetWrongChangePasswordInputFlash(ctx.W, ctx.R)

		r.PageTitle = msgr.ChangePasswordPageTitle

		r.Body = Div(
			Link(StyleCSSURL).Type("text/css").Rel("stylesheet"),
			Script("").Src(ZxcvbnJSURL),
			DefaultViewCommon.Notice(vh, msgr, ctx.W, ctx.R),
			Div(
				H1(msgr.ChangePasswordTitle).Class(DefaultViewCommon.TitleClass),
				Form(
					Div(
						Label(msgr.ChangePasswordOldLabel).Class(DefaultViewCommon.LabelClass).For("old_password"),
						DefaultViewCommon.PasswordInputWithRevealFunction("old_password", msgr.ChangePasswordOldPlaceholder, "old_password", wIn.OldPassword),
					),
					Div(
						Label(msgr.ChangePasswordNewLabel).Class(DefaultViewCommon.LabelClass).For("password"),
						DefaultViewCommon.PasswordInputWithRevealFunction("password", msgr.ChangePasswordNewPlaceholder, "password", wIn.NewPassword),
						DefaultViewCommon.PasswordStrengthMeter("password"),
					).Class("mt-6"),
					Div(
						Label(msgr.ChangePasswordNewConfirmLabel).Class(DefaultViewCommon.LabelClass).For("confirm_password"),
						DefaultViewCommon.PasswordInputWithRevealFunction("confirm_password", msgr.ChangePasswordNewConfirmPlaceholder, "confirm_password", wIn.ConfirmPassword),
					).Class("mt-6"),
					If(vh.TOTPEnabled(),
						Div(
							Label(msgr.TOTPValidateCodeLabel).Class(DefaultViewCommon.LabelClass).For("otp"),
							Input("otp").Placeholder(msgr.TOTPValidateCodePlaceholder).
								Class(DefaultViewCommon.InputClass).
								Value(wIn.TOTP),
						).Class("mt-6"),
					),
					Div(
						Button(msgr.Confirm).Class(DefaultViewCommon.ButtonClass),
					).Class("mt-6"),
				).Method(http.MethodPost).Action(vh.ChangePasswordURL()),
			).Class(DefaultViewCommon.WrapperClass),
		)
		return
	}
}

func defaultTOTPSetupPage(vh *ViewHelper) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		msgr := i18n.MustGetModuleMessages(ctx.R, I18nLoginKey, Messages_en_US).(*Messages)

		user := GetCurrentUser(ctx.R)
		u := user.(UserPasser)

		var QRCode bytes.Buffer

		// Generate key from TOTPSecret
		var key *otp.Key
		totpSecret := u.GetTOTPSecret()
		if len(totpSecret) == 0 {
			r.Body = DefaultViewCommon.ErrorBody("need setup totp")
			return
		}
		key, err = otp.NewKeyFromURL(
			fmt.Sprintf("otpauth://totp/%s:%s?issuer=%s&secret=%s",
				url.PathEscape(vh.TOTPIssuer()),
				url.PathEscape(u.GetAccountName()),
				url.QueryEscape(vh.TOTPIssuer()),
				url.QueryEscape(totpSecret),
			),
		)

		img, err := key.Image(200, 200)
		if err != nil {
			r.Body = DefaultViewCommon.ErrorBody(err.Error())
			return
		}

		err = png.Encode(&QRCode, img)
		if err != nil {
			r.Body = DefaultViewCommon.ErrorBody(err.Error())
			return
		}

		r.PageTitle = msgr.TOTPSetupPageTitle
		r.Body = Div(
			Link(StyleCSSURL).Type("text/css").Rel("stylesheet"),
			DefaultViewCommon.Notice(vh, msgr, ctx.W, ctx.R),
			Div(
				Div(
					H1(msgr.TOTPSetupTitle).
						Class(DefaultViewCommon.TitleClass),
					Label(msgr.TOTPSetupScanPrompt),
				),
				Div(
					Img(fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(QRCode.Bytes()))),
				).Class("my-2 flex items-center justify-center"),
				Div(
					Label(msgr.TOTPSetupSecretPrompt),
				),
				Div(Label(u.GetTOTPSecret()).Class("text-sm font-bold")).Class("my-4"),
				Form(
					Label(msgr.TOTPSetupEnterCodePrompt),
					Input("otp").Placeholder(msgr.TOTPSetupCodePlaceholder).
						Class(DefaultViewCommon.InputClass).
						Class("mt-6"),
					Div(
						Button(msgr.Verify).Class(DefaultViewCommon.ButtonClass),
					).Class("mt-6"),
				).Method(http.MethodPost).Action(vh.ValidateTOTPURL()),
			).Class(DefaultViewCommon.WrapperClass).Class("text-center"),
		)

		return
	}
}

func defaultTOTPValidatePage(vh *ViewHelper) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		msgr := i18n.MustGetModuleMessages(ctx.R, I18nLoginKey, Messages_en_US).(*Messages)

		r.PageTitle = msgr.TOTPValidatePageTitle
		r.Body = Div(
			Link(StyleCSSURL).Type("text/css").Rel("stylesheet"),
			DefaultViewCommon.Notice(vh, msgr, ctx.W, ctx.R),
			Div(
				Div(
					H1(msgr.TOTPValidateTitle).
						Class(DefaultViewCommon.TitleClass),
					Label(msgr.TOTPValidateEnterCodePrompt),
				),
				Form(
					Input("otp").Placeholder(msgr.TOTPValidateCodePlaceholder).
						Class(DefaultViewCommon.InputClass).
						Class("mt-6").
						Attr("autofocus", true),
					Div(
						Button(msgr.Verify).Class(DefaultViewCommon.ButtonClass),
					).Class("mt-6"),
				).Method(http.MethodPost).Action(vh.ValidateTOTPURL()),
			).Class(DefaultViewCommon.WrapperClass).Class("text-center"),
		)

		return
	}
}
