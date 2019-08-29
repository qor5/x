package presets

import (
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
	"goji.io/pat"
)

type DetailingBuilder struct {
	mb         *ModelBuilder
	fieldNames []string
	actions    []*ActionBuilder
	pageFunc   ui.PageFunc
	fetcher    FetchFunc
	FieldBuilders
}

func (b *ModelBuilder) Detailing(vs ...string) (r *DetailingBuilder) {
	r = b.detailing
	b.hasDetailing = true
	if len(vs) == 0 {
		return
	}

	r.fieldNames = vs
	var newfields []*FieldBuilder
	for _, f := range vs {
		newfields = append(newfields, r.Field(f))
	}
	r.fields = newfields
	return r
}

func (b *DetailingBuilder) PageFunc(pf ui.PageFunc) (r *DetailingBuilder) {
	b.pageFunc = pf
	return b
}

func (b *DetailingBuilder) Fetcher(v FetchFunc) (r *DetailingBuilder) {
	b.fetcher = v
	return b
}

func (b *DetailingBuilder) GetPageFunc() ui.PageFunc {
	if b.pageFunc != nil {
		return b.pageFunc
	}
	return b.defaultPageFunc
}

func (b *DetailingBuilder) defaultPageFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	id := pat.Param(ctx.R, "id")
	r.Schema = VContainer(h.Text(id))

	var obj = b.mb.newModel()

	if len(id) == 0 {
		panic("not found")
	}

	obj, err = b.fetcher(obj, id, ctx)
	if err != nil {
		return
	}

	var notice h.HTMLComponent
	if msg, ok := ctx.Flash.(string); ok {
		notice = VSnackbar(h.Text(msg)).Value(true).Top(true).Color("success").Value(true)
	}

	var comps []h.HTMLComponent
	for _, f := range b.fields {
		if f.compFunc == nil {
			continue
		}
		comps = append(comps, f.compFunc(obj, &FieldContext{
			Name:  f.name,
			Label: b.mb.getLabel(f.NameLabel),
		}, ctx))
	}

	r.Schema = VContainer(
		notice,
		ui.LazyPortal().Name(deleteConfirmPortalName),
	).AppendChildren(comps...).Fluid(true)
	return
}
