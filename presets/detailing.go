package presets

import (
	"net/url"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets/actions"
	. "github.com/goplaid/x/vuetify"
	"github.com/jinzhu/inflection"
	h "github.com/theplant/htmlgo"
	"goji.io/pat"
)

type DetailingBuilder struct {
	mb         *ModelBuilder
	fieldNames []string
	actions    []*ActionBuilder
	pageFunc   web.PageFunc
	fetcher    FetchFunc
	FieldsBuilder
}

type pageTitle interface {
	PageTitle() string
}

func (mb *ModelBuilder) Detailing(vs ...string) (r *DetailingBuilder) {
	r = mb.detailing
	mb.hasDetailing = true
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

	id := pat.Param(ctx.R, "id")
	r.Body = VContainer(h.Text(id))

	var obj = b.mb.NewModel()

	if id == "" {
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
			ModelInfo: b.mb.Info(),
			Name:      f.name,
			Label:     b.mb.getLabel(f.NameLabel),
		}, ctx))
	}

	r.Body = VContainer(
		notice,
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
	action := getAction(b.actions, ctx.R.FormValue(ParamAction))
	if action == nil {
		panic("action required")
	}
	id := ctx.R.FormValue(ParamID)
	err1 := action.updateFunc([]string{id}, ctx)
	if err1 != nil || ctx.Flash != nil {
		if ctx.Flash == nil {
			ctx.Flash = err1
		}

		r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
			Name: rightDrawerContentPortalName,
			Body: b.actionForm(action, ctx),
		})
		return
	}

	r.PushState = web.Location(url.Values{})
	r.VarsScript = closeRightDrawerVarScript

	return
}

func (b *DetailingBuilder) formDrawerAction(ctx *web.EventContext) (r web.EventResponse, err error) {
	action := getAction(b.actions, ctx.R.FormValue(ParamAction))
	if action == nil {
		panic("action required")
	}

	b.mb.p.rightDrawer(&r, b.actionForm(action, ctx), "")
	return
}

func (b *DetailingBuilder) actionForm(action *ActionBuilder, ctx *web.EventContext) h.HTMLComponent {
	msgr := MustGetMessages(ctx.R)

	id := ctx.R.FormValue(ParamID)
	if id == "" {
		panic("id required")
	}

	return VContainer(
		VCard(
			VCardText(
				action.compFunc([]string{id}, ctx),
			),
			VCardActions(
				VSpacer(),
				VBtn(msgr.Update).
					Dark(true).
					Color("primary").
					Attr("@click", web.Plaid().
						EventFunc(actions.DoAction).
						Query(ParamID, id).
						Query(ParamAction, ctx.R.FormValue(ParamAction)).
						URL(b.mb.Info().DetailingHref(id)).
						Go()),
			),
		).Flat(true),
	).Fluid(true)
}
