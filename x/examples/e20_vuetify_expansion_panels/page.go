package e20_vuetify_expansion_panels

import (
	"time"

	"github.com/sunfmin/bran/web"
	s "github.com/sunfmin/bran/x/stripeui"
	. "github.com/sunfmin/bran/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type Event struct {
	Title     string
	CreatedAt time.Time
}

func ExpansionPanelDemo(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Schema = VApp(
		VContent(
			VExpansionPanels(
				VExpansionPanel(
					VExpansionPanelHeader(
						h.Text("VISA •••• 4242	11 / 2028"),
						web.Slot(
							VIcon("search"),
						).Name("actions"),
					).DisableIconRotate(true),
					VExpansionPanelContent(
						VDivider(),
						s.DetailInfo(
							s.DetailColumn(
								s.DetailField(s.OptionalText("FENGMIN SUN").ZeroLabel("No Name")).Label("Name"),
								s.DetailField(s.OptionalText("•••• 4242").ZeroLabel("No Number")).Label("Number"),
								s.DetailField(s.OptionalText("QlfGjXhL3I1xfKVV").ZeroLabel("No Fingerprint")).Label("Fingerprint"),
								s.DetailField(s.OptionalText("11 / 2028").ZeroLabel("No Expires")).Label("Expires"),
								s.DetailField(s.OptionalText("Visa credit card").ZeroLabel("No Type")).Label("Type"),
								s.DetailField(s.OptionalText("card_1EJtLGAqkzzGorqLeFb6h2YV").ZeroLabel("No Type")).Label("ID"),
							),
						).Class("pa-0"),
					),
				),

				VExpansionPanel(
					VExpansionPanelHeader(
						h.Text("VISA •••• 2121	11 / 2028"),
					),
					VExpansionPanelContent(
						VDivider(),
						s.DetailInfo(
							s.DetailColumn(
								s.DetailField(s.OptionalText("FENGMIN SUN").ZeroLabel("No Name")).Label("Name"),
								s.DetailField(s.OptionalText("•••• 4242").ZeroLabel("No Number")).Label("Number"),
								s.DetailField(s.OptionalText("QlfGjXhL3I1xfKVV").ZeroLabel("No Fingerprint")).Label("Fingerprint"),
								s.DetailField(s.OptionalText("11 / 2028").ZeroLabel("No Expires")).Label("Expires"),
								s.DetailField(s.OptionalText("Visa credit card").ZeroLabel("No Type")).Label("Type"),
								s.DetailField(s.OptionalText("card_1EJtLGAqkzzGorqLeFb6h2YV").ZeroLabel("No Type")).Label("ID"),
							),
						).Class("pa-0"),
					),
				),
			),
		),
	)
	return
}
