package e18_filter_component

import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	"github.com/goplaid/x/vuetifyx"
)

func FilterComponent(ctx *web.EventContext) (pr web.PageResponse, err error) {

	fd := vuetifyx.FilterData([]*vuetifyx.FilterItem{
		{
			Key:          "invoiceDate",
			Label:        "Invoice Date",
			ItemType:     vuetifyx.ItemTypeDatetimeRange,
			SQLCondition: "InvoiceDate %s datetime(?, 'unixepoch')",
			Selected:     true,
		},
		{
			Key:          "country",
			Label:        "Country",
			ItemType:     vuetifyx.ItemTypeSelect,
			SQLCondition: "upper(BillingCountry) %s upper(?)",
			Options: []*vuetifyx.SelectItem{
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
			ItemType:     vuetifyx.ItemTypeNumber,
			SQLCondition: "Total %s ?",
		},
	})

	fd.SetByQueryString(ctx.R.URL.RawQuery)

	pr.Body = VApp(
		VMain(
			vuetifyx.VXFilter(fd),
		),
	)
	return
}
