package e18_filter_component

import (
	fp "github.com/sunfmin/bran/filterpanel"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
)

func FilterComponent(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

	fd := fp.FilterData([]*fp.FilterItem{
		{
			Key:          "invoiceDate",
			Label:        "Invoice Date",
			ItemType:     fp.ItemTypeDate,
			SQLCondition: "InvoiceDate %s datetime(?, 'unixepoch')",
			Selected:     true,
		},
		{
			Key:          "country",
			Label:        "Country",
			ItemType:     fp.ItemTypeSelect,
			SQLCondition: "upper(BillingCountry) %s upper(?)",
			Options: []*fp.SelectItem{
				{
					Key:   "US",
					Label: "United States",
				},
				{
					Key:   "CN",
					Label: "China",
				},
			},
		},
		{
			Key:          "totalAmount",
			Label:        "Total Amount",
			ItemType:     fp.ItemTypeNumber,
			SQLCondition: "Total %s ?",
		},
	})

	fd.SetByQueryString(ctx.R.URL.RawQuery)

	pr.Schema = VApp(
		VContent(
			fp.Filter(fd),
		),
	)
	return
}
