package presets_guide

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var EditingCustomizations = Components(
	md.Markdown(`
Editing an object will be always in a drawer popup. select which fields can edit for each model
by using the ~.Only~ func of ~EditingBuilder~, There are different ways to configure the type 
of component that is used to do the editing.

`),
	utils.Anchor(H2(""), "Configure field for a single model"),
	md.Markdown(`
Use a customized component is as simple as add the extra asset to the preset instance.
And configure the component func on the field:
`),
	ch.Code(examples.PresetsEditingCustomizationDescriptionSample).Language("go"),
	utils.Demo("", e21_presents.PresetsEditingCustomizationDescriptionPath+"/customers"),
	md.Markdown(`
- Added the tiptap javascript and css component pack as an extra asset
- Configure the description field to use the component func that returns the ~tiptap.TipTapEditor()~ component
- Set the field name and value of the component
`),
	utils.Anchor(H2(""), "Configure field type for all models"),
)
