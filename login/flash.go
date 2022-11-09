package login

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

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

const failCodeFlashCookieName = "qor5_fc_flash"
const warnCodeFlashCookieName = "qor5_wc_flash"
const infoCodeFlashCookieName = "qor5_ic_flash"

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

const customErrorMessageFlashCookieName = "qor5_cem_flash"

func setCustomErrorMessageFlash(w http.ResponseWriter, f string) {
	http.SetCookie(w, &http.Cookie{
		Name:     customErrorMessageFlashCookieName,
		Value:    f,
		Path:     "/",
		HttpOnly: true,
	})
}

const wrongLoginInputFlashCookieName = "qor5_wli_flash"

type WrongLoginInputFlash struct {
	Account  string
	Password string
}

func setWrongLoginInputFlash(w http.ResponseWriter, f WrongLoginInputFlash) {
	bf, _ := json.Marshal(&f)
	http.SetCookie(w, &http.Cookie{
		Name:     wrongLoginInputFlashCookieName,
		Value:    base64.StdEncoding.EncodeToString(bf),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
}

const wrongForgetPasswordInputFlashCookieName = "qor5_wfpi_flash"

type WrongForgetPasswordInputFlash struct {
	Account string
	TOTP    string
}

func setWrongForgetPasswordInputFlash(w http.ResponseWriter, f WrongForgetPasswordInputFlash) {
	bf, _ := json.Marshal(&f)
	http.SetCookie(w, &http.Cookie{
		Name:     wrongForgetPasswordInputFlashCookieName,
		Value:    base64.StdEncoding.EncodeToString(bf),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
}

const wrongResetPasswordInputFlashCookieName = "qor5_wrpi_flash"

type WrongResetPasswordInputFlash struct {
	Password        string
	ConfirmPassword string
	TOTP            string
}

func setWrongResetPasswordInputFlash(w http.ResponseWriter, f WrongResetPasswordInputFlash) {
	bf, _ := json.Marshal(&f)
	http.SetCookie(w, &http.Cookie{
		Name:     wrongResetPasswordInputFlashCookieName,
		Value:    base64.StdEncoding.EncodeToString(bf),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
}

const wrongChangePasswordInputFlashCookieName = "qor5_wcpi_flash"

type WrongChangePasswordInputFlash struct {
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
	TOTP            string
}

func setWrongChangePasswordInputFlash(w http.ResponseWriter, f WrongChangePasswordInputFlash) {
	bf, _ := json.Marshal(&f)
	http.SetCookie(w, &http.Cookie{
		Name:     wrongChangePasswordInputFlashCookieName,
		Value:    base64.StdEncoding.EncodeToString(bf),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
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
