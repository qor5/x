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
	mb       *ModelBuilder
	pageFunc ui.PageFunc
	fetcher  FetchOpFunc
	setter   SetterFunc
	saver    SaveOpFunc
	deleter  DeleteOpFunc
	FieldBuilders
}

func (b *ModelBuilder) Editing(vs ...string) (r *EditingBuilder) {
	r = b.editing
	if len(vs) == 0 {
		return
	}

	var newFields []*FieldBuilder
	for _, f := range vs {
		field := b.writeFields.GetField(f)
		if field == nil {
			newFields = append(newFields, NewField(f))
		} else {
			newFields = append(newFields, field.Clone())
		}
	}

	r.fields = newFields
	return r
}

func (b *EditingBuilder) CloneForCreating(vs ...string) (r *EditingBuilder) {

	if b.mb.creating == nil {
		b.mb.creating = &EditingBuilder{
			mb:       b.mb,
			pageFunc: b.pageFunc,
			fetcher:  b.fetcher,
			setter:   b.setter,
			saver:    b.saver,
			deleter:  b.deleter,
		}
	}
	r = b.mb.creating

	var editingFields []string
	for _, f := range b.fields {
		editingFields = append(editingFields, f.name)
	}

	if len(vs) == 0 {
		vs = editingFields
	}

	var newFields []*FieldBuilder

	for _, name := range vs {
		field := b.GetField(name)

		if field == nil {
			field = b.mb.writeFields.GetField(name)
		}

		if field == nil {
			newFields = append(newFields, NewField(name))
		} else {
			newFields = append(newFields, field.Clone())
		}
	}

	r.fields = newFields
	return r
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
	msgs := MustGetMessages(ctx.R)
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
	msgs := MustGetMessages(ctx.R)
	creatingB := b
	if b.mb.creating != nil {
		creatingB = b.mb.creating
	}

	er, err = creatingB.editFormFor(msgs.CreatingObjectTitle(inflection.Singular(b.mb.label)), msgs.Create)(ctx)
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
	msgs := MustGetMessages(ctx.R)
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

		fields := b.getFieldsWithDefaults()

		for _, f := range fields {
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

func (b *EditingBuilder) getFieldsWithDefaults() []*FieldBuilder {
	fields := b.fields
	if len(fields) == 0 {
		fields = b.mb.writeFields.fields
	}
	return fields
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

	usingB := b
	if b.mb.creating != nil && len(id) == 0 {
		usingB = b.mb.creating
	}

	if len(id) > 0 {
		obj, err = usingB.fetcher(obj, id)
		if err != nil {
			return
		}
	}

	fields := usingB.getFieldsWithDefaults()

	for _, f := range fields {
		err = reflectutils.Set(obj, f.name, reflectutils.MustGet(newObj, f.name))
		if err != nil {
			panic(err)
		}
	}

	if usingB.setter != nil {
		usingB.setter(obj, ctx.R.MultipartForm, ctx)
	}

	err = usingB.saver(obj, id)
	if err != nil {
		panic(err)
	}

	msgr := MustGetMessages(ctx.R)
	ctx.Flash = msgr.SuccessfullyUpdated

	r.PushState = ui.PushState(nil)
	return
}
