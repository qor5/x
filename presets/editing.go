package presets

import (
	"github.com/jinzhu/inflection"
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
	fetcher     FetchOpFunc
	saver       SaveOpFunc
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

func (b *EditingBuilder) Fetcher(v FetchOpFunc) (r *EditingBuilder) {
	b.fetcher = v
	return b
}

func (b *EditingBuilder) Saver(v SaveOpFunc) (r *EditingBuilder) {
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
	ctx.Event = &ui.Event{
		Params: []string{id},
	}
	msgs := b.mb.p.messagesFunc(ctx)

	var er ui.EventResponse
	er, err = b.editFormFor(msgs.EditingObjectTitle(inflection.Singular(b.mb.label)), msgs.Update)(ctx)
	if err != nil {
		return
	}
	r.Schema = er.Schema
	return
}

func (b *EditingBuilder) formDrawerNew(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	var er ui.EventResponse
	msgs := b.mb.p.messagesFunc(ctx)
	er, err = b.editFormFor(msgs.CreatingObjectTitle(inflection.Singular(b.mb.label)), msgs.Create)(ctx)
	if err != nil {
		return
	}

	r.UpdatePortals = append(r.UpdatePortals, &ui.PortalUpdate{
		Name: "rightDrawer",
		Schema: VNavigationDrawer(
			er.Schema.(h.HTMLComponent),
		).Attr("v-model", "boolean1").
			Bottom(true).
			Right(true).
			Absolute(true).
			Width(600).
			Temporary(true),
		AfterLoaded: `setTimeout(function(){ comp.boolean1 = true }, 100)`,
	})
	return
}

func (b *EditingBuilder) formDrawerEdit(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	var er ui.EventResponse
	msgs := b.mb.p.messagesFunc(ctx)
	er, err = b.editFormFor(msgs.EditingObjectTitle(inflection.Singular(b.mb.label)), msgs.Update)(ctx)
	if err != nil {
		return
	}

	r.UpdatePortals = append(r.UpdatePortals, &ui.PortalUpdate{
		Name: "rightDrawer",
		Schema: VNavigationDrawer(
			er.Schema.(h.HTMLComponent),
		).Attr("v-model", "boolean1").
			Bottom(true).
			Right(true).
			Absolute(true).
			Width(600).
			Temporary(true),
		AfterLoaded: `setTimeout(function(){ comp.boolean1 = true }, 100)`,
	})
	return
}

func (b *EditingBuilder) editFormFor(title, buttonLabel string) ui.EventFunc {
	return func(ctx *ui.EventContext) (r ui.EventResponse, err error) {
		id := ctx.Event.Params[0]
		ctx.Hub.RegisterEventFunc("update", b.defaultUpdate)
		var obj = b.mb.newModel()

		if len(id) > 0 {
			obj, err = b.fetcher(obj, id)
			if err != nil {
				return
			}
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
			comps = append(comps, f.compFunc(obj, &Field{
				Name:  f.name,
				Label: b.mb.getLabel(f),
			}, ctx))
		}

		r.Schema = VContainer(
			notice,
			h.H2(title).Class("title pb-3"),
			VCard(
				VCardText(
					comps...,
				),
				VCardActions(
					VSpacer(),
					VBtn(buttonLabel).
						Dark(true).
						Color(b.mb.p.primaryColor).
						OnClick("update", id),
				),
			).Flat(true),
		).Fluid(true)

		return
	}
}

func (b *EditingBuilder) defaultUpdate(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	id := ctx.Event.Params[0]
	var newObj = b.mb.newModel()
	ctx.MustUnmarshalForm(newObj)

	var obj = b.mb.newModel()
	if len(id) > 0 {
		obj, err = b.fetcher(obj, id)
		if err != nil {
			return
		}
	}

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
	msgs := b.mb.p.messagesFunc(ctx)
	ctx.Flash = msgs.SuccessfullyUpdated

	r.Reload = true
	return
}