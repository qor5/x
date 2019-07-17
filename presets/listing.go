package presets

import (
	"fmt"

	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
)

type ListingBuilder struct {
	mb            *ModelBuilder
	fields        []*FieldBuilder
	bulkActions   []*BulkActionBuilder
	filtering     *FilteringBuilder
	pageFunc      ui.PageFunc
	searcher      Searcher
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

func (b *ListingBuilder) Searcher(v Searcher) (r *ListingBuilder) {
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
	var objs interface{}
	objs, err = b.searcher(b.mb.newModelArray(), &SearchParams{
		KeywordColumns: b.searchColumns,
		Keyword:        ctx.R.URL.Query().Get("keyword"),
	})
	if err != nil {
		return
	}
	r.Schema = VBtn(fmt.Sprint(objs))
	return
}
