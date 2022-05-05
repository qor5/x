package e21_presents

// @snippet_begin(NotificationCenterSample)
import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/docs/examples/utils"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/gorm2op"
	v "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func PresetsNotificationCenterSample(b *presets.Builder) {
	db := utils.InitDB()
	b.URIPrefix(NotificationCenterSamplePath).
		DataOperator(gorm2op.DataOperator(db))

	db.AutoMigrate(&utils.Page{})
	b.Model(&utils.Page{})

	b.NotificationFunc(NotifierComponent(), NotifierCount())

	return
}

func NotifierComponent() func(ctx *web.EventContext) h.HTMLComponent {
	return func(ctx *web.EventContext) h.HTMLComponent {
		return v.VList(
			v.VListItem(
				v.VListItemContent(h.A(h.Label("New Notice:"),
					h.Text("unread notes: 3")),
				),
			),
		)
	}
}

func NotifierCount() func(ctx *web.EventContext) int {
	return func(ctx *web.EventContext) int {
		// Use your own count calculation logic here
		return 3
	}
}

// @snippet_end
const NotificationCenterSamplePath = "/samples/notification_center"
