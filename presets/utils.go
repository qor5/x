package presets

import (
	"github.com/sunfmin/bran/stripeui"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

func EditDeleteMenuItems(ctx *ui.EventContext, url string, id string, editExtraParams ...string) []h.HTMLComponent {
	msgr := MustGetMessages(ctx.R)
	return []h.HTMLComponent{
		ui.Bind(VListItem(
			VListItemIcon(VIcon("edit")),
			VListItemTitle(h.Text(msgr.Edit)),
		)).OnClick("formDrawerEdit", append([]string{id}, editExtraParams...)...).URL(url),

		ui.Bind(VListItem(
			VListItemIcon(VIcon("delete")),
			VListItemTitle(h.Text(msgr.Delete)),
		)).OnClick("deleteConfirmation", id).URL(url),
	}
}

func EditDeleteRowMenuItemsFunc(ctx *ui.EventContext, url string, editExtraParams ...string) stripeui.RowMenuItemsFunc {
	return func(obj interface{}, id string, ctx *ui.EventContext) []h.HTMLComponent {
		return EditDeleteMenuItems(ctx, url, id, editExtraParams...)
	}
}
