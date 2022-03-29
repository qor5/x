package basics

import (
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var NotificationCenter = Doc(
	Markdown(`
To enable notification center: Call ~~~NotificationFunc~~~ on ~~~presets.Builder~~~ With 2 function parameters
like this ~~~builder.NotificationFunc(NotifierComponent(db), NotifierCount(db))~~~

The first function is for rendering the content of the popup after user clicked the "bell icon". 
The second function is for rendering the number at the top right corner of the "bell icon". 

`),

	ch.Code(`
presets.New().NotificationFunc(NotifierComponent(db), NotifierCount(db))

func NotifierComponent(db *gorm.DB) func(ctx *web.EventContext) h.HTMLComponent {
	return func(ctx *web.EventContext) h.HTMLComponent {
		return v.VList(
			v.VListItem(
				v.VListItemContent(h.A(h.Label("Demand Video:"),
					h.Text("unread notes").Href("/admin/demand-videos?active_filter_tab=hasUnreadNotes&hasUnreadNotes=1")),
			),
			v.VListItem(
				v.VListItemContent(h.A(h.Label("Supply Video:"),
					h.Text("unread notes").Href("/admin/demand-videos?active_filter_tab=hasUnreadNotes&hasUnreadNotes=1")),
			),
		)
	}
}

func NotifierCount(db *gorm.DB) func(ctx *web.EventContext) int {
	return func(ctx *web.EventContext) int {
		// Use your own count calculation logic here
		return GetUnreadCount(ctx, db)
	}
}
`),
).Slug("basics/notification-center").Title("Notification Center")
