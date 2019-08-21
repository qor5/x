package e19_stripeui_key_info

import (
	"time"

	s "github.com/sunfmin/bran/stripeui"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

func KeyInfoDemo(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

	pr.Schema = VApp(
		VContent(
			s.Card(
				s.KeyInfo(
					s.KeyField(h.Text(time.Now().Format("Jan _2, 15:04 PM"))).Label("Date"),
					s.KeyField(h.A().Href("https://google.com").Text("customer0077N52")).Label("Customer"),
					s.KeyField(h.Text("•••• 4242")).Label("Payment method").Icon(VIcon("credit_card")),
					s.KeyField(h.Text("Normal")).Label("Risk evaluation").Icon(VChip(h.Text("43")).Small(true)),
				),
			).SystemBar(
				VIcon("link"),
				h.Text("Hello"),
				VSpacer(),
				h.Text("ch_1EJtQMAqkzzGorqLtIjCEPU5"),
			).Header(
				h.Text("$100.00USD"),
				VChip(h.Text("Refunded"), VIcon("reply").Small(true)).Small(true),
			).Actions(
				VBtn("Edit").Depressed(true),
			),
		),
	)
	return
}
