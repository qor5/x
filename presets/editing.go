package presets

import (
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	"github.com/sunfmin/reflectutils"
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

//type updateState struct {
//	Obj       interface{}
//	OKMessage string
//}

func (b *EditingBuilder) defaultPageFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	id := pat.Param(ctx.R, "id")
	var obj interface{}

	obj, err = b.fetcher(b.mb.newModel(), id)
	if err != nil {
		return
	}

	ctx.StateOrInit(obj)

	var comps []h.HTMLComponent
	//if len(state.OKMessage) > 0 {
	//	comps = append(comps, VAlert(h.Text(state.OKMessage)).Type("success"))
	//}
	for _, f := range b.fields {
		if f.compFunc == nil {
			continue
		}
		comps = append(comps, f.compFunc(obj, &Field{Name: f.name, Label: b.mb.getLabel(f)}, ctx))
	}
	comps = append(comps, VBtn("Update").Color("primary").OnClick(ctx.Hub, "update", b.defaultUpdate))
	r.Schema = VContainer(comps...)
	return
}

func (b *EditingBuilder) defaultUpdate(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	id := pat.Param(ctx.R, "id")
	var obj interface{}

	obj, err = b.fetcher(b.mb.newModel(), id)
	if err != nil {
		return
	}

	newObj := ctx.State

	for _, f := range b.fields {
		err = reflectutils.Set(obj, f.name, reflectutils.MustGet(newObj, f.name))
		if err != nil {
			panic(err)
		}
	}

	err = b.saver(obj, id)
	if err != nil {
		panic(err)
	}

	r.Reload = true
	return
}
