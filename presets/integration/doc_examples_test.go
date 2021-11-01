package integration_test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/goplaid/multipartestutils"
	"github.com/goplaid/x/docs"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/theplant/gofixtures"
	"gorm.io/gorm"
)

func TestDocExamples(t *testing.T) {
	var emptyProductsData = gofixtures.Data(gofixtures.Sql(``, []string{"products"}))

	var mycases = []reqCase{
		{
			name: "Custom MyFile Type",
			reqFunc: func(db *sql.DB) *http.Request {
				emptyProductsData.TruncatePut(db)
				return multipartestutils.NewMultipartBuilder().
					EventFunc("presets_Update").
					AddReader("MainImage_NewFile", "myfile.png", strings.NewReader("Hello")).
					PageURL(e21_presents.PresetsEditingCustomizationFileTypePath + "/products").
					BuildEventFuncRequest()
			},
			eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
				var u = &e21_presents.Product{}
				err := db.Find(u).Error
				if err != nil {
					t.Error(err)
				}
				if !strings.HasPrefix(string(u.MainImage), "http://transfer.sh") {
					t.Error(u)
				}
				return
			},
		},
		{
			name: "Custom MyFile Type Without File",
			reqFunc: func(db *sql.DB) *http.Request {
				emptyProductsData.TruncatePut(db)
				return multipartestutils.NewMultipartBuilder().
					EventFunc("presets_Update").
					PageURL(e21_presents.PresetsEditingCustomizationFileTypePath + "/products").
					BuildEventFuncRequest()
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
			rawDB, _ := db.DB()
			r := c.reqFunc(rawDB)
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
