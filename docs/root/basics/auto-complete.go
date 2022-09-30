package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e10_vuetify_autocomplete"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var AutoComplete = Doc(
	Markdown(`
AutoComplete is a more advanced component that vuetify provides, We extend it
So that it can fetch remote options from an event func. here we show these examples:

- An auto complete that you can select multiple with static data
- An auto complete that you can select multiple with remote fetched dynamic data
- A static normal select component

`),
	ch.Code(examples.VuetifyAutoCompleteSample).Language("go"),
	utils.Demo("Vuetify AutoComplete", e10_vuetify_autocomplete.VuetifyAutoCompletePath, "e10_vuetify_autocomplete/page.go"),
).Title("Auto Complete").
	Slug("vuetify-components/auto-complete")
