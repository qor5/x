package presets

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/qor/inflection"
	s "github.com/sunfmin/bran/stripeui"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
	"github.com/thoas/go-funk"
)

type ListingBuilder struct {
	mb             *ModelBuilder
	fields         []*FieldBuilder
	bulkActions    []*BulkActionBuilder
	filterDataFunc FilterDataFunc
	filterTabsFunc FilterTabsFunc
	pageFunc       ui.PageFunc
	searcher       SearchOpFunc
	searchColumns  []string
	perPage        int64
}

func (b *ModelBuilder) Listing(vs ...string) (r *ListingBuilder) {
	r = b.listing
	if len(vs) == 0 {
		return
	}

	var newfields []*FieldBuilder
	for _, f := range vs {
		newfields = append(newfields, r.Field(f))
	}
	r.fields = newfields
	return r
}

func (b *ListingBuilder) Field(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	r = &FieldBuilder{}
	r.name = name
	b.fields = append(b.fields, r)
	return
}

func (b *ListingBuilder) PageFunc(pf ui.PageFunc) (r *ListingBuilder) {
	b.pageFunc = pf
	return b
}

func (b *ListingBuilder) Searcher(v SearchOpFunc) (r *ListingBuilder) {
	b.searcher = v
	return b
}

func (b *ListingBuilder) SearchColumns(vs ...string) (r *ListingBuilder) {
	b.searchColumns = vs
	return b
}

func (b *ListingBuilder) PerPage(v int64) (r *ListingBuilder) {
	b.perPage = v
	return b
}

func (b *ListingBuilder) GetPageFunc() ui.PageFunc {
	if b.pageFunc != nil {
		return b.pageFunc
	}
	return b.defaultPageFunc
}

const selectedParamName = "selected"
const bulkPanelOpenParamName = "bulkOpen"
const bulkPanelPortalName = "bulkPanel"
const deleteConfirmPortalName = "deleteConfirm"

func (b *ListingBuilder) defaultPageFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("formDrawerNew", b.mb.editing.formDrawerNew)
	ctx.Hub.RegisterEventFunc("formDrawerEdit", b.mb.editing.formDrawerEdit)
	ctx.Hub.RegisterEventFunc("update", b.mb.editing.defaultUpdate)
	ctx.Hub.RegisterEventFunc("deleteConfirmation", b.deleteConfirmation)
	ctx.Hub.RegisterEventFunc("doDelete", b.mb.editing.doDelete)
	ctx.Hub.RegisterEventFunc("doBulkAction", b.doBulkAction)

	msgr := b.mb.p.messagesFunc(ctx)
	title := msgr.ListingObjectTitle(inflection.Plural(b.mb.label))
	r.PageTitle = fmt.Sprintf("%s - %s", title, b.mb.p.brandTitle)

	perPage := b.perPage
	if perPage == 0 {
		perPage = 50
	}

	//time.Sleep(1 * time.Second)
	urlQuery := ctx.R.URL.Query()
	searchParams := &SearchParams{
		KeywordColumns: b.searchColumns,
		Keyword:        urlQuery.Get("keyword"),
		PerPage:        perPage,
	}

	searchParams.Page, _ = strconv.ParseInt(urlQuery.Get("page"), 10, 64)
	if searchParams.Page == 0 {
		searchParams.Page = 1
	}

	var fd FilterData
	if b.filterDataFunc != nil {
		fd = b.filterDataFunc(ctx)

		cond, args := fd.SetByQueryString(ctx.R.URL.RawQuery)

		searchParams.SQLConditions = append(searchParams.SQLConditions, &SQLCondition{
			Query: cond,
			Args:  args,
		})
	}

	var objs interface{}
	var totalCount int
	objs, totalCount, err = b.searcher(b.mb.newModelArray(), searchParams)
	if err != nil {
		return
	}

	haveCheckboxes := len(b.bulkActions) > 0

	selected := getSelectedIds(ctx)

	var toolbar h.HTMLComponent
	var bulkPanel h.HTMLComponent
	bulkName := ctx.R.URL.Query().Get(bulkPanelOpenParamName)
	bulk := b.getBulkAction(bulkName)
	if bulk == nil {
		if haveCheckboxes && len(selected) > 0 {
			toolbar = b.bulkActionsToolbar(msgr, ctx)
		} else {
			toolbar = b.newAndFilterToolbar(msgr, ctx, fd)
		}
	} else {
		bulkPanel = ui.LazyPortal(b.bulkPanel(bulk, selected, ctx)).Name(bulkPanelPortalName)
	}

	pagesCount := int(int64(totalCount)/searchParams.PerPage + 1)
	if int64(totalCount)%searchParams.PerPage == 0 {
		pagesCount--
	}

	dataTable := s.DataTable(objs).
		CellWrapperFunc(func(cell h.MutableAttrHTMLComponent, id string) h.HTMLComponent {
			tdbind := ui.Bind(cell)
			if b.mb.hasDetailing {
				tdbind.On("click.self").PushStateURL(b.mb.Info().DetailingHref(id))
			} else {
				tdbind.On("click.self").EventFunc("formDrawerEdit", id)
			}
			return tdbind
		}).
		RowMenuItemsFunc(func(obj interface{}, id string, ctx *ui.EventContext) []h.HTMLComponent {
			return []h.HTMLComponent{
				ui.Bind(VListItem(
					VListItemIcon(VIcon("edit")),
					VListItemTitle(h.Text(msgr.Edit)),
				)).OnClick("formDrawerEdit", id),
				ui.Bind(VListItem(
					VListItemIcon(VIcon("delete")),
					VListItemTitle(h.Text(msgr.Delete)),
				)).OnClick("deleteConfirmation", id),
			}
		}).
		Selectable(haveCheckboxes).
		SelectionParamName(selectedParamName)

	for _, f := range b.fields {
		dataTable.Column(f.name).
			Title(b.mb.getLabel(f.NameLabel)).
			CellComponentFunc(b.cellComponentFunc(f))
	}

	r.Schema = VContainer(

		h.H2(title).Class("title pb-3"),
		b.filterTabs(msgr, ctx),
		bulkPanel,
		VCard(
			toolbar,
			VDivider(),
			VCardText(
				ui.LazyPortal().Name(deleteConfirmPortalName),
				dataTable,
			).Class("pa-0"),
		),

		h.If(pagesCount > 1, h.Components(
			VPagination().Length(pagesCount).Value(int(searchParams.Page)),
		)),
	).Fluid(true)

	return
}

func (b *ListingBuilder) cellComponentFunc(f *FieldBuilder) s.CellComponentFunc {
	return func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent {
		return f.compFunc(obj, b.mb.getComponentFuncField(f), ctx)
	}
}

func getSelectedIds(ctx *ui.EventContext) (selected []string) {
	selectedValue := ctx.R.URL.Query().Get(selectedParamName)
	if len(selectedValue) > 0 {
		selected = strings.Split(selectedValue, ",")
	}
	return selected
}

func (b *ListingBuilder) bulkPanel(bulk *BulkActionBuilder, selectedIds []string, ctx *ui.EventContext) (r h.HTMLComponent) {
	msgr := b.mb.p.messagesFunc(ctx)

	return VCard(
		VCardText(
			bulk.compFunc(selectedIds, ctx),
		),
		VCardActions(
			VSpacer(),
			ui.Bind(VBtn(msgr.Cancel).
				Depressed(true).
				Class("ml-2")).PushStateQuery(url.Values{bulkPanelOpenParamName: []string{""}}).MergeQuery(true),

			VBtn(msgr.OK).
				Color(b.mb.p.primaryColor).
				Depressed(true).
				Dark(true).
				OnClick("doBulkAction", bulk.name, strings.Join(selectedIds, ",")),
		),
	).Class("mb-5")
}

func (b *ListingBuilder) deleteConfirmation(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	msgr := b.mb.p.messagesFunc(ctx)
	id := ctx.Event.Params[0]

	r.UpdatePortals = append(r.UpdatePortals, &ui.PortalUpdate{
		Name: deleteConfirmPortalName,
		Schema: VDialog(
			VCard(
				VCardTitle(h.Text(msgr.DeleteConfirmationText(id))),
				VCardActions(
					VSpacer(),
					VBtn(msgr.Cancel).
						Depressed(true).
						Class("ml-2").
						On("click", "vars.deleteConfirmation = false"),

					VBtn(msgr.Delete).
						Color(b.mb.p.primaryColor).
						Depressed(true).
						Dark(true).
						OnClick("doDelete", id),
				),
			),
		).MaxWidth("600px").
			Attr("v-model", "vars.deleteConfirmation").
			Attr("v-init-context-vars", `{deleteConfirmation: false}`),
		AfterLoaded: "setTimeout(function(){ comp.vars.deleteConfirmation = true }, 100)",
	})
	return
}

func (b *ListingBuilder) doBulkAction(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	bulk := b.getBulkAction(ctx.Event.Params[0])
	if bulk == nil {
		panic("bulk required")
	}
	selectedIds := strings.Split(ctx.Event.Params[1], ",")
	err1 := bulk.updateFunc(selectedIds, ctx.R.MultipartForm, ctx)
	if err1 != nil || ctx.Flash != nil {
		r.UpdatePortals = append(r.UpdatePortals, &ui.PortalUpdate{
			Name:   bulkPanelPortalName,
			Schema: b.bulkPanel(bulk, selectedIds, ctx),
		})
		return
	}

	r.PushState = ui.PushState(url.Values{bulkPanelOpenParamName: []string{}, selectedParamName: []string{}}).MergeQuery(true)

	return
}

func (b *ListingBuilder) bulkActionsToolbar(msgr *Messages, ctx *ui.EventContext) h.HTMLComponent {
	var toolbar = VToolbar(
		VSpacer(),
	).Flat(true)

	for _, ba := range b.bulkActions {
		toolbar.AppendChildren(
			ui.Bind(VBtn(b.mb.getLabel(ba.NameLabel)).
				Color(b.mb.p.primaryColor).
				Depressed(true).
				Dark(true).
				Class("ml-2")).PushStateQuery(url.Values{bulkPanelOpenParamName: []string{ba.name}}).MergeQuery(true),
		)
	}
	return toolbar
}

func (b *ListingBuilder) filterTabs(msgr *Messages, ctx *ui.EventContext) (r h.HTMLComponent) {
	if b.filterTabsFunc == nil {
		return
	}

	tabs := VTabs().Class("mb-3").Grow(true).Value(2)
	tabsData := b.filterTabsFunc(ctx)
	value := -1
	rawQuery := ctx.R.URL.RawQuery
	for i, td := range tabsData {
		if strings.Index(rawQuery, td.Query.Encode()) >= 0 {
			value = i
		}
		tabs.AppendChildren(
			ui.Bind(
				VTab(h.Text(td.Label)),
			).PushStateQuery(td.Query),
		)
	}
	return tabs.Value(value)
}

func (b *ListingBuilder) newAndFilterToolbar(msgr *Messages, ctx *ui.EventContext, fd FilterData) h.HTMLComponent {
	var toolbar = VToolbar(
		VSpacer(),
		VBtn(msgr.New).
			Color(b.mb.p.primaryColor).
			Depressed(true).
			Dark(true).
			OnClick("formDrawerNew", ""),
	).Flat(true)
	if fd != nil {
		toolbar.PrependChildren(Filter(fd))
	}
	return toolbar
}

func allSelected(selectedInURL []string, pageSelected []string) bool {
	for _, ps := range pageSelected {
		if !funk.ContainsString(selectedInURL, ps) {
			return false
		}
	}
	return true
}
