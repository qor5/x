package presets

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/goplaid/web"
	"github.com/goplaid/x/i18n"
	"github.com/goplaid/x/perm"
	. "github.com/goplaid/x/vuetify"
	"github.com/goplaid/x/vuetifyx"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	h "github.com/theplant/htmlgo"
	"go.uber.org/zap"
	goji "goji.io"
	"goji.io/pat"
	"golang.org/x/text/language"
)

type Builder struct {
	prefix              string
	models              []*ModelBuilder
	mux                 *goji.Mux
	builder             *web.Builder
	i18nBuilder         *i18n.Builder
	logger              *zap.Logger
	permissionBuilder   *perm.Builder
	verifier            *perm.Verifier
	dataOperator        DataOperator
	messagesFunc        MessagesFunc
	homePageFunc        web.PageFunc
	brandFunc           ComponentFunc
	profileFunc         ComponentFunc
	brandTitle          string
	vuetifyOptions      string
	progressBarColor    string
	rightDrawerWidth    int
	writeFieldDefaults  *FieldDefaults
	listFieldDefaults   *FieldDefaults
	detailFieldDefaults *FieldDefaults
	extraAssets         []*extraAsset
	assetFunc           AssetFunc
	menuGroups          MenuGroups
	menuOrder           []interface{}
}

type AssetFunc func(ctx *web.EventContext)

type extraAsset struct {
	path        string
	contentType string
	body        web.ComponentsPack
	refTag      string
}

const (
	CoreI18nModuleKey   i18n.ModuleKey = "CoreI18nModuleKey"
	ModelsI18nModuleKey i18n.ModuleKey = "ModelsI18nModuleKey"
)

func New() *Builder {
	l, _ := zap.NewDevelopment()
	return &Builder{
		logger:  l,
		builder: web.New(),
		i18nBuilder: i18n.New().
			RegisterForModule(language.English, CoreI18nModuleKey, Messages_en_US).
			RegisterForModule(language.SimplifiedChinese, CoreI18nModuleKey, Messages_zh_CN),
		writeFieldDefaults:  NewFieldDefaults(WRITE),
		listFieldDefaults:   NewFieldDefaults(LIST),
		detailFieldDefaults: NewFieldDefaults(DETAIL),
		progressBarColor:    "amber",
		brandTitle:          "Admin",
		rightDrawerWidth:    600,
		verifier:            perm.NewVerifier(PermModule, nil),
	}
}

func (b *Builder) I18n() (r *i18n.Builder) {
	return b.i18nBuilder
}

func (b *Builder) Permission(v *perm.Builder) (r *Builder) {
	b.permissionBuilder = v
	b.verifier = perm.NewVerifier(PermModule, v)
	return b
}

func (b *Builder) GetPermission() (r *perm.Builder) {
	return b.permissionBuilder
}

func (b *Builder) URIPrefix(v string) (r *Builder) {
	b.prefix = strings.TrimRight(v, "/")
	return b
}

func (b *Builder) Builder(v *web.Builder) (r *Builder) {
	b.builder = v
	return b
}

func (b *Builder) GetWebBuilder() (r *web.Builder) {
	return b.builder
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

func (b *Builder) ProfileFunc(v ComponentFunc) (r *Builder) {
	b.profileFunc = v
	return b
}

func (b *Builder) BrandTitle(v string) (r *Builder) {
	b.brandTitle = v
	return b
}

func (b *Builder) VuetifyOptions(v string) (r *Builder) {
	b.vuetifyOptions = v
	return b
}

func (b *Builder) RightDrawerWidth(v int) (r *Builder) {
	b.rightDrawerWidth = v
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

func (b *Builder) MenuGroup(name string) *MenuGroupBuilder {
	mgb := b.menuGroups.MenuGroup(name)
	if !b.isMenuGroupInOrder(mgb) {
		b.menuOrder = append(b.menuOrder, mgb)
	}
	return mgb
}

func (b *Builder) isMenuGroupInOrder(mgb *MenuGroupBuilder) bool {
	for _, v := range b.menuOrder {
		if v == mgb {
			return true
		}
	}
	return false
}

func (b *Builder) removeMenuGroupInOrder(mgb *MenuGroupBuilder) {
	for i, om := range b.menuOrder {
		if om == mgb {
			b.menuOrder = append(b.menuOrder[:i], b.menuOrder[i+1:]...)
			break
		}
	}

}

// item can be URI name, model name, *MenuGroupBuilder
// the underlying logic is using URI name,
// so if the URI name is customized, item must be the URI name
// example:
// b.MenuOrder(
// 	b.MenuGroup("Product Management").SubItems(
// 		"products",
// 		"Variant",
// 	),
// 	"customized-uri",
// )
func (b *Builder) MenuOrder(items ...interface{}) {
	for _, item := range items {
		switch v := item.(type) {
		case string:
			b.menuOrder = append(b.menuOrder, v)
		case *MenuGroupBuilder:
			if b.isMenuGroupInOrder(v) {
				b.removeMenuGroupInOrder(v)
			}
			b.menuOrder = append(b.menuOrder, v)
		default:
			panic(fmt.Sprintf("unknown menu order item type: %T\n", item))
		}
	}
}

func (b *Builder) menuItem(ctx *web.EventContext, m *ModelBuilder, isSub bool) (r h.HTMLComponent) {
	menuIcon := m.menuIcon
	if isSub {
		menuIcon = ""
	}

	href := m.Info().ListingHref()
	item := VListItem(
		VListItemAction(
			VIcon(menuIcon),
		),
		VListItemContent(
			VListItemTitle(
				h.Text(i18n.T(ctx.R, ModelsI18nModuleKey, m.label)),
			),
		),
	).Attr("@click", web.Plaid().PushStateURL(href).Go())
	if b.isMenuItemActive(ctx, m) {
		item = item.Class("v-list-item--active")
	}
	return item
}

func (b *Builder) isMenuItemActive(ctx *web.EventContext, m *ModelBuilder) bool {
	href := m.Info().ListingHref()
	if strings.HasPrefix(ctx.R.URL.Path, href) {
		return true
	}

	return false
}

func (b *Builder) createMenus(ctx *web.EventContext) (r h.HTMLComponent) {
	mMap := make(map[string]*ModelBuilder)
	for _, m := range b.models {
		mMap[m.uriName] = m
	}

	inOrderMap := make(map[string]struct{})
	var menus []h.HTMLComponent
	for _, om := range b.menuOrder {
		switch v := om.(type) {
		case *MenuGroupBuilder:
			ver := b.verifier.Do(PermList).On("menu:groups").SnakeOn(v.name).WithReq(ctx.R)
			if ver.IsAllowed() != nil {
				continue
			}
			var subMenus = []h.HTMLComponent{
				VListItem(
					VListItemContent(
						VListItemTitle(h.Text(i18n.T(ctx.R, ModelsI18nModuleKey, v.name))),
					),
				).Slot("activator").Class("pa-0"),
			}
			subCount := 0
			hasActiveMenuItem := false
			for _, subOm := range v.subMenuItems {
				m, ok := mMap[subOm]
				if !ok {
					m = mMap[inflection.Plural(strcase.ToKebab(subOm))]
				}
				if m == nil {
					continue
				}
				m.menuGroupName = v.name
				if m.notInMenu {
					continue
				}
				if ver.SnakeOn(m.uriName).IsAllowed() != nil {
					continue
				}

				subMenus = append(subMenus, b.menuItem(ctx, m, true))
				subCount++
				inOrderMap[m.uriName] = struct{}{}
				if b.isMenuItemActive(ctx, m) {
					hasActiveMenuItem = true
				}
			}
			if subCount == 0 {
				continue
			}
			menus = append(menus, VListGroup(
				subMenus...).
				PrependIcon(v.icon).
				Value(hasActiveMenuItem),
			)
		case string:
			m, ok := mMap[v]
			if !ok {
				m = mMap[inflection.Plural(strcase.ToKebab(v))]
			}
			if m == nil {
				continue
			}
			if b.verifier.Do(PermList).On("menu").SnakeOn(m.uriName).WithReq(ctx.R).IsAllowed() != nil {
				continue
			}

			if m.notInMenu {
				continue
			}

			menus = append(menus, b.menuItem(ctx, m, false))
			inOrderMap[m.uriName] = struct{}{}
		}
	}

	for _, m := range b.models {
		_, ok := inOrderMap[m.uriName]
		if ok {
			continue
		}

		if b.verifier.Do(PermList).On("menu").SnakeOn(m.uriName).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}

		if m.notInMenu {
			continue
		}
		menus = append(menus, b.menuItem(ctx, m, false))
	}

	r = VList(menus...).Class("primary--text")
	return
}

func (b *Builder) runBrandFunc(ctx *web.EventContext) (r h.HTMLComponent) {
	if b.brandFunc != nil {
		return b.brandFunc(ctx)
	}

	return VToolbarTitle(i18n.T(ctx.R, ModelsI18nModuleKey, b.brandTitle))
}

type contextKey int

const (
	presetsKey contextKey = iota
)

func MustGetMessages(r *http.Request) *Messages {
	return i18n.MustGetModuleMessages(r, CoreI18nModuleKey, Messages_en_US).(*Messages)
}

const rightDrawerName = "rightDrawer"
const rightDrawerContentPortalName = "rightDrawerContentPortalName"

func (b *Builder) rightDrawer(r *web.EventResponse, comp h.HTMLComponent) {
	r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
		Name: rightDrawerName,
		Body: VNavigationDrawer(
			web.Portal(comp).Name(rightDrawerContentPortalName),
		).
			Class("v-navigation-drawer--temporary").
			Attr("v-model", "vars.rightDrawer").
			Right(true).
			Fixed(true).
			Width(b.rightDrawerWidth).
			Bottom(false).
			Attr(":height", `"100%"`).
			//Temporary(true).
			//HideOverlay(true).
			//Floating(true).
			Attr(web.InitContextVars, `{rightDrawer: false}`),
	})
	r.VarsScript = "setTimeout(function(){ vars.rightDrawer = true }, 100)"
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

		if len(os.Getenv("DEV_PRESETS")) > 0 {
			ctx.Injector.TailHTML(`
<script src='http://localhost:3080/js/chunk-vendors.js'></script>
<script src='http://localhost:3080/js/app.js'></script>
<script src='http://localhost:3100/js/chunk-vendors.js'></script>
<script src='http://localhost:3100/js/app.js'></script>
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

		var profile h.HTMLComponent
		if b.profileFunc != nil {
			profile = b.profileFunc(ctx)
		}

		msgr := i18n.MustGetModuleMessages(ctx.R, CoreI18nModuleKey, Messages_en_US).(*Messages)

		pr.PageTitle = fmt.Sprintf("%s - %s", innerPr.PageTitle, i18n.T(ctx.R, ModelsI18nModuleKey, b.brandTitle))
		pr.Body = VApp(

			VNavigationDrawer(
				b.createMenus(ctx),
			).App(true).
				Clipped(true).
				Value(true).
				Attr("v-model", "vars.navDrawer").
				Attr(web.InitContextVars, `{navDrawer: null}`),

			web.Portal().Name(rightDrawerName),

			VAppBar(
				VAppBarNavIcon().On("click.stop", "vars.navDrawer = !vars.navDrawer"),
				b.runBrandFunc(ctx),
				VSpacer(),
				VLayout(
					// h.Form(
					VTextField().
						SoloInverted(true).
						PrependIcon("search").
						Label(msgr.Search).
						Flat(true).
						Clearable(true).
						HideDetails(true).
						Value(ctx.R.URL.Query().Get("keyword")).
						Attr("@keyup.enter", web.Plaid().
							Query("keyword", web.Var("[$event.target.value]")).
							Go()),
					// ).Method("GET"),
				).AlignCenter(true).Attr("style", "max-width: 650px"),
				profile,
			).Dark(true).
				Color("primary").
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
			Vuetify(b.vuetifyOptions),
			JSComponentsPack(),
			vuetifyx.JSComponentsPack(),
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
			b.wrap(m, b.defaultLayout(m.listing.GetPageFunc())),
		)
		log.Println("mounted url", routePath)
		if m.hasDetailing {
			routePath = fmt.Sprintf("%s/%s/:id", b.prefix, pluralUri)
			mux.Handle(
				pat.New(routePath),
				b.wrap(m, b.defaultLayout(m.detailing.GetPageFunc())),
			)
			log.Println("mounted url", routePath)
		}
	}

	b.mux = mux
}

func (b *Builder) wrap(m *ModelBuilder, pf web.PageFunc) http.Handler {
	p := b.builder.Page(pf)
	if m != nil {
		m.ensureEventFuncs(p)
	}
	return b.I18n().EnsureLanguage(
		p,
	)
}

func (b *Builder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if b.mux == nil {
		b.initMux()
	}
	b.mux.ServeHTTP(w, r)
}
