package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/theplant/gofixtures"

	"github.com/jinzhu/gorm"

	"github.com/goplaid/x/docs"
	"github.com/goplaid/x/docs/examples/e21_presents"
)

func TestDocExamples(t *testing.T) {
	var emptyProductsData = gofixtures.Data(gofixtures.Sql(``, []string{"products"}))

	var mycases = []reqCase{
		{
			name: "Custom MyFile Type",
			reqFunc: func(db *gorm.DB) *http.Request {
				emptyProductsData.TruncatePut(db)
				body := bytes.NewBuffer(nil)

				mw := multipart.NewWriter(body)
				_ = mw.WriteField("__event_data__", `{"eventFuncId":{"id":"update","pushState":null, "params":[""]},"event":{"value":""}}
		`)
				fw, _ := mw.CreateFormFile("MainImage_NewFile", "myfile.png")
				_, _ = fw.Write([]byte("Hello"))

				_ = mw.Close()

				r := httptest.NewRequest("POST", e21_presents.PresetsEditingCustomizationFileTypePath+"/products?__execute_event__=update", body)
				r.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", mw.Boundary()))
				return r
			},
			eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
				var u = &e21_presents.Product{}
				err := db.Find(u).Error
				if err != nil {
					t.Error(err)
				}
				if !strings.HasPrefix(string(u.MainImage), "https://transfer.sh") {
					t.Error(u)
				}
				return
			},
		},
		{
			name: "Custom MyFile Type Without File",
			reqFunc: func(db *gorm.DB) *http.Request {
				emptyProductsData.TruncatePut(db)
				body := bytes.NewBuffer(nil)

				mw := multipart.NewWriter(body)
				_ = mw.WriteField("__event_data__", `{"eventFuncId":{"id":"update","pushState":null, "params":[""]},"event":{"value":""}}
	`)
				_ = mw.Close()

				r := httptest.NewRequest("POST", e21_presents.PresetsEditingCustomizationFileTypePath+"/products?__execute_event__=update", body)
				r.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", mw.Boundary()))
				return r
			},
			eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
				var u = &e21_presents.Product{}
				err := db.Find(u).Error
				if err != nil {
					t.Error(err)
				}
				if string(u.MainImage) != "" {
					t.Error(u)
				}
				return
			},
		},
	}

	h := docs.SamplesHandler("")
	db := e21_presents.DB

	for _, c := range mycases {
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := c.reqFunc(db)
			h.ServeHTTP(w, r)

			if c.eventResponseMatch != nil {
				var er testEventResponse
				err := json.NewDecoder(w.Body).Decode(&er)
				if err != nil {
					panic(err)
				}
				c.eventResponseMatch(&er, db, t)
			}

			if c.pageMatch != nil {
				c.pageMatch(w.Body, db, t)
			}
		})

	}
}
