package presets

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/presets/actions"
	"github.com/goplaid/x/stripeui"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func EditDeleteMenuItems(ctx *web.EventContext, url string, id string, editExtraParams ...string) []h.HTMLComponent {
	msgr := MustGetMessages(ctx.R)
	return []h.HTMLComponent{
		VListItem(
			VListItemIcon(VIcon("edit")),
			VListItemTitle(h.Text(msgr.Edit)),
		).Attr("@click", web.Plaid().
			EventFunc(actions.DrawerEdit, append([]string{id}, editExtraParams...)...).
			URL(url).
			Go()),

		VListItem(
			VListItemIcon(VIcon("delete")),
			VListItemTitle(h.Text(msgr.Delete)),
		).Attr("@click", web.Plaid().
			EventFunc(actions.DeleteConfirmation, id).
			URL(url).
			Go()),
	}
}

func EditDeleteRowMenuItemsFunc(ctx *web.EventContext, url string, editExtraParams ...string) stripeui.RowMenuItemsFunc {
	return func(obj interface{}, id string, ctx *web.EventContext) []h.HTMLComponent {
		return EditDeleteMenuItems(ctx, url, id, editExtraParams...)
	}
}
