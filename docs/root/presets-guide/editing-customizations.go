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
	utils.Demo("Presets Editing Customization Description Field", e21_presents.PresetsEditingCustomizationDescriptionPath+"/customers", "e21_presents/editing.go"),
	md.Markdown(`
- Added the tiptap javascript and css component pack as an extra asset
- Configure the description field to use the component func that returns the ~tiptap.TipTapEditor()~ component
- Set the field name and value of the component
`),
	utils.Anchor(H2(""), "Configure field type for all models"),
	md.Markdown(`
Set a global field type to component func like the following:
`),
	ch.Code(examples.PresetsEditingCustomizationFileTypeSample).Language("go"),
	utils.Demo("Presets Editing Customization File Type", e21_presents.PresetsEditingCustomizationFileTypePath+"/products", "e21_presents/editing.go"),
	md.Markdown(`
- We define ~MyFile~ to actually be a string
- We set ~FieldDefaults~ for writing, which is the editing drawer popup to be a customized component
- The component show an img tag with the string as src if it's not empty
- The component add a file input for user to upload new file
- The ~SetterFunc~ is called before save the object, it uploads the file to transfer.sh, and get the url back,
  then set the value to ~MainImage~ field

With ~FieldDefaults~ we can write libraries that add customized type for different models to reuse. It can take care
of how to display the edit controls, and How to save the object.

`),
	utils.Anchor(H2(""), "Validation"),
	md.Markdown(`
Field level validation and display on field can be added by implement ~ValidateFunc~,
and set the ~web.ValidationErrors~ result:
`),
	ch.Code(examples.PresetsEditingCustomizationValidationSample).Language("go"),
	utils.Demo("Presets Editing Customization Validation", e21_presents.PresetsEditingCustomizationValidationPath+"/customers", "e21_presents/editing.go"),
	md.Markdown(`
- We validate the ~Name~ of the customer must be longer than 10
- If the error happens, If will show below the field

`),
)
