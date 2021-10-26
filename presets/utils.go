package presets

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/presets/actions"
	"github.com/goplaid/x/stripeui"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func EditDeleteRowMenuItemFuncs(m *ModelInfo, url string, editExtraParams ...string) []stripeui.RowMenuItemFunc {
	return []stripeui.RowMenuItemFunc{
		editRowMenuItemFunc(m, url, editExtraParams...),
		deleteRowMenuItemFunc(m, url, editExtraParams...),
	}
}

func editRowMenuItemFunc(m *ModelInfo, url string, editExtraParams ...string) stripeui.RowMenuItemFunc {
	return func(obj interface{}, id string, ctx *web.EventContext) h.HTMLComponent {
		msgr := MustGetMessages(ctx.R)
		if m.Verifier().Do(PermUpdate).ObjectOn(obj).WithReq(ctx.R).IsAllowed() != nil {
			return nil
		}
		return VListItem(
			VListItemIcon(VIcon("edit")),
			VListItemTitle(h.Text(msgr.Edit)),
		).Attr("@click", web.Plaid().
			EventFunc(actions.DrawerEdit, append([]string{id}, editExtraParams...)...).
			URL(url).
			Go())
	}
}

func deleteRowMenuItemFunc(m *ModelInfo, url string, editExtraParams ...string) stripeui.RowMenuItemFunc {
	return func(obj interface{}, id string, ctx *web.EventContext) h.HTMLComponent {
		msgr := MustGetMessages(ctx.R)
		if m.Verifier().Do(PermDelete).ObjectOn(obj).WithReq(ctx.R).IsAllowed() != nil {
			return nil
		}
		return VListItem(
			VListItemIcon(VIcon("delete")),
			VListItemTitle(h.Text(msgr.Delete)),
		).Attr("@click", web.Plaid().
			EventFunc(actions.DeleteConfirmation, id).
			URL(url).
			Go())
	}
}
