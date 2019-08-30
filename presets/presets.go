package presets

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/qor/inflection"
	"github.com/sunfmin/bran"
	"github.com/sunfmin/bran/core"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
	"go.uber.org/zap"
	goji "goji.io"
	"goji.io/pat"
)

type Builder struct {
	prefix             string
	models             []*ModelBuilder
	mux                *goji.Mux
	builder            *bran.Builder
	logger             *zap.Logger
	dataOperator       DataOperator
	messagesFunc       MessagesFunc
	homePageFunc       ui.PageFunc
	brandFunc          ComponentFunc
	brandTitle         string
	primaryColor       string
	writeFieldDefaults *FieldDefaults
	listFieldDefaults  *FieldDefaults
	MenuGroups
}

func New() *Builder {
	l, _ := zap.NewDevelopment()
	return &Builder{
		logger:             l,
		builder:            bran.New(),
		messagesFunc:       defaultMessageFunc,
		writeFieldDefaults: NewFieldDefaults(WRITE),
		listFieldDefaults:  NewFieldDefaults(LIST),
		primaryColor:       "indigo",
		brandTitle:         "Admin",
	}
}

func (b *Builder) URIPrefix(v string) (r *Builder) {
	b.prefix = v
	return b
}

func (b *Builder) Builder(v *bran.Builder) (r *Builder) {
	b.builder = v
	return b
}

func (b *Builder) Logger(v *zap.Logger) (r *Builder) {
	b.logger = v
	return b
}

func (b *Builder) MessagesFunc(v MessagesFunc) (r *Builder) {
	b.messagesFunc = v
	return b
}

func (b *Builder) HomePageFunc(v ui.PageFunc) (r *Builder) {
	b.homePageFunc = v
	return b
}

func (b *Builder) BrandFunc(v ComponentFunc) (r *Builder) {
	b.brandFunc = v
	return b
}

func (b *Builder) BrandTitle(v string) (r *Builder) {
	b.brandTitle = v
	return b
}

func (b *Builder) PrimaryColor(v string) (r *Builder) {
	b.primaryColor = v
	return b
}

func (b *Builder) FieldDefaults(v FieldMode) (r *FieldDefaults) {
	if v == WRITE {
		return b.writeFieldDefaults
	}

	if v == LIST {
		return b.listFieldDefaults
	}

	return r
}

func (b *Builder) Model(v interface{}) (r *ModelBuilder) {
	r = NewModelBuilder(b, v)
	b.models = append(b.models, r)
	return r
}

func (b *Builder) DataOperator(v DataOperator) (r *ModelBuilder) {
	b.dataOperator = v
	return r
}

func modelNames(ms []*ModelBuilder) (r []string) {
	for _, m := range ms {
		r = append(r, m.uriName)
	}
	return
}

func (b *Builder) defaultBrandFunc(ctx *ui.EventContext) (r h.HTMLComponent) {
	return
}

func (b *Builder) createMenus(ctx *ui.EventContext) (r h.HTMLComponent) {

	var menus []h.HTMLComponent
	for _, mg := range b.menuGroups {
		var subMenus = []h.HTMLComponent{
			VListItem(
				VListItemContent(
					VListItemTitle(h.Text(mg.label)),
				),
			).Slot("activator").Class("pa-0"),
		}
		for _, m := range mg.models {
			if m.notInMenu {
				continue
			}
			href := m.Info().ListingHref()
			subMenus = append(subMenus,
				ui.Bind(VListItem(
					VListItemAction(
						VIcon(""),
					),
					VListItemContent(
						VListItemTitle(
							h.Text(m.label),
						),
					),
				).Class(activeClass(ctx, href))).PushStateURL(href),
			)
		}
		menus = append(menus, VListGroup(
			subMenus...).
			PrependIcon(mg.icon).
			Value(true).
			Color(b.primaryColor),
		)
	}

	for _, m := range b.models {
		if m.inGroup {
			continue
		}
		if m.notInMenu {
			continue
		}

		href := m.Info().ListingHref()
		menus = append(menus,
			ui.Bind(VListItem(
				VListItemAction(
					VIcon(m.menuIcon),
				),
				VListItemContent(
					VListItemTitle(
						h.Text(m.label),
					),
				),
			).Class(activeClass(ctx, href)).Color(b.primaryColor)).PushStateURL(href),
		)
	}

	r = VList(menus...)
	return
}

func activeClass(ctx *ui.EventContext, url string) string {
	if strings.HasPrefix(ctx.R.URL.Path, url) {
		return "v-list-item--active"
	}
	return ""
}

func (b *Builder) runBrandFunc(ctx *ui.EventContext) (r h.HTMLComponent) {
	if b.brandFunc != nil {
		return b.brandFunc(ctx)
	}

	return VAppBar(
		VToolbarTitle("Admin"),
	)
}

type contextKey int

const (
	messagesKey contextKey = iota
	modelInfoKey
)

func MustGetMessages(r *http.Request) *Messages {
	return r.Context().Value(messagesKey).(*Messages)
}

func GetModelInfo(req *http.Request) (r *ModelInfo) {
	r, _ = req.Context().Value(modelInfoKey).(*ModelInfo)
	return
}

func (b *Builder) putMessages(in http.Handler) (out http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgr := b.messagesFunc(r)
		in.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), messagesKey, msgr)))
	})
}

func putModelInfo(mi *ModelInfo, in http.Handler) (out http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), modelInfoKey, mi)))
	})
}

const rightDrawerName = "rightDrawer"
const rightDrawerPortalName = "rightDrawerPortalName"

func rightDrawer(r *ui.EventResponse, comp h.HTMLComponent) {
	r.UpdatePortals = append(r.UpdatePortals, &ui.PortalUpdate{
		Name: rightDrawerName,
		Schema: VNavigationDrawer(
			ui.LazyPortal(comp).Name(rightDrawerPortalName),
		).Attr("v-model", "vars.rightDrawer").
			Bottom(true).
			Right(true).
			Absolute(true).
			Width(600).
			Temporary(true).
			Attr("v-init-context-vars", `{rightDrawer: false}`),
		AfterLoaded: `setTimeout(function(){ comp.vars.rightDrawer = true }, 100)`,
	})
}

func (b *Builder) defaultLayout(in ui.PageFunc) (out ui.PageFunc) {
	return func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

		ctx.Injector.PutHeadHTML(`
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono">
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500">
			<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
			<link rel="stylesheet" href="/assets/main.css">
			<script src='/assets/vue.js'></script>
			<style>
				[v-cloak] {
					display: none;
				}
			</style>
		`)

		if len(os.Getenv("DEV")) > 0 {
			ctx.Injector.PutTailHTML(`
			<script src='http://localhost:3080/app.js'></script>
			<script src='http://localhost:3100/app.js'></script>
			`)

		} else {
			ctx.Injector.PutTailHTML(`
			<script src='/assets/main.js'></script>
			`)
		}

		var innerPr ui.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}
		pr.PageTitle = innerPr.PageTitle
		pr.Schema = VApp(

			VNavigationDrawer(
				b.createMenus(ctx),
			).App(true).
				Clipped(true).
				Value(true).
				Attr("v-model", "vars.navDrawer").
				Attr("v-init-context-vars", `{navDrawer: null}`),

			ui.LazyPortal().EventFunc("").Name(rightDrawerName),

			VAppBar(
				VAppBarNavIcon().On("click.stop", "vars.navDrawer = !vars.navDrawer"),
				b.runBrandFunc(ctx),
				VSpacer(),
				VLayout(
					// h.Form(
					VTextField().
						SoloInverted(true).
						PrependIcon("search").
						Label("Search").
						Flat(true).
						Clearable(true).
						HideDetails(true).
						Value(ctx.R.URL.Query().Get("keyword")).
						Attr("@keyup.enter", `topage({ query: {keyword: [$event.target.value]}})`),
					// ).Method("GET"),
				).AlignCenter(true).Attr("style", "max-width: 650px"),
			).Dark(true).
				Color(b.primaryColor).
				App(true).
				ClippedLeft(true),

			VContent(
				innerPr.Schema.(h.HTMLComponent),
			),
		).Id("vt-app")

		return
	}
}

func (b *Builder) defaultHomePageFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	r.Schema = h.Div().Text("home")
	return
}

func (b *Builder) getHomePageFunc() ui.PageFunc {
	if b.homePageFunc != nil {
		return b.homePageFunc
	}
	return b.defaultHomePageFunc
}

func (b *Builder) initMux() {
	b.logger.Info("initializing mux for", zap.Reflect("models", modelNames(b.models)))
	mux := goji.NewMux()
	ub := b.builder

	mux.Handle(pat.Get("/assets/main.js"),
		ub.PacksHandler("text/javascript",
			JSComponentsPack(),
			core.JSComponentsPack(),
		),
	)

	mux.Handle(pat.Get("/assets/vue.js"),
		ub.PacksHandler("text/javascript",
			core.JSVueComponentsPack(),
		),
	)

	mux.Handle(pat.Get("/assets/main.css"),
		ub.PacksHandler("text/css",
			CSSComponentsPack(),
		),
	)

	mux.Handle(
		pat.New(b.prefix),
		b.wrap(nil, b.defaultLayout(b.getHomePageFunc())),
	)

	for _, m := range b.models {
		muri := inflection.Plural(m.uriName)
		info := m.Info()
		routePath := info.ListingHref()
		mux.Handle(
			pat.New(routePath),
			b.wrap(info, b.defaultLayout(m.listing.GetPageFunc())),
		)
		log.Println("mounted url", routePath)
		if m.hasDetailing {
			routePath = fmt.Sprintf("%s/%s/:id", b.prefix, muri)
			mux.Handle(
				pat.New(routePath),
				b.wrap(info, b.defaultLayout(m.detailing.GetPageFunc())),
			)
			log.Println("mounted url", routePath)
		}

		routePath = fmt.Sprintf("%s/%s/:id/edit", b.prefix, muri)
		mux.Handle(
			pat.New(routePath),
			b.wrap(info, b.defaultLayout(m.editing.GetPageFunc())),
		)
		log.Println("mounted url", routePath)

		routePath = fmt.Sprintf("%s/%s/new", b.prefix, muri)
		mux.Handle(
			pat.New(fmt.Sprintf("%s/%s/new", b.prefix, muri)),
			b.wrap(info, b.defaultLayout(m.editing.GetPageFunc())),
		)
		log.Println("mounted url", routePath)
	}

	b.mux = mux
}

func (b *Builder) wrap(mi *ModelInfo, pf ui.PageFunc) http.Handler {
	return putModelInfo(
		mi,
		b.putMessages(
			b.builder.Page(pf),
		),
	)
}

func (b *Builder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if b.mux == nil {
		b.initMux()
	}
	b.mux.ServeHTTP(w, r)
}
