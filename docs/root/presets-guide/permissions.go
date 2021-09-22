package presets_guide

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var Permissions = Components(
	md.Markdown(`
Permissions sample:
`),
	ch.Code(examples.PresetsPermissionsSample).Language("go"),
	utils.Demo("Permissions Demo", e21_presents.PresetsPermissionsPath+"/customers", "e21_presents/permissions.go"),
)
