package stripeui

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/xid"

	"github.com/sunfmin/bran/ui"
	"github.com/sunfmin/reflectutils"
	"github.com/thoas/go-funk"

	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

type CellWrapperFunc func(cell h.MutableAttrHTMLComponent, id string) h.HTMLComponent

type DataTableBuilder struct {
	data               interface{}
	selectable         bool
	withoutHeaders     bool
	selectionParamName string
	cellWrapper        CellWrapperFunc
	rowMenuItemsFunc   RowMenuItemsFunc
	columns            []*DataTableColumnBuilder
	primaryField       string
	loadMoreCount      int
	loadMoreLabel      string
}

func DataTable(data interface{}) (r *DataTableBuilder) {
	r = &DataTableBuilder{
		data:               data,
		selectionParamName: "selected",
		primaryField:       "ID",
	}
	return
}

func (b *DataTableBuilder) LoadMoreAt(count int, label string) (r *DataTableBuilder) {
	b.loadMoreCount = count
	b.loadMoreLabel = label
	return b
}

func (b *DataTableBuilder) Selectable(v bool) (r *DataTableBuilder) {
	b.selectable = v
	return b
}

func (b *DataTableBuilder) Data(v interface{}) (r *DataTableBuilder) {
	b.data = v
	return b
}

func (b *DataTableBuilder) SelectionParamName(v string) (r *DataTableBuilder) {
	b.selectionParamName = v
	return b
}

func (b *DataTableBuilder) PrimaryField(v string) (r *DataTableBuilder) {
	b.primaryField = v
	return b
}

func (b *DataTableBuilder) WithoutHeader(v bool) (r *DataTableBuilder) {
	b.withoutHeaders = v
	return b
}

func (b *DataTableBuilder) CellWrapperFunc(v CellWrapperFunc) (r *DataTableBuilder) {
	b.cellWrapper = v
	return b
}

func (b *DataTableBuilder) RowMenuItemsFunc(v RowMenuItemsFunc) (r *DataTableBuilder) {
	b.rowMenuItemsFunc = v
	return b
}

func (b *DataTableBuilder) MarshalHTML(c context.Context) (r []byte, err error) {
	ctx := ui.MustGetEventContext(c)

	selected := getSelectedIds(ctx, b.selectionParamName)

	loadMoreVarName := fmt.Sprintf("v_%s", xid.New().String())

	haveRowMenus := b.rowMenuItemsFunc != nil

	var rows []h.HTMLComponent
	var idsOfPage []string

	i := 0
	tdCount := 0
	funk.ForEach(b.data, func(obj interface{}) {

		idRaw, _ := reflectutils.Get(obj, b.primaryField)
		id := fmt.Sprint(idRaw)

		idsOfPage = append(idsOfPage, id)
		inputValue := ""
		if funk.ContainsString(selected, id) {
			inputValue = id
		}
		var tds []h.HTMLComponent
		if b.selectable {
			tds = append(tds, h.Td(
				VCheckbox().
					Class("mt-0").
					FieldName(b.selectionParamName).
					LoadPageWithArrayOp(true).
					InputValue(inputValue).
					TrueValue(id).
					FalseValue("").
					HideDetails(true),
			).Class("pr-0"))
		}

		for _, f := range b.columns {
			tds = append(tds, f.cellComponentFunc(obj, f.name, ctx))
		}

		var bindTds []h.HTMLComponent
		for _, td := range tds {
			std, ok := td.(h.MutableAttrHTMLComponent)
			if !ok {
				bindTds = append(bindTds, td)
				continue
			}

			var tdWrapped h.HTMLComponent = std
			if b.cellWrapper != nil {
				tdWrapped = b.cellWrapper(std, id)
			}

			bindTds = append(bindTds, tdWrapped)
		}

		if haveRowMenus {
			bindTds = append(bindTds, h.Td(
				VMenu(
					ui.Slot(
						VBtn("").Children(
							VIcon("more_horiz"),
						).Attr("v-on", "on").Text(true).Fab(true).Small(true),
					).Name("activator").Scope("{ on }"),

					VList(
						b.rowMenuItemsFunc(obj, id, ctx)...,
					).Dense(true),
				),
			).Style("width: 48px;").Class("pl-0"))
		}

		tdCount = len(bindTds)
		row := h.Tr(bindTds...)
		if b.loadMoreCount > 0 && i >= b.loadMoreCount {
			row.Attr("v-if", fmt.Sprintf("vars.%s", loadMoreVarName))
		}

		rows = append(rows, row)
		i++
	})

	var thead h.HTMLComponent

	if !b.withoutHeaders {

		var heads []h.HTMLComponent

		if b.selectable {
			allInputValue := ""
			idsOfPageComma := strings.Join(idsOfPage, ",")
			if allSelected(selected, idsOfPage) {
				allInputValue = idsOfPageComma
			}

			heads = append(heads, h.Th("").Children(
				VCheckbox().
					Class("mt-0").
					TrueValue(idsOfPageComma).
					InputValue(allInputValue).
					FieldName(b.selectionParamName).
					LoadPageWithArrayOp(true).
					HideDetails(true),
			).Style("width: 48px;").Class("pr-0"))
		}

		for _, f := range b.columns {
			heads = append(heads, h.Th(f.title))
		}

		if haveRowMenus {
			heads = append(heads, h.Th(" ").Style("width: 56px;")) // Edit, Delete menu
		}
		thead = h.Thead(
			h.Tr(heads...),
		).Class("grey lighten-5")
	}

	var tfoot h.HTMLComponent
	if b.loadMoreCount > 0 {
		tfoot = h.Tfoot(
			h.Tr(
				h.Td(
					VDivider(),
					VBtn(b.loadMoreLabel).
						Text(true).
						Small(true).
						Class("mt-1").
						On("click",
							fmt.Sprintf("vars.%s = !vars.%s", loadMoreVarName, loadMoreVarName)),
				).Class("text-center px-0 pt-0").Attr("colspan", fmt.Sprint(tdCount)),
			),
		).Attr("v-if", fmt.Sprintf("!vars.%s", loadMoreVarName))
	}

	table := VSimpleTable(
		thead,
		h.Tbody(
			rows...,
		),
		tfoot,
	)

	if b.loadMoreCount > 0 {
		table.Attr("v-init-context-vars", fmt.Sprintf(`{ %s : false }`, loadMoreVarName))
	}
	return table.MarshalHTML(c)
}

func getSelectedIds(ctx *ui.EventContext, selectedParamName string) (selected []string) {
	selectedValue := ctx.R.URL.Query().Get(selectedParamName)
	if len(selectedValue) > 0 {
		selected = strings.Split(selectedValue, ",")
	}
	return selected
}

func allSelected(selectedInURL []string, pageSelected []string) bool {
	for _, ps := range pageSelected {
		if !funk.ContainsString(selectedInURL, ps) {
			return false
		}
	}
	return true
}

func (b *DataTableBuilder) Column(name string) (r *DataTableColumnBuilder) {
	r = &DataTableColumnBuilder{}
	for _, c := range b.columns {
		if c.name == name {
			return c
		}
	}

	r.Name(name).CellComponentFunc(defaultCellComponentFunc)
	b.columns = append(b.columns, r)
	return
}

type CellComponentFunc func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent
type RowMenuItemsFunc func(obj interface{}, id string, ctx *ui.EventContext) []h.HTMLComponent

func defaultCellComponentFunc(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent {
	return h.Td(h.Text(fmt.Sprint(reflectutils.MustGet(obj, fieldName))))
}

type DataTableColumnBuilder struct {
	name              string
	title             string
	cellComponentFunc CellComponentFunc
}

func (b *DataTableColumnBuilder) Name(v string) (r *DataTableColumnBuilder) {
	b.name = v
	return b
}

func (b *DataTableColumnBuilder) Title(v string) (r *DataTableColumnBuilder) {
	b.title = v
	return b
}

func (b *DataTableColumnBuilder) CellComponentFunc(v CellComponentFunc) (r *DataTableColumnBuilder) {
	b.cellComponentFunc = v
	return b
}
