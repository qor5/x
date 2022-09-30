package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var NotificationCenter = Doc(
	Markdown(`
To enable notification center: Call ~NotificationFunc~ on ~presets.Builder~ With 2 function parameters
like this ~builder.NotificationFunc(NotifierComponent(), NotifierCount())~

The first function is for rendering the content of the popup after user clicked the "bell icon".
The second function is for rendering the number at the top right corner of the "bell icon".
`),

	ch.Code(examples.NotificationCenterSample).Language("go"),
	utils.Demo("Notification Center", e21_presents.NotificationCenterSamplePath+"/pages", "e00_basics/notification-center.go"),
).Slug("basics/notification-center").Title("Notification Center")
