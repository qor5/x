package presets

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/goplaid/web"
	"github.com/goplaid/x/i18n"
	"github.com/goplaid/x/perm"
	"github.com/goplaid/x/presets/actions"
	s "github.com/goplaid/x/stripeui"
	. "github.com/goplaid/x/vuetify"
	"github.com/goplaid/x/vuetifyx"
	h "github.com/theplant/htmlgo"
)

type ListingBuilder struct {
	mb              *ModelBuilder
	bulkActions     []*ActionBuilder
	rowMenu         *RowMenuBuilder
	filterDataFunc  FilterDataFunc
	filterTabsFunc  FilterTabsFunc
	pageFunc        web.PageFunc
	searcher        SearchFunc
	searchColumns   []string
	perPage         int64
	totalVisible    int64
	orderBy         string
	orderableFields []*OrderableField
	FieldBuilders
}

func (b *ModelBuilder) Listing(vs ...string) (r *ListingBuilder) {
	r = b.listing
	if len(vs) == 0 {
		return
	}

	r.Only(vs...)
	return r
}

func (b *ListingBuilder) Only(vs ...string) (r *ListingBuilder) {
	r = b
	r.FieldBuilders = *r.FieldBuilders.Only(vs...)
	return
}

func (b *ListingBuilder) PageFunc(pf web.PageFunc) (r *ListingBuilder) {
	b.pageFunc = pf
	return b
}

func (b *ListingBuilder) Searcher(v SearchFunc) (r *ListingBuilder) {
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

func (b *ListingBuilder) TotalVisible(v int64) (r *ListingBuilder) {
	b.totalVisible = v
	return b
}

func (b *ListingBuilder) OrderBy(v string) (r *ListingBuilder) {
	b.orderBy = v
	return b
}

type OrderableField struct {
	FieldName string
	DBColumn  string
}

func (b *ListingBuilder) OrderableFields(v []*OrderableField) (r *ListingBuilder) {
	b.orderableFields = v
	return b
}

func (b *ListingBuilder) GetPageFunc() web.PageFunc {
	if b.pageFunc != nil {
		return b.pageFunc
	}
	return b.defaultPageFunc
}

const selectedParamName = "selected"
const bulkPanelOpenParamName = "bulkOpen"
const bulkPanelPortalName = "bulkPanel"
const deleteConfirmPortalName = "deleteConfirm"

func (b *ListingBuilder) defaultPageFunc(ctx *web.EventContext) (r web.PageResponse, err error) {

	if b.mb.Info().Verifier().Do(PermList).WithReq(ctx.R).IsAllowed() != nil {
		err = perm.PermissionDenied
		return
	}

	msgr := MustGetMessages(ctx.R)
	title := msgr.ListingObjectTitle(i18n.T(ctx.R, ModelsI18nModuleKey, b.mb.label))
	r.PageTitle = title

	var requestPerPage int64
	qPerPageStr := ctx.R.URL.Query().Get("per_page")
	qPerPage, _ := strconv.ParseInt(qPerPageStr, 10, 64)
	if qPerPage != 0 {
		setLocalPerPage(ctx, b.mb, qPerPage)
		requestPerPage = qPerPage
	} else if cPerPage := getLocalPerPage(ctx, b.mb); cPerPage != 0 {
		requestPerPage = cPerPage
	}
	perPage := b.perPage
	if requestPerPage != 0 {
		perPage = requestPerPage
	}
	if perPage == 0 {
		perPage = 50
	}

	totalVisible := b.totalVisible
	if totalVisible == 0 {
		totalVisible = 10
	}

	// time.Sleep(1 * time.Second)
	urlQuery := ctx.R.URL.Query()
	var orderBySQL string
	orderBys := s.GetOrderBysFromQuery(urlQuery)
	// map[FieldName]DBColumn
	orderableFieldMap := make(map[string]string)
	for _, v := range b.orderableFields {
		orderableFieldMap[v.FieldName] = v.DBColumn
	}
	for _, ob := range orderBys {
		dbCol, ok := orderableFieldMap[ob.FieldName]
		if !ok {
			continue
		}
		orderBySQL += fmt.Sprintf("%s %s,", dbCol, ob.OrderBy)
	}
	if orderBySQL != "" {
		orderBySQL = orderBySQL[:len(orderBySQL)-1]
	}
	if orderBySQL == "" {
		if b.orderBy != "" {
			orderBySQL = b.orderBy
		} else {
			orderBySQL = fmt.Sprintf("%s DESC", b.mb.primaryField)
		}
	}
	searchParams := &SearchParams{
		KeywordColumns: b.searchColumns,
		Keyword:        urlQuery.Get("keyword"),
		PerPage:        perPage,
		OrderBy:        orderBySQL,
	}

	searchParams.Page, _ = strconv.ParseInt(urlQuery.Get("page"), 10, 64)
	if searchParams.Page == 0 {
		searchParams.Page = 1
	}

	var fd vuetifyx.FilterData
	if b.filterDataFunc != nil {
		fd = b.filterDataFunc(ctx)

		cond, args := fd.SetByQueryString(ctx.R.URL.RawQuery)

		searchParams.SQLConditions = append(searchParams.SQLConditions, &SQLCondition{
			Query: cond,
			Args:  args,
		})
	}

	if b.searcher == nil || b.mb.p.dataOperator == nil {
		panic("presets.New().DataOperator(...) required")
	}

	var objs interface{}
	var totalCount int
	objs, totalCount, err = b.searcher(b.mb.NewModelSlice(), searchParams, ctx)
	if err != nil {
		panic(err)
	}

	haveCheckboxes := len(b.bulkActions) > 0

	selected := getSelectedIds(ctx)

	var toolbar h.HTMLComponent
	var bulkPanel h.HTMLComponent
	bulkName := ctx.R.URL.Query().Get(bulkPanelOpenParamName)
	bulk := getAction(b.bulkActions, bulkName)
	if bulk == nil {
		if haveCheckboxes && len(selected) > 0 {
			toolbar = b.bulkActionsToolbar(msgr, ctx)
		} else {
			toolbar = b.newAndFilterToolbar(msgr, ctx, fd)
		}
	} else {
		bulkPanel = web.Portal(b.bulkPanel(bulk, selected, ctx)).Name(bulkPanelPortalName)
	}

	pagesCount := int(int64(totalCount)/searchParams.PerPage + 1)
	if int64(totalCount)%searchParams.PerPage == 0 {
		pagesCount--
	}

	dataTable := s.DataTable(objs).
		CellWrapperFunc(func(cell h.MutableAttrHTMLComponent, id string, obj interface{}) h.HTMLComponent {
			tdbind := cell
			if b.mb.hasDetailing {
				tdbind.SetAttr("@click.self", web.Plaid().
					PushStateURL(
						b.mb.Info().
							DetailingHref(id)).
					Go())
			} else {
				if b.mb.Info().Verifier().Do(PermUpdate).ObjectOn(obj).WithReq(ctx.R).IsAllowed() == nil {

					tdbind.SetAttr("@click.self",
						web.Plaid().
							EventFunc(actions.Edit).
							Query(ParamID, id).
							Go())
				}

			}
			return tdbind
		}).
		RowMenuItemFuncs(b.RowMenu().listingItemFuncs(ctx)...).
		Selectable(haveCheckboxes).
		SelectionParamName(selectedParamName)

	for _, f := range b.fields {
		_, ok := orderableFieldMap[f.name]
		dataTable.Column(f.name).
			Title(i18n.PT(ctx.R, ModelsI18nModuleKey, b.mb.label, b.mb.getLabel(f.NameLabel))).
			CellComponentFunc(b.cellComponentFunc(f)).
			Orderable(ok)
	}

	r.Body = VContainer(

		b.filterTabs(msgr, ctx),
		bulkPanel,
		VCard(
			toolbar,
			VDivider(),
			VCardText(
				web.Portal().Name(deleteConfirmPortalName),
				dataTable,
			).Class("pa-0"),
		),

		vTablePagination(
			msgr,
			int64(totalCount),
			searchParams.Page,
			searchParams.PerPage,
			[]int64{b.perPage},
		),
	).Fluid(true)

	return
}

func (b *ListingBuilder) cellComponentFunc(f *FieldBuilder) s.CellComponentFunc {
	return func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent {
		return f.compFunc(obj, b.mb.getComponentFuncField(f), ctx)
	}
}

func getSelectedIds(ctx *web.EventContext) (selected []string) {
	selectedValue := ctx.R.URL.Query().Get(selectedParamName)
	if len(selectedValue) > 0 {
		selected = strings.Split(selectedValue, ",")
	}
	return selected
}

func (b *ListingBuilder) bulkPanel(bulk *ActionBuilder, selectedIds []string, ctx *web.EventContext) (r h.HTMLComponent) {
	msgr := MustGetMessages(ctx.R)

	return VCard(
		VCardText(
			bulk.compFunc(selectedIds, ctx),
		),
		VCardActions(
			VSpacer(),
			VBtn(msgr.Cancel).
				Depressed(true).
				Class("ml-2").
				Attr("@click", web.Plaid().
					Queries(url.Values{bulkPanelOpenParamName: []string{""}}).
					MergeQuery(true).
					PushState(true).
					Go()),

			VBtn(msgr.OK).
				Color("primary").
				Depressed(true).
				Dark(true).
				Attr("@click", web.Plaid().EventFunc(actions.DoBulkAction).
					Query(ParamBulkActionName, bulk.name).
					Query(ParamSelectedIds, strings.Join(selectedIds, ",")).
					Go(),
				),
		),
	).Class("mb-5")
}

func (b *ListingBuilder) deleteConfirmation(ctx *web.EventContext) (r web.EventResponse, err error) {
	msgr := MustGetMessages(ctx.R)
	id := ctx.R.FormValue(ParamID)

	r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
		Name: deleteConfirmPortalName,
		Body: VDialog(
			VCard(
				VCardTitle(h.Text(msgr.DeleteConfirmationText(id))),
				VCardActions(
					VSpacer(),
					VBtn(msgr.Cancel).
						Depressed(true).
						Class("ml-2").
						On("click", "vars.deleteConfirmation = false"),

					VBtn(msgr.Delete).
						Color("primary").
						Depressed(true).
						Dark(true).
						Attr("@click", web.Plaid().
							EventFunc(actions.DoDelete).
							Query(ParamID, id).
							URL(ctx.R.URL.Path).
							Go()),
				),
			),
		).MaxWidth("600px").
			Attr("v-model", "vars.deleteConfirmation").
			Attr(web.InitContextVars, `{deleteConfirmation: false}`),
	})

	r.VarsScript = "setTimeout(function(){ vars.deleteConfirmation = true }, 100)"
	return
}

func (b *ListingBuilder) doBulkAction(ctx *web.EventContext) (r web.EventResponse, err error) {
	bulk := getAction(b.bulkActions, ctx.R.FormValue(ParamBulkActionName))
	if bulk == nil {
		panic("bulk required")
	}

	if b.mb.Info().Verifier().SnakeDo("bulk_actions", bulk.name).WithReq(ctx.R).IsAllowed() != nil {
		err = perm.PermissionDenied
		return
	}

	selectedIds := strings.Split(ctx.R.FormValue(ParamSelectedIds), ",")
	err1 := bulk.updateFunc(selectedIds, ctx)
	if err1 != nil || ctx.Flash != nil {
		r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
			Name: bulkPanelPortalName,
			Body: b.bulkPanel(bulk, selectedIds, ctx),
		})
		return
	}

	r.PushState = web.Location(url.Values{bulkPanelOpenParamName: []string{}, selectedParamName: []string{}}).MergeQuery(true)

	return
}

func (b *ListingBuilder) bulkActionsToolbar(msgr *Messages, ctx *web.EventContext) h.HTMLComponent {
	var toolbar = VToolbar(
		VSpacer(),
	).Flat(true)

	for _, ba := range b.bulkActions {
		if b.mb.Info().Verifier().SnakeDo("bulk_actions", ba.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}

		toolbar.AppendChildren(
			VBtn(b.mb.getLabel(ba.NameLabel)).
				Color("primary").
				Depressed(true).
				Dark(true).
				Class("ml-2").
				Attr("@click", web.Plaid().
					Queries(url.Values{bulkPanelOpenParamName: []string{ba.name}}).
					MergeQuery(true).
					Go()),
		)
	}
	return toolbar
}

func (b *ListingBuilder) filterTabs(msgr *Messages, ctx *web.EventContext) (r h.HTMLComponent) {
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
			VTab(h.Text(td.Label)).
				Attr("@click", web.Plaid().Queries(td.Query).
					PushState(true).Go()),
		)
	}
	return tabs.Value(value)
}

func (b *ListingBuilder) newAndFilterToolbar(msgr *Messages, ctx *web.EventContext, fd vuetifyx.FilterData) h.HTMLComponent {
	ft := vuetifyx.FilterTranslations{}
	ft.Filters = msgr.Filters
	ft.Filter = msgr.Filter
	ft.Done = msgr.FiltersDone
	ft.Clear = msgr.FiltersClear

	ft.Date.InTheLast = msgr.FiltersDateInTheLast
	ft.Date.Days = msgr.FiltersDateDays
	ft.Date.Months = msgr.FiltersDateMonths
	ft.Date.And = msgr.FiltersDateAnd
	ft.Date.Between = msgr.FiltersDateBetween
	ft.Date.Equals = msgr.FiltersDateEquals
	ft.Date.IsAfter = msgr.FiltersDateIsAfter
	ft.Date.IsAfterOrOn = msgr.FiltersDateIsAfterOrOn
	ft.Date.IsBeforeOrOn = msgr.FiltersDateIsBeforeOrOn
	ft.Date.IsBefore = msgr.FiltersDateIsBefore

	ft.Number.And = msgr.FiltersNumberAnd
	ft.Number.Equals = msgr.FiltersNumberEquals
	ft.Number.Between = msgr.FiltersNumberBetween
	ft.Number.GreaterThan = msgr.FiltersNumberGreaterThan
	ft.Number.LessThan = msgr.FiltersNumberLessThan

	ft.String.Equals = msgr.FiltersStringEquals
	ft.String.Contains = msgr.FiltersStringContains

	disableNewBtn := b.mb.Info().Verifier().Do(PermCreate).WithReq(ctx.R).IsAllowed() != nil

	var toolbar = VToolbar(
		VSpacer(),
		VBtn(msgr.New).
			Color("primary").
			Depressed(true).
			Dark(true).
			Disabled(disableNewBtn).
			Attr("@click", web.Plaid().EventFunc(actions.New).
				Go()),
	).Flat(true)
	if fd != nil {
		toolbar.PrependChildren(vuetifyx.VXFilter(fd).Translations(ft))
	}
	return toolbar
}

func getLocalPerPage(
	ctx *web.EventContext,
	mb *ModelBuilder,
) int64 {
	c, err := ctx.R.Cookie("_perPage")
	if err != nil {
		return 0
	}
	vals := strings.Split(c.Value, "$")
	for _, v := range vals {
		vvs := strings.Split(v, "#")
		if len(vvs) != 2 {
			continue
		}
		if vvs[0] == mb.uriName {
			r, _ := strconv.ParseInt(vvs[1], 10, 64)
			return r
		}
	}

	return 0
}

func setLocalPerPage(
	ctx *web.EventContext,
	mb *ModelBuilder,
	v int64,
) {
	var oldVals []string
	{
		c, err := ctx.R.Cookie("_perPage")
		if err == nil {
			oldVals = strings.Split(c.Value, "$")
		}
	}
	newVals := []string{fmt.Sprintf("%s#%d", mb.uriName, v)}
	for _, v := range oldVals {
		vvs := strings.Split(v, "#")
		if len(vvs) != 2 {
			continue
		}
		if vvs[0] == mb.uriName {
			continue
		}
		newVals = append(newVals, v)
	}
	http.SetCookie(ctx.W, &http.Cookie{
		Name:  "_perPage",
		Value: strings.Join(newVals, "$"),
	})
}

func vTablePagination(
	msgr *Messages,
	total int64,
	currPage int64,
	perPage int64,
	customPerPages []int64,
) h.HTMLComponent {
	var sItems []string
	{
		perPagesM := map[int64]struct{}{
			10:  {},
			15:  {},
			20:  {},
			50:  {},
			100: {},
		}
		if perPage > 0 {
			perPagesM[perPage] = struct{}{}
		}
		for _, v := range customPerPages {
			if v <= 0 {
				continue
			}
			perPagesM[v] = struct{}{}
		}
		perPages := make([]int, 0, len(perPagesM))
		for k, _ := range perPagesM {
			perPages = append(perPages, int(k))
		}
		sort.Ints(perPages)
		for _, v := range perPages {
			sItems = append(sItems, fmt.Sprint(v))
		}
	}

	currPageStart := (currPage-1)*perPage + 1
	currPageEnd := currPage * perPage
	if currPageEnd > total {
		currPageEnd = total
	}

	canNext := false
	canPrev := false
	if currPage*perPage < total {
		canNext = true
	}
	if currPage > 1 {
		canPrev = true
	}
	var nextIconStyle string
	var prevIconStyle string
	if canNext {
		nextIconStyle = "cursor: pointer;"
	}
	if canPrev {
		prevIconStyle = "cursor: pointer;"
	}

	rowsPerPageText := "Rows per page: "
	if msgr.PaginationRowsPerPage != "" {
		rowsPerPageText = msgr.PaginationRowsPerPage
	}
	return VRow().Justify("end").Align("center").Class("mt-3 mr-3").
		Children(
			h.Div(
				h.Text(rowsPerPageText),
			),
			h.Div(
				VSelect().Items(sItems).Value(fmt.Sprint(perPage)).
					Attr("@input", web.Plaid().
						PushState(true).
						Query("per_page", web.Var("[$event]")).
						MergeQuery(true).
						Go()),
			).Style("width: 60px;").Class("ml-6"),
			h.Div(
				h.Text(fmt.Sprintf("%d-%d of %d", currPageStart, currPageEnd, total)),
			).Class("ml-6"),
			h.Div(
				h.Span("").Style(prevIconStyle).Children(
					VIcon("navigate_before").Size(32).Disabled(!canPrev).
						Attr("@click", web.Plaid().
							PushState(true).
							Query("page", currPage-1).
							MergeQuery(true).
							Go()),
				),
				h.Span("").Style(nextIconStyle).Children(
					VIcon("navigate_next").Size(32).Disabled(!canNext).
						Attr("@click", web.Plaid().
							PushState(true).
							Query("page", currPage+1).
							MergeQuery(true).
							Go()),
				).Class("ml-3"),
			).Class("ml-6"),
		)
}
