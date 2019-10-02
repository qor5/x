package e21_presents

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/tiptap"
	"github.com/jinzhu/gorm"
	h "github.com/theplant/htmlgo"
)

// @snippet_begin(PresetsEditingCustomizationDescriptionSample)

func PresetsEditingCustomizationDescription(b *presets.Builder) (
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cl, ce, db = PresetsListingCustomizationBulkActions(b)
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
