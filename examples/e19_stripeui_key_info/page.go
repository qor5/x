package e19_stripeui_key_info

import (
	"fmt"
	"time"

	"github.com/sunfmin/reflectutils"

	s "github.com/sunfmin/bran/stripeui"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

type Event struct {
	Title     string
	CreatedAt time.Time
}

func KeyInfoDemo(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

	dt := s.DataTable([]*Event{
		{
			"<span><strong>¥5,000</strong> was refunded from a <strong>¥236,170</strong> payment</span>",
			time.Now(),
		},
		{
			"<span><strong>¥207,626</strong> was refunded from a <strong>¥236,170</strong> payment</span>",
			time.Now(),
		},
		{
			"<span><strong>¥7,848</strong> was refunded from a <strong>¥236,170</strong> payment</span>",
			time.Now(),
		},
		{
			"<span><strong>¥5,000</strong> was refunded from a <strong>¥236,170</strong> payment</span>",
			time.Now(),
		},
		{
			"<span><strong>¥207,626</strong> was refunded from a <strong>¥236,170</strong> payment</span>",
			time.Now(),
		},
		{
			"<span><strong>¥7,848</strong> was refunded from a <strong>¥236,170</strong> payment</span>",
			time.Now(),
		},
	}).WithoutHeaders(true)

	dt.Column("Title").CellComponentFunc(func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent {
		return h.Td(h.RawHTML(fmt.Sprint(reflectutils.MustGet(obj, fieldName))))
	})

	dt.Column("CreatedAt").CellComponentFunc(func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent {
		t := reflectutils.MustGet(obj, fieldName).(time.Time)
		return h.Td(h.Text(t.Format("01/02/06, 15:04:05 PM"))).Class("text-right")
	})

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
			).Class("mb-4"),

			s.Card(dt).HeaderTitle("Events"),
		),
	)
	return
}
