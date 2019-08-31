package presets

import (
	"github.com/jinzhu/inflection"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

type EditingBuilder struct {
	mb        *ModelBuilder
	fetcher   FetchFunc
	setter    SetterFunc
	saver     SaveFunc
	deleter   DeleteFunc
	validator ValidateFunc
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
			mb:        b.mb,
			fetcher:   b.fetcher,
			setter:    b.setter,
			saver:     b.saver,
			deleter:   b.deleter,
			validator: b.validator,
		}
	}
	r = b.mb.creating

	r.FieldBuilders = *b.mb.writeFields.Only(vs...)

	return r
}

func (b *EditingBuilder) FetchFunc(v FetchFunc) (r *EditingBuilder) {
	b.fetcher = v
	return b
}

func (b *EditingBuilder) SaveFunc(v SaveFunc) (r *EditingBuilder) {
	b.saver = v
	return b
}

func (b *EditingBuilder) DeleteFunc(v DeleteFunc) (r *EditingBuilder) {
	b.deleter = v
	return b
}

func (b *EditingBuilder) ValidateFunc(v ValidateFunc) (r *EditingBuilder) {
	b.validator = v
	return b
}

func (b *EditingBuilder) SetterFunc(v SetterFunc) (r *EditingBuilder) {
	b.setter = v
	return b
}

const formDrawerNew = "formDrawerNew"

func (b *EditingBuilder) formDrawerNew(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	creatingB := b
	if b.mb.creating != nil {
		creatingB = b.mb.creating
	}

	rightDrawer(&r, creatingB.editFormFor(nil, ctx))
	return
}

const formDrawerEdit = "formDrawerEdit"

func (b *EditingBuilder) formDrawerEdit(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	rightDrawer(&r, b.editFormFor(nil, ctx))
	return
}

func (b *EditingBuilder) editFormFor(obj interface{}, ctx *ui.EventContext) h.HTMLComponent {
	msgr := MustGetMessages(ctx.R)

	id := ctx.Event.Params[0]
	ctx.Hub.RegisterEventFunc("update", b.defaultUpdate)

	var buttonLabel = msgr.Create
	var title = msgr.CreatingObjectTitle(inflection.Singular(b.mb.label))
	if len(id) > 0 {
		if obj == nil {
			var err error
			obj, err = b.fetcher(b.mb.newModel(), id, ctx)
			if err != nil {
				panic(err)
			}
		}
		buttonLabel = msgr.Update
		title = msgr.EditingObjectTitle(inflection.Singular(b.mb.label), getPageTitle(obj, id))
	}

	if obj == nil {
		obj = b.mb.newModel()
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
		err = b.deleter(obj, id, ctx)
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
		obj, err1 := usingB.fetcher(obj, id, ctx)
		if err1 != nil {
			b.renderFormWithError(&r, err1, obj, ctx)
			return
		}
	}

	usingB.MustSet(obj, newObj)

	if usingB.setter != nil {
		usingB.setter(obj, ctx)
	}

	if usingB.validator != nil {
		if vErr := usingB.validator(obj, ctx); vErr.HaveErrors() {
			b.renderFormWithError(&r, &vErr, obj, ctx)
			return
		}
	}

	err1 := usingB.saver(obj, id, ctx)
	if err1 != nil {
		b.renderFormWithError(&r, err1, obj, ctx)
		return
	}

	msgr := MustGetMessages(ctx.R)
	ctx.Flash = msgr.SuccessfullyUpdated

	r.PushState = ui.PushState(nil)
	return
}

func (b *EditingBuilder) renderFormWithError(r *ui.EventResponse, err error, obj interface{}, ctx *ui.EventContext) {
	ctx.Flash = err

	if _, ok := err.(*ValidationErrors); !ok {
		vErr := &ValidationErrors{}
		ctx.Flash = vErr.GlobalError(err.Error())
	}

	r.UpdatePortals = append(r.UpdatePortals, &ui.PortalUpdate{
		Name:   rightDrawerPortalName,
		Schema: b.editFormFor(obj, ctx),
	})

}
