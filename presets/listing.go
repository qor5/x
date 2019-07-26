package presets

import (
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
	filtering     *FilteringBuilder
	pageFunc      ui.PageFunc
	searcher      SearchOpFunc
	searchColumns []string
}

func (b *ModelBuilder) Listing(vs ...string) (r *ListingBuilder) {
	r = b.listing
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
	ctx.Hub.RegisterEventFunc("formNew", b.mb.editing.editForm)

	var objs interface{}
	objs, err = b.searcher(b.mb.newModelArray(), &SearchParams{
		KeywordColumns: b.searchColumns,
		Keyword:        ctx.R.URL.Query().Get("keyword"),
	})
	if err != nil {
		return
	}

	var rows []h.HTMLComponent

	funk.ForEach(objs, func(obj interface{}) {
		var tds []h.HTMLComponent
		for _, f := range b.fields {
			tds = append(tds, f.compFunc(obj, b.mb.getComponentFuncField(f), ctx))
		}
		rows = append(rows, h.Tr(tds...))
	})

	var heads []h.HTMLComponent

	for _, f := range b.fields {
		label := b.mb.getLabel(f)
		heads = append(heads, h.Th(label).Role("columnheader").
			Class("column text-xs-left").
			Attr(
				"scope", "col",
				"aria-label", label,
			),
		)
	}
	msgs := b.mb.p.messagesFunc(ctx)

	r.Schema = VContainer(

		VLayout(
			h.H2(msgs.ListingObjectTitle(inflection.Plural(b.mb.label))).Class("title"),
			VSpacer(),
			//overlay.Drawer(
			//	ui.LazyPortal("formNew", "").LoadWhenParentVisible(),
			//).TriggerElement(
			VBtn(msgs.New).Color("primary mr-0"),
			//).Width(500).GetContainer("#vt-app"),
		).Class("align-center"),

		VCard(
			VCardText(
				h.Div(
					h.Table(
						h.Thead(
							h.Tr(heads...),
						),
						h.Tbody(
							rows...,
						),
					).Class("v-datatable v-table theme--light"),
				).Class("v-table__overflow"),
			).Class("pa-0"),
		),
	).Fluid(true)

	return
}
