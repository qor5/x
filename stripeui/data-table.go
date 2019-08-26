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

type CellComponentFunc func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent
type CellWrapperFunc func(cell h.MutableAttrHTMLComponent, id string) h.HTMLComponent
type RowMenuItemsFunc func(obj interface{}, id string, ctx *ui.EventContext) []h.HTMLComponent
type RowComponentFunc func(obj interface{}, ctx *ui.EventContext) h.HTMLComponent

type DataTableBuilder struct {
	data               interface{}
	selectable         bool
	withoutHeaders     bool
	selectionParamName string
	cellWrapper        CellWrapperFunc
	rowMenuItemsFunc   RowMenuItemsFunc
	rowExpandFunc      RowComponentFunc
	columns            []*DataTableColumnBuilder
	primaryField       string
	loadMoreCount      int
	loadMoreLabel      string
	loadMoreURL        string
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

func (b *DataTableBuilder) LoadMoreURL(url string) (r *DataTableBuilder) {
	b.loadMoreURL = url
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

func (b *DataTableBuilder) RowExpandFunc(v RowComponentFunc) (r *DataTableBuilder) {
	b.rowExpandFunc = v
	return b
}

func (b *DataTableBuilder) MarshalHTML(c context.Context) (r []byte, err error) {
	ctx := ui.MustGetEventContext(c)

	selected := getSelectedIds(ctx, b.selectionParamName)

	dataTableId := xid.New().String()
	loadMoreVarName := fmt.Sprintf("loadmore_%s", dataTableId)
	expandVarName := fmt.Sprintf("expand_%s", dataTableId)

	initContextVarsMap := map[string]bool{}

	haveRowMenus := b.rowMenuItemsFunc != nil

	var rows []h.HTMLComponent
	var idsOfPage []string

	inPlaceLoadMore := b.loadMoreCount > 0 && len(b.loadMoreURL) == 0

	hasExpand := b.rowExpandFunc != nil

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
		if hasExpand {
			initContextVarsMap[fmt.Sprintf("%s_%d", expandVarName, i)] = false
			tds = append(tds, h.Td(
				VIcon("$vuetify.icons.expand").
					Attr(":class", fmt.Sprintf("{\"v-data-table__expand-icon--active\": vars.%s_%d, \"v-data-table__expand-icon\": true}", expandVarName, i)).
					On("click", fmt.Sprintf("vars.%s_%d = !vars.%s_%d", expandVarName, i, expandVarName, i)),
			).Class("pr-0").Style("width: 48px;"))
		}

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
			if len(b.loadMoreURL) > 0 {
				return
			} else {
				row.Attr("v-if", fmt.Sprintf("vars.%s", loadMoreVarName))
			}
		}

		rows = append(rows, row)

		if hasExpand {
			rows = append(rows, VExpandTransition(h.Tr(
				h.Td(b.rowExpandFunc(obj, ctx)).Attr("colspan", fmt.Sprint(tdCount)),
			).Attr("v-show", fmt.Sprintf("vars.%s_%d", expandVarName, i))))
		}
		i++
	})

	var thead h.HTMLComponent

	if !b.withoutHeaders {

		var heads []h.HTMLComponent

		if hasExpand {
			heads = append(heads, h.Th(" "))
		}

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
		var btn h.HTMLComponent

		if inPlaceLoadMore {
			btn = VBtn(b.loadMoreLabel).
				Text(true).
				Small(true).
				Class("mt-2").
				On("click",
					fmt.Sprintf("vars.%s = !vars.%s", loadMoreVarName, loadMoreVarName))
		} else {
			btn = VBtn(b.loadMoreLabel).
				Text(true).
				Small(true).
				Link(true).
				Class("mt-2").
				Href(b.loadMoreURL)
		}

		tfoot = h.Tfoot(
			h.Tr(
				h.Td(
					h.If(!hasExpand, VDivider()),
					btn,
				).Class("text-center pa-0").Attr("colspan", fmt.Sprint(tdCount)),
			),
		).Attr("v-if", fmt.Sprintf("!vars.%s", loadMoreVarName))
	}

	table := VSimpleTable(
		thead,
		h.Tbody(rows...),
		tfoot,
	)

	if inPlaceLoadMore {
		initContextVarsMap[loadMoreVarName] = false
	}

	if len(initContextVarsMap) > 0 {
		table.Attr("v-init-context-vars", h.JSONString(initContextVarsMap))
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
