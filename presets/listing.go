package presets

import (
	"fmt"

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
	msgs := b.mb.p.messagesFunc(ctx)
	//time.Sleep(1 * time.Second)
	searchParams := &SearchParams{
		KeywordColumns: b.searchColumns,
		Keyword:        ctx.R.URL.Query().Get("keyword"),
	}

	var toolbar = VToolbar(
		VSpacer(),
		VBtn(msgs.New).
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
	objs, err = b.searcher(b.mb.newModelArray(), searchParams)
	if err != nil {
		return
	}

	var rows []h.HTMLComponent

	funk.ForEach(objs, func(obj interface{}) {
		var tds []h.HTMLComponent
		for _, f := range b.fields {
			tds = append(tds, f.compFunc(obj, b.mb.getComponentFuncField(f), ctx))
		}
		v, err := reflectutils.Get(obj, "ID")
		if err != nil {
			panic(err)
		}
		rows = append(rows, ui.Bind(h.Tr(tds...)).PushStateLink(
			b.mb.detailingHref(fmt.Sprint(v)),
		))
	})

	var heads []h.HTMLComponent

	for _, f := range b.fields {
		label := b.mb.getLabel(f)
		heads = append(heads, h.Th(label))
	}

	r.Schema = VContainer(

		h.H2(msgs.ListingObjectTitle(inflection.Plural(b.mb.label))).Class("title pb-3"),

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
	).Fluid(true)

	return
}
