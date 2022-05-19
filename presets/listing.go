package presets

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
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
	mb                *ModelBuilder
	bulkActions       []*ActionBuilder
	actions           []*ActionBuilder
	actionsAsMenu     bool
	rowMenu           *RowMenuBuilder
	filterDataFunc    FilterDataFunc
	filterTabsFunc    FilterTabsFunc
	newBtnFunc        ComponentFunc
	pageFunc          web.PageFunc
	searcher          SearchFunc
	searchColumns     []string
	perPage           int64
	totalVisible      int64
	orderBy           string
	orderableFields   []*OrderableField
	selectableColumns bool
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
	ivs := make([]interface{}, 0, len(vs))
	for _, v := range vs {
		ivs = append(ivs, v)
	}
	r.FieldsBuilder = *r.FieldsBuilder.Only(ivs...)
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

func (b *ListingBuilder) NewButtonFunc(v ComponentFunc) (r *ListingBuilder) {
	b.newBtnFunc = v
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

func (b *ListingBuilder) SelectableColumns(v bool) (r *ListingBuilder) {
	b.selectableColumns = v
	return b
}

func (b *ListingBuilder) GetPageFunc() web.PageFunc {
	if b.pageFunc != nil {
		return b.pageFunc
	}
	return b.defaultPageFunc
}

const bulkPanelOpenParamName = "bulkOpen"
const actionPanelOpenParamName = "actionOpen"
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

func (b *ListingBuilder) actionPanel(action *ActionBuilder, ctx *web.EventContext) (r h.HTMLComponent) {
	msgr := MustGetMessages(ctx.R)

	var errComp h.HTMLComponent
	if vErr, ok := ctx.Flash.(*web.ValidationErrors); ok {
		errComp = VAlert(h.Text(vErr.GetGlobalError())).
			Border("left").
			Type("error").
			Elevation(2).
			ColoredBorder(true)
	}

	return VCard(
		VCardTitle(
			h.Text(action.NameLabel.label),
		),
		VCardText(
			errComp,
			action.compFunc([]string{}, ctx), // because action and bulk action shared the same func, so pass blank slice here
		),
		VCardActions(
			VSpacer(),
			VBtn(msgr.Cancel).
				Depressed(true).
				Class("ml-2").
				Attr("@click", web.Plaid().
					Queries(url.Values{actionPanelOpenParamName: []string{""}}).
					MergeQuery(true).
					PushState(true).
					Go()),

			VBtn(msgr.OK).
				Color("primary").
				Depressed(true).
				Dark(true).
				Attr("@click", web.Plaid().EventFunc(actions.DoListingAction).
					Query(ParamListingActionName, action.name).
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

func (b *ListingBuilder) openActionDialog(ctx *web.EventContext) (r web.EventResponse, err error) {
	actionName := ctx.R.URL.Query().Get(actionPanelOpenParamName)
	action := getAction(b.actions, actionName)
	if action == nil {
		err = errors.New("cannot find requested action")
		return
	}

	b.mb.p.dialog(
		&r,
		b.actionPanel(action, ctx),
		action.dialogWidth,
	)
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

	// If selectedIdsProcessorFunc is not nil, process the request in it and skip the confirmation dialog
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

	if b.mb.Info().Verifier().SnakeDo(PermBulkActions, bulk.name).WithReq(ctx.R).IsAllowed() != nil {
		ShowMessage(&r, perm.PermissionDenied.Error(), "warning")
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

func (b ListingBuilder) doListingAction(ctx *web.EventContext) (r web.EventResponse, err error) {
	action := getAction(b.actions, ctx.R.FormValue(ParamListingActionName))
	if action == nil {
		panic("action required")
	}

	if b.mb.Info().Verifier().SnakeDo(PermListingActions, action.name).WithReq(ctx.R).IsAllowed() != nil {
		ShowMessage(&r, perm.PermissionDenied.Error(), "warning")
		return
	}

	err1 := action.updateFunc([]string{}, ctx)

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
			Body: b.actionPanel(action, ctx),
		})
		return
	}

	msgr := MustGetMessages(ctx.R)
	ShowMessage(&r, msgr.SuccessfullyUpdated, "")

	r.PushState = web.Location(url.Values{actionPanelOpenParamName: []string{}}).MergeQuery(true)

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

type selectColumns struct {
	DisplayColumns []string       `json:"displayColumns,omitempty"`
	SortedColumns  []sortedColumn `json:"sortedColumns,omitempty"`
}
type sortedColumn struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

func (b *ListingBuilder) selectColumnsBtn(pageURL *url.URL, ctx *web.EventContext) (btn h.HTMLComponent, displaySortedFields []*FieldBuilder) {
	var (
		_, respath         = path.Split(pageURL.Path)
		displayColumnsName = fmt.Sprintf("%s_display_columns", respath)
		sortedColumnsName  = fmt.Sprintf("%s_sorted_columns", respath)
		originalColumns    []string
		displayColumns     []string
		sortedColumns      []string
	)

	for _, f := range b.fields {
		if b.mb.Info().Verifier().Do(PermList).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}
		originalColumns = append(originalColumns, f.name)
	}

	// get the columns setting from url params or cookie data
	if urldata := pageURL.Query().Get(displayColumnsName); urldata != "" {
		if urlColumns := strings.Split(urldata, ","); len(urlColumns) > 0 {
			displayColumns = urlColumns
		}
	}

	if urldata := pageURL.Query().Get(sortedColumnsName); urldata != "" {
		if urlColumns := strings.Split(urldata, ","); len(urlColumns) > 0 {
			sortedColumns = urlColumns
		}
	}

	// get the columns setting from  cookie data
	if len(displayColumns) == 0 {
		cookiedata, err := ctx.R.Cookie(displayColumnsName)
		if err == nil {
			if cookieColumns := strings.Split(cookiedata.Value, ","); len(cookieColumns) > 0 {
				displayColumns = cookieColumns
			}
		}
	}

	if len(sortedColumns) == 0 {
		cookiedata, err := ctx.R.Cookie(sortedColumnsName)
		if err == nil {
			if cookieColumns := strings.Split(cookiedata.Value, ","); len(cookieColumns) > 0 {
				sortedColumns = cookieColumns
			}
		}
	}

	// check if listing fileds is changed. if yes, use the original columns
	var originalFiledsChanged bool

	if len(sortedColumns) > 0 && len(originalColumns) != len(sortedColumns) {
		originalFiledsChanged = true
	}

	if len(sortedColumns) > 0 && !originalFiledsChanged {
		for _, sortedColumn := range sortedColumns {
			var find bool
			for _, originalColumn := range originalColumns {
				if sortedColumn == originalColumn {
					find = true
					break
				}
			}
			if !find {
				originalFiledsChanged = true
				break
			}
		}
	}

	if len(displayColumns) > 0 && !originalFiledsChanged {
		for _, displayColumn := range displayColumns {
			var find bool
			for _, originalColumn := range originalColumns {
				if displayColumn == originalColumn {
					find = true
					break
				}
			}
			if !find {
				originalFiledsChanged = true
				break
			}
		}
	}

	// save display columns setting on cookie
	if !originalFiledsChanged && len(displayColumns) > 0 {
		http.SetCookie(ctx.W, &http.Cookie{
			Name:  displayColumnsName,
			Value: strings.Join(displayColumns, ","),
		})
	}

	// save sorted columns setting on cookie
	if !originalFiledsChanged && len(sortedColumns) > 0 {
		http.SetCookie(ctx.W, &http.Cookie{
			Name:  sortedColumnsName,
			Value: strings.Join(sortedColumns, ","),
		})
	}

	// set the data for displaySortedFields on data table
	if originalFiledsChanged || (len(sortedColumns) == 0 && len(displayColumns) == 0) {
		displaySortedFields = b.fields
	}

	if originalFiledsChanged || len(displayColumns) == 0 {
		displayColumns = originalColumns
	}

	if originalFiledsChanged || len(sortedColumns) == 0 {
		sortedColumns = originalColumns
	}

	if len(displaySortedFields) == 0 {
		for _, sortedColumn := range sortedColumns {
			for _, displayColumn := range displayColumns {
				if sortedColumn == displayColumn {
					displaySortedFields = append(displaySortedFields, b.Field(sortedColumn))
					break
				}
			}
		}
	}

	// set the data for selected columns on toolbar
	selectColumns := selectColumns{
		DisplayColumns: displayColumns,
	}
	for _, sc := range sortedColumns {
		selectColumns.SortedColumns = append(selectColumns.SortedColumns, sortedColumn{
			Name:  sc,
			Label: i18n.PT(ctx.R, ModelsI18nModuleKey, b.mb.label, b.mb.getLabel(b.Field(sc).NameLabel)),
		})
	}

	msgr := MustGetMessages(ctx.R)
	// add the HTML component of columns setting into toolbar
	btn = VMenu(
		web.Slot(
			VBtn("").Children(VIcon("settings")).Attr("v-on", "on").Text(true).Fab(true).Small(true),
		).Name("activator").Scope("{ on }"),

		web.Scope(VList(
			h.Tag("vx-draggable").Attr("v-model", "locals.sortedColumns", "draggable", ".vx_column_item", "animation", "300").Children(
				h.Div(
					VListItem(
						VListItemContent(
							VListItemTitle(
								VSwitch().Dense(true).Attr("v-model", "locals.displayColumns", ":value", "column.name", ":label", "column.label", "@click", "event.preventDefault()"),
							),
						),
						VListItemIcon(
							VIcon("reorder"),
						).Attr("style", "margin-top: 28px"),
					),
					VDivider(),
				).Attr("v-for", "(column, index) in locals.sortedColumns", ":key", "column.name", "class", "vx_column_item"),
			),
			VListItem(
				VListItemAction(VBtn(msgr.Cancel).Elevation(0).Attr("@click", web.Plaid().Reload().Go())),
				VListItemAction(VBtn(msgr.OK).Elevation(0).Color("primary").Attr("@click", web.Plaid().Query(displayColumnsName, web.Var("locals.displayColumns")).Query(sortedColumnsName, web.Var("locals.sortedColumns.map(column => column.name )")).MergeQuery(true).Go()))),
		).Dense(true)).
			Init(h.JSONString(selectColumns)).
			VSlot("{ locals }"),
	).OffsetY(true).CloseOnClick(false).CloseOnContentClick(false)
	return
}

func (b *ListingBuilder) tableToolbar(msgr *Messages, pageURL *url.URL, ctx *web.EventContext, fd vuetifyx.FilterData) (toolbar *VToolbarBuilder, displayFields []*FieldBuilder) {
	ft := vuetifyx.FilterTranslations{}
	ft.Filters = msgr.Filters
	ft.Filter = msgr.Filter
	ft.Done = msgr.FiltersDone
	ft.Clear = msgr.FiltersClear

	ft.Date.To = msgr.FiltersDateTo

	ft.Number.And = msgr.FiltersNumberAnd
	ft.Number.Equals = msgr.FiltersNumberEquals
	ft.Number.Between = msgr.FiltersNumberBetween
	ft.Number.GreaterThan = msgr.FiltersNumberGreaterThan
	ft.Number.LessThan = msgr.FiltersNumberLessThan

	ft.String.Equals = msgr.FiltersStringEquals
	ft.String.Contains = msgr.FiltersStringContains

	ft.MultipleSelect.In = msgr.FiltersMultipleSelectIn
	ft.MultipleSelect.NotIn = msgr.FiltersMultipleSelectNotIn

	toolbar = VToolbar(
		VSpacer(),
	).Flat(true)

	toolbar.AppendChildren(b.actionsComponent(msgr, ctx))

	if b.newBtnFunc != nil {
		toolbar.AppendChildren(b.newBtnFunc(ctx))
	} else {
		disableNewBtn := b.mb.Info().Verifier().Do(PermCreate).WithReq(ctx.R).IsAllowed() != nil
		if !disableNewBtn {
			toolbar.AppendChildren(VBtn(msgr.New).
				Color("primary").
				Depressed(true).
				Dark(true).Class("ml-2").
				Disabled(disableNewBtn).
				Attr("@click", web.Plaid().EventFunc(actions.New).
					Go()))
		}
	}

	displayFields = b.fields
	if b.selectableColumns {
		var btn h.HTMLComponent
		btn, displayFields = b.selectColumnsBtn(pageURL, ctx)
		toolbar.PrependChildren(btn)
	}
	if fd != nil {
		toolbar.PrependChildren(
			vuetifyx.VXFilter(fd).Translations(ft),
			h.If(b.selectableColumns, VDivider().Vertical(true).Inset(true).Light(true).Attr("style", "margin-left:8px;")),
		)
	}

	return
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
	var displayFields = b.fields
	toolbar, displayFields = b.tableToolbar(msgr, pageURL, ctx, fd)

	pagesCount := int(int64(totalCount)/searchParams.PerPage + 1)
	if int64(totalCount)%searchParams.PerPage == 0 {
		pagesCount--
	}

	dataTable = s.DataTable(objs).
		CellWrapperFunc(func(cell h.MutableAttrHTMLComponent, id string, obj interface{}, dataTableID string) h.HTMLComponent {
			tdbind := cell
			if b.mb.hasDetailing && !b.mb.detailing.drawer {
				tdbind.SetAttr("@click.self", web.Plaid().
					PushStateURL(
						b.mb.Info().
							DetailingHref(id)).
					Go())
			} else {
				event := actions.Edit
				if b.mb.hasDetailing {
					event = actions.DetailingDrawer
				}
				tdbind.SetAttr("@click.self",
					web.Plaid().
						EventFunc(event).
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
		if b.mb.Info().Verifier().Do(PermList).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}
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

	// Render bulk actions
	for _, ba := range b.bulkActions {
		if b.mb.Info().Verifier().SnakeDo(PermBulkActions, ba.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}

		var btn h.HTMLComponent
		if ba.buttonCompFunc != nil {
			btn = ba.buttonCompFunc(ctx)
		} else {
			buttonColor := ba.buttonColor
			if buttonColor == "" {
				buttonColor = ColorSecondary
			}
			btn = VBtn(b.mb.getLabel(ba.NameLabel)).
				Color(buttonColor).
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

	// Render actions
	for _, ba := range b.actions {
		if b.mb.Info().Verifier().SnakeDo(PermActions, ba.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}

		var btn h.HTMLComponent
		if ba.buttonCompFunc != nil {
			btn = ba.buttonCompFunc(ctx)
		} else {
			buttonColor := ba.buttonColor
			if buttonColor == "" {
				buttonColor = ColorPrimary
			}

			btn = VBtn(b.mb.getLabel(ba.NameLabel)).
				Color(buttonColor).
				Depressed(true).
				Dark(true).
				Class("ml-2").
				Attr("@click", web.Plaid().EventFunc(actions.OpenActionDialog).
					Queries(url.Values{actionPanelOpenParamName: []string{ba.name}}).
					MergeQuery(true).
					Go())
		}

		actionBtns = append(actionBtns, btn)
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
