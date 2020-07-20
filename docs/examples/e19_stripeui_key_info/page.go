package e19_stripeui_key_info

import (
	"fmt"
	"time"

	"github.com/goplaid/web"
	s "github.com/goplaid/x/stripeui"
	. "github.com/goplaid/x/vuetify"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type Event struct {
	Title     string
	CreatedAt time.Time
}

func KeyInfoDemo(ctx *web.EventContext) (pr web.PageResponse, err error) {
	data := []*Event{
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
	}

	dt := s.DataTable(data).WithoutHeader(true).LoadMoreAt(3, "Show More")

	dt.Column("Title").CellComponentFunc(func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent {
		return h.Td(h.RawHTML(fmt.Sprint(reflectutils.MustGet(obj, fieldName))))
	})

	dt.Column("CreatedAt").CellComponentFunc(func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent {
		t := reflectutils.MustGet(obj, fieldName).(time.Time)
		return h.Td(h.Text(t.Format("01/02/06, 15:04:05 PM"))).Class("text-right")
	})

	logsDt := s.DataTable(data).
		WithoutHeader(true).
		LoadMoreAt(3, "Show More").
		LoadMoreURL("/e20_vuetify_expansion_panels").
		RowExpandFunc(func(obj interface{}, ctx *web.EventContext) h.HTMLComponent {
			return h.Div().Text(h.JSONString(obj)).Class("pa-5")

		})

	logsDt.Column("Title").CellComponentFunc(func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent {
		return h.Td(h.RawHTML(fmt.Sprint(reflectutils.MustGet(obj, fieldName))))
	})

	logsDt.Column("CreatedAt").CellComponentFunc(func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent {
		t := reflectutils.MustGet(obj, fieldName).(time.Time)
		return h.Td(h.Text(t.Format("01/02/06, 15:04:05 PM"))).Class("text-right")
	})

	pr.Body = VApp(
		VMain(
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

			s.Card(s.DetailInfo(
				s.DetailColumn(
					s.DetailField(s.OptionalText("cus_EnUK8WcwQkuKQP")).Label("ID"),
					s.DetailField(s.OptionalText(time.Now().Format("2006/01/02 15:04"))).Label("Created"),
					s.DetailField(s.OptionalText("hello@example.com")).Label("Email"),
					s.DetailField(s.OptionalText("customer0077N52")).Label("Description"),
					s.DetailField(s.OptionalText("B0E69DBD")).Label("Invoice prefix"),
					s.DetailField(s.OptionalText("").ZeroLabel("No VAT number")).Label("VAT number"),
					s.DetailField(s.OptionalText("Normal")).Label("Risk evaluation").Icon(VChip(h.Text("43")).Small(true)),
				).Header("ACCOUNT INFORMATION"),
				s.DetailColumn(
					s.DetailField(s.OptionalText("").ZeroLabel("No address")).Label("Address"),
					s.DetailField(s.OptionalText("").ZeroLabel("No phone number")).Label("Phone number"),
				).Header("BILLING INFORMATION"),
			)).HeaderTitle("Details").
				Actions(VBtn("Update details").Depressed(true)).
				Class("mb-4"),

			s.Card(dt).HeaderTitle("Events").Class("mb-4"),

			s.Card(logsDt).HeaderTitle("Logs").Class("mb-4"),
		),
	)
	return
}
