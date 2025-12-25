package vuetifyx

import (
	"context"
	"fmt"
	"strings"

	"github.com/qor5/web/v3"
	v "github.com/qor5/x/v3/ui/vuetify"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type (
	CellComponentFunc   func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent
	CellWrapperFunc     func(cell h.MutableAttrHTMLComponent, id string, obj interface{}, dataTableID string) h.HTMLComponent
	HeadCellWrapperFunc func(cell h.MutableAttrHTMLComponent, field string, title string) h.HTMLComponent
	RowWrapperFunc      func(row h.MutableAttrHTMLComponent, id string, obj interface{}, dataTableID string) h.HTMLComponent
	RowMenuItemFunc     func(obj interface{}, id string, ctx *web.EventContext) h.HTMLComponent
	RowComponentFunc    func(obj interface{}, ctx *web.EventContext) h.HTMLComponent
)

type DataTableBuilder struct {
	data               interface{}
	withoutHeaders     bool
	selectedIds        []string
	onSelectionChanged string // function(selectedIds) { console.log(selectedIds) }
	cellWrapper        CellWrapperFunc
	headCellWrapper    HeadCellWrapperFunc
	rowWrapper         RowWrapperFunc
	rowMenuHead        h.HTMLComponent
	rowMenuItemFuncs   []RowMenuItemFunc
	rowExpandFunc      RowComponentFunc
	columns            []*DataTableColumnBuilder
	loadMoreCount      int
	loadMoreLabel      string
	loadMoreURL        string
	// e.g. {count} records are selected.
	selectedCountLabel  string
	clearSelectionLabel string
	tfootChildren       []h.HTMLComponent
	hover               bool
	hoverClass          string
}

func DataTable(data interface{}) (r *DataTableBuilder) {
	r = &DataTableBuilder{
		data:                data,
		selectedCountLabel:  "{count} records are selected.",
		clearSelectionLabel: "clear selection",
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

func (b *DataTableBuilder) Hover(v bool) (r *DataTableBuilder) {
	b.hover = v
	return b
}

func (b *DataTableBuilder) HoverClass(v string) (r *DataTableBuilder) {
	b.hoverClass = v
	return b
}

func (b *DataTableBuilder) Tfoot(children ...h.HTMLComponent) (r *DataTableBuilder) {
	b.tfootChildren = children
	return b
}

func (b *DataTableBuilder) Data(v interface{}) (r *DataTableBuilder) {
	b.data = v
	return b
}

func (b *DataTableBuilder) SelectedIds(vs []string) (r *DataTableBuilder) {
	b.selectedIds = vs
	return b
}

// OnSelectionChanged
// example: function(selectedIds) { console.log(selectedIds) }
func (b *DataTableBuilder) OnSelectionChanged(v string) (r *DataTableBuilder) {
	b.onSelectionChanged = v
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

func (b *DataTableBuilder) RowMenuHead(v h.HTMLComponent) (r *DataTableBuilder) {
	b.rowMenuHead = v
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

func (b *DataTableBuilder) ClearSelectionLabel(v string) (r *DataTableBuilder) {
	b.clearSelectionLabel = v
	return b
}

type primarySlugger interface {
	PrimarySlug() string
}

func (b *DataTableBuilder) MarshalHTML(c context.Context) (r []byte, err error) {
	ctx := web.MustGetEventContext(c)

	expandVarName := "expand"

	// map[obj_id]{rowMenus}
	objRowMenusMap := make(map[string][]h.HTMLComponent)
	reflectutils.ForEach(b.data, func(obj interface{}) {
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

	hasRowMenuCol := len(objRowMenusMap) > 0 || b.rowMenuHead != nil

	var rows []h.HTMLComponent
	var idsOfPage []string

	inPlaceLoadMore := b.loadMoreCount > 0 && b.loadMoreURL == ""

	hasExpand := b.rowExpandFunc != nil

	i := 0
	tdCount := 0
	haveMoreRecord := false
	reflectutils.ForEach(b.data, func(obj interface{}) {
		id := ObjectID(obj)

		idsOfPage = append(idsOfPage, id)
		var tds []h.HTMLComponent
		if hasExpand {
			localsExpandVarName := fmt.Sprintf("_dataTableLocals_.%s_%d", expandVarName, i)
			tds = append(tds, h.Td(
				v.VIcon("").
					Attr(":icon", fmt.Sprintf(`%s?"mdi-chevron-up-circle":"mdi-chevron-down"`, localsExpandVarName)).
					Attr(":class", fmt.Sprintf(`{"v-data-table__expand-icon--active": _dataTableLocals_.%s_%d, "v-data-table__expand-icon": true}`, expandVarName, i)).
					On("click", fmt.Sprintf("_dataTableLocals_.%s_%d = !_dataTableLocals_.%s_%d", expandVarName, i, expandVarName, i)),
			).Class("pr-0").Style("width: 40px;"))
		}

		if b.onSelectionChanged != "" {
			tds = append(tds, h.Td().Class("pr-0").Children(
				v.VCheckbox().
					Density(v.DensityCompact).
					Class("mt-0").
					Value(id).
					HideDetails(true).
					Attr("@click.native.stop", true).
					Attr("v-model", "_dataTableLocals_.selectedIds"),
			))
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
				tdWrapped = b.cellWrapper(std, id, obj, "")
			}

			bindTds = append(bindTds, tdWrapped)
		}

		if hasRowMenuCol {
			var td h.HTMLComponent
			rowMenus, ok := objRowMenusMap[id]
			if ok {
				td = h.Td(
					v.VMenu(
						web.Slot(
							v.VBtn("").Children(
								v.VIcon("mdi-dots-horizontal"),
							).Attr("v-bind", "props").Variant("text").Size("small"),
						).Name("activator").Scope("{ props }"),

						v.VList(
							rowMenus...,
						),
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
				row.Attr("v-if", "_dataTableLocals_.loadmore")
			}
			haveMoreRecord = true
		}

		if b.rowWrapper != nil {
			rows = append(rows, b.rowWrapper(row, id, obj, ""))
		} else {
			rows = append(rows, row)
		}

		if hasExpand {
			rows = append(rows,
				h.Tr(
					h.Td(
						v.VExpandTransition(
							h.Div(
								b.rowExpandFunc(obj, ctx),
								v.VDivider(),
							).Attr("v-if", fmt.Sprintf("_dataTableLocals_.%s_%d", expandVarName, i)).
								Class("bg-grey-lighten-5"), // bg-grey-lighten-5 | grey lighten-5
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

		if b.onSelectionChanged != "" {
			idsOfPageJSON := h.JSONString(idsOfPage)
			heads = append(heads, h.Th("").Children(
				web.Scope().VSlot("{ locals: _head0Locals_ }").Init(fmt.Sprintf(`{ idsOfPage : %s || []} `, idsOfPageJSON)).Children(
					v.VCheckbox().
						Density(v.DensityCompact).
						Class("mt-0").
						HideDetails(true).
						Attr("@click.native.stop", true).
						Attr(":model-value", "_head0Locals_.idsOfPage.every(id => _dataTableLocals_.selectedIds.includes(id))").
						Attr("@update:model-value", `(value) => {
								let arr = _dataTableLocals_.selectedIds;
								arr = value ? arr.concat(_head0Locals_.idsOfPage) : arr.filter(id => !_head0Locals_.idsOfPage.includes(id))
								_dataTableLocals_.selectedIds = arr.filter((item, index) => arr.indexOf(item) === index);
							}`,
						),
				),
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
			heads = append(heads, h.Th("").Children(b.rowMenuHead).Style("width: 64px;").Class("pl-0")) // Edit, Delete menu
		}
		thead = h.Thead(
			h.Tr(heads...),
		).Class("bg-grey-lighten-5")
	}

	var tfoot h.HTMLComponent
	if len(b.tfootChildren) > 0 {
		tfoot = h.Tfoot(b.tfootChildren...)
	}
	if b.loadMoreCount > 0 && haveMoreRecord {
		var btn h.HTMLComponent

		if inPlaceLoadMore {
			btn = v.VBtn(b.loadMoreLabel).
				Variant("text").
				Size("small").
				Class("mt-2").
				On("click", "_dataTableLocals_.loadmore = !_dataTableLocals_.loadmore")
		} else {
			btn = v.VBtn(b.loadMoreLabel).
				Variant("text").
				Size("small").
				Class("mt-2").
				Href(b.loadMoreURL)
		}

		tfoot = h.Tfoot(
			h.Tr(
				h.Td(
					h.If(!hasExpand, v.VDivider()),
					btn,
				).Class("text-center pa-0").Attr("colspan", fmt.Sprint(tdCount)),
			),
		).Attr("v-if", "!_dataTableLocals_.loadmore")
	}

	var selectedCountCompo h.HTMLComponent
	if b.onSelectionChanged != "" {
		var selectedCountNotice h.HTMLComponents
		ss := strings.Split(b.selectedCountLabel, "{count}")
		if len(ss) == 1 {
			selectedCountNotice = append(selectedCountNotice, h.Text(ss[0]))
		} else {
			selectedCountNotice = append(selectedCountNotice,
				h.Text(ss[0]),
				h.Strong("{{_dataTableLocals_.selectedIds.length}}"),
				h.Text(ss[1]),
			)
		}
		selectedCountCompo = h.Div().Attr("v-show", "_dataTableLocals_.selectedIds.length > 0").Children(
			h.Div().Class("bg-grey-lighten-3 d-flex justify-center align-center ga-1 py-2").Children(
				selectedCountNotice,
				v.VBtn(b.clearSelectionLabel).Variant(v.VariantPlain).Size(v.SizeSmall).On("click", "_dataTableLocals_.selectedIds = [];"),
			),
		)
	}

	selectedIdsJSON := h.JSONString(b.selectedIds)
	onSelectionChanged := b.onSelectionChanged
	if onSelectionChanged == "" {
		onSelectionChanged = "function(v){}"
	}
	return web.Scope().
		VSlot("{ locals:_dataTableLocals_ }").
		Init(fmt.Sprintf(`{ loadmore : false, selectedIds: %s || [], onSelectionChanged: %s, }`, selectedIdsJSON, onSelectionChanged)).
		Children(
			h.Div().Style("display: none;").Attr("v-on-mounted", `({watch}) => {
				watch(() => _dataTableLocals_.selectedIds, (val) => {
					_dataTableLocals_.onSelectionChanged([...val]);
				})
			}`),
			selectedCountCompo,
			v.VTable(
				thead,
				h.Tbody(rows...).ClassIf(b.hoverClass, b.hover && b.hoverClass != ""),
				tfoot,
			).Hover(b.hover),
		).MarshalHTML(c)
}

func ScriptDataTableSwitchSelectedIds(ids ...string) string {
	return fmt.Sprintf(`() => {
		let arr = _dataTableLocals_.selectedIds;
		for(const id of %s) {
			arr = arr.includes(id) ? arr.filter(item => item !== id) : [...arr, id];
		}
		_dataTableLocals_.selectedIds = arr;
	}`, h.JSONString(ids))
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
