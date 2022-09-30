package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e22_vuetify_variant_sub_form"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var VariantSubForm = Doc(
	Markdown(`
VSelect changes, the form below it will change to a new form accordingly.

By use of ~web.Portal()~ and ~VSelect~'s ~OnInput~
`),
	ch.Code(examples.VuetifyVariantSubForm).Language("go"),
	utils.Demo("Vuetify Variant Sub Form", e22_vuetify_variant_sub_form.VuetifyVariantSubFormPath, "e22_vuetify_variant_sub_form/page.go"),
).Title("Variant Sub Form").
	Slug("vuetify-components/variant-sub-form")
