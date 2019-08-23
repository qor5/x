package presets

import (
	"fmt"

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
	setter      SetterFunc
	saver       SaveOpFunc
	deleter     DeleteOpFunc
}

func (b *ModelBuilder) Editing(vs ...string) (r *EditingBuilder) {
	r = b.editing
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

func (b *EditingBuilder) Field(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	r = NewField(name)
	b.fields = append(b.fields, r)
	return
}

func (b *EditingBuilder) PageFunc(pf ui.PageFunc) (r *EditingBuilder) {
	b.pageFunc = pf
	return b
}

func (b *EditingBuilder) FetchFunc(v FetchOpFunc) (r *EditingBuilder) {
	b.fetcher = v
	return b
}

func (b *EditingBuilder) SaveFunc(v SaveOpFunc) (r *EditingBuilder) {
	b.saver = v
	return b
}

func (b *EditingBuilder) DeleteFunc(v DeleteOpFunc) (r *EditingBuilder) {
	b.deleter = v
	return b
}

func (b *EditingBuilder) SetterFunc(v SetterFunc) (r *EditingBuilder) {
	b.setter = v
	return b
}

func (b *EditingBuilder) GetPageFunc() ui.PageFunc {
	if b.pageFunc != nil {
		return b.pageFunc
	}
	return b.defaultPageFunc
}

func (b *EditingBuilder) defaultPageFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	msgs := b.mb.p.messagesFunc(ctx)
	title := msgs.EditingObjectTitle(inflection.Singular(b.mb.label))
	r.PageTitle = fmt.Sprintf("%s - %s", title, b.mb.p.brandTitle)

	id := pat.Param(ctx.R, "id")
	ctx.Event = &ui.Event{
		Params: []string{id},
	}

	var er ui.EventResponse
	er, err = b.editFormFor(title, msgs.Update)(ctx)
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
		).Attr("v-model", "vars.formDrawerNew").
			Bottom(true).
			Right(true).
			Absolute(true).
			Width(600).
			Temporary(true).
			Attr("v-init-context-vars", `{formDrawerNew: false}`),
		AfterLoaded: `setTimeout(function(){ comp.vars.formDrawerNew = true }, 100)`,
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
		).Attr("v-model", "vars.formDrawerEdit").
			Bottom(true).
			Right(true).
			Absolute(true).
			Width(600).
			Temporary(true).
			Attr("v-init-context-vars", `{formDrawerEdit: false}`),
		AfterLoaded: `setTimeout(function(){ comp.vars.formDrawerEdit = true }, 100)`,
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
			comps = append(comps, f.compFunc(obj, &FieldContext{
				Name:  f.name,
				Label: b.mb.getLabel(f.NameLabel),
			}, ctx))
		}

		r.Schema = VContainer(
			notice,
			h.H2(title).Class("title px-4 py-2"),
			VCard(
				VCardText(
					comps...,
				),
				VCardActions(
					VSpacer(),
					ui.Bind(VBtn(buttonLabel).
						Dark(true).
						Color(b.mb.p.primaryColor)).
						OnClick("update",
							ctx.Event.Params...).
						URL(b.mb.Info().ListingHref()),
				),
			).Flat(true),
		).Fluid(true)

		return
	}
}

func (b *EditingBuilder) doDelete(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	id := ctx.Event.Params[0]
	var obj = b.mb.newModel()
	if len(id) > 0 {
		err = b.deleter(obj, id)
		if err != nil {
			return
		}
	}

	r.PushState = ui.PushState(nil)
	return
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

	if b.setter != nil {
		b.setter(obj, ctx.R.MultipartForm, ctx)
	}

	err = b.saver(obj, id)
	if err != nil {
		panic(err)
	}
	msgs := b.mb.p.messagesFunc(ctx)
	ctx.Flash = msgs.SuccessfullyUpdated

	r.PushState = ui.PushState(nil)
	return
}
