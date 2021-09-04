package e21_presents

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/tiptap"
	"github.com/jinzhu/gorm"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

// @snippet_begin(PresetsEditingCustomizationDescriptionSample)

func PresetsEditingCustomizationDescription(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsListingCustomizationBulkActions(b)
	b.URIPrefix(PresetsEditingCustomizationDescriptionPath)
	b.ExtraAsset("/tiptap.js", "text/javascript", tiptap.JSComponentsPack())
	b.ExtraAsset("/tiptap.css", "text/css", tiptap.CSSComponentsPack())

	ce.Only("Name", "CompanyID", "Description")

	ce.Field("Description").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		return tiptap.TipTapEditor().
			FieldName(field.Name).
			Value(field.Value(obj).(string))
	})
	return
}

const PresetsEditingCustomizationDescriptionPath = "/samples/presets-editing-customization-description"

// @snippet_end

// @snippet_begin(PresetsEditingCustomizationFileTypeSample)

type MyFile string

type Product struct {
	ID        int
	Title     string
	MainImage MyFile
}

func PresetsEditingCustomizationFileType(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsEditingCustomizationDescription(b)
	err := db.AutoMigrate(&Product{}).Error
	if err != nil {
		panic(err)
	}

	b.URIPrefix(PresetsEditingCustomizationFileTypePath)
	b.FieldDefaults(presets.WRITE).
		FieldType(MyFile("")).
		ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
			val := field.Value(obj).(MyFile)
			var img h.HTMLComponent
			if len(string(val)) > 0 {
				img = h.Img(string(val))
			}
			var er h.HTMLComponent
			if len(field.Errors) > 0 {
				er = h.Div().Text(field.Errors[0]).Style("color:red")
			}
			return h.Div(
				img,
				er,
				h.Input("").Type("file").Attr(web.VFieldName(fmt.Sprintf("%s_NewFile", field.Name))...),
			)
		}).
		SetterFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) (err error) {
			ff, _, _ := ctx.R.FormFile(fmt.Sprintf("%s_NewFile", field.Name))

			if ff == nil {
				return
			}

			req, err := http.NewRequest("PUT", "http://transfer.sh/myfile.png", ff)
			if err != nil {
				return
			}
			var res *http.Response
			res, err = http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}
			var b []byte
			b, err = ioutil.ReadAll(res.Body)
			if err != nil {
				return
			}
			if res.StatusCode == 500 {
				err = fmt.Errorf("%s", string(b))
				return
			}
			err = reflectutils.Set(obj, field.Name, MyFile(b))
			return
		})

	mb := b.Model(&Product{})
	mb.Editing("Title", "MainImage")
	return
}

const PresetsEditingCustomizationFileTypePath = "/samples/presets-editing-customization-file-type"

// @snippet_end

// @snippet_begin(PresetsEditingCustomizationValidationSample)

func PresetsEditingCustomizationValidation(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsEditingCustomizationDescription(b)
	b.URIPrefix(PresetsEditingCustomizationValidationPath)

	ce.ValidateFunc(func(obj interface{}, ctx *web.EventContext) (err web.ValidationErrors) {
		cus := obj.(*Customer)
		if len(cus.Name) < 10 {
			err.FieldError("Name", "name is too short")
		}
		return
	})
	return
}

const PresetsEditingCustomizationValidationPath = "/samples/presets-editing-customization-validation"

// @snippet_end
