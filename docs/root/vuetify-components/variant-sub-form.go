package vuetify_components

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e22_vuetify_variant_sub_form"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var VariantSubForm = Components(
	md.Markdown(`
VSelect changes, the form below it will change to a new form accordingly.

By use of ~web.Portal()~ and ~VSelect~'s ~OnInput~
`),
	ch.Code(examples.VuetifyVariantSubForm).Language("go"),
	utils.Demo("Vuetify Variant Sub Form", e22_vuetify_variant_sub_form.VuetifyVariantSubFormPath, "e22_vuetify_variant_sub_form/page.go"),
)
