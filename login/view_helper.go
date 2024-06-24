package login

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/qor5/x/v3/i18n"
)

type ViewHelper struct {
	b *Builder
}

func (vh *ViewHelper) I18n() *i18n.Builder {
	return vh.b.i18nBuilder
}

func (vh *ViewHelper) OAuthEnabled() bool {
	return vh.b.oauthEnabled
}

func (vh *ViewHelper) RecaptchaEnabled() bool {
	return vh.b.recaptchaEnabled
}

func (vh *ViewHelper) UserPassEnabled() bool {
	return vh.b.userPassEnabled
}

func (vh *ViewHelper) TOTPEnabled() bool {
	return vh.b.totpEnabled
}

func (vh *ViewHelper) NoForgetPasswordLink() bool {
	return vh.b.noForgetPasswordLink
}

func (vh *ViewHelper) OAuthProviders() []*Provider {
	return vh.b.providers
}

func (vh *ViewHelper) OAuthBeginURL() string {
	return vh.b.oauthBeginURL
}

func (vh *ViewHelper) PasswordLoginURL() string {
	return vh.b.passwordLoginURL
}

func (vh *ViewHelper) ForgetPasswordPageURL() string {
	return vh.b.forgetPasswordPageURL
}

func (vh *ViewHelper) SendResetPasswordLinkURL() string {
	return vh.b.sendResetPasswordLinkURL
}

func (vh *ViewHelper) ResetPasswordURL() string {
	return vh.b.resetPasswordURL
}

func (vh *ViewHelper) ChangePasswordURL() string {
	return vh.b.changePasswordURL
}

func (vh *ViewHelper) ValidateTOTPURL() string {
	return vh.b.validateTOTPURL
}

func (vh *ViewHelper) RecaptchaSiteKey() string {
	return vh.b.recaptchaConfig.SiteKey
}

func (vh *ViewHelper) TOTPIssuer() string {
	return vh.b.totpConfig.Issuer
}

func (vh *ViewHelper) FindUserByID(id string) (user interface{}, err error) {
	return vh.b.findUserByID(id)
}

func (vh *ViewHelper) GetFailCodeFlash(w http.ResponseWriter, r *http.Request) FailCode {
	c, err := r.Cookie(failCodeFlashCookieName)
	if err != nil {
		return 0
	}
	http.SetCookie(w, &http.Cookie{
		Name:     failCodeFlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	v, _ := strconv.Atoi(c.Value)
	return FailCode(v)
}

func (vh *ViewHelper) GetWarnCodeFlash(w http.ResponseWriter, r *http.Request) WarnCode {
	c, err := r.Cookie(warnCodeFlashCookieName)
	if err != nil {
		return 0
	}
	http.SetCookie(w, &http.Cookie{
		Name:     warnCodeFlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	v, _ := strconv.Atoi(c.Value)
	return WarnCode(v)
}

func (vh *ViewHelper) GetInfoCodeFlash(w http.ResponseWriter, r *http.Request) InfoCode {
	c, err := r.Cookie(infoCodeFlashCookieName)
	if err != nil {
		return 0
	}
	http.SetCookie(w, &http.Cookie{
		Name:     infoCodeFlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	v, _ := strconv.Atoi(c.Value)
	return InfoCode(v)
}

func (vh *ViewHelper) GetNoticeFlash(w http.ResponseWriter, r *http.Request) *NoticeError {
	c, err := r.Cookie(noticeFlashCookieName)
	if err != nil {
		return nil
	}
	var level NoticeLevel
	var message string
	{
		vs := strings.SplitN(c.Value, "#", 2)
		if len(vs) != 2 {
			return nil
		}
		n, err := strconv.Atoi(vs[0])
		if err != nil {
			return nil
		}
		level = NoticeLevel(n)
		message = vs[1]
	}
	http.SetCookie(w, &http.Cookie{
		Name:     noticeFlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	return &NoticeError{
		Level:   level,
		Message: message,
	}
}

func (vh *ViewHelper) GetWrongLoginInputFlash(w http.ResponseWriter, r *http.Request) WrongLoginInputFlash {
	c, err := r.Cookie(wrongLoginInputFlashCookieName)
	if err != nil {
		return WrongLoginInputFlash{}
	}
	http.SetCookie(w, &http.Cookie{
		Name:     wrongLoginInputFlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   vh.b.cookieConfig.Secure,
	})
	v, _ := base64.StdEncoding.DecodeString(c.Value)
	wi := WrongLoginInputFlash{}
	json.Unmarshal([]byte(v), &wi)
	return wi
}

func (vh *ViewHelper) GetWrongForgetPasswordInputFlash(w http.ResponseWriter, r *http.Request) WrongForgetPasswordInputFlash {
	c, err := r.Cookie(wrongForgetPasswordInputFlashCookieName)
	if err != nil {
		return WrongForgetPasswordInputFlash{}
	}
	http.SetCookie(w, &http.Cookie{
		Name:     wrongForgetPasswordInputFlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   vh.b.cookieConfig.Secure,
	})
	v, _ := base64.StdEncoding.DecodeString(c.Value)
	f := WrongForgetPasswordInputFlash{}
	json.Unmarshal([]byte(v), &f)
	return f
}

func (vh *ViewHelper) GetWrongResetPasswordInputFlash(w http.ResponseWriter, r *http.Request) WrongResetPasswordInputFlash {
	c, err := r.Cookie(wrongResetPasswordInputFlashCookieName)
	if err != nil {
		return WrongResetPasswordInputFlash{}
	}
	http.SetCookie(w, &http.Cookie{
		Name:     wrongResetPasswordInputFlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   vh.b.cookieConfig.Secure,
	})
	v, _ := base64.StdEncoding.DecodeString(c.Value)
	f := WrongResetPasswordInputFlash{}
	json.Unmarshal([]byte(v), &f)
	return f
}

func (vh *ViewHelper) GetWrongChangePasswordInputFlash(w http.ResponseWriter, r *http.Request) WrongChangePasswordInputFlash {
	c, err := r.Cookie(wrongChangePasswordInputFlashCookieName)
	if err != nil {
		return WrongChangePasswordInputFlash{}
	}
	http.SetCookie(w, &http.Cookie{
		Name:     wrongChangePasswordInputFlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   vh.b.cookieConfig.Secure,
	})
	v, _ := base64.StdEncoding.DecodeString(c.Value)
	f := WrongChangePasswordInputFlash{}
	json.Unmarshal([]byte(v), &f)
	return f
}

func (vh *ViewHelper) GetSecondsToRedoFlash(w http.ResponseWriter, r *http.Request) int {
	c, err := r.Cookie(secondsToRedoFlashCookieName)
	if err != nil {
		return 0
	}
	http.SetCookie(w, &http.Cookie{
		Name:     secondsToRedoFlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	v, _ := strconv.Atoi(c.Value)
	return v
}

func (vh *ViewHelper) GetFailFlashMessage(msgr *Messages, w http.ResponseWriter, r *http.Request) string {
	code := vh.GetFailCodeFlash(w, r)
	switch code {
	case FailCodeSystemError:
		return msgr.ErrorSystemError
	case FailCodeCompleteUserAuthFailed:
		return msgr.ErrorCompleteUserAuthFailed
	case FailCodeUserNotFound:
		return msgr.ErrorUserNotFound
	case FailCodeIncorrectAccountNameOrPassword:
		return msgr.ErrorIncorrectAccountNameOrPassword
	case FailCodeUserLocked:
		return msgr.ErrorUserLocked
	case FailCodeAccountIsRequired:
		return msgr.ErrorAccountIsRequired
	case FailCodePasswordCannotBeEmpty:
		return msgr.ErrorPasswordCannotBeEmpty
	case FailCodePasswordNotMatch:
		return msgr.ErrorPasswordNotMatch
	case FailCodeIncorrectPassword:
		return msgr.ErrorIncorrectPassword
	case FailCodeInvalidToken:
		return msgr.ErrorInvalidToken
	case FailCodeTokenExpired:
		return msgr.ErrorTokenExpired
	case FailCodeIncorrectTOTPCode:
		return msgr.ErrorIncorrectTOTPCode
	case FailCodeTOTPCodeHasBeenUsed:
		return msgr.ErrorTOTPCodeReused
	case FailCodeIncorrectRecaptchaToken:
		return msgr.ErrorIncorrectRecaptchaToken
	}

	return ""
}

func (vh *ViewHelper) GetWarnFlashMessage(msgr *Messages, w http.ResponseWriter, r *http.Request) string {
	code := vh.GetWarnCodeFlash(w, r)
	switch code {
	case WarnCodePasswordHasBeenChanged:
		return msgr.WarnPasswordHasBeenChanged
	}
	return ""
}

func (vh *ViewHelper) GetInfoFlashMessage(msgr *Messages, w http.ResponseWriter, r *http.Request) string {
	code := vh.GetInfoCodeFlash(w, r)
	switch code {
	case InfoCodePasswordSuccessfullyReset:
		return msgr.InfoPasswordSuccessfullyReset
	case InfoCodePasswordSuccessfullyChanged:
		return msgr.InfoPasswordSuccessfullyChanged
	}
	return ""
}
