package login

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type NoticeLevel int

const (
	NoticeLevel_Info NoticeLevel = iota
	NoticeLevel_Warn
	NoticeLevel_Error
)

type NoticeError struct {
	Level   NoticeLevel
	Message string
}

func (e *NoticeError) Error() string {
	return e.Message
}

type FailCode int

const (
	FailCodeSystemError FailCode = iota + 1
	FailCodeCompleteUserAuthFailed
	FailCodeUserNotFound
	FailCodeIncorrectAccountNameOrPassword
	FailCodeUserLocked
	FailCodeAccountIsRequired
	FailCodePasswordCannotBeEmpty
	FailCodePasswordNotMatch
	FailCodeIncorrectPassword
	FailCodeInvalidToken
	FailCodeTokenExpired
	FailCodeIncorrectTOTPCode
	FailCodeTOTPCodeHasBeenUsed
	FailCodeIncorrectRecaptchaToken
)

type WarnCode int

const (
	WarnCodePasswordHasBeenChanged = iota + 1
)

type InfoCode int

const (
	InfoCodePasswordSuccessfullyReset InfoCode = iota + 1
	InfoCodePasswordSuccessfullyChanged
)

const (
	failCodeFlashCookieName = "qor5_fc_flash"
	warnCodeFlashCookieName = "qor5_wc_flash"
	infoCodeFlashCookieName = "qor5_ic_flash"
)

func setFailCodeFlash(w http.ResponseWriter, c FailCode) {
	http.SetCookie(w, &http.Cookie{
		Name:     failCodeFlashCookieName,
		Value:    fmt.Sprint(c),
		Path:     "/",
		HttpOnly: true,
	})
}

func setWarnCodeFlash(w http.ResponseWriter, c WarnCode) {
	http.SetCookie(w, &http.Cookie{
		Name:     warnCodeFlashCookieName,
		Value:    fmt.Sprint(c),
		Path:     "/",
		HttpOnly: true,
	})
}

func setInfoCodeFlash(w http.ResponseWriter, c InfoCode) {
	http.SetCookie(w, &http.Cookie{
		Name:     infoCodeFlashCookieName,
		Value:    fmt.Sprint(c),
		Path:     "/",
		HttpOnly: true,
	})
}

const noticeFlashCookieName = "qor5_notice_flash"

func setNoticeFlash(w http.ResponseWriter, ne *NoticeError) {
	if ne == nil {
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     noticeFlashCookieName,
		Value:    fmt.Sprintf("%d#%s", ne.Level, ne.Message),
		Path:     "/",
		HttpOnly: true,
	})
}

func setNoticeOrFailCodeFlash(w http.ResponseWriter, err error, c FailCode) {
	if err == nil {
		return
	}
	ne, ok := err.(*NoticeError)
	if ok {
		setNoticeFlash(w, ne)
		return
	}
	setFailCodeFlash(w, c)
}

func setNoticeOrPanic(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	var ne *NoticeError
	ok := errors.As(err, &ne)
	if !ok {
		panic(err)
	}
	setNoticeFlash(w, ne)
}

const wrongLoginInputFlashCookieName = "qor5_wli_flash"

type WrongLoginInputFlash struct {
	Account  string
	Password string
}

func (b *Builder) setWrongLoginInputFlash(w http.ResponseWriter, f WrongLoginInputFlash) {
	bf, _ := json.Marshal(&f)
	http.SetCookie(w, &http.Cookie{
		Name:     wrongLoginInputFlashCookieName,
		Value:    base64.StdEncoding.EncodeToString(bf),
		Path:     "/",
		HttpOnly: true,
		Secure:   b.cookieConfig.Secure,
	})
}

const wrongForgetPasswordInputFlashCookieName = "qor5_wfpi_flash"

type WrongForgetPasswordInputFlash struct {
	Account string
	TOTP    string
}

func (b *Builder) setWrongForgetPasswordInputFlash(w http.ResponseWriter, f WrongForgetPasswordInputFlash) {
	bf, _ := json.Marshal(&f)
	http.SetCookie(w, &http.Cookie{
		Name:     wrongForgetPasswordInputFlashCookieName,
		Value:    base64.StdEncoding.EncodeToString(bf),
		Path:     "/",
		HttpOnly: true,
		Secure:   b.cookieConfig.Secure,
	})
}

const wrongResetPasswordInputFlashCookieName = "qor5_wrpi_flash"

type WrongResetPasswordInputFlash struct {
	Password        string
	ConfirmPassword string
	TOTP            string
}

func (b *Builder) setWrongResetPasswordInputFlash(w http.ResponseWriter, f WrongResetPasswordInputFlash) {
	bf, _ := json.Marshal(&f)
	http.SetCookie(w, &http.Cookie{
		Name:     wrongResetPasswordInputFlashCookieName,
		Value:    base64.StdEncoding.EncodeToString(bf),
		Path:     "/",
		HttpOnly: true,
		Secure:   b.cookieConfig.Secure,
	})
}

const wrongChangePasswordInputFlashCookieName = "qor5_wcpi_flash"

type WrongChangePasswordInputFlash struct {
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
	TOTP            string
}

func (b *Builder) setWrongChangePasswordInputFlash(w http.ResponseWriter, f WrongChangePasswordInputFlash) {
	bf, _ := json.Marshal(&f)
	http.SetCookie(w, &http.Cookie{
		Name:     wrongChangePasswordInputFlashCookieName,
		Value:    base64.StdEncoding.EncodeToString(bf),
		Path:     "/",
		HttpOnly: true,
		Secure:   b.cookieConfig.Secure,
	})
}

const secondsToRedoFlashCookieName = "qor5_stre_flash"

func setSecondsToRedoFlash(w http.ResponseWriter, c int) {
	http.SetCookie(w, &http.Cookie{
		Name:     secondsToRedoFlashCookieName,
		Value:    fmt.Sprint(c),
		Path:     "/",
		HttpOnly: true,
	})
}
