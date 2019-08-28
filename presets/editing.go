package presets

import (
	"fmt"

	"github.com/jinzhu/inflection"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
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

	r.FieldBuilders = *r.FieldBuilders.Only(vs...)
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

	r.FieldBuilders = *b.mb.writeFields.Only(vs...)

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
	msgr := MustGetMessages(ctx.R)
	title := msgr.EditingObjectTitle(inflection.Singular(b.mb.label))
	r.PageTitle = fmt.Sprintf("%s - %s", title, b.mb.p.brandTitle)

	id := pat.Param(ctx.R, "id")
	ctx.Event = &ui.Event{
		Params: []string{id},
	}

	r.Schema = b.editFormFor(title, msgr.Update, ctx)
	return
}

func (b *EditingBuilder) formDrawerNew(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	msgr := MustGetMessages(ctx.R)
	creatingB := b
	if b.mb.creating != nil {
		creatingB = b.mb.creating
	}

	form := creatingB.editFormFor(msgr.CreatingObjectTitle(inflection.Singular(b.mb.label)), msgr.Create, ctx)

	r.UpdatePortals = append(r.UpdatePortals, &ui.PortalUpdate{
		Name: "rightDrawer",
		Schema: VNavigationDrawer(
			form,
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
	msgr := MustGetMessages(ctx.R)
	form := b.editFormFor(msgr.EditingObjectTitle(inflection.Singular(b.mb.label)), msgr.Update, ctx)

	r.UpdatePortals = append(r.UpdatePortals, &ui.PortalUpdate{
		Name: "rightDrawer",
		Schema: VNavigationDrawer(
			form,
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

func (b *EditingBuilder) editFormFor(title, buttonLabel string, ctx *ui.EventContext) h.HTMLComponent {
	id := ctx.Event.Params[0]
	ctx.Hub.RegisterEventFunc("update", b.defaultUpdate)
	var obj = b.mb.newModel()

	if len(id) > 0 {
		var err error
		obj, err = b.fetcher(obj, id)
		if err != nil {
			panic(err)
		}
	}

	var notice h.HTMLComponent
	if msg, ok := ctx.Flash.(string); ok {
		notice = VSnackbar(h.Text(msg)).Value(true).Top(true).Color("success").Value(true)
	}

	vErr, _ := ctx.Flash.(*ValidationErrors)

	return VContainer(
		notice,
		h.H2(title).Class("title px-4 py-2"),
		VCard(
			VCardText(
				b.ToComponent(obj, vErr, ctx),
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

	usingB.MustSet(obj, newObj)

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
