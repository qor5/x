package e18_filter_component

import (
	"github.com/sunfmin/bran/web"
	. "github.com/sunfmin/bran/x/vuetify"
)

func FilterComponent(ctx *web.EventContext) (pr web.PageResponse, err error) {

	fd := FilterData([]*FilterItem{
		{
			Key:          "invoiceDate",
			Label:        "Invoice Date",
			ItemType:     ItemTypeDate,
			SQLCondition: "InvoiceDate %s datetime(?, 'unixepoch')",
			Selected:     true,
		},
		{
			Key:          "country",
			Label:        "Country",
			ItemType:     ItemTypeSelect,
			SQLCondition: "upper(BillingCountry) %s upper(?)",
			Options: []*SelectItem{
				{
					Value: "US",
					Text:  "United States",
				},
				{
					Value: "CN",
					Text:  "China",
				},
			},
		},
		{
			Key:          "totalAmount",
			Label:        "Total Amount",
			ItemType:     ItemTypeNumber,
			SQLCondition: "Total %s ?",
		},
	})

	fd.SetByQueryString(ctx.R.URL.RawQuery)

	pr.Schema = VApp(
		VContent(
			Filter(fd),
		),
	)
	return
}
