package stripeui

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	"github.com/rs/xid"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
	"github.com/thoas/go-funk"
)

type CellComponentFunc func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent
type CellWrapperFunc func(cell h.MutableAttrHTMLComponent, id string, obj interface{}) h.HTMLComponent
type RowMenuItemFunc func(obj interface{}, id string, ctx *web.EventContext) h.HTMLComponent
type RowComponentFunc func(obj interface{}, ctx *web.EventContext) h.HTMLComponent

type DataTableBuilder struct {
	data               interface{}
	selectable         bool
	withoutHeaders     bool
	selectionParamName string
	cellWrapper        CellWrapperFunc
	rowMenuItemFuncs   []RowMenuItemFunc
	rowExpandFunc      RowComponentFunc
	columns            []*DataTableColumnBuilder
	loadMoreCount      int
	loadMoreLabel      string
	loadMoreURL        string
}

func DataTable(data interface{}) (r *DataTableBuilder) {
	r = &DataTableBuilder{
		data:               data,
		selectionParamName: "selected",
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

func (b *DataTableBuilder) WithoutHeader(v bool) (r *DataTableBuilder) {
	b.withoutHeaders = v
	return b
}

func (b *DataTableBuilder) CellWrapperFunc(v CellWrapperFunc) (r *DataTableBuilder) {
	b.cellWrapper = v
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

type primarySlugger interface {
	PrimarySlug() string
}

func (b *DataTableBuilder) MarshalHTML(c context.Context) (r []byte, err error) {
	ctx := web.MustGetEventContext(c)

	selected := getSelectedIds(ctx, b.selectionParamName)

	dataTableId := xid.New().String()
	loadMoreVarName := fmt.Sprintf("loadmore_%s", dataTableId)
	expandVarName := fmt.Sprintf("expand_%s", dataTableId)

	initContextVarsMap := map[string]bool{}

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
			tds = append(tds, h.Td(
				VCheckbox().
					Class("mt-0").
					InputValue(inputValue).
					TrueValue(id).
					FalseValue("").
					HideDetails(true).
					Attr("@change", web.Plaid().
						PushState(true).
						MergeQuery(true).
						Query(b.selectionParamName,
							web.Var(fmt.Sprintf(`{value: %s, add: $event, remove: !$event}`, h.JSONString(id))),
						).RunPushState(),
					),
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
				tdWrapped = b.cellWrapper(std, id, obj)
			}

			bindTds = append(bindTds, tdWrapped)
		}

		if len(objRowMenusMap) > 0 {
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

		rows = append(rows, row)

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

			heads = append(heads, h.Th("").Children(
				VCheckbox().
					Class("mt-0").
					TrueValue(idsOfPageComma).
					InputValue(allInputValue).
					HideDetails(true).
					Attr("@change", web.Plaid().
						PushState(true).
						MergeQuery(true).
						Query(b.selectionParamName,
							web.Var(fmt.Sprintf(`{value: %s, add: $event, remove: !$event}`,
								h.JSONString(idsOfPage))),
						).Go(),
					),
			).Style("width: 48px;").Class("pr-0"))
		}

		orderBys := GetOrderBysFromQuery(ctx.R.URL.Query())
		for _, f := range b.columns {
			head := h.Th(f.title)
			if b.Column(f.name).orderable {
				var orderBy string
				var orderByIdx int
				for i, ob := range orderBys {
					if ob.FieldName == f.name {
						orderBy = ob.OrderBy
						orderByIdx = i + 1
						break
					}
				}
				head = h.Th("").Style("cursor: pointer; white-space: nowrap;").
					Children(
						h.Span(f.title).
							Style("text-decoration: underline;"),
						h.If(orderBy == "ASC",
							VIcon("arrow_drop_up").Small(true),
							h.Span(fmt.Sprint(orderByIdx)),
						).ElseIf(orderBy == "DESC",
							VIcon("arrow_drop_down").Small(true),
							h.Span(fmt.Sprint(orderByIdx)),
						).Else(
							// take up place
							h.Span("").Style("visibility: hidden;").Children(
								VIcon("arrow_drop_down").Small(true),
								h.Span(fmt.Sprint(orderByIdx)),
							),
						),
					).
					Attr("@click", web.Plaid().
						PushState(true).
						Queries(newQueryWithFieldToggleOrderBy(ctx.R.URL.Query(), f.name)).
						Go())
			}
			heads = append(heads, head)
		}

		if len(objRowMenusMap) > 0 {
			heads = append(heads, h.Th(" ").Style("width: 56px;")) // Edit, Delete menu
		}
		thead = h.Thead(
			h.Tr(heads...),
		).Class("grey lighten-5")
	}

	var tfoot h.HTMLComponent
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

	table := VSimpleTable(
		thead,
		h.Tbody(rows...),
		tfoot,
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
	orderable         bool
}

func (b *DataTableColumnBuilder) Name(v string) (r *DataTableColumnBuilder) {
	b.name = v
	return b
}

func (b *DataTableColumnBuilder) Title(v string) (r *DataTableColumnBuilder) {
	b.title = v
	return b
}

func (b *DataTableColumnBuilder) Orderable(v bool) (r *DataTableColumnBuilder) {
	b.orderable = v
	return b
}

func (b *DataTableColumnBuilder) CellComponentFunc(v CellComponentFunc) (r *DataTableColumnBuilder) {
	b.cellComponentFunc = v
	return b
}

type ColOrderBy struct {
	FieldName string
	// ASC, DESC
	OrderBy string
}

func GetOrderBysFromQuery(query url.Values) []*ColOrderBy {
	r := make([]*ColOrderBy, 0)
	qs := strings.Split(query.Get("order_by"), ",")
	for _, q := range qs {
		ss := strings.Split(q, "_")
		ssl := len(ss)
		if ssl == 1 {
			continue
		}
		if ss[ssl-1] != "ASC" && ss[ssl-1] != "DESC" {
			continue
		}
		r = append(r, &ColOrderBy{
			FieldName: strings.Join(ss[:ssl-1], "_"),
			OrderBy:   ss[ssl-1],
		})
	}

	return r
}

func newQueryWithFieldToggleOrderBy(query url.Values, fieldName string) url.Values {
	oldOrderBys := GetOrderBysFromQuery(query)
	newOrderBysQueryValue := []string{}
	existed := false
	for _, oob := range oldOrderBys {
		if oob.FieldName == fieldName {
			existed = true
			if oob.OrderBy == "ASC" {
				newOrderBysQueryValue = append(newOrderBysQueryValue, oob.FieldName+"_DESC")
			}
			continue
		}
		newOrderBysQueryValue = append(newOrderBysQueryValue, oob.FieldName+"_"+oob.OrderBy)
	}
	if !existed {
		newOrderBysQueryValue = append(newOrderBysQueryValue, fieldName+"_ASC")
	}

	newQuery := make(url.Values)
	for k, v := range query {
		if k == "__execute_event__" {
			continue
		}
		newQuery[k] = v
	}
	newQuery.Set("order_by", strings.Join(newOrderBysQueryValue, ","))
	return newQuery
}
