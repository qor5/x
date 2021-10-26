package presets

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/i18n"
	"github.com/goplaid/x/perm"
	"github.com/goplaid/x/presets/actions"
	. "github.com/goplaid/x/vuetify"
	"github.com/jinzhu/inflection"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type EditingBuilder struct {
	mb        *ModelBuilder
	fetcher   FetchFunc
	setter    SetterFunc
	saver     SaveFunc
	deleter   DeleteFunc
	validator ValidateFunc
	sidePanel ComponentFunc
	FieldBuilders
}

func (b *ModelBuilder) Editing(vs ...string) (r *EditingBuilder) {
	r = b.editing
	if len(vs) == 0 {
		return
	}

	r.Only(vs...)
	return r
}

func (b *EditingBuilder) Only(vs ...string) (r *EditingBuilder) {
	r = b
	r.FieldBuilders = *r.FieldBuilders.Only(vs...)
	return
}

func (b *EditingBuilder) Creating(vs ...string) (r *EditingBuilder) {

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

func (b *EditingBuilder) SidePanelFunc(v ComponentFunc) (r *EditingBuilder) {
	b.sidePanel = v
	return b
}

func (b *EditingBuilder) formDrawerNew(ctx *web.EventContext) (r web.EventResponse, err error) {
	creatingB := b
	if b.mb.creating != nil {
		creatingB = b.mb.creating
	}

	b.mb.p.rightDrawer(&r, creatingB.editFormFor(nil, ctx), b.mb.rightDrawerWidth)
	return
}

func (b *EditingBuilder) formDrawerEdit(ctx *web.EventContext) (r web.EventResponse, err error) {
	b.mb.p.rightDrawer(&r, b.editFormFor(nil, ctx), b.mb.rightDrawerWidth)
	return
}

func (b *EditingBuilder) editFormFor(obj interface{}, ctx *web.EventContext) h.HTMLComponent {
	msgr := MustGetMessages(ctx.R)

	id := ctx.Event.Params[0]

	var buttonLabel = msgr.Create
	var disableUpdateBtn bool
	var title = msgr.CreatingObjectTitle(
		i18n.T(ctx.R, ModelsI18nModuleKey, inflection.Singular(b.mb.label)),
	)
	if len(id) > 0 {
		if obj == nil {
			var err error
			obj, err = b.fetcher(b.mb.newModel(), id, ctx)
			if err != nil {
				panic(err)
			}
		}
		disableUpdateBtn = b.mb.Info().Verifier().Do(PermUpdate).ObjectOn(obj).WithReq(ctx.R).IsAllowed() != nil
		buttonLabel = msgr.Update
		title = msgr.EditingObjectTitle(
			i18n.T(ctx.R, ModelsI18nModuleKey, inflection.Singular(b.mb.label)),
			getPageTitle(obj, id))
	}

	if obj == nil {
		obj = b.mb.newModel()
	}

	var notice h.HTMLComponent
	if msg, ok := ctx.Flash.(string); ok {
		notice = VSnackbar(h.Text(msg)).Value(true).Top(true).Color("success").Value(true)
	}

	vErr, _ := ctx.Flash.(*web.ValidationErrors)

	var panel h.HTMLComponent

	var formContent h.HTMLComponent = h.Components(
		VCardText(
			b.ToComponent(b.mb, obj, vErr, ctx),
		),
		VCardActions(
			VSpacer(),
			VBtn(buttonLabel).
				Dark(true).
				Color("primary").
				Disabled(disableUpdateBtn).
				Attr("@click", web.Plaid().
					EventFunc(actions.Update, ctx.Event.Params...).
					URL(b.mb.Info().ListingHref()).
					Go()),
		),
	)

	if b.sidePanel != nil {
		panel = b.sidePanel(ctx)
		formContent = VContainer(
			VRow(
				VCol(
					formContent,
				).Cols(8),
				VCol(
					panel,
				).Cols(4),
			),
		)
	}

	return h.Components(
		VAppBar(
			VToolbarTitle(title).Class("pl-2"),
			VSpacer(),
			VBtn("").Icon(true).Children(
				VIcon("close"),
			).Attr("@click.stop", "vars.rightDrawer = false"),
		).Color("white").Elevation(0).Dense(true),

		VContainer(
			formContent,
			notice,
			VCard().Flat(true),
		).Fluid(true),
	)
}

func (b *EditingBuilder) doDelete(ctx *web.EventContext) (r web.EventResponse, err error) {
	id := ctx.Event.Params[0]
	var obj = b.mb.newModel()
	if len(id) > 0 {
		err = b.deleter(obj, id, ctx)
		if err != nil {
			return
		}
	}

	r.PushState = web.PushState(nil)
	return
}

func (b *EditingBuilder) defaultUpdate(ctx *web.EventContext) (r web.EventResponse, err error) {
	id := ctx.Event.Params[0]
	var newObj = b.mb.newModel()
	// don't panic for fields that set in SetterFunc
	_ = ctx.UnmarshalForm(newObj)

	if len(id) == 0 {
		if b.mb.Info().Verifier().Do(PermCreate).ObjectOn(newObj).WithReq(ctx.R).IsAllowed() != nil {
			b.renderFormWithError(&r, perm.PermissionDenied, newObj, ctx)
			return
		}
	}

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
		if b.mb.Info().Verifier().Do(PermUpdate).ObjectOn(obj).WithReq(ctx.R).IsAllowed() != nil {
			b.renderFormWithError(&r, perm.PermissionDenied, obj, ctx)
			return
		}
	}

	if usingB.setter != nil {
		usingB.setter(obj, ctx)
	}

	var vErr web.ValidationErrors
	for _, f := range usingB.fields {

		if b.mb.Info().Verifier().Do(PermUpdate).ObjectOn(obj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}

		if f.setterFunc == nil {
			_ = reflectutils.Set(obj, f.name, reflectutils.MustGet(newObj, f.name))
			continue
		}

		err1 := f.setterFunc(obj, &FieldContext{
			ModelInfo: b.mb.Info(),
			Name:      f.name,
			Label:     b.getLabel(f.NameLabel),
		}, ctx)
		if err1 != nil {
			vErr.FieldError(f.name, err1.Error())
		}
	}

	if vErr.HaveErrors() {
		b.renderFormWithError(&r, &vErr, obj, ctx)
		return
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

	r.PushState = web.PushState(nil)
	r.VarsScript = `vars.rightDrawer = false`
	return
}

func (b *EditingBuilder) renderFormWithError(r *web.EventResponse, err error, obj interface{}, ctx *web.EventContext) {
	ctx.Flash = err

	if _, ok := err.(*web.ValidationErrors); !ok {
		vErr := &web.ValidationErrors{}
		vErr.GlobalError(err.Error())
		ctx.Flash = vErr
	}

	r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
		Name: rightDrawerContentPortalName,
		Body: b.editFormFor(obj, ctx),
	})

}
