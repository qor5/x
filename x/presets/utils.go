package presets

import (
	"github.com/sunfmin/bran/web"
	"github.com/sunfmin/bran/x/stripeui"
	. "github.com/sunfmin/bran/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func EditDeleteMenuItems(ctx *web.EventContext, url string, id string, editExtraParams ...string) []h.HTMLComponent {
	msgr := MustGetMessages(ctx.R)
	return []h.HTMLComponent{
		web.Bind(VListItem(
			VListItemIcon(VIcon("edit")),
			VListItemTitle(h.Text(msgr.Edit)),
		)).OnClick("formDrawerEdit", append([]string{id}, editExtraParams...)...).URL(url),

		web.Bind(VListItem(
			VListItemIcon(VIcon("delete")),
			VListItemTitle(h.Text(msgr.Delete)),
		)).OnClick("deleteConfirmation", id).URL(url),
	}
}

func EditDeleteRowMenuItemsFunc(ctx *web.EventContext, url string, editExtraParams ...string) stripeui.RowMenuItemsFunc {
	return func(obj interface{}, id string, ctx *web.EventContext) []h.HTMLComponent {
		return EditDeleteMenuItems(ctx, url, id, editExtraParams...)
	}
}
