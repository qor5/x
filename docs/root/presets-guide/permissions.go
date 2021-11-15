package presets_guide

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var Permissions = Doc(
	Markdown(`
Permissions sample:
`),
	ch.Code(examples.PresetsPermissionsSample).Language("go"),
	utils.Demo("Permissions Demo", e21_presents.PresetsPermissionsPath+"/customers", "e21_presents/permissions.go"),
).Title("Permissions").
	Slug("presets-guide/permissions")
