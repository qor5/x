package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitter"
	"github.com/qor5/x/v3/login"
	. "github.com/theplant/htmlgo"
	"github.com/theplant/osenv"
	"github.com/theplant/testingutils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	login.UserPass
	login.OAuthInfo
	login.SessionSecure
}

var (
	loginGoogleKey     = osenv.Get("LOGIN_GOOGLE_KEY", "Google client key for Login with Google", "")
	loginGoogleSecret  = osenv.Get("LOGIN_GOOGLE_SECRET", "Google client secret for Login with Google", "")
	loginTwitterKey    = osenv.Get("LOGIN_TWITTER_KEY", "Twitter client key for Login with Twitter", "")
	loginTwitterSecret = osenv.Get("LOGIN_TWITTER_KEY", "Twitter client secret for Login with Twitter", "")
)

func main() {
	var err error
	db, err := gorm.Open(sqlite.Open("/tmp/test_login.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	user := &User{
		UserPass: login.UserPass{
			Account:  "user@example.com",
			Password: "123",
		},
	}
	user.EncryptPassword()
	db.Create(user)

	b := login.New().
		DB(db).
		UserModel(&User{}).
		Secret("123").
		OAuthProviders(
			&login.Provider{
				Goth: google.New(loginGoogleKey, loginGoogleSecret,
					"http://localhost:9500/auth/callback?provider=google"),
				Key:  "google",
				Text: "Login with Google",
				Logo: RawHTML(`<svg xmlns="http://www.w3.org/2000/svg" class="inline w-4 h-4 mr-3 text-gray-900 fill-current" viewBox="0 0 48 48" width="48px" height="48px"><path fill="#fbc02d" d="M43.611,20.083H42V20H24v8h11.303c-1.649,4.657-6.08,8-11.303,8c-6.627,0-12-5.373-12-12	s5.373-12,12-12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C12.955,4,4,12.955,4,24s8.955,20,20,20	s20-8.955,20-20C44,22.659,43.862,21.35,43.611,20.083z"></path><path fill="#e53935" d="M6.306,14.691l6.571,4.819C14.655,15.108,18.961,12,24,12c3.059,0,5.842,1.154,7.961,3.039	l5.657-5.657C34.046,6.053,29.268,4,24,4C16.318,4,9.656,8.337,6.306,14.691z"></path><path fill="#4caf50" d="M24,44c5.166,0,9.86-1.977,13.409-5.192l-6.19-5.238C29.211,35.091,26.715,36,24,36	c-5.202,0-9.619-3.317-11.283-7.946l-6.522,5.025C9.505,39.556,16.227,44,24,44z"></path><path fill="#1565c0" d="M43.611,20.083L43.595,20L42,20H24v8h11.303c-0.792,2.237-2.231,4.166-4.087,5.571	c0.001-0.001,0.002-0.001,0.003-0.002l6.19,5.238C36.971,39.205,44,34,44,24C44,22.659,43.862,21.35,43.611,20.083z"></path></svg>`),
			},
			&login.Provider{
				Goth: twitter.New(loginTwitterKey, loginTwitterSecret,
					"http://localhost:9500/auth/callback?provider=twitter"),
				Key:  "twitter",
				Text: "Login with Twitter",
				Logo: RawHTML(`<svg class="inline w-4 h-4 mr-3 text-gray-900 fill-current" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"	 viewBox="0 0 248 204" style="enable-background:new 0 0 248 204;" xml:space="preserve"><style type="text/css">	.st0{fill:#1D9BF0;}</style><g id="Logo_1_">	<path id="white_background" class="st0" d="M221.95,51.29c0.15,2.17,0.15,4.34,0.15,6.53c0,66.73-50.8,143.69-143.69,143.69v-0.04		C50.97,201.51,24.1,193.65,1,178.83c3.99,0.48,8,0.72,12.02,0.73c22.74,0.02,44.83-7.61,62.72-21.66		c-21.61-0.41-40.56-14.5-47.18-35.07c7.57,1.46,15.37,1.16,22.8-0.87C27.8,117.2,10.85,96.5,10.85,72.46c0-0.22,0-0.43,0-0.64		c7.02,3.91,14.88,6.08,22.92,6.32C11.58,63.31,4.74,33.79,18.14,10.71c25.64,31.55,63.47,50.73,104.08,52.76		c-4.07-17.54,1.49-35.92,14.61-48.25c20.34-19.12,52.33-18.14,71.45,2.19c11.31-2.23,22.15-6.38,32.07-12.26		c-3.77,11.69-11.66,21.62-22.2,27.93c10.01-1.18,19.79-3.86,29-7.95C240.37,35.29,231.83,44.14,221.95,51.29z"/></g></svg>`),
			},
		).
		HomeURLFunc(func(r *http.Request, user interface{}) string {
			return "/admin"
		}).
		BeforeSetPassword(func(r *http.Request, user interface{}, extraVals ...interface{}) error {
			password := extraVals[0].(string)
			if len(password) <= 2 {
				return &login.NoticeError{
					Message: "password length cannot be less than 2",
				}
			}
			return nil
		}).
		AfterConfirmSendResetPasswordLink(func(r *http.Request, user interface{}, extraVals ...interface{}) error {
			link := extraVals[0]
			fmt.Println("#########################################start")
			testingutils.PrintlnJson(link)
			fmt.Println("#########################################end")
			return nil
		}).
		TOTP(false)

	h := http.NewServeMux()
	b.Mount(h)
	h.Handle("/admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "%#+v", login.GetCurrentUser(r))
	}))

	mux := http.NewServeMux()
	mux.Handle("/", b.Middleware()(h))

	log.Println("serving at http://localhost:9500")
	log.Fatal(http.ListenAndServe(":9500", mux))
}
