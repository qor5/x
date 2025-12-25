package login

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/qor5/web/v3"
	h "github.com/theplant/htmlgo"
	"golang.org/x/text/language"
	"gorm.io/gorm"

	"github.com/qor5/x/v3/i18n"
)

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrPasswordChanged     = errors.New("password changed")
	ErrWrongPassword       = errors.New("wrong password")
	ErrUserLocked          = errors.New("user locked")
	ErrUserGetLocked       = errors.New("user get locked")
	ErrWrongTOTPCode       = errors.New("wrong totp code")
	ErrTOTPCodeHasBeenUsed = errors.New("totp code has been used")
	ErrEmptyPassword       = errors.New("empty password")
	ErrPasswordNotMatch    = errors.New("password not match")
)

type (
	HomeURLFunc func(r *http.Request, user interface{}) string
	HookFunc    func(r *http.Request, user interface{}, extraVals ...interface{}) error
)

var NopHookFunc HookFunc = func(r *http.Request, user interface{}, extraVals ...interface{}) error {
	return nil
}

type Provider struct {
	Goth goth.Provider
	Key  string
	Text string
	Logo h.HTMLComponent
}

type CookieConfig struct {
	Path     string
	Domain   string
	Secure   bool
	SameSite http.SameSite
}

type TOTPConfig struct {
	Issuer string
}

type RecaptchaConfig struct {
	SiteKey   string
	SecretKey string
}

type Builder struct {
	secret                string
	providers             []*Provider
	authCookieName        string
	authSecureCookieName  string
	continueUrlCookieName string
	// seconds
	sessionMaxAge        int
	cookieConfig         CookieConfig
	totpEnabled          bool
	totpConfig           TOTPConfig
	recaptchaEnabled     bool
	recaptchaConfig      RecaptchaConfig
	autoExtendSession    bool
	maxRetryCount        int
	noForgetPasswordLink bool
	i18nBuilder          *i18n.Builder

	// Common URLs
	homePageURLFunc HomeURLFunc
	loginPageURL    string
	LogoutURL       string

	// TOTP URLs
	validateTOTPURL     string
	totpSetupPageURL    string
	totpValidatePageURL string

	// OAuth URLs
	oauthBeginURL            string
	oauthCallbackURL         string
	oauthCallbackCompleteURL string

	// UserPass URLs
	passwordLoginURL             string
	resetPasswordURL             string
	resetPasswordPageURL         string
	changePasswordURL            string
	changePasswordPageURL        string
	forgetPasswordPageURL        string
	sendResetPasswordLinkURL     string
	resetPasswordLinkSentPageURL string

	loginPageFunc                 web.PageFunc
	forgetPasswordPageFunc        web.PageFunc
	resetPasswordLinkSentPageFunc web.PageFunc
	resetPasswordPageFunc         web.PageFunc
	changePasswordPageFunc        web.PageFunc
	totpSetupPageFunc             web.PageFunc
	totpValidatePageFunc          web.PageFunc

	beforeSetPasswordHook HookFunc
	beforeTOTPFlowHook    HookFunc

	afterLoginHook                        HookFunc
	afterFailedToLoginHook                HookFunc
	afterUserLockedHook                   HookFunc
	afterLogoutHook                       HookFunc
	afterConfirmSendResetPasswordLinkHook HookFunc
	afterResetPasswordHook                HookFunc
	afterChangePasswordHook               HookFunc
	afterExtendSessionHook                HookFunc
	afterTOTPCodeReusedHook               HookFunc
	afterOAuthCompleteHook                HookFunc

	db                   *gorm.DB
	userModel            interface{}
	snakePrimaryField    string
	tUser                reflect.Type
	userPassEnabled      bool
	oauthEnabled         bool
	sessionSecureEnabled bool
	// key is provider
	oauthIdentifiers map[string]OAuthIdentifier
}

func New() *Builder {
	r := &Builder{
		authCookieName:        "auth",
		authSecureCookieName:  "qor5_auth_secure",
		continueUrlCookieName: "qor5_continue_url",

		homePageURLFunc: func(r *http.Request, user interface{}) string {
			return "/"
		},
		loginPageURL: "/auth/login",
		LogoutURL:    "/auth/logout",

		validateTOTPURL:     "/auth/2fa/totp/do",
		totpSetupPageURL:    "/auth/2fa/totp/setup",
		totpValidatePageURL: "/auth/2fa/totp/validate",

		oauthBeginURL:            "/auth/begin",
		oauthCallbackURL:         "/auth/callback",
		oauthCallbackCompleteURL: "/auth/callback-complete",

		passwordLoginURL:             "/auth/userpass/login",
		resetPasswordURL:             "/auth/do-reset-password",
		resetPasswordPageURL:         "/auth/reset-password",
		changePasswordURL:            "/auth/do-change-password",
		changePasswordPageURL:        "/auth/change-password",
		forgetPasswordPageURL:        "/auth/forget-password",
		sendResetPasswordLinkURL:     "/auth/send-reset-password-link",
		resetPasswordLinkSentPageURL: "/auth/reset-password-link-sent",

		sessionMaxAge: 60 * 60,
		cookieConfig: CookieConfig{
			Path:     "/",
			Domain:   "",
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		},
		autoExtendSession: true,
		maxRetryCount:     5,
		totpEnabled:       true,
		totpConfig: TOTPConfig{
			Issuer: "QOR5",
		},
		oauthIdentifiers: make(map[string]OAuthIdentifier),
	}

	i18nB := i18n.New()
	i18nB.SupportLanguages(language.English, language.SimplifiedChinese, language.Japanese)
	r.I18n(i18nB)

	vh := r.ViewHelper()
	r.loginPageFunc = defaultLoginPage(vh)
	r.forgetPasswordPageFunc = defaultForgetPasswordPage(vh)
	r.resetPasswordLinkSentPageFunc = defaultResetPasswordLinkSentPage(vh)
	r.resetPasswordPageFunc = defaultResetPasswordPage(vh)
	r.changePasswordPageFunc = defaultChangePasswordPage(vh)
	r.totpSetupPageFunc = defaultTOTPSetupPage(vh)
	r.totpValidatePageFunc = defaultTOTPValidatePage(vh)

	return r
}

func (b *Builder) Secret(v string) (r *Builder) {
	b.secret = v
	return b
}

func (b *Builder) CookieSecure(v bool) (r *Builder) {
	b.cookieConfig.Secure = v
	return b
}

func (b *Builder) CookieConfig(v CookieConfig) (r *Builder) {
	b.cookieConfig = v
	return b
}

// Google reCAPTCHA.
func (b *Builder) Recaptcha(enable bool, config ...RecaptchaConfig) (r *Builder) {
	b.recaptchaEnabled = enable
	if len(config) > 0 {
		b.recaptchaConfig = config[0]
	}
	if enable {
		if b.recaptchaConfig.SiteKey == "" {
			panic("SiteKey is empty")
		}
		if b.recaptchaConfig.SecretKey == "" {
			panic("SecretKey is empty")
		}
	}
	return b
}

func (b *Builder) OAuthProviders(vs ...*Provider) (r *Builder) {
	if len(vs) == 0 {
		return b
	}
	b.oauthEnabled = true
	b.providers = vs
	var gothProviders []goth.Provider
	for _, v := range vs {
		gothProviders = append(gothProviders, v.Goth)
	}
	goth.UseProviders(gothProviders...)
	return b
}

func (b *Builder) AuthCookieName(v string) (r *Builder) {
	b.authCookieName = v
	return b
}

func (b *Builder) AuthSecureCookieName(v string) *Builder {
	b.authSecureCookieName = v
	return b
}

func (b *Builder) ContinueUrlCookieName(v string) *Builder {
	b.continueUrlCookieName = v
	return b
}

func (b *Builder) HomeURLFunc(v HomeURLFunc) (r *Builder) {
	b.homePageURLFunc = v
	return b
}

func (b *Builder) URIPrefix(v string) (r *Builder) {
	prefix := strings.TrimRight(v, "/")

	b.loginPageURL = prefix + b.loginPageURL
	b.LogoutURL = prefix + b.LogoutURL
	b.validateTOTPURL = prefix + b.validateTOTPURL
	b.totpSetupPageURL = prefix + b.totpSetupPageURL
	b.totpValidatePageURL = prefix + b.totpValidatePageURL
	b.oauthBeginURL = prefix + b.oauthBeginURL
	b.oauthCallbackURL = prefix + b.oauthCallbackURL
	b.oauthCallbackCompleteURL = prefix + b.oauthCallbackCompleteURL
	b.passwordLoginURL = prefix + b.passwordLoginURL
	b.resetPasswordURL = prefix + b.resetPasswordURL
	b.resetPasswordPageURL = prefix + b.resetPasswordPageURL
	b.changePasswordURL = prefix + b.changePasswordURL
	b.changePasswordPageURL = prefix + b.changePasswordPageURL
	b.forgetPasswordPageURL = prefix + b.forgetPasswordPageURL
	b.sendResetPasswordLinkURL = prefix + b.sendResetPasswordLinkURL
	b.resetPasswordLinkSentPageURL = prefix + b.resetPasswordLinkSentPageURL

	return b
}

func (b *Builder) LoginPageURL(v string) (r *Builder) {
	b.loginPageURL = v
	return b
}

func (b *Builder) ResetPasswordPageURL(v string) (r *Builder) {
	b.resetPasswordPageURL = v
	return b
}

func (b *Builder) ChangePasswordPageURL(v string) (r *Builder) {
	b.changePasswordPageURL = v
	return b
}

func (b *Builder) ForgetPasswordPageURL(v string) (r *Builder) {
	b.forgetPasswordPageURL = v
	return b
}

func (b *Builder) ResetPasswordLinkSentPageURL(v string) (r *Builder) {
	b.resetPasswordLinkSentPageURL = v
	return b
}

func (b *Builder) TOTPSetupPageURL(v string) (r *Builder) {
	b.totpSetupPageURL = v
	return b
}

func (b *Builder) TOTPValidatePageURL(v string) (r *Builder) {
	b.totpValidatePageURL = v
	return b
}

func (b *Builder) LoginPageFunc(v web.PageFunc) (r *Builder) {
	b.loginPageFunc = v
	return b
}

func (b *Builder) ForgetPasswordPageFunc(v web.PageFunc) (r *Builder) {
	b.forgetPasswordPageFunc = v
	return b
}

func (b *Builder) ResetPasswordLinkSentPageFunc(v web.PageFunc) (r *Builder) {
	b.resetPasswordLinkSentPageFunc = v
	return b
}

func (b *Builder) ResetPasswordPageFunc(v web.PageFunc) (r *Builder) {
	b.resetPasswordPageFunc = v
	return b
}

func (b *Builder) ChangePasswordPageFunc(v web.PageFunc) (r *Builder) {
	b.changePasswordPageFunc = v
	return b
}

func (b *Builder) TOTPSetupPageFunc(v web.PageFunc) (r *Builder) {
	b.totpSetupPageFunc = v
	return b
}

func (b *Builder) TOTPValidatePageFunc(v web.PageFunc) (r *Builder) {
	b.totpValidatePageFunc = v
	return b
}

func (b *Builder) wrapHook(v HookFunc) HookFunc {
	if v == nil {
		return nil
	}

	return func(r *http.Request, user interface{}, extraVals ...interface{}) error {
		if user != nil && GetCurrentUser(r) == nil {
			r = r.WithContext(context.WithValue(r.Context(), UserKey, user))
		}
		return v(r, user, extraVals...)
	}
}

// extra vals:
// - password
func (b *Builder) BeforeSetPassword(v HookFunc) (r *Builder) {
	b.beforeSetPasswordHook = v
	return b
}

func (b *Builder) WrapBeforeSetPassword(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.beforeSetPasswordHook == nil {
		b.beforeSetPasswordHook = w(NopHookFunc)
	} else {
		b.beforeSetPasswordHook = w(b.beforeSetPasswordHook)
	}
	return b
}

func (b *Builder) BeforeTOTPFlow(v HookFunc) (r *Builder) {
	b.beforeTOTPFlowHook = v
	return b
}

func (b *Builder) WrapBeforeTOTPFlow(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.beforeTOTPFlowHook == nil {
		b.beforeTOTPFlowHook = w(NopHookFunc)
	} else {
		b.beforeTOTPFlowHook = w(b.beforeTOTPFlowHook)
	}
	return b
}

func (b *Builder) AfterLogin(v HookFunc) (r *Builder) {
	b.afterLoginHook = v
	return b
}

func (b *Builder) WrapAfterLogin(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterLoginHook == nil {
		b.afterLoginHook = w(NopHookFunc)
	} else {
		b.afterLoginHook = w(b.afterLoginHook)
	}
	return b
}

// extra vals:
// - login error
func (b *Builder) AfterFailedToLogin(v HookFunc) (r *Builder) {
	b.afterFailedToLoginHook = v
	return b
}

func (b *Builder) WrapAfterFailedToLogin(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterFailedToLoginHook == nil {
		b.afterFailedToLoginHook = w(NopHookFunc)
	} else {
		b.afterFailedToLoginHook = w(b.afterFailedToLoginHook)
	}
	return b
}

func (b *Builder) AfterUserLocked(v HookFunc) (r *Builder) {
	b.afterUserLockedHook = v
	return b
}

func (b *Builder) WrapAfterUserLocked(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterUserLockedHook == nil {
		b.afterUserLockedHook = w(NopHookFunc)
	} else {
		b.afterUserLockedHook = w(b.afterUserLockedHook)
	}
	return b
}

func (b *Builder) AfterLogout(v HookFunc) (r *Builder) {
	b.afterLogoutHook = v
	return b
}

func (b *Builder) WrapAfterLogout(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterLogoutHook == nil {
		b.afterLogoutHook = w(NopHookFunc)
	} else {
		b.afterLogoutHook = w(b.afterLogoutHook)
	}
	return b
}

// extra vals:
// - reset link
func (b *Builder) AfterConfirmSendResetPasswordLink(v HookFunc) (r *Builder) {
	b.afterConfirmSendResetPasswordLinkHook = v
	return b
}

func (b *Builder) WrapAfterConfirmSendResetPasswordLink(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterConfirmSendResetPasswordLinkHook == nil {
		b.afterConfirmSendResetPasswordLinkHook = w(NopHookFunc)
	} else {
		b.afterConfirmSendResetPasswordLinkHook = w(b.afterConfirmSendResetPasswordLinkHook)
	}
	return b
}

func (b *Builder) AfterResetPassword(v HookFunc) (r *Builder) {
	b.afterResetPasswordHook = v
	return b
}

func (b *Builder) WrapAfterResetPassword(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterResetPasswordHook == nil {
		b.afterResetPasswordHook = w(NopHookFunc)
	} else {
		b.afterResetPasswordHook = w(b.afterResetPasswordHook)
	}
	return b
}

func (b *Builder) AfterChangePassword(v HookFunc) (r *Builder) {
	b.afterChangePasswordHook = v
	return b
}

func (b *Builder) WrapAfterChangePassword(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterChangePasswordHook == nil {
		b.afterChangePasswordHook = w(NopHookFunc)
	} else {
		b.afterChangePasswordHook = w(b.afterChangePasswordHook)
	}
	return b
}

// extra vals:
// - old session token
func (b *Builder) AfterExtendSession(v HookFunc) (r *Builder) {
	b.afterExtendSessionHook = v
	return b
}

func (b *Builder) WrapAfterExtendSession(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterExtendSessionHook == nil {
		b.afterExtendSessionHook = w(NopHookFunc)
	} else {
		b.afterExtendSessionHook = w(b.afterExtendSessionHook)
	}
	return b
}

func (b *Builder) AfterTOTPCodeReused(v HookFunc) (r *Builder) {
	b.afterTOTPCodeReusedHook = v
	return b
}

func (b *Builder) WrapAfterTOTPCodeReused(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterTOTPCodeReusedHook == nil {
		b.afterTOTPCodeReusedHook = w(NopHookFunc)
	} else {
		b.afterTOTPCodeReusedHook = w(b.afterTOTPCodeReusedHook)
	}
	return b
}

// user is goth.User
func (b *Builder) AfterOAuthComplete(v HookFunc) (r *Builder) {
	b.afterOAuthCompleteHook = v
	return b
}

func (b *Builder) WrapAfterOAuthComplete(w func(in HookFunc) HookFunc) (r *Builder) {
	if b.afterOAuthCompleteHook == nil {
		b.afterOAuthCompleteHook = w(NopHookFunc)
	} else {
		b.afterOAuthCompleteHook = w(b.afterOAuthCompleteHook)
	}
	return b
}

// seconds
// default 1h
func (b *Builder) SessionMaxAge(v int) (r *Builder) {
	b.sessionMaxAge = v
	return b
}

// extend the session if successfully authenticated
// default true
func (b *Builder) AutoExtendSession(v bool) (r *Builder) {
	b.autoExtendSession = v
	return b
}

// default 5
// MaxRetryCount <= 0 means no max retry count limit
func (b *Builder) MaxRetryCount(v int) (r *Builder) {
	b.maxRetryCount = v
	return b
}

func (b *Builder) TOTP(enable bool, config ...TOTPConfig) (r *Builder) {
	b.totpEnabled = enable
	if len(config) > 0 {
		b.totpConfig = config[0]
	}
	if enable {
		if b.totpConfig.Issuer == "" {
			panic("Issuer is empty")
		}
	}
	return b
}

func (b *Builder) NoForgetPasswordLink(v bool) (r *Builder) {
	b.noForgetPasswordLink = v
	return b
}

func (b *Builder) DB(v *gorm.DB) (r *Builder) {
	b.db = v
	return b
}

func (b *Builder) I18n(v *i18n.Builder) (r *Builder) {
	v.RegisterForModule(language.English, I18nLoginKey, Messages_en_US).
		RegisterForModule(language.SimplifiedChinese, I18nLoginKey, Messages_zh_CN).
		RegisterForModule(language.Japanese, I18nLoginKey, Messages_ja_JP)
	b.i18nBuilder = v
	return b
}

type OAuthIdentifier int

const (
	OAuthIdentifierEmail OAuthIdentifier = iota
	OAuthIdentifierName
	OAuthIdentifierNickName
	OAuthIdentifierUserID
)

// OAuthIdentifier is an externally-facing account identifier, such as an email address for a Google account.
// default is email, fallback is userID
func (b *Builder) OAuthIdentifier(provider string, identifier OAuthIdentifier) (r *Builder) {
	b.oauthIdentifiers[provider] = identifier
	return b
}

func (b *Builder) oauthIdentifierValue(ouser goth.User) string {
	var val string
	idt := b.oauthIdentifiers[ouser.Provider]
	switch idt {
	case OAuthIdentifierEmail:
		val = ouser.Email
	case OAuthIdentifierName:
		val = ouser.Name
	case OAuthIdentifierNickName:
		val = ouser.NickName
	case OAuthIdentifierUserID:
		val = ouser.UserID
	default:
		panic(fmt.Sprintf("unknown identifier type %v", idt))
	}
	if val == "" {
		val = ouser.UserID
	}
	return val
}

func (b *Builder) GetSessionMaxAge() int {
	return b.sessionMaxAge
}

func (b *Builder) ViewHelper() *ViewHelper {
	return &ViewHelper{
		b: b,
	}
}

func (b *Builder) UserModel(m interface{}) (r *Builder) {
	b.userModel = m
	b.tUser = underlyingReflectType(reflect.TypeOf(m))
	b.snakePrimaryField = snakePrimaryField(m)
	if _, ok := m.(UserPasser); ok {
		b.userPassEnabled = true
	}
	if _, ok := m.(OAuthUser); ok {
		b.oauthEnabled = true
	}
	if _, ok := m.(SessionSecurer); ok {
		b.sessionSecureEnabled = true
	}
	return b
}

func (b *Builder) newUserObject() interface{} {
	return reflect.New(b.tUser).Interface()
}

func (b *Builder) findUserByID(id string) (user interface{}, err error) {
	m := b.newUserObject()
	err = b.db.Where(fmt.Sprintf("%s = ?", b.snakePrimaryField), id).
		First(m).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return m, nil
}

// completeUserAuthCallback is for url "/auth/{provider}/callback"
func (b *Builder) completeUserAuthCallback(w http.ResponseWriter, r *http.Request) {
	if b.cookieConfig.SameSite != http.SameSiteStrictMode {
		b.completeUserAuthCallbackComplete(w, r)
		return
	}

	q := r.URL.Query()
	callbackURL := b.oauthCallbackCompleteURL
	if len(q) > 0 {
		callbackURL = callbackURL + "?" + q.Encode()
	}

	// Use html/template to safely handle user input and prevent XSS attacks
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl := template.Must(template.New("callback").Parse(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Auth Redirect</title>
</head>
<body>
    <script>
    window.location.href="{{.CompleteURL}}";
    </script>
    <p>If the page doesn't redirect automatically, <a href="{{.CompleteURL}}">click here</a> to continue.</p>
</body>
</html>
`))

	data := struct {
		CompleteURL string
	}{
		CompleteURL: callbackURL,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (b *Builder) completeUserAuthCallbackComplete(w http.ResponseWriter, r *http.Request) {
	var err error
	var user interface{}
	failRedirectURL := b.LogoutURL
	defer func() {
		if perr := recover(); perr != nil {
			panic(perr)
		}
		if err != nil {
			if b.afterFailedToLoginHook != nil {
				if herr := b.wrapHook(b.afterFailedToLoginHook)(r, user, err); herr != nil {
					setNoticeOrPanic(w, herr)
				}
			}
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
		}
	}()

	var ouser goth.User
	ouser, err = gothic.CompleteUserAuth(w, r)
	if err != nil {
		setFailCodeFlash(w, FailCodeCompleteUserAuthFailed)
		return
	}

	if b.afterOAuthCompleteHook != nil {
		if err = b.wrapHook(b.afterOAuthCompleteHook)(r, ouser); err != nil {
			setNoticeOrPanic(w, err)
			return
		}
	}

	userID := ouser.UserID
	if userID == "" {
		setFailCodeFlash(w, FailCodeCompleteUserAuthFailed)
		return
	}

	if b.userModel != nil {
		user, err = b.userModel.(OAuthUser).FindUserByOAuthUserID(b.db, b.newUserObject(), ouser.Provider, ouser.UserID)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				panic(err)
			}
			identifier := b.oauthIdentifierValue(ouser)
			user, err = b.userModel.(OAuthUser).FindUserByOAuthIdentifier(b.db, b.newUserObject(), ouser.Provider, identifier)
			if err != nil {
				if err != gorm.ErrRecordNotFound {
					panic(err)
				}
				setFailCodeFlash(w, FailCodeUserNotFound)
				return
			}
			err = user.(OAuthUser).InitOAuthUserID(b.db, b.newUserObject(), ouser.Provider, identifier, ouser.UserID)
			if err != nil {
				panic(err)
			}
			// refetch user by oauth user id to prevent fake identifier
			user, err = b.userModel.(OAuthUser).FindUserByOAuthUserID(b.db, b.newUserObject(), ouser.Provider, ouser.UserID)
			if err != nil {
				if err != gorm.ErrRecordNotFound {
					panic(err)
				}
				setFailCodeFlash(w, FailCodeUserNotFound)
				return
			}
		}
		userID = objectID(user)
	}

	claims := UserClaims{
		Provider:         ouser.Provider,
		Email:            ouser.Email,
		Name:             ouser.Name,
		UserID:           userID,
		AvatarURL:        ouser.AvatarURL,
		RegisteredClaims: b.genBaseSessionClaim(userID),
	}
	if user == nil {
		user = &claims
	}

	if b.afterLoginHook != nil {
		setCookieForRequest(r, &http.Cookie{Name: b.authCookieName, Value: b.mustGetSessionToken(claims)})
		if err = b.wrapHook(b.afterLoginHook)(r, user); err != nil {
			setNoticeOrPanic(w, err)
			return
		}
	}

	if err = b.setSecureCookiesByClaims(w, user, claims); err != nil {
		panic(err)
	}

	redirectURL := b.homePageURLFunc(r, user)
	if v := b.getContinueURL(w, r); v != "" {
		redirectURL = v
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
	return
}

// return user if account exists even if there is an error returned
func (b *Builder) authUserPass(account string, password string) (user interface{}, err error) {
	user, err = b.userModel.(UserPasser).FindUser(b.db, b.newUserObject(), account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	u := user.(UserPasser)
	if u.GetLocked() {
		return user, ErrUserLocked
	}

	if !u.IsPasswordCorrect(password) {
		if b.maxRetryCount > 0 {
			if err = u.IncreaseRetryCount(b.db, b.newUserObject()); err != nil {
				return user, err
			}
			if u.GetLoginRetryCount() >= b.maxRetryCount {
				if err = u.LockUser(b.db, b.newUserObject()); err != nil {
					return user, err
				}
				return user, ErrUserGetLocked
			}
		}

		return user, ErrWrongPassword
	}

	if u.GetLoginRetryCount() != 0 {
		if err = u.UnlockUser(b.db, b.newUserObject()); err != nil {
			return user, err
		}
	}
	return user, nil
}

func (b *Builder) userpassLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check reCAPTCHA token
	if b.recaptchaEnabled {
		token := r.FormValue("token")
		if !recaptchaTokenCheck(b, token) {
			setFailCodeFlash(w, FailCodeIncorrectRecaptchaToken)
			http.Redirect(w, r, b.loginPageURL, http.StatusFound)
			return
		}
	}

	var err error
	var user interface{}
	failRedirectURL := b.LogoutURL
	defer func() {
		if perr := recover(); perr != nil {
			panic(perr)
		}
		if err != nil {
			if b.afterFailedToLoginHook != nil {
				if herr := b.wrapHook(b.afterFailedToLoginHook)(r, user, err); herr != nil {
					setNoticeOrPanic(w, herr)
				}
			}
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
		}
	}()

	account := r.FormValue("account")
	password := r.FormValue("password")
	user, err = b.authUserPass(account, password)
	if err != nil {
		if err == ErrUserGetLocked && b.afterUserLockedHook != nil {
			if herr := b.wrapHook(b.afterUserLockedHook)(r, user); herr != nil {
				setNoticeOrPanic(w, herr)
				return
			}
		}

		var code FailCode
		switch err {
		case ErrWrongPassword, ErrUserNotFound:
			code = FailCodeIncorrectAccountNameOrPassword
		case ErrUserLocked, ErrUserGetLocked:
			code = FailCodeUserLocked
		default:
			panic(err)
		}
		setFailCodeFlash(w, code)
		b.setWrongLoginInputFlash(w, WrongLoginInputFlash{
			Account:  account,
			Password: password,
		})
		return
	}

	u := user.(UserPasser)
	userID := objectID(user)
	claims := UserClaims{
		UserID:           userID,
		PassUpdatedAt:    u.GetPasswordUpdatedAt(),
		RegisteredClaims: b.genBaseSessionClaim(userID),
	}

	if !b.totpEnabled {
		if b.afterLoginHook != nil {
			setCookieForRequest(r, &http.Cookie{Name: b.authCookieName, Value: b.mustGetSessionToken(claims)})
			if err = b.wrapHook(b.afterLoginHook)(r, user); err != nil {
				setNoticeOrPanic(w, err)
				return
			}
		}
	}

	if err = b.setSecureCookiesByClaims(w, user, claims); err != nil {
		panic(err)
	}

	if b.totpEnabled {
		if b.beforeTOTPFlowHook != nil {
			if err = b.wrapHook(b.beforeTOTPFlowHook)(r, user); err != nil {
				setNoticeOrPanic(w, err)
				return
			}
		}

		if u.GetIsTOTPSetup() {
			http.Redirect(w, r, b.totpValidatePageURL, http.StatusFound)
			return
		}

		var key *otp.Key
		if key, err = totp.Generate(
			totp.GenerateOpts{
				Issuer:      b.totpConfig.Issuer,
				AccountName: u.GetAccountName(),
			},
		); err != nil {
			panic(err)
		}

		if err = u.SetTOTPSecret(b.db, b.newUserObject(), key.Secret()); err != nil {
			panic(err)
		}

		http.Redirect(w, r, b.totpSetupPageURL, http.StatusFound)
		return
	}

	redirectURL := b.homePageURLFunc(r, user)
	if v := b.getContinueURL(w, r); v != "" {
		redirectURL = v
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
	return
}

func (b *Builder) genBaseSessionClaim(id string) jwt.RegisteredClaims {
	return genBaseClaims(id, b.sessionMaxAge)
}

func (b *Builder) mustGetSessionToken(claims UserClaims) string {
	return mustSignClaims(claims, b.secret)
}

func (b *Builder) setAuthCookiesFromUserClaims(w http.ResponseWriter, claims *UserClaims, secureSalt string) {
	http.SetCookie(w, &http.Cookie{
		Name:     b.authCookieName,
		Value:    b.mustGetSessionToken(*claims),
		Path:     b.cookieConfig.Path,
		Domain:   b.cookieConfig.Domain,
		MaxAge:   b.sessionMaxAge,
		Expires:  time.Now().Add(time.Duration(b.sessionMaxAge) * time.Second),
		HttpOnly: true,
		Secure:   b.cookieConfig.Secure,
		SameSite: b.cookieConfig.SameSite,
	})

	if secureSalt != "" {
		http.SetCookie(w, &http.Cookie{
			Name:     b.authSecureCookieName,
			Value:    mustSignClaims(&claims.RegisteredClaims, b.secret+secureSalt),
			Path:     b.cookieConfig.Path,
			Domain:   b.cookieConfig.Domain,
			MaxAge:   b.sessionMaxAge,
			Expires:  time.Now().Add(time.Duration(b.sessionMaxAge) * time.Second),
			HttpOnly: true,
			Secure:   b.cookieConfig.Secure,
			SameSite: b.cookieConfig.SameSite,
		})
	}
}

func (b *Builder) cleanAuthCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     b.authCookieName,
		Value:    "",
		Path:     b.cookieConfig.Path,
		Domain:   b.cookieConfig.Domain,
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
		HttpOnly: true,
		Secure:   b.cookieConfig.Secure,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     b.authSecureCookieName,
		Value:    "",
		Path:     b.cookieConfig.Path,
		Domain:   b.cookieConfig.Domain,
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
		HttpOnly: true,
		Secure:   b.cookieConfig.Secure,
	})
}

func (b *Builder) setContinueURL(w http.ResponseWriter, r *http.Request) {
	continueURL := r.RequestURI
	if strings.Contains(continueURL, "?__execute_event__=") {
		continueURL = r.Referer()
	}
	ignore := false
	{
		ignoreURLs := map[string]struct{}{
			b.loginPageURL:                 {},
			b.resetPasswordPageURL:         {},
			b.forgetPasswordPageURL:        {},
			b.resetPasswordLinkSentPageURL: {},
			b.totpSetupPageURL:             {},
			b.totpValidatePageURL:          {},
			b.LogoutURL:                    {},
		}
		u, err := url.Parse(continueURL)
		if err != nil {
			ignore = true
		} else {
			if _, ok := ignoreURLs[u.Path]; ok {
				ignore = true
			}
		}
	}
	if ignore {
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     b.continueUrlCookieName,
		Value:    continueURL,
		Path:     b.cookieConfig.Path,
		Domain:   b.cookieConfig.Domain,
		HttpOnly: true,
	})
}

func (b *Builder) getContinueURL(w http.ResponseWriter, r *http.Request) string {
	c, err := r.Cookie(b.continueUrlCookieName)
	if err != nil || c.Value == "" {
		return ""
	}

	http.SetCookie(w, &http.Cookie{
		Name:     b.continueUrlCookieName,
		Value:    "",
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
		Path:     b.cookieConfig.Path,
		Domain:   b.cookieConfig.Domain,
		HttpOnly: true,
	})

	return c.Value
}

func (b *Builder) setSecureCookiesByClaims(w http.ResponseWriter, user interface{}, claims UserClaims) (err error) {
	var secureSalt string
	if b.sessionSecureEnabled {
		if user.(SessionSecurer).GetSecure() == "" {
			err = user.(SessionSecurer).UpdateSecure(b.db, b.newUserObject(), objectID(user))
			if err != nil {
				return err
			}
		}
		secureSalt = user.(SessionSecurer).GetSecure()
	}
	b.setAuthCookiesFromUserClaims(w, &claims, secureSalt)
	return nil
}

func (b *Builder) consumeTOTPCode(r *http.Request, up UserPasser, passcode string) error {
	if !totp.Validate(passcode, up.GetTOTPSecret()) {
		return ErrWrongTOTPCode
	}
	lastCode, usedAt := up.GetLastUsedTOTPCode()
	if usedAt != nil && time.Since(*usedAt) > 90*time.Second {
		lastCode = ""
	}
	if passcode == lastCode {
		if b.afterTOTPCodeReusedHook != nil {
			if herr := b.wrapHook(b.afterTOTPCodeReusedHook)(r, GetCurrentUser(r)); herr != nil {
				return herr
			}
		}
		return ErrTOTPCodeHasBeenUsed
	}
	if err := up.SetLastUsedTOTPCode(b.db, b.newUserObject(), passcode); err != nil {
		return err
	}
	return nil
}

// logout is for url "/logout/{provider}"
func (b *Builder) logout(w http.ResponseWriter, r *http.Request) {
	err := gothic.Logout(w, r)
	if err != nil {
		//
	}

	b.cleanAuthCookies(w)

	if b.afterLogoutHook != nil {
		user := GetCurrentUser(r)
		if user != nil {
			if herr := b.wrapHook(b.afterLogoutHook)(r, user); herr != nil {
				setNoticeOrPanic(w, herr)
				http.Redirect(w, r, b.loginPageURL, http.StatusFound)
				return
			}
		}
	}

	http.Redirect(w, r, b.loginPageURL, http.StatusFound)
}

// beginAuth is for url "/auth/{provider}"
func (b *Builder) beginAuth(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (b *Builder) sendResetPasswordLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	failRedirectURL := b.forgetPasswordPageURL

	// check reCAPTCHA token
	if b.recaptchaEnabled {
		token := r.FormValue("token")
		if !recaptchaTokenCheck(b, token) {
			setFailCodeFlash(w, FailCodeIncorrectRecaptchaToken)
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}
	}

	account := strings.TrimSpace(r.FormValue("account"))
	passcode := r.FormValue("otp")
	doTOTP := r.URL.Query().Get("totp") == "1"

	if doTOTP {
		failRedirectURL = MustSetQuery(failRedirectURL, "totp", "1")
	}

	if account == "" {
		setFailCodeFlash(w, FailCodeAccountIsRequired)
		http.Redirect(w, r, failRedirectURL, http.StatusFound)
		return
	}

	u, err := b.userModel.(UserPasser).FindUser(b.db, b.newUserObject(), account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Redirect(w, r, fmt.Sprintf("%s?a=%s", b.resetPasswordLinkSentPageURL, account), http.StatusFound)
			return
		}
		panic(err)
	}

	_, createdAt, _ := u.(UserPasser).GetResetPasswordToken()
	if createdAt != nil {
		v := 60 - int(time.Since(*createdAt).Seconds())
		if v > 0 {
			setSecondsToRedoFlash(w, v)
			b.setWrongForgetPasswordInputFlash(w, WrongForgetPasswordInputFlash{
				Account: account,
			})
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}
	}

	if u.(UserPasser).GetIsTOTPSetup() {
		if !doTOTP {
			b.setWrongForgetPasswordInputFlash(w, WrongForgetPasswordInputFlash{
				Account: account,
			})
			failRedirectURL = MustSetQuery(failRedirectURL, "totp", "1")
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}

		if err = b.consumeTOTPCode(r, u.(UserPasser), passcode); err != nil {
			var fc FailCode
			switch err {
			case ErrWrongTOTPCode:
				fc = FailCodeIncorrectTOTPCode
			case ErrTOTPCodeHasBeenUsed:
				fc = FailCodeTOTPCodeHasBeenUsed
			default:
				panic(err)
			}
			setNoticeOrFailCodeFlash(w, err, fc)
			b.setWrongForgetPasswordInputFlash(w, WrongForgetPasswordInputFlash{
				Account: account,
				TOTP:    passcode,
			})
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}
	}

	token, err := u.(UserPasser).GenerateResetPasswordToken(b.db, b.newUserObject())
	if err != nil {
		panic(err)
	}

	scheme := "https"
	if r.TLS == nil {
		scheme = "http"
	}
	link := fmt.Sprintf("%s://%s%s?id=%s&token=%s", scheme, r.Host, b.resetPasswordPageURL, objectID(u), token)
	if doTOTP {
		link = MustSetQuery(link, "totp", "1")
	}
	if b.afterConfirmSendResetPasswordLinkHook != nil {
		if herr := b.wrapHook(b.afterConfirmSendResetPasswordLinkHook)(r, u, link); herr != nil {
			setNoticeOrPanic(w, herr)
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}
	}

	http.Redirect(w, r, fmt.Sprintf("%s?a=%s", b.resetPasswordLinkSentPageURL, account), http.StatusFound)
	return
}

func (b *Builder) doResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userID := r.FormValue("user_id")
	token := r.FormValue("token")
	passcode := r.FormValue("otp")
	doTOTP := r.URL.Query().Get("totp") == "1"

	failRedirectURL := fmt.Sprintf("%s?id=%s&token=%s", b.resetPasswordPageURL, userID, token)
	if doTOTP {
		failRedirectURL = MustSetQuery(failRedirectURL, "totp", "1")
	}
	if userID == "" {
		setFailCodeFlash(w, FailCodeUserNotFound)
		http.Redirect(w, r, failRedirectURL, http.StatusFound)
		return
	}
	if token == "" {
		setFailCodeFlash(w, FailCodeInvalidToken)
		http.Redirect(w, r, failRedirectURL, http.StatusFound)
		return
	}

	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")
	if password == "" {
		setFailCodeFlash(w, FailCodePasswordCannotBeEmpty)
		http.Redirect(w, r, failRedirectURL, http.StatusFound)
		return
	}
	if confirmPassword != password {
		setFailCodeFlash(w, FailCodePasswordNotMatch)
		b.setWrongResetPasswordInputFlash(w, WrongResetPasswordInputFlash{
			Password:        password,
			ConfirmPassword: confirmPassword,
		})
		http.Redirect(w, r, failRedirectURL, http.StatusFound)
		return
	}

	u, err := b.findUserByID(userID)
	if err != nil {
		if err == ErrUserNotFound {
			setFailCodeFlash(w, FailCodeUserNotFound)
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}
		panic(err)
	}

	storedToken, _, expired := u.(UserPasser).GetResetPasswordToken()
	if expired {
		setFailCodeFlash(w, FailCodeTokenExpired)
		http.Redirect(w, r, failRedirectURL, http.StatusFound)
		return
	}
	if token != storedToken {
		setFailCodeFlash(w, FailCodeInvalidToken)
		http.Redirect(w, r, failRedirectURL, http.StatusFound)
		return
	}

	if b.beforeSetPasswordHook != nil {
		if herr := b.wrapHook(b.beforeSetPasswordHook)(r, u, password); herr != nil {
			setNoticeOrPanic(w, herr)
			b.setWrongResetPasswordInputFlash(w, WrongResetPasswordInputFlash{
				Password:        password,
				ConfirmPassword: confirmPassword,
			})
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}
	}

	if u.(UserPasser).GetIsTOTPSetup() {
		if !doTOTP {
			b.setWrongResetPasswordInputFlash(w, WrongResetPasswordInputFlash{
				Password:        password,
				ConfirmPassword: confirmPassword,
			})
			failRedirectURL = MustSetQuery(failRedirectURL, "totp", "1")
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}

		if err = b.consumeTOTPCode(r, u.(UserPasser), passcode); err != nil {
			var fc FailCode
			switch err {
			case ErrWrongTOTPCode:
				fc = FailCodeIncorrectTOTPCode
			case ErrTOTPCodeHasBeenUsed:
				fc = FailCodeTOTPCodeHasBeenUsed
			default:
				panic(err)
			}
			setFailCodeFlash(w, fc)
			b.setWrongResetPasswordInputFlash(w, WrongResetPasswordInputFlash{
				Password:        password,
				ConfirmPassword: confirmPassword,
				TOTP:            passcode,
			})
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}
	}

	err = u.(UserPasser).ConsumeResetPasswordToken(b.db, b.newUserObject())
	if err != nil {
		panic(err)
	}

	err = u.(UserPasser).SetPassword(b.db, b.newUserObject(), password)
	if err != nil {
		panic(err)
	}

	if b.afterResetPasswordHook != nil {
		if herr := b.wrapHook(b.afterResetPasswordHook)(r, u); herr != nil {
			setNoticeOrPanic(w, herr)
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
			return
		}
	}

	setInfoCodeFlash(w, InfoCodePasswordSuccessfullyReset)
	http.Redirect(w, r, b.loginPageURL, http.StatusFound)
	return
}

// NoticeError
// ErrWrongPassword
// ErrEmptyPassword
// ErrPasswordNotMatch
// ErrWrongTOTPCode
// ErrTOTPCodeHasBeenUsed
func (b *Builder) ChangePassword(
	r *http.Request,
	oldPassword string,
	password string,
	confirmPassword string,
	otp string,
) error {
	user := GetCurrentUser(r).(UserPasser)

	if !user.IsPasswordCorrect(oldPassword) {
		return ErrWrongPassword
	}

	if password == "" {
		return ErrEmptyPassword
	}
	if confirmPassword != password {
		return ErrPasswordNotMatch
	}

	if b.beforeSetPasswordHook != nil {
		if herr := b.wrapHook(b.beforeSetPasswordHook)(r, user, password); herr != nil {
			return herr
		}
	}

	if b.totpEnabled {
		if err := b.consumeTOTPCode(r, user, otp); err != nil {
			return err
		}
	}

	err := user.SetPassword(b.db, b.newUserObject(), password)
	if err != nil {
		return err
	}

	if b.afterChangePasswordHook != nil {
		if herr := b.wrapHook(b.afterChangePasswordHook)(r, user); herr != nil {
			return herr
		}
	}

	return nil
}

func (b *Builder) doFormChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oldPassword := r.FormValue("old_password")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")
	otp := r.FormValue("otp")

	redirectURL := b.changePasswordPageURL

	err := b.ChangePassword(r, oldPassword, password, confirmPassword, otp)
	if err != nil {
		var ne *NoticeError
		if errors.As(err, &ne) {
			setNoticeFlash(w, ne)
		} else {
			var fc FailCode
			switch err {
			case ErrWrongPassword:
				fc = FailCodeIncorrectPassword
			case ErrEmptyPassword:
				fc = FailCodePasswordCannotBeEmpty
			case ErrPasswordNotMatch:
				fc = FailCodePasswordNotMatch
			case ErrWrongTOTPCode:
				fc = FailCodeIncorrectTOTPCode
			case ErrTOTPCodeHasBeenUsed:
				fc = FailCodeTOTPCodeHasBeenUsed
			default:
				panic(err)
			}
			setFailCodeFlash(w, fc)
		}

		b.setWrongChangePasswordInputFlash(w, WrongChangePasswordInputFlash{
			OldPassword:     oldPassword,
			NewPassword:     password,
			ConfirmPassword: confirmPassword,
			TOTP:            otp,
		})
		http.Redirect(w, r, redirectURL, http.StatusFound)
		return
	}

	setInfoCodeFlash(w, InfoCodePasswordSuccessfullyChanged)
	http.Redirect(w, r, b.loginPageURL, http.StatusFound)
}

func (b *Builder) totpDo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var claims *UserClaims
	claims, err := parseUserClaimsFromCookie(r, b.authCookieName, b.secret)
	if err != nil {
		http.Redirect(w, r, b.LogoutURL, http.StatusFound)
		return
	}

	user, err := b.findUserByID(claims.UserID)
	if err != nil {
		if err == ErrUserNotFound {
			setFailCodeFlash(w, FailCodeUserNotFound)
			http.Redirect(w, r, b.LogoutURL, http.StatusFound)
			return
		}
		panic(err)
	}

	failRedirectURL := b.LogoutURL
	defer func() {
		if perr := recover(); perr != nil {
			panic(perr)
		}
		if err != nil {
			if b.afterFailedToLoginHook != nil {
				if herr := b.wrapHook(b.afterFailedToLoginHook)(r, user, err); herr != nil {
					setNoticeOrPanic(w, herr)
				}
			}
			http.Redirect(w, r, failRedirectURL, http.StatusFound)
		}
	}()

	u := user.(UserPasser)

	otp := r.FormValue("otp")
	isTOTPSetup := u.GetIsTOTPSetup()

	if err = b.consumeTOTPCode(r, u, otp); err != nil {
		var fc FailCode
		switch err {
		case ErrWrongTOTPCode:
			fc = FailCodeIncorrectTOTPCode
		case ErrTOTPCodeHasBeenUsed:
			fc = FailCodeTOTPCodeHasBeenUsed
		default:
			panic(err)
		}
		setFailCodeFlash(w, fc)
		failRedirectURL = b.totpValidatePageURL
		if !isTOTPSetup {
			failRedirectURL = b.totpSetupPageURL
		}
		return
	}

	if !isTOTPSetup {
		if err = u.SetIsTOTPSetup(b.db, b.newUserObject(), true); err != nil {
			panic(err)
		}
	}

	claims.TOTPValidated = true
	if b.afterLoginHook != nil {
		setCookieForRequest(r, &http.Cookie{Name: b.authCookieName, Value: b.mustGetSessionToken(*claims)})
		if err = b.wrapHook(b.afterLoginHook)(r, user); err != nil {
			setNoticeOrPanic(w, err)
			return
		}
	}

	err = b.setSecureCookiesByClaims(w, user, *claims)
	if err != nil {
		panic(err)
	}

	redirectURL := b.homePageURLFunc(r, user)
	if v := b.getContinueURL(w, r); v != "" {
		redirectURL = v
	}
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func (b *Builder) Mount(mux *http.ServeMux) {
	b.MountAPI(mux)

	// pages
	wb := web.New()
	mux.Handle(b.loginPageURL, b.i18nBuilder.EnsureLanguage(wb.Page(b.loginPageFunc)))
	if b.userPassEnabled {
		mux.Handle(b.resetPasswordPageURL, b.i18nBuilder.EnsureLanguage(wb.Page(b.resetPasswordPageFunc)))
		mux.Handle(b.changePasswordPageURL, b.i18nBuilder.EnsureLanguage(wb.Page(b.changePasswordPageFunc)))
		if !b.noForgetPasswordLink {
			mux.Handle(b.forgetPasswordPageURL, b.i18nBuilder.EnsureLanguage(wb.Page(b.forgetPasswordPageFunc)))
			mux.Handle(b.resetPasswordLinkSentPageURL, b.i18nBuilder.EnsureLanguage(wb.Page(b.resetPasswordLinkSentPageFunc)))
		}
		if b.totpEnabled {
			mux.Handle(b.totpSetupPageURL, b.i18nBuilder.EnsureLanguage(wb.Page(b.totpSetupPageFunc)))
			mux.Handle(b.totpValidatePageURL, b.i18nBuilder.EnsureLanguage(wb.Page(b.totpValidatePageFunc)))
		}
	}

	// assets
	assetsSubFS, err := fs.Sub(assetsFS, "assets")
	if err != nil {
		panic(err)
	}
	mux.Handle(assetsPathPrefix, http.StripPrefix(assetsPathPrefix, http.FileServer(http.FS(assetsSubFS))))
}

func (b *Builder) MountAPI(mux *http.ServeMux) {
	if len(b.secret) == 0 {
		panic("secret is empty")
	}
	if b.userModel != nil {
		if b.db == nil {
			panic("db is required")
		}
	}

	mux.HandleFunc(b.LogoutURL, b.logout)
	if b.userPassEnabled {
		mux.HandleFunc(b.passwordLoginURL, b.userpassLogin)
		mux.HandleFunc(b.resetPasswordURL, b.doResetPassword)
		mux.HandleFunc(b.changePasswordURL, b.doFormChangePassword)
		if !b.noForgetPasswordLink {
			mux.HandleFunc(b.sendResetPasswordLinkURL, b.sendResetPasswordLink)
		}
		if b.totpEnabled {
			mux.HandleFunc(b.validateTOTPURL, b.totpDo)
		}
	}
	if b.oauthEnabled {
		mux.HandleFunc(b.oauthBeginURL, b.beginAuth)
		mux.HandleFunc(b.oauthCallbackURL, b.completeUserAuthCallback)
		mux.HandleFunc(b.oauthCallbackCompleteURL, b.completeUserAuthCallbackComplete)
	}
}
