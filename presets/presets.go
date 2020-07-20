package presets

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	"github.com/jinzhu/inflection"
	h "github.com/theplant/htmlgo"
	"go.uber.org/zap"
	goji "goji.io"
	"goji.io/pat"
)

type Builder struct {
	prefix              string
	models              []*ModelBuilder
	mux                 *goji.Mux
	builder             *web.Builder
	logger              *zap.Logger
	dataOperator        DataOperator
	messagesFunc        MessagesFunc
	homePageFunc        web.PageFunc
	brandFunc           ComponentFunc
	brandTitle          string
	primaryColor        string
	progressBarColor    string
	writeFieldDefaults  *FieldDefaults
	listFieldDefaults   *FieldDefaults
	detailFieldDefaults *FieldDefaults
	extraAssets         []*extraAsset
	assetFunc           AssetFunc
	MenuGroups
}

type AssetFunc func(ctx *web.EventContext)

type extraAsset struct {
	path        string
	contentType string
	body        web.ComponentsPack
	refTag      string
}

func New() *Builder {
	l, _ := zap.NewDevelopment()
	return &Builder{
		logger:              l,
		builder:             web.New(),
		messagesFunc:        defaultMessageFunc,
		writeFieldDefaults:  NewFieldDefaults(WRITE),
		listFieldDefaults:   NewFieldDefaults(LIST),
		detailFieldDefaults: NewFieldDefaults(DETAIL),
		primaryColor:        "indigo",
		progressBarColor:    "amber",
		brandTitle:          "Admin",
	}
}

func (b *Builder) URIPrefix(v string) (r *Builder) {
	b.prefix = strings.TrimRight(v, "/")
	return b
}

func (b *Builder) Builder(v *web.Builder) (r *Builder) {
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

func (b *Builder) HomePageFunc(v web.PageFunc) (r *Builder) {
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

func (b *Builder) ProgressBarColor(v string) (r *Builder) {
	b.progressBarColor = v
	return b
}

func (b *Builder) AssetFunc(v AssetFunc) (r *Builder) {
	b.assetFunc = v
	return b
}

func (b *Builder) ExtraAsset(path string, contentType string, body web.ComponentsPack, refTag ...string) (r *Builder) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	var theOne *extraAsset
	for _, ea := range b.extraAssets {
		if ea.path == path {
			theOne = ea
			break
		}
	}

	if theOne == nil {
		theOne = &extraAsset{path: path, contentType: contentType, body: body}
		b.extraAssets = append(b.extraAssets, theOne)
	} else {
		theOne.contentType = contentType
		theOne.body = body
	}

	if len(refTag) > 0 {
		theOne.refTag = refTag[0]
	}

	return b
}

func (b *Builder) FieldDefaults(v FieldMode) (r *FieldDefaults) {
	if v == WRITE {
		return b.writeFieldDefaults
	}

	if v == LIST {
		return b.listFieldDefaults
	}

	if v == DETAIL {
		return b.detailFieldDefaults
	}

	return r
}

func (b *Builder) Model(v interface{}) (r *ModelBuilder) {
	r = NewModelBuilder(b, v)
	b.models = append(b.models, r)
	return r
}

func (b *Builder) DataOperator(v DataOperator) (r *Builder) {
	b.dataOperator = v
	return b
}

func modelNames(ms []*ModelBuilder) (r []string) {
	for _, m := range ms {
		r = append(r, m.uriName)
	}
	return
}

func (b *Builder) defaultBrandFunc(ctx *web.EventContext) (r h.HTMLComponent) {
	return
}

func (b *Builder) createMenus(ctx *web.EventContext) (r h.HTMLComponent) {

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
				web.Bind(VListItem(
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
			web.Bind(VListItem(
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

func activeClass(ctx *web.EventContext, url string) string {
	if strings.HasPrefix(ctx.R.URL.Path, url) {
		return "v-list-item--active"
	}
	return ""
}

func (b *Builder) runBrandFunc(ctx *web.EventContext) (r h.HTMLComponent) {
	if b.brandFunc != nil {
		return b.brandFunc(ctx)
	}

	return VToolbarTitle(b.brandTitle)
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

func rightDrawer(r *web.EventResponse, comp h.HTMLComponent) {
	r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
		Name: rightDrawerName,
		Body: VNavigationDrawer(
			web.Portal(comp).Name(rightDrawerPortalName),
		).Attr("v-model", "vars.rightDrawer").
			Bottom(true).
			Right(true).
			Absolute(true).
			Width(600).
			Temporary(true).
			Attr(web.InitContextVars, `{rightDrawer: false}`),
		AfterLoaded: `setTimeout(function(){ comp.vars.rightDrawer = true }, 100)`,
	})
}

func (b *Builder) defaultLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {

		ctx.Injector.HeadHTML(strings.Replace(`
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono">
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500">
			<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
			<link rel="stylesheet" href="{{prefix}}/assets/main.css">
			<script src='{{prefix}}/assets/vue.js'></script>
			<style>
				[v-cloak] {
					display: none;
				}
			</style>
		`, "{{prefix}}", b.prefix, -1))

		for _, ea := range b.extraAssets {
			if len(ea.refTag) > 0 {
				ctx.Injector.HeadHTML(ea.refTag)
				continue
			}

			if strings.HasSuffix(ea.path, "css") {
				ctx.Injector.HeadHTML(fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s\">", b.extraFullPath(ea)))
				continue
			}

			ctx.Injector.HeadHTML(fmt.Sprintf("<script src=\"%s\"></script>", b.extraFullPath(ea)))
		}

		if len(os.Getenv("DEV")) > 0 {
			ctx.Injector.TailHTML(`
			<script src='http://localhost:3080/app.js'></script>
			<script src='http://localhost:3100/app.js'></script>
			`)

		} else {
			ctx.Injector.TailHTML(strings.Replace(`
			<script src='{{prefix}}/assets/main.js'></script>
			`, "{{prefix}}", b.prefix, -1))
		}

		if b.assetFunc != nil {
			b.assetFunc(ctx)
		}

		var innerPr web.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		pr.PageTitle = fmt.Sprintf("%s - %s", innerPr.PageTitle, b.brandTitle)
		pr.Body = VApp(

			VNavigationDrawer(
				b.createMenus(ctx),
			).App(true).
				Clipped(true).
				Value(true).
				Attr("v-model", "vars.navDrawer").
				Attr(web.InitContextVars, `{navDrawer: null}`),

			web.Portal().EventFunc("").Name(rightDrawerName),

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

			VProgressLinear().
				Attr(":active", "isFetching").
				Attr("style", "position: fixed; z-index: 99").
				Indeterminate(true).
				Height(2).
				Color(b.progressBarColor),

			VMain(
				innerPr.Body.(h.HTMLComponent),
			),
		).Id("vt-app")

		return
	}
}

func (b *Builder) defaultHomePageFunc(ctx *web.EventContext) (r web.PageResponse, err error) {
	r.Body = h.Div().Text("home")
	return
}

func (b *Builder) getHomePageFunc() web.PageFunc {
	if b.homePageFunc != nil {
		return b.homePageFunc
	}
	return b.defaultHomePageFunc
}

func (b *Builder) extraFullPath(ea *extraAsset) string {
	return b.prefix + "/extra" + ea.path
}

func (b *Builder) initMux() {
	b.logger.Info("initializing mux for", zap.Reflect("models", modelNames(b.models)), zap.String("prefix", b.prefix))
	mux := goji.NewMux()
	ub := b.builder

	mainJSPath := b.prefix + "/assets/main.js"
	mux.Handle(pat.Get(mainJSPath),
		ub.PacksHandler("text/javascript",
			JSComponentsPack(),
			web.JSComponentsPack(),
		),
	)
	log.Println("mounted url", mainJSPath)

	vueJSPath := b.prefix + "/assets/vue.js"
	mux.Handle(pat.Get(vueJSPath),
		ub.PacksHandler("text/javascript",
			web.JSVueComponentsPack(),
		),
	)
	log.Println("mounted url", vueJSPath)

	mainCSSPath := b.prefix + "/assets/main.css"
	mux.Handle(pat.Get(mainCSSPath),
		ub.PacksHandler("text/css",
			CSSComponentsPack(),
		),
	)
	log.Println("mounted url", mainCSSPath)

	for _, ea := range b.extraAssets {
		fullPath := b.extraFullPath(ea)
		mux.Handle(pat.Get(fullPath), ub.PacksHandler(
			ea.contentType,
			ea.body,
		))
		log.Println("mounted url", fullPath)
	}

	mux.Handle(
		pat.New(b.prefix),
		b.wrap(nil, b.defaultLayout(b.getHomePageFunc())),
	)

	for _, m := range b.models {
		pluralUri := inflection.Plural(m.uriName)
		info := m.Info()
		routePath := info.ListingHref()
		mux.Handle(
			pat.New(routePath),
			b.wrap(info, b.defaultLayout(m.listing.GetPageFunc())),
		)
		log.Println("mounted url", routePath)
		if m.hasDetailing {
			routePath = fmt.Sprintf("%s/%s/:id", b.prefix, pluralUri)
			mux.Handle(
				pat.New(routePath),
				b.wrap(info, b.defaultLayout(m.detailing.GetPageFunc())),
			)
			log.Println("mounted url", routePath)
		}
	}

	b.mux = mux
}

func (b *Builder) wrap(mi *ModelInfo, pf web.PageFunc) http.Handler {
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
