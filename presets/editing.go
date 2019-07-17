package presets

import (
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
	"goji.io/pat"
)

type EditingBuilder struct {
	mb          *ModelBuilder
	fields      []*FieldBuilder
	bulkActions []*BulkActionBuilder
	filters     []string
	pageFunc    ui.PageFunc
	fetcher     Fetcher
	saver       Saver
}

func (b *ModelBuilder) Editing(vs ...string) (r *EditingBuilder) {
	r = b.editing
	var newfields []*FieldBuilder
	for _, f := range vs {
		newfields = append(newfields, r.Field(f))
	}
	r.fields = newfields
	return r
}

func (b *EditingBuilder) Field(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	r = &FieldBuilder{name: name}
	b.fields = append(b.fields, r)
	return
}

func (b *EditingBuilder) PageFunc(pf ui.PageFunc) (r *EditingBuilder) {
	b.pageFunc = pf
	return b
}

func (b *EditingBuilder) Fetcher(v Fetcher) (r *EditingBuilder) {
	b.fetcher = v
	return b
}

func (b *EditingBuilder) Saver(v Saver) (r *EditingBuilder) {
	b.saver = v
	return b
}

func (b *EditingBuilder) GetPageFunc() ui.PageFunc {
	if b.pageFunc != nil {
		return b.pageFunc
	}
	return b.defaultPageFunc
}

func (b *EditingBuilder) defaultPageFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	id := pat.Param(ctx.R, "id")
	var obj interface{}

	obj, err = b.fetcher(b.mb.newModel(), id)
	if err != nil {
		return
	}

	var comps []h.HTMLComponent
	for _, f := range b.fields {
		if f.compFunc == nil {
			continue
		}
		comps = append(comps, f.compFunc(obj, &Field{Name: f.name, Label: b.mb.getLabel(f)}, ctx))
	}
	r.Schema = VContainer(comps...)
	return
}
