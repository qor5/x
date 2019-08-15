package presets

import (
	"fmt"
	"strconv"

	"github.com/sunfmin/reflectutils"

	"github.com/qor/inflection"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
	"github.com/thoas/go-funk"
)

type ListingBuilder struct {
	mb            *ModelBuilder
	fields        []*FieldBuilder
	bulkActions   []*BulkActionBuilder
	filterData    FilterData
	pageFunc      ui.PageFunc
	searcher      SearchOpFunc
	searchColumns []string
	perPage       int64
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
	r = &FieldBuilder{name: name}
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

func (b *ListingBuilder) defaultPageFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("formDrawerNew", b.mb.editing.formDrawerNew)
	ctx.Hub.RegisterEventFunc("formDrawerEdit", b.mb.editing.formDrawerEdit)
	ctx.Hub.RegisterEventFunc("update", b.mb.editing.defaultUpdate)
	msgr := b.mb.p.messagesFunc(ctx)

	perPage := b.perPage
	if perPage == 0 {
		perPage = 50
	}

	//time.Sleep(1 * time.Second)
	searchParams := &SearchParams{
		KeywordColumns: b.searchColumns,
		Keyword:        ctx.R.URL.Query().Get("keyword"),
		PerPage:        perPage,
	}

	searchParams.Page, _ = strconv.ParseInt(ctx.R.URL.Query().Get("page"), 10, 64)
	if searchParams.Page == 0 {
		searchParams.Page = 1
	}

	var toolbar = VToolbar(
		VSpacer(),
		VBtn(msgr.New).
			Color(b.mb.p.primaryColor).
			Depressed(true).
			Dark(true).
			OnClick("formDrawerNew", ""),
	).Flat(true)

	if b.filterData != nil {
		fd := b.filterData.Clone()

		cond, args := fd.SetByQueryString(ctx.R.URL.RawQuery)

		searchParams.SQLConditions = append(searchParams.SQLConditions, &SQLCondition{
			Query: cond,
			Args:  args,
		})

		toolbar.PrependChildren(Filter(fd))
	}

	var objs interface{}
	var totalCount int
	objs, totalCount, err = b.searcher(b.mb.newModelArray(), searchParams)
	if err != nil {
		return
	}

	pagesCount := int(int64(totalCount)/searchParams.PerPage + 1)
	if int64(totalCount)%searchParams.PerPage == 0 {
		pagesCount--
	}

	var rows []h.HTMLComponent

	haveCheckboxes := len(b.bulkActions) > 0

	funk.ForEach(objs, func(obj interface{}) {
		var tds []h.HTMLComponent
		if haveCheckboxes {
			tds = append(tds, h.Td(VCheckbox().Class("mt-0").FieldName("selected").HideDetails(true)).Class("pr-0"))
		}
		for _, f := range b.fields {
			tds = append(tds, f.compFunc(obj, b.mb.getComponentFuncField(f), ctx))
		}
		if err != nil {
			panic(err)
		}
		id := fmt.Sprint(reflectutils.MustGet(obj, "ID"))
		var bindTds []h.HTMLComponent
		for _, td := range tds {
			std, ok := td.(h.MutableAttrHTMLComponent)
			if !ok {
				bindTds = append(bindTds, td)
				continue
			}

			tdbind := ui.Bind(std)
			if b.mb.hasDetailing {
				tdbind.On("click.self").PushStateLink(b.mb.Info().DetailingHref(id))
			} else {
				tdbind.On("click.self").EventFunc("formDrawerEdit", id)
			}

			bindTds = append(bindTds, tdbind)
		}

		rows = append(rows, h.Tr(bindTds...))
	})

	var heads []h.HTMLComponent

	if haveCheckboxes {
		heads = append(heads, h.Th("").Children(VCheckbox().Class("mt-0").HideDetails(true)).Style("width: 48px;").Class("pr-0"))
	}
	for _, f := range b.fields {
		label := b.mb.getLabel(f)
		heads = append(heads, h.Th(label))
	}

	r.Schema = VContainer(

		h.H2(msgr.ListingObjectTitle(inflection.Plural(b.mb.label))).Class("title pb-3"),

		VCard(
			toolbar,
			VDivider(),
			VCardText(
				VSimpleTable(
					h.Thead(
						h.Tr(heads...),
					),
					h.Tbody(
						rows...,
					),
				),
			).Class("pa-0"),

			//VCardText(
			//	VBtn("").
			//		Color(b.mb.p.primaryColor).
			//		Fab(true).
			//		Bottom(true).
			//		Right(true).
			//		Dark(true).
			//		Absolute(true).
			//		Children(
			//			VIcon("add"),
			//		).OnClick("formDrawerNew", ""),
			//).Attr("style", "position: relative"),
		),

		h.If(pagesCount > 1, h.Components(
			VPagination().Length(pagesCount).Value(int(searchParams.Page)),
		)),
	).Fluid(true)

	return
}
