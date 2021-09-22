package presets

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/presets/actions"
	"github.com/goplaid/x/stripeui"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func EditDeleteRowMenuItemsFunc(m *ModelInfo, url string, editExtraParams ...string) stripeui.RowMenuItemsFunc {
	return func(obj interface{}, id string, ctx *web.EventContext) []h.HTMLComponent {
		msgr := MustGetMessages(ctx.R)
		var r []h.HTMLComponent
		if m.Verifier().Do(PermUpdate).ObjectOn(obj).WithReq(ctx.R).IsAllowed() == nil {
			r = append(r,
				VListItem(
					VListItemIcon(VIcon("edit")),
					VListItemTitle(h.Text(msgr.Edit)),
				).Attr("@click", web.Plaid().
					EventFunc(actions.DrawerEdit, append([]string{id}, editExtraParams...)...).
					URL(url).
					Go()),
			)
		}

		if m.Verifier().Do(PermDelete).ObjectOn(obj).WithReq(ctx.R).IsAllowed() == nil {
			r = append(r,
				VListItem(
					VListItemIcon(VIcon("delete")),
					VListItemTitle(h.Text(msgr.Delete)),
				).Attr("@click", web.Plaid().
					EventFunc(actions.DeleteConfirmation, id).
					URL(url).
					Go()),
			)
		}

		return r
	}
}
