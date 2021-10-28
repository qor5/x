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
	mb          *ModelBuilder
	Fetcher     FetchFunc
	Setter      SetterFunc
	Saver       SaveFunc
	Deleter     DeleteFunc
	Validator   ValidateFunc
	sidePanel   ComponentFunc
	actionsFunc ComponentFunc
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
			Fetcher:   b.Fetcher,
			Setter:    b.Setter,
			Saver:     b.Saver,
			Deleter:   b.Deleter,
			Validator: b.Validator,
		}
	}
	r = b.mb.creating

	r.FieldBuilders = *b.mb.writeFields.Only(vs...)

	return r
}

func (b *EditingBuilder) FetchFunc(v FetchFunc) (r *EditingBuilder) {
	b.Fetcher = v
	return b
}

func (b *EditingBuilder) SaveFunc(v SaveFunc) (r *EditingBuilder) {
	b.Saver = v
	return b
}

func (b *EditingBuilder) DeleteFunc(v DeleteFunc) (r *EditingBuilder) {
	b.Deleter = v
	return b
}

func (b *EditingBuilder) ValidateFunc(v ValidateFunc) (r *EditingBuilder) {
	b.Validator = v
	return b
}

func (b *EditingBuilder) SetterFunc(v SetterFunc) (r *EditingBuilder) {
	b.Setter = v
	return b
}

func (b *EditingBuilder) SidePanelFunc(v ComponentFunc) (r *EditingBuilder) {
	b.sidePanel = v
	return b
}

func (b *EditingBuilder) ActionsFunc(v ComponentFunc) (r *EditingBuilder) {
	b.actionsFunc = v
	return b
}

func (b *EditingBuilder) formNew(ctx *web.EventContext) (r web.EventResponse, err error) {
	creatingB := b
	if b.mb.creating != nil {
		creatingB = b.mb.creating
	}

	b.mb.p.rightDrawer(&r, creatingB.editFormFor(nil, ctx), b.mb.rightDrawerWidth)
	return
}

func (b *EditingBuilder) formEdit(ctx *web.EventContext) (r web.EventResponse, err error) {
	b.mb.p.rightDrawer(&r, b.editFormFor(nil, ctx), b.mb.rightDrawerWidth)
	return
}

func (b *EditingBuilder) editFormFor(obj interface{}, ctx *web.EventContext) h.HTMLComponent {
	msgr := MustGetMessages(ctx.R)

	overlayType := ctx.Event.Params[0]
	id := ctx.Event.Params[1]

	var buttonLabel = msgr.Create
	var disableUpdateBtn bool
	var title = msgr.CreatingObjectTitle(
		i18n.T(ctx.R, ModelsI18nModuleKey, inflection.Singular(b.mb.label)),
	)
	if len(id) > 0 {
		if obj == nil {
			var err error
			obj, err = b.Fetcher(b.mb.NewModel(), id, ctx)
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
		obj = b.mb.NewModel()
	}

	var notice h.HTMLComponent
	if msg, ok := ctx.Flash.(string); ok {
		if len(msg) > 0 {
			notice = VAlert(h.Text(msg)).
				Border("left").
				Type("success").
				Elevation(2).
				ColoredBorder(true)
		}
	}

	vErr, ok := ctx.Flash.(*web.ValidationErrors)
	if ok {
		gErr := vErr.GetGlobalError()
		if len(gErr) > 0 {
			notice = VAlert(h.Text(gErr)).
				Border("left").
				Type("error").
				Elevation(2).
				ColoredBorder(true)
		}
	}

	var actionButtons h.HTMLComponent = h.Components(
		VSpacer(),
		VBtn(buttonLabel).
			Color("primary").
			Disabled(disableUpdateBtn).
			Attr("@click", web.Plaid().
				EventFunc(actions.Update, ctx.Event.Params...).
				URL(b.mb.Info().ListingHref()).
				Go()),
	)

	if b.actionsFunc != nil {
		actionButtons = b.actionsFunc(ctx)
	}

	var formContent h.HTMLComponent = h.Components(
		VCardText(
			notice,
			b.ToComponent(b.mb, obj, vErr, ctx),
		),
		VCardActions(actionButtons),
	)

	var sidePanel h.HTMLComponent
	if b.sidePanel != nil {
		sidePanel = b.sidePanel(ctx)
		formContent = VContainer(
			VRow(
				VCol(
					formContent,
				).Cols(8),
				VCol(
					sidePanel,
				).Cols(4),
			),
		)
	}

	closeBtnVarScript := closeRightDrawerVarScript
	if overlayType == string(actions.Dialog) {
		closeBtnVarScript = closeDialogVarScript
	}

	return h.Components(
		VAppBar(
			VToolbarTitle(title).Class("pl-2"),
			VSpacer(),
			VBtn("").Icon(true).Children(
				VIcon("close"),
			).Attr("@click.stop", closeBtnVarScript),
		).Color("white").Elevation(0).Dense(true),

		VSheet(
			VCard(formContent).Flat(true),
		).Class("pa-2"),
	)
}

func (b *EditingBuilder) doDelete(ctx *web.EventContext) (r web.EventResponse, err error) {
	id := ctx.Event.Params[0]
	var obj = b.mb.NewModel()
	if len(id) > 0 {
		err = b.Deleter(obj, id, ctx)
		if err != nil {
			return
		}
	}

	r.PushState = web.PushState(nil)
	return
}

func (b *EditingBuilder) defaultUpdate(ctx *web.EventContext) (r web.EventResponse, err error) {
	id := ctx.Event.Params[1]
	var newObj = b.mb.NewModel()
	// don't panic for fields that set in SetterFunc
	_ = ctx.UnmarshalForm(newObj)

	if len(id) == 0 {
		if b.mb.Info().Verifier().Do(PermCreate).ObjectOn(newObj).WithReq(ctx.R).IsAllowed() != nil {
			b.UpdateOverlayContent(ctx, &r, newObj, "", perm.PermissionDenied)
			return
		}
	}

	var obj = b.mb.NewModel()
	usingB := b
	if b.mb.creating != nil && len(id) == 0 {
		usingB = b.mb.creating
	}

	if len(id) > 0 {
		obj, err1 := usingB.Fetcher(obj, id, ctx)
		if err1 != nil {
			b.UpdateOverlayContent(ctx, &r, obj, "", err1)
			return
		}
		if b.mb.Info().Verifier().Do(PermUpdate).ObjectOn(obj).WithReq(ctx.R).IsAllowed() != nil {
			b.UpdateOverlayContent(ctx, &r, obj, "", perm.PermissionDenied)
			return
		}
	}

	if usingB.Setter != nil {
		usingB.Setter(obj, ctx)
	}

	var vErr web.ValidationErrors
	for _, f := range usingB.fields {

		if b.mb.Info().Verifier().Do(PermUpdate).ObjectOn(obj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}

		if f.setterFunc == nil {
			val, err1 := reflectutils.Get(newObj, f.name)
			if err1 != nil {
				continue
			}
			_ = reflectutils.Set(obj, f.name, val)
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
		b.UpdateOverlayContent(ctx, &r, obj, "", &vErr)
		return
	}

	if usingB.Validator != nil {
		if vErr := usingB.Validator(obj, ctx); vErr.HaveErrors() {
			b.UpdateOverlayContent(ctx, &r, obj, "", &vErr)
			return
		}
	}

	err1 := usingB.Saver(obj, id, ctx)
	if err1 != nil {
		b.UpdateOverlayContent(ctx, &r, obj, "", err1)
		return
	}

	msgr := MustGetMessages(ctx.R)
	ShowMessage(&r, msgr.SuccessfullyUpdated, "")
	r.PushState = web.PushState(nil)
	r.VarsScript = r.VarsScript + ";" + closeRightDrawerVarScript
	return
}

func (b *EditingBuilder) UpdateOverlayContent(
	ctx *web.EventContext,
	r *web.EventResponse,
	obj interface{},
	successMessage string,
	err error,
) {
	ctx.Flash = err

	if err != nil {
		if _, ok := err.(*web.ValidationErrors); !ok {
			vErr := &web.ValidationErrors{}
			vErr.GlobalError(err.Error())
			ctx.Flash = vErr
		}
	}

	if ctx.Flash == nil {
		ctx.Flash = successMessage
	}

	overlayType := ctx.Event.Params[0]
	p := rightDrawerContentPortalName

	if overlayType == string(actions.Dialog) {
		p = dialogContentPortalName
	}

	r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
		Name: p,
		Body: b.editFormFor(obj, ctx),
	})

}
