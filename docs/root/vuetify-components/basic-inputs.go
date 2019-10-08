package vuetify_components

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e11_vuetify_basic_inputs"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var BasicInputs = Components(
	md.Markdown(`
Vuetify provides many form basic inputs, and also with error messages display on fields.

Here is one example:
`),
	ch.Code(examples.VuetifyBasicInputsSample).Language("go"),
	utils.Demo("", e11_vuetify_basic_inputs.VuetifyBasicInputsPath),
)
