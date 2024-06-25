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

type DataTableBuilderX struct {
	data             interface{}
	withoutHeaders   bool
	varSelectedIDs   string
	cellWrapper      CellWrapperFunc
	headCellWrapper  HeadCellWrapperFunc
	rowWrapper       RowWrapperFunc
	rowMenuHead      h.HTMLComponent
	rowMenuItemFuncs []RowMenuItemFunc
	rowExpandFunc    RowComponentFunc
	columns          []*DataTableColumnBuilder
	loadMoreCount    int
	loadMoreLabel    string
	loadMoreURL      string
	// e.g. {count} records are selected.
	selectedCountLabel  string
	clearSelectionLabel string
	tfootChildren       []h.HTMLComponent
}

func DataTableX(data interface{}) (r *DataTableBuilderX) {
	r = &DataTableBuilderX{
		data:                data,
		selectedCountLabel:  "{count} records are selected.",
		clearSelectionLabel: "clear selection",
	}
	return
}

func (b *DataTableBuilderX) LoadMoreAt(count int, label string) (r *DataTableBuilderX) {
	b.loadMoreCount = count
	b.loadMoreLabel = label
	return b
}

func (b *DataTableBuilderX) LoadMoreURL(url string) (r *DataTableBuilderX) {
	b.loadMoreURL = url
	return b
}

func (b *DataTableBuilderX) Tfoot(children ...h.HTMLComponent) (r *DataTableBuilderX) {
	b.tfootChildren = children
	return b
}

func (b *DataTableBuilderX) Data(v interface{}) (r *DataTableBuilderX) {
	b.data = v
	return b
}

func (b *DataTableBuilderX) VarSelectedIDs(v string) (r *DataTableBuilderX) {
	b.varSelectedIDs = v
	return b
}

func (b *DataTableBuilderX) WithoutHeader(v bool) (r *DataTableBuilderX) {
	b.withoutHeaders = v
	return b
}

func (b *DataTableBuilderX) CellWrapperFunc(v CellWrapperFunc) (r *DataTableBuilderX) {
	b.cellWrapper = v
	return b
}

func (b *DataTableBuilderX) HeadCellWrapperFunc(v HeadCellWrapperFunc) (r *DataTableBuilderX) {
	b.headCellWrapper = v
	return b
}

func (b *DataTableBuilderX) RowWrapperFunc(v RowWrapperFunc) (r *DataTableBuilderX) {
	b.rowWrapper = v
	return b
}

func (b *DataTableBuilderX) RowMenuHead(v h.HTMLComponent) (r *DataTableBuilderX) {
	b.rowMenuHead = v
	return b
}

func (b *DataTableBuilderX) RowMenuItemFuncs(vs ...RowMenuItemFunc) (r *DataTableBuilderX) {
	b.rowMenuItemFuncs = vs
	return b
}

func (b *DataTableBuilderX) RowMenuItemFunc(v RowMenuItemFunc) (r *DataTableBuilderX) {
	b.rowMenuItemFuncs = append(b.rowMenuItemFuncs, v)
	return b
}

func (b *DataTableBuilderX) RowExpandFunc(v RowComponentFunc) (r *DataTableBuilderX) {
	b.rowExpandFunc = v
	return b
}

func (b *DataTableBuilderX) SelectedCountLabel(v string) (r *DataTableBuilderX) {
	b.selectedCountLabel = v
	return b
}

func (b *DataTableBuilderX) ClearSelectionLabel(v string) (r *DataTableBuilderX) {
	b.clearSelectionLabel = v
	return b
}

func (b *DataTableBuilderX) MarshalHTML(c context.Context) (r []byte, err error) {
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

		if b.varSelectedIDs != "" {
			tds = append(tds, h.Td().Class("pr-0").Children(
				v.VCheckbox().
					Density(v.DensityCompact).
					Class("mt-0").
					Value(id).
					HideDetails(true).
					Attr("v-model", b.varSelectedIDs),
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

		if b.varSelectedIDs != "" {
			idsOfPageJSON := h.JSONString(idsOfPage)
			heads = append(heads, h.Th("").Children(
				web.Scope().VSlot("{ locals: _head0Locals_ }").Init(fmt.Sprintf(`{ ids_of_page : %s || []} `, idsOfPageJSON)).Children(
					v.VCheckbox().
						Density(v.DensityCompact).
						Class("mt-0").
						HideDetails(true).
						Attr(":model-value", fmt.Sprintf("_head0Locals_.ids_of_page.every(element => %s.includes(element))", b.varSelectedIDs)).
						Attr("@update:model-value", fmt.Sprintf(`value => {
								const arr = value ? %s.concat(_head0Locals_.ids_of_page) : %s.filter(id => !_head0Locals_.ids_of_page.includes(id)); 
								%s = arr.filter((item, index) => arr.indexOf(item) === index);
							}`,
							b.varSelectedIDs, b.varSelectedIDs,
							b.varSelectedIDs,
						)),
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
	if b.varSelectedIDs != "" {
		var selectedCountNotice h.HTMLComponents
		ss := strings.Split(b.selectedCountLabel, "{count}")
		if len(ss) == 1 {
			selectedCountNotice = append(selectedCountNotice, h.Text(ss[0]))
		} else {
			selectedCountNotice = append(selectedCountNotice,
				h.Text(ss[0]),
				h.Strong(fmt.Sprintf("{{%s.length}}", b.varSelectedIDs)),
				h.Text(ss[1]),
			)
		}
		selectedCountCompo = h.Div().Attr("v-show", fmt.Sprintf("%s.length > 0", b.varSelectedIDs)).Children(
			h.Div().Class("bg-grey-lighten-3 d-flex justify-center align-center ga-1 py-2").Children(
				selectedCountNotice,
				v.VBtn(b.clearSelectionLabel).Variant(v.VariantPlain).Size(v.SizeSmall).On("click", b.varSelectedIDs+" = [];"),
			),
		)
	}

	return web.Scope().VSlot("{ locals:_dataTableLocals_ }").Init(`{ loadmore : false }`).Children(
		selectedCountCompo,
		v.VTable(
			thead,
			h.Tbody(rows...),
			tfoot,
		),
	).MarshalHTML(c)
}

func (b *DataTableBuilderX) Column(name string) (r *DataTableColumnBuilder) {
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
