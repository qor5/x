package login

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/huandu/go-clone"
)

type ContextUserKey int

const (
	UserKey ContextUserKey = iota
	loginWIPKey
)

type MiddlewareConfig interface {
	middlewareConfig()
}

// LoginNotRequired executes the next handler regardless of whether the user is logged in or not
type LoginNotRequired struct{}

func (*LoginNotRequired) middlewareConfig() {}

// DisableAutoRedirectToHomePage makes it possible to visit login page when user is logged in
type DisableAutoRedirectToHomePage struct{}

func (*DisableAutoRedirectToHomePage) middlewareConfig() {}

func MockCurrentUser(user any) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserKey, user)))
		})
	}
}

var ErrShouldNotExtend = errors.New("should not extend session")

func (b *Builder) Middleware(cfgs ...MiddlewareConfig) func(next http.Handler) http.Handler {
	mustLogin := true
	autoRedirectToHomePage := true
	for _, cfg := range cfgs {
		switch cfg.(type) {
		case *LoginNotRequired:
			mustLogin = false
		case *DisableAutoRedirectToHomePage:
			autoRedirectToHomePage = false
		}
	}

	whiteList := map[string]struct{}{
		b.oauthBeginURL:                {},
		b.oauthCallbackURL:             {},
		b.oauthCallbackCompleteURL:     {},
		b.passwordLoginURL:             {},
		b.forgetPasswordPageURL:        {},
		b.sendResetPasswordLinkURL:     {},
		b.resetPasswordLinkSentPageURL: {},
		b.resetPasswordURL:             {},
		b.resetPasswordPageURL:         {},
		b.validateTOTPURL:              {},
	}

	staticFileRe := regexp.MustCompile(`\.(css|js|gif|jpg|jpeg|png|ico|svg|ttf|eot|woff|woff2|js\.map)$`)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if staticFileRe.MatchString(strings.ToLower(r.URL.Path)) {
				next.ServeHTTP(w, r)
				return
			}
			if _, ok := whiteList[r.URL.Path]; ok {
				next.ServeHTTP(w, r)
				return
			}

			path := strings.TrimRight(r.URL.Path, "/")

			claims, err := parseUserClaimsFromCookie(r, b.authCookieName, b.secret)
			if err != nil {
				if !mustLogin {
					next.ServeHTTP(w, r)
					return
				}
				if r.Method == http.MethodGet {
					b.setContinueURL(w, r)
				}
				if path == b.loginPageURL {
					next.ServeHTTP(w, r)
				} else {
					http.Redirect(w, r, b.loginPageURL, http.StatusFound)
				}
				return
			}

			var user interface{}
			var secureSalt string
			if b.userModel != nil {
				var err error
				user, err = b.findUserByID(claims.UserID)
				if err == nil {
					if claims.Provider == "" {
						if user.(UserPasser).GetPasswordUpdatedAt() != claims.PassUpdatedAt {
							err = ErrPasswordChanged
						}
						if user.(UserPasser).GetLocked() {
							err = ErrUserLocked
						}
					} else {
						user.(OAuthUser).SetAvatar(claims.AvatarURL)
					}
				}
				if err != nil {
					if !mustLogin {
						next.ServeHTTP(w, r)
						return
					}
					switch err {
					case ErrUserNotFound:
						setFailCodeFlash(w, FailCodeUserNotFound)
					case ErrUserLocked:
						setFailCodeFlash(w, FailCodeUserLocked)
					case ErrPasswordChanged:
						isSelfChange := false
						if c, err := r.Cookie(infoCodeFlashCookieName); err == nil {
							v, _ := strconv.Atoi(c.Value)
							if InfoCode(v) == InfoCodePasswordSuccessfullyChanged {
								isSelfChange = true
							}
						}
						if !isSelfChange {
							setWarnCodeFlash(w, WarnCodePasswordHasBeenChanged)
						}
					default:
						panic(err)
					}
					if path == b.LogoutURL {
						next.ServeHTTP(w, r)
					} else {
						http.Redirect(w, r, b.LogoutURL, http.StatusFound)
					}
					return
				}

				if b.sessionSecureEnabled {
					secureSalt = user.(SessionSecurer).GetSecure()
					_, err := parseBaseClaimsFromCookie(r, b.authSecureCookieName, b.secret+secureSalt)
					if err != nil {
						if !mustLogin {
							next.ServeHTTP(w, r)
							return
						}
						if path == b.LogoutURL {
							next.ServeHTTP(w, r)
						} else {
							http.Redirect(w, r, b.LogoutURL, http.StatusFound)
						}
						return
					}
				}
			} else {
				user = claims
			}

			if b.autoExtendSession && time.Since(claims.IssuedAt.Time).Seconds() > float64(b.sessionMaxAge)/10 {
				if b.afterExtendSessionHook != nil {
					oldSessionToken := b.mustGetSessionToken(*claims)

					newClaims := clone.Clone(claims).(*UserClaims)
					newClaims.RegisteredClaims = b.genBaseSessionClaim(claims.UserID)
					newSessionToken := b.mustGetSessionToken(*newClaims)
					if herr := b.wrapHook(b.afterExtendSessionHook)(r, user, oldSessionToken, newSessionToken); herr != nil {
						if !errors.Is(herr, ErrShouldNotExtend) {
							if !mustLogin {
								next.ServeHTTP(w, r)
								return
							}
							setNoticeOrPanic(w, herr)
							http.Redirect(w, r, b.LogoutURL, http.StatusFound)
							return
						}
					} else {
						claims = newClaims
						b.setAuthCookiesFromUserClaims(w, claims, secureSalt)
						setCookieForRequest(r, &http.Cookie{Name: b.authCookieName, Value: b.mustGetSessionToken(*claims)})
					}
				} else {
					claims.RegisteredClaims = b.genBaseSessionClaim(claims.UserID)
					b.setAuthCookiesFromUserClaims(w, claims, secureSalt)
					setCookieForRequest(r, &http.Cookie{Name: b.authCookieName, Value: b.mustGetSessionToken(*claims)})
				}
			}

			r = r.WithContext(context.WithValue(r.Context(), UserKey, user))

			if path == b.LogoutURL {
				next.ServeHTTP(w, r)
				return
			}

			if claims.Provider == "" && b.totpEnabled {
				if !user.(UserPasser).GetIsTOTPSetup() {
					if path == b.loginPageURL {
						next.ServeHTTP(w, r)
						return
					}
					r = r.WithContext(context.WithValue(r.Context(), loginWIPKey, true))
					if path == b.totpSetupPageURL {
						next.ServeHTTP(w, r)
						return
					}
					http.Redirect(w, r, b.totpSetupPageURL, http.StatusFound)
					return
				}

				if !claims.TOTPValidated {
					if path == b.loginPageURL {
						next.ServeHTTP(w, r)
						return
					}
					r = r.WithContext(context.WithValue(r.Context(), loginWIPKey, true))
					if path == b.totpValidatePageURL {
						next.ServeHTTP(w, r)
						return
					}
					http.Redirect(w, r, b.totpValidatePageURL, http.StatusFound)
					return
				}
			}

			if autoRedirectToHomePage {
				if path == b.loginPageURL || path == b.totpSetupPageURL || path == b.totpValidatePageURL {
					http.Redirect(w, r, b.homePageURLFunc(r, user), http.StatusFound)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

func GetCurrentUser(r *http.Request) (u interface{}) {
	return r.Context().Value(UserKey)
}

// IsLoginWIP indicates whether the user is in an intermediate step of login process,
// such as on the TOTP validation page
func IsLoginWIP(r *http.Request) bool {
	v, ok := r.Context().Value(loginWIPKey).(bool)
	if !ok {
		return false
	}
	return v
}
