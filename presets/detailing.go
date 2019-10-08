package presets

import (
	"net/url"

	"github.com/goplaid/web"

	"github.com/jinzhu/inflection"

	"github.com/goplaid/x/presets/actions"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
	"goji.io/pat"
)

type DetailingBuilder struct {
	mb         *ModelBuilder
	fieldNames []string
	actions    []*ActionBuilder
	pageFunc   web.PageFunc
	fetcher    FetchFunc
	FieldBuilders
}

type pageTitle interface {
	PageTitle() string
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

func (b *DetailingBuilder) PageFunc(pf web.PageFunc) (r *DetailingBuilder) {
	b.pageFunc = pf
	return b
}

func (b *DetailingBuilder) Fetcher(v FetchFunc) (r *DetailingBuilder) {
	b.fetcher = v
	return b
}

func (b *DetailingBuilder) GetPageFunc() web.PageFunc {
	if b.pageFunc != nil {
		return b.pageFunc
	}
	return b.defaultPageFunc
}

func (b *DetailingBuilder) defaultPageFunc(ctx *web.EventContext) (r web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc(actions.DrawerAction, b.formDrawerAction)
	ctx.Hub.RegisterEventFunc("doAction", b.doAction)

	id := pat.Param(ctx.R, "id")
	r.Body = VContainer(h.Text(id))

	var obj = b.mb.newModel()

	if len(id) == 0 {
		panic("not found")
	}

	obj, err = b.fetcher(obj, id, ctx)
	if err != nil {
		return
	}

	msgr := MustGetMessages(ctx.R)
	r.PageTitle = msgr.DetailingObjectTitle(inflection.Singular(b.mb.label), getPageTitle(obj, id))

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

	r.Body = VContainer(
		notice,
		web.Portal().Name(deleteConfirmPortalName),
	).AppendChildren(comps...).Fluid(true)
	return
}

func getPageTitle(obj interface{}, id string) string {
	title := id
	if pt, ok := obj.(pageTitle); ok {
		title = pt.PageTitle()
	}
	return title
}

func (b *DetailingBuilder) doAction(ctx *web.EventContext) (r web.EventResponse, err error) {
	action := getAction(b.actions, ctx.Event.Params[0])
	if action == nil {
		panic("action required")
	}
	id := ctx.Event.Params[1]
	err1 := action.updateFunc([]string{id}, ctx)
	if err1 != nil || ctx.Flash != nil {
		if ctx.Flash == nil {
			ctx.Flash = err1
		}

		r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
			Name: rightDrawerPortalName,
			Body: b.actionForm(action, ctx),
		})
		return
	}

	r.PushState = web.PushState(url.Values{})

	return
}

func (b *DetailingBuilder) formDrawerAction(ctx *web.EventContext) (r web.EventResponse, err error) {
	action := getAction(b.actions, ctx.Event.Params[0])
	if action == nil {
		panic("action required")
	}

	rightDrawer(&r, b.actionForm(action, ctx))
	return
}

func (b *DetailingBuilder) actionForm(action *ActionBuilder, ctx *web.EventContext) h.HTMLComponent {
	msgr := MustGetMessages(ctx.R)

	id := ctx.Event.Params[1]
	if len(id) == 0 {
		panic("id required")
	}

	return VContainer(
		VCard(
			VCardText(
				action.compFunc([]string{id}, ctx),
			),
			VCardActions(
				VSpacer(),
				web.Bind(VBtn(msgr.Update).
					Dark(true).
					Color(b.mb.p.primaryColor)).
					OnClick("doAction",
						ctx.Event.Params...).
					URL(b.mb.Info().DetailingHref(id)),
			),
		).Flat(true),
	).Fluid(true)
}
