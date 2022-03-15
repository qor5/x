package presets

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/goplaid/web"
	"github.com/goplaid/x/i18n"
	"github.com/goplaid/x/perm"
	"github.com/goplaid/x/presets/actions"
	"github.com/goplaid/x/stripeui"
	s "github.com/goplaid/x/stripeui"
	. "github.com/goplaid/x/vuetify"
	"github.com/goplaid/x/vuetifyx"
	h "github.com/theplant/htmlgo"
)

type ListingBuilder struct {
	mb               *ModelBuilder
	bulkActions      []*ActionBuilder
	actions          []*ActionBuilder
	actionsAsMenu    bool
	rowMenu          *RowMenuBuilder
	filterDataFunc   FilterDataFunc
	filterTabsFunc   FilterTabsFunc
	pageFunc         web.PageFunc
	searcher         SearchFunc
	searchColumns    []string
	perPage          int64
	totalVisible     int64
	orderBy          string
	orderableFields  []*OrderableField
	customizedColumn bool
	FieldsBuilder
}

func (mb *ModelBuilder) Listing(vs ...string) (r *ListingBuilder) {
	r = mb.listing
	if len(vs) == 0 {
		return
	}

	r.Only(vs...)
	return r
}

func (b *ListingBuilder) Only(vs ...string) (r *ListingBuilder) {
	r = b
	r.FieldsBuilder = *r.FieldsBuilder.Only(vs...)
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

func (b *ListingBuilder) ActionsAsMenu(v bool) (r *ListingBuilder) {
	b.actionsAsMenu = v
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

func (b *ListingBuilder) EnableCustomizedColumn() (r *ListingBuilder) {
	b.customizedColumn = true
	return b
}

func (b *ListingBuilder) GetPageFunc() web.PageFunc {
	if b.pageFunc != nil {
		return b.pageFunc
	}
	return b.defaultPageFunc
}

const bulkPanelOpenParamName = "bulkOpen"
const bulkPanelPortalName = "bulkPanel"
const deleteConfirmPortalName = "deleteConfirm"
const dataTablePortalName = "dataTable"
const dataTableAdditionsPortalName = "dataTableAdditions"

func (b *ListingBuilder) defaultPageFunc(ctx *web.EventContext) (r web.PageResponse, err error) {
	if b.mb.Info().Verifier().Do(PermList).WithReq(ctx.R).IsAllowed() != nil {
		err = perm.PermissionDenied
		return
	}

	msgr := MustGetMessages(ctx.R)
	title := msgr.ListingObjectTitle(i18n.T(ctx.R, ModelsI18nModuleKey, b.mb.label))
	r.PageTitle = title

	bulkPanel, toolbar, dataTable, dataTableAdditions := b.getComponents(ctx, ctx.R.URL)

	r.Body = VContainer(
		b.filterTabs(msgr, ctx),
		bulkPanel,
		VCard(
			toolbar,
			VDivider(),
			VCardText(
				web.Portal().Name(deleteConfirmPortalName),
				web.Portal(dataTable).Name(dataTablePortalName),
			).Class("pa-0"),
		),

		web.Portal(dataTableAdditions).Name(dataTableAdditionsPortalName),
	).Fluid(true).
		Attr(web.InitContextVars, `{currEditingListItemID: ''}`)

	return
}

func (b *ListingBuilder) cellComponentFunc(f *FieldBuilder) s.CellComponentFunc {
	return func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent {
		return f.compFunc(obj, b.mb.getComponentFuncField(f), ctx)
	}
}

func getSelectedIds(ctx *web.EventContext) (selected []string) {
	selectedValue := ctx.R.URL.Query().Get(ParamSelectedIds)
	if len(selectedValue) > 0 {
		selected = strings.Split(selectedValue, ",")
	}
	return selected
}

func (b *ListingBuilder) bulkPanel(
	bulk *ActionBuilder,
	selectedIds []string,
	processedSelectedIds []string,
	ctx *web.EventContext,
) (r h.HTMLComponent) {
	msgr := MustGetMessages(ctx.R)

	var errComp h.HTMLComponent
	if vErr, ok := ctx.Flash.(*web.ValidationErrors); ok {
		errComp = VAlert(h.Text(vErr.GetGlobalError())).
			Border("left").
			Type("error").
			Elevation(2).
			ColoredBorder(true)
	}
	var processSelectedIdsNotice h.HTMLComponent
	if len(processedSelectedIds) < len(selectedIds) {
		unactionables := make([]string, 0, len(selectedIds))
		{
			processedSelectedIdsM := make(map[string]struct{})
			for _, v := range processedSelectedIds {
				processedSelectedIdsM[v] = struct{}{}
			}
			for _, v := range selectedIds {
				if _, ok := processedSelectedIdsM[v]; !ok {
					unactionables = append(unactionables, v)
				}
			}
		}

		if len(unactionables) > 0 {
			var noticeText string
			if bulk.selectedIdsProcessorNoticeFunc != nil {
				noticeText = bulk.selectedIdsProcessorNoticeFunc(selectedIds, processedSelectedIds, unactionables)
			} else {
				var idsText string
				if len(unactionables) <= 10 {
					idsText = strings.Join(unactionables, ", ")
				} else {
					idsText = fmt.Sprintf("%s...(+%d)", strings.Join(unactionables[:10], ", "), len(unactionables)-10)
				}
				noticeText = msgr.BulkActionSelectedIdsProcessNotice(idsText)
			}
			processSelectedIdsNotice = VAlert(h.Text(noticeText)).
				Type("warning")
		}
	}

	return VCard(
		VCardTitle(
			h.Text(bulk.NameLabel.label),
		),
		VCardText(
			errComp,
			processSelectedIdsNotice,
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
					MergeQuery(true).
					Go(),
				),
		),
	)
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

func (b *ListingBuilder) openBulkActionDialog(ctx *web.EventContext) (r web.EventResponse, err error) {
	msgr := MustGetMessages(ctx.R)
	selected := getSelectedIds(ctx)
	bulkName := ctx.R.URL.Query().Get(bulkPanelOpenParamName)
	bulk := getAction(b.bulkActions, bulkName)

	if bulk == nil {
		err = errors.New("cannot find requested action")
		return
	}

	if len(selected) == 0 {
		ShowMessage(&r, "Please select record", "warning")
		return
	}

	var processedSelectedIds []string
	if bulk.selectedIdsProcessorFunc != nil {
		processedSelectedIds, err = bulk.selectedIdsProcessorFunc(selected, ctx)
		if err != nil {
			return
		}
		if len(processedSelectedIds) == 0 {
			if bulk.selectedIdsProcessorNoticeFunc != nil {
				ShowMessage(&r, bulk.selectedIdsProcessorNoticeFunc(selected, processedSelectedIds, selected), "warning")
			} else {
				ShowMessage(&r, msgr.BulkActionNoAvailableRecords, "warning")
			}
			return
		}
	} else {
		processedSelectedIds = selected
	}

	b.mb.p.dialog(
		&r,
		b.bulkPanel(bulk, selected, processedSelectedIds, ctx),
		bulk.dialogWidth,
	)
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

	var selectedIds []string
	if v := ctx.R.FormValue(ParamSelectedIds); v != "" {
		selectedIds = strings.Split(v, ",")
	}

	var err1 error
	var processedSelectedIds []string
	if bulk.selectedIdsProcessorFunc != nil {
		processedSelectedIds, err1 = bulk.selectedIdsProcessorFunc(selectedIds, ctx)
	} else {
		processedSelectedIds = selectedIds
	}

	if err1 == nil {
		err1 = bulk.updateFunc(processedSelectedIds, ctx)
	}

	if err1 != nil {
		if _, ok := err1.(*web.ValidationErrors); !ok {
			vErr := &web.ValidationErrors{}
			vErr.GlobalError(err1.Error())
			ctx.Flash = vErr
		}
	}

	if ctx.Flash != nil {
		r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
			Name: dialogContentPortalName,
			Body: b.bulkPanel(bulk, selectedIds, processedSelectedIds, ctx),
		})
		return
	}

	msgr := MustGetMessages(ctx.R)
	ShowMessage(&r, msgr.SuccessfullyUpdated, "")

	r.PushState = web.Location(url.Values{bulkPanelOpenParamName: []string{}}).MergeQuery(true)

	return
}

const ActiveFilterTabQueryKey = "active_filter_tab"

func (b *ListingBuilder) filterTabs(msgr *Messages, ctx *web.EventContext) (r h.HTMLComponent) {
	if b.filterTabsFunc == nil {
		return
	}

	tabs := VTabs().Class("mb-3").Grow(true).ShowArrows(true)
	tabsData := b.filterTabsFunc(ctx)
	for i, tab := range tabsData {
		if tab.ID == "" {
			tab.ID = fmt.Sprintf("tab%d", i)
		}
	}
	value := -1
	activeTabValue := ctx.R.URL.Query().Get(ActiveFilterTabQueryKey)

	for i, td := range tabsData {
		// Find selected tab by active_filter_tab=xx in the url query
		if activeTabValue == td.ID {
			value = i
		}

		tabContent := h.Text(td.Label)
		if td.AdvancedLabel != nil {
			tabContent = td.AdvancedLabel
		}

		totalQuery := url.Values{}
		totalQuery.Set(ActiveFilterTabQueryKey, td.ID)
		for k, v := range td.Query {
			totalQuery[k] = v
		}

		tabs.AppendChildren(
			VTab(tabContent).
				Attr("@click", web.Plaid().Queries(totalQuery).
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

	ft.MultipleSelect.In = msgr.FiltersMultipleSelectIn
	ft.MultipleSelect.NotIn = msgr.FiltersMultipleSelectNotIn

	disableNewBtn := b.mb.Info().Verifier().Do(PermCreate).WithReq(ctx.R).IsAllowed() != nil

	var toolbar = VToolbar(
		VSpacer(),
	).Flat(true)

	toolbar.AppendChildren(b.actionsComponent(msgr, ctx))

	if !disableNewBtn {
		toolbar.AppendChildren(VBtn(msgr.New).
			Color("primary").
			Depressed(true).
			Dark(true).Class("ml-2").
			Disabled(disableNewBtn).
			Attr("@click", web.Plaid().EventFunc(actions.New).
				Go()))
	}

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

func (b *ListingBuilder) getComponents(
	ctx *web.EventContext,
	pageURL *url.URL,
) (
	bulkPanel h.HTMLComponent,
	toolbar h.HTMLComponent,
	dataTable h.HTMLComponent,
	// pagination, no-record message
	datatableAdditions h.HTMLComponent,
) {
	msgr := MustGetMessages(ctx.R)

	var requestPerPage int64
	qPerPageStr := pageURL.Query().Get("per_page")
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
	if perPage > 1000 {
		perPage = 1000
	}

	totalVisible := b.totalVisible
	if totalVisible == 0 {
		totalVisible = 10
	}

	var orderBySQL string
	orderBys := s.GetOrderBysFromQuery(pageURL.Query())
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
		Keyword:        pageURL.Query().Get("keyword"),
		PerPage:        perPage,
		OrderBy:        orderBySQL,
	}

	searchParams.Page, _ = strconv.ParseInt(pageURL.Query().Get("page"), 10, 64)
	if searchParams.Page == 0 {
		searchParams.Page = 1
	}

	var fd vuetifyx.FilterData
	if b.filterDataFunc != nil {
		fd = b.filterDataFunc(ctx)
		cond, args := fd.SetByQueryString(pageURL.RawQuery)

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
	var err error

	objs, totalCount, err = b.searcher(b.mb.NewModelSlice(), searchParams, ctx)

	if err != nil {
		panic(err)
	}

	haveCheckboxes := len(b.bulkActions) > 0
	toolbar = b.newAndFilterToolbar(msgr, ctx, fd)

	pagesCount := int(int64(totalCount)/searchParams.PerPage + 1)
	if int64(totalCount)%searchParams.PerPage == 0 {
		pagesCount--
	}

	displayFields := b.fields
	if b.customizedColumn {
		var customizedColumns []string
		var customizedFields []*FieldBuilder
		// get the columns setting from url params or cookie data
		if urldata := pageURL.Query().Get("columns"); urldata != "" {
			if urlColumns := strings.Split(urldata, ","); len(urlColumns) > 0 {
				customizedColumns = urlColumns
			}
		}

		if len(customizedColumns) == 0 {
			cookiedata, err := ctx.R.Cookie(strings.TrimPrefix(strings.ToLower(fmt.Sprintf("%T_columns", b.mb.model)), "*"))
			if err == nil {
				if cookieColumns := strings.Split(cookiedata.Value, ","); len(cookieColumns) > 0 {
					customizedColumns = cookieColumns
				}
			}
		}

		if len(customizedColumns) > 0 {
			for _, f := range b.fields {
				for _, customizedColumn := range customizedColumns {
					if customizedColumn == f.name {
						customizedFields = append(customizedFields, f)
					}
				}
			}

			// save columns setting on cookie
			http.SetCookie(ctx.W, &http.Cookie{
				Name:  strings.TrimPrefix(strings.ToLower(fmt.Sprintf("%T_columns", b.mb.model)), "*"),
				Value: strings.Join(customizedColumns, ","),
			})
		}

		if len(customizedFields) > 0 {
			displayFields = customizedFields
		}

		var columns []string
		for _, f := range displayFields {
			columns = append(columns, f.name)
		}

		// add the HTML component of columns setting into toolbar
		columnList := VList().Dense(true)
		for _, column := range b.fields {
			columnList.AppendChildren(
				VListItem(
					VSwitch().Label(i18n.PT(ctx.R, ModelsI18nModuleKey, b.mb.label, b.mb.getLabel(column.NameLabel))).
						Value(column.name).
						Attr("v-model", "locals.columns").
						On("click", web.Plaid().
							PushState(true).
							Query("columns", web.Var("locals.columns")).
							MergeQuery(true).
							Go()),
				),
			)
		}

		toolbar.(*VToolbarBuilder).AppendChildren(
			VDivider().Vertical(true).Inset(true).Light(true).Attr("style", "margin-left:8px;"),
			VMenu(
				web.Slot(
					VBtn("").Children(
						VIcon("settings"),
					).Attr("v-on", "on").Text(true).Fab(true).Small(true),
				).Name("activator").Scope("{ on }"),
				web.Scope(columnList).Init(fmt.Sprintf(`{columns: %v}`, h.JSONString(columns))).VSlot("{ locals }"),
			),
		)
	}

	dataTable = s.DataTable(objs).
		CellWrapperFunc(func(cell h.MutableAttrHTMLComponent, id string, obj interface{}, dataTableID string) h.HTMLComponent {
			tdbind := cell
			if b.mb.hasDetailing {
				tdbind.SetAttr("@click.self", web.Plaid().
					PushStateURL(
						b.mb.Info().
							DetailingHref(id)).
					Go())
			} else {
				tdbind.SetAttr("@click.self",
					web.Plaid().
						EventFunc(actions.Edit).
						Query(ParamID, id).
						Go()+fmt.Sprintf(`; vars.currEditingListItemID="%s-%s"`, dataTableID, id))
			}
			return tdbind
		}).
		RowWrapperFunc(func(row h.MutableAttrHTMLComponent, id string, obj interface{}, dataTableID string) h.HTMLComponent {
			row.SetAttr(":class", fmt.Sprintf(`{"blue lighten-5": vars.presetsRightDrawer && vars.currEditingListItemID==="%s-%s"}`, dataTableID, id))
			return row
		}).
		RowMenuItemFuncs(b.RowMenu().listingItemFuncs(ctx)...).
		Selectable(haveCheckboxes).
		SelectionParamName(ParamSelectedIds).
		SelectedCountLabel(msgr.ListingSelectedCountNotice)

	for _, f := range displayFields {
		_, ok := orderableFieldMap[f.name]
		dataTable.(*stripeui.DataTableBuilder).Column(f.name).
			Title(i18n.PT(ctx.R, ModelsI18nModuleKey, b.mb.label, b.mb.getLabel(f.NameLabel))).
			CellComponentFunc(b.cellComponentFunc(f)).
			Orderable(ok)
	}

	if totalCount > 0 {
		datatableAdditions = vuetifyx.VXTablePagination().
			Total(int64(totalCount)).
			CurrPage(searchParams.Page).
			PerPage(searchParams.PerPage).
			CustomPerPages([]int64{b.perPage}).
			PerPageText(msgr.PaginationRowsPerPage)
	} else {
		datatableAdditions = h.Div(h.Text(msgr.ListingNoRecordToShow)).Class("mt-10 text-center grey--text text--darken-2")
	}

	return
}

func (b *ListingBuilder) ReloadList(
	ctx *web.EventContext,
	r *web.EventResponse,
	pageURL *url.URL,
) {
	_, _, dataTable, dataTableAdditions := b.getComponents(ctx, pageURL)
	r.UpdatePortals = append(r.UpdatePortals,
		&web.PortalUpdate{
			Name: dataTablePortalName,
			Body: dataTable,
		},
		&web.PortalUpdate{
			Name: dataTableAdditionsPortalName,
			Body: dataTableAdditions,
		},
	)
}

func (b *ListingBuilder) actionsComponent(msgr *Messages, ctx *web.EventContext) h.HTMLComponent {
	var actionBtns []h.HTMLComponent
	for _, ba := range b.bulkActions {
		if b.mb.Info().Verifier().SnakeDo("bulk_actions", ba.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}

		var btn h.HTMLComponent
		if ba.buttonCompFunc != nil {
			btn = ba.buttonCompFunc(ctx)
		} else {
			btn = VBtn(b.mb.getLabel(ba.NameLabel)).
				Color("secondary").
				Depressed(true).
				Dark(true).
				Class("ml-2").
				Attr("@click", web.Plaid().EventFunc(actions.OpenBulkActionDialog).
					Queries(url.Values{bulkPanelOpenParamName: []string{ba.name}}).
					MergeQuery(true).
					Go())
		}

		actionBtns = append(actionBtns, btn)
	}

	for _, ba := range b.actions {
		if b.mb.Info().Verifier().SnakeDo("actions", ba.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}

		var button h.HTMLComponent = VBtn(b.mb.getLabel(ba.NameLabel)).
			Color("primary").
			Depressed(true).
			Dark(true).
			Class("ml-2")
		if ba.buttonCompFunc != nil {
			button = ba.buttonCompFunc(ctx)
		}

		actionBtns = append(actionBtns, button)
	}

	if b.actionsAsMenu {
		var listItems []h.HTMLComponent
		for _, btn := range actionBtns {
			listItems = append(listItems, VListItem(btn))
		}
		return VMenu(
			web.Slot(
				VBtn("Actions").
					Attr("v-bind", "attrs").
					Attr("v-on", "on"),
			).Name("activator").Scope("{ on, attrs }"),
			VList(listItems...),
		).OpenOnHover(true).
			OffsetY(true).
			AllowOverflow(true)
	}
	return h.Components(actionBtns...)
}
