package stripeui

import (
	"context"
	"fmt"
	"strings"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	"github.com/rs/xid"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
	"github.com/thoas/go-funk"
)

type CellComponentFunc func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent
type CellWrapperFunc func(cell h.MutableAttrHTMLComponent, id string, obj interface{}, dataTableID string) h.HTMLComponent
type HeadCellWrapperFunc func(cell h.MutableAttrHTMLComponent, field string, title string) h.HTMLComponent
type RowWrapperFunc func(row h.MutableAttrHTMLComponent, id string, obj interface{}, dataTableID string) h.HTMLComponent
type RowMenuItemFunc func(obj interface{}, id string, ctx *web.EventContext) h.HTMLComponent
type RowComponentFunc func(obj interface{}, ctx *web.EventContext) h.HTMLComponent
type OnSelectFunc func(id string, ctx *web.EventContext) string
type OnSelectAllFunc func(idsOfPage []string, ctx *web.EventContext) string

type DataTableBuilder struct {
	data               interface{}
	selectable         bool
	withoutHeaders     bool
	selectionParamName string
	cellWrapper        CellWrapperFunc
	headCellWrapper    HeadCellWrapperFunc
	rowWrapper         RowWrapperFunc
	rowMenuItemFuncs   []RowMenuItemFunc
	rowExpandFunc      RowComponentFunc
	columns            []*DataTableColumnBuilder
	loadMoreCount      int
	loadMoreLabel      string
	loadMoreURL        string
	// e.g. {count} records are selected.
	selectedCountLabel   string
	tfootChildren        []h.HTMLComponent
	selectableColumnsBtn h.HTMLComponent
	onSelectFunc         OnSelectFunc
	onSelectAllFunc      OnSelectAllFunc
}

func DataTable(data interface{}) (r *DataTableBuilder) {
	r = &DataTableBuilder{
		data:               data,
		selectionParamName: "selected_ids",
		selectedCountLabel: "{count} records are selected.",
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

func (b *DataTableBuilder) Tfoot(children ...h.HTMLComponent) (r *DataTableBuilder) {
	b.tfootChildren = children
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

func (b *DataTableBuilder) WithoutHeader(v bool) (r *DataTableBuilder) {
	b.withoutHeaders = v
	return b
}

func (b *DataTableBuilder) CellWrapperFunc(v CellWrapperFunc) (r *DataTableBuilder) {
	b.cellWrapper = v
	return b
}

func (b *DataTableBuilder) HeadCellWrapperFunc(v HeadCellWrapperFunc) (r *DataTableBuilder) {
	b.headCellWrapper = v
	return b
}

func (b *DataTableBuilder) RowWrapperFunc(v RowWrapperFunc) (r *DataTableBuilder) {
	b.rowWrapper = v
	return b
}

func (b *DataTableBuilder) RowMenuItemFuncs(vs ...RowMenuItemFunc) (r *DataTableBuilder) {
	b.rowMenuItemFuncs = vs
	return b
}

func (b *DataTableBuilder) RowMenuItemFunc(v RowMenuItemFunc) (r *DataTableBuilder) {
	b.rowMenuItemFuncs = append(b.rowMenuItemFuncs, v)
	return b
}

func (b *DataTableBuilder) RowExpandFunc(v RowComponentFunc) (r *DataTableBuilder) {
	b.rowExpandFunc = v
	return b
}

func (b *DataTableBuilder) SelectedCountLabel(v string) (r *DataTableBuilder) {
	b.selectedCountLabel = v
	return b
}

func (b *DataTableBuilder) SelectableColumnsBtn(v h.HTMLComponent) (r *DataTableBuilder) {
	b.selectableColumnsBtn = v
	return b
}

func (b *DataTableBuilder) OnSelectAllFunc(v OnSelectAllFunc) (r *DataTableBuilder) {
	b.onSelectAllFunc = v
	return b
}

func (b *DataTableBuilder) OnSelectFunc(v OnSelectFunc) (r *DataTableBuilder) {
	b.onSelectFunc = v
	return b
}

type primarySlugger interface {
	PrimarySlug() string
}

func (b *DataTableBuilder) MarshalHTML(c context.Context) (r []byte, err error) {
	ctx := web.MustGetEventContext(c)

	selected := getSelectedIds(ctx, b.selectionParamName)

	dataTableId := xid.New().String()
	loadMoreVarName := fmt.Sprintf("loadmore_%s", dataTableId)
	expandVarName := fmt.Sprintf("expand_%s", dataTableId)
	selectedCountVarName := fmt.Sprintf("selected_count_%s", dataTableId)

	initContextVarsMap := map[string]interface{}{
		selectedCountVarName: len(selected),
	}

	// map[obj_id]{rowMenus}
	objRowMenusMap := make(map[string][]h.HTMLComponent)
	funk.ForEach(b.data, func(obj interface{}) {
		id := ObjectID(obj)
		var opMenuItems []h.HTMLComponent
		for _, f := range b.rowMenuItemFuncs {
			item := f(obj, id, ctx)
			if item == nil {
				continue
			}
			opMenuItems = append(opMenuItems, item)
		}
		if len(opMenuItems) > 0 {
			objRowMenusMap[id] = opMenuItems
		}
	})

	hasRowMenuCol := len(objRowMenusMap) > 0 || b.selectableColumnsBtn != nil

	var rows []h.HTMLComponent
	var idsOfPage []string

	inPlaceLoadMore := b.loadMoreCount > 0 && b.loadMoreURL == ""

	hasExpand := b.rowExpandFunc != nil

	i := 0
	tdCount := 0
	haveMoreRecord := false
	funk.ForEach(b.data, func(obj interface{}) {

		id := ObjectID(obj)

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
			).Class("pr-0").Style("width: 40px;"))
		}

		if b.selectable {
			onChange := web.Plaid().
				PushState(true).
				MergeQuery(true).
				Query(b.selectionParamName,
					web.Var(fmt.Sprintf(`{value: %s, add: $event, remove: !$event}`, h.JSONString(id))),
				).RunPushState()
			if b.onSelectFunc != nil {
				onChange = b.onSelectFunc(id, ctx)
			}
			tds = append(tds, h.Td(
				VCheckbox().
					Class("mt-0").
					InputValue(inputValue).
					TrueValue(id).
					FalseValue("").
					HideDetails(true).
					Attr("@change", onChange+fmt.Sprintf(";vars.%s+=($event?1:-1)", selectedCountVarName)),
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
				tdWrapped = b.cellWrapper(std, id, obj, dataTableId)
			}

			bindTds = append(bindTds, tdWrapped)
		}

		if hasRowMenuCol {
			var td h.HTMLComponent
			rowMenus, ok := objRowMenusMap[id]
			if ok {
				td = h.Td(
					VMenu(
						web.Slot(
							VBtn("").Children(
								VIcon("more_horiz"),
							).Attr("v-on", "on").Text(true).Fab(true).Small(true),
						).Name("activator").Scope("{ on }"),

						VList(
							rowMenus...,
						).Dense(true),
					),
				).Style("width: 64px;").Class("pl-0")
			} else {
				td = h.Td().Style("width: 64px;").Class("pl-0")
			}
			bindTds = append(bindTds, td)
		}

		tdCount = len(bindTds)
		row := h.Tr(bindTds...)
		if b.loadMoreCount > 0 && i >= b.loadMoreCount {
			if len(b.loadMoreURL) > 0 {
				return
			} else {
				row.Attr("v-if", fmt.Sprintf("vars.%s", loadMoreVarName))
			}
			haveMoreRecord = true
		}

		if b.rowWrapper != nil {
			rows = append(rows, b.rowWrapper(row, id, obj, dataTableId))
		} else {
			rows = append(rows, row)
		}

		if hasExpand {
			rows = append(rows,
				h.Tr(
					h.Td(
						VExpandTransition(
							h.Div(
								b.rowExpandFunc(obj, ctx),
								VDivider(),
							).Attr("v-if", fmt.Sprintf("vars.%s_%d", expandVarName, i)).
								Class("grey lighten-5"),
						),
					).Attr("colspan", fmt.Sprint(tdCount)).Class("pa-0").Style("height: auto; border-bottom: none"),
				).Class("v-data-table__expand-row"),
			)
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

			onChange := web.Plaid().
				PushState(true).
				MergeQuery(true).
				Query(b.selectionParamName,
					web.Var(fmt.Sprintf(`{value: %s, add: $event, remove: !$event}`,
						h.JSONString(idsOfPage))),
				).Go()
			if b.onSelectAllFunc != nil {
				onChange = b.onSelectAllFunc(idsOfPage, ctx)
			}
			heads = append(heads, h.Th("").Children(
				VCheckbox().
					Class("mt-0").
					TrueValue(idsOfPageComma).
					InputValue(allInputValue).
					HideDetails(true).
					Attr("@change", onChange),
			).Style("width: 48px;").Class("pr-0"))
		}

		for _, f := range b.columns {
			var head h.HTMLComponent
			th := h.Th(f.title)
			head = th
			if b.headCellWrapper != nil {
				head = b.headCellWrapper(
					(h.MutableAttrHTMLComponent)(th),
					f.name,
					f.title,
				)
			}
			heads = append(heads, head)
		}

		if hasRowMenuCol {
			heads = append(heads, h.Th("").Children(b.selectableColumnsBtn).Style("width: 64px;").Class("pl-0")) // Edit, Delete menu
		}
		thead = h.Thead(
			h.Tr(heads...),
		).Class("grey lighten-5")
	}

	var tfoot h.HTMLComponent
	if len(b.tfootChildren) > 0 {
		tfoot = h.Tfoot(b.tfootChildren...)
	}
	if b.loadMoreCount > 0 && haveMoreRecord {
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

	var selectedCountNotice []h.HTMLComponent
	{
		ss := strings.Split(b.selectedCountLabel, "{count}")
		if len(ss) == 1 {
			selectedCountNotice = []h.HTMLComponent{h.Text(ss[0])}
		} else {
			selectedCountNotice = []h.HTMLComponent{
				h.Text(ss[0]),
				h.Strong(fmt.Sprintf("{{vars.%s}}", selectedCountVarName)),
				h.Text(ss[1]),
			}
		}
	}
	table := h.Div(
		h.Div(selectedCountNotice...).
			Class("grey lighten-3 text-center pt-3 pb-3").
			Attr("v-show", fmt.Sprintf("vars.%s > 0", selectedCountVarName)),
		VSimpleTable(
			thead,
			h.Tbody(rows...),
			tfoot,
		),
	)

	if inPlaceLoadMore {
		initContextVarsMap[loadMoreVarName] = false
	}

	if len(initContextVarsMap) > 0 {
		table.Attr(web.InitContextVars, h.JSONString(initContextVarsMap))
	}

	return table.MarshalHTML(c)
}

func ObjectID(obj interface{}) string {
	var id string
	if slugger, ok := obj.(primarySlugger); ok {
		id = slugger.PrimarySlug()
	} else {
		id = fmt.Sprint(reflectutils.MustGet(obj, "ID"))
	}
	return id
}

func getSelectedIds(ctx *web.EventContext, selectedParamName string) (selected []string) {
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

func defaultCellComponentFunc(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent {
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
