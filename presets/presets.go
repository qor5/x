package presets

import (
	"fmt"
	"net/http"

	"github.com/qor/inflection"
	"github.com/sunfmin/bran"
	"github.com/sunfmin/bran/core"
	branoverlay "github.com/sunfmin/bran/overlay"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
	"go.uber.org/zap"
	goji "goji.io"
	"goji.io/pat"
)

type Builder struct {
	prefix       string
	models       []*ModelBuilder
	mux          *goji.Mux
	builder      *bran.Builder
	logger       *zap.Logger
	dataOperator DataOperator
	messagesFunc MessagesFunc
	FieldTypes
}

type DataOperator interface {
	Search(obj interface{}, params *SearchParams) (r interface{}, err error)
	Fetch(obj interface{}, id string) (r interface{}, err error)
	UpdateField(obj interface{}, id string, fieldName string, value interface{}) (err error)
	Save(obj interface{}, id string) (err error)
}

func New() *Builder {
	l, _ := zap.NewDevelopment()
	return &Builder{
		logger:       l,
		builder:      bran.New(),
		messagesFunc: defaultMessageFunc,
		FieldTypes:   builtInFieldTypes(),
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

func (b *Builder) defaultLayout(in ui.PageFunc) (out ui.PageFunc) {
	return func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

		ctx.Injector.Title("Hello")
		ctx.Injector.PutHeadHTML(`
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" async>
			<link rel="stylesheet" href="/assets/main.css">
			<script src='/assets/vue.js'></script>
			<style>
				[v-cloak] {
					display: none;
				}
			</style>
		`)

		ctx.Injector.PutTailHTML(`
			<script src='/assets/main.js'></script>
		`)

		var innerPr ui.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		pr.Schema = VApp(
			VNavigationDrawer(
				VToolbar(
					VToolbarTitle("Hello"),
				),
			).App(true),
			VToolbar(
				h.Form(
					VTextField().
						SoloInverted(true).
						PrependInnerIcon("search").
						Label("Search").
						Flat(true).
						Clearable(true).HideDetails(true),
				).Method("GET"),

				VSpacer(),
				VToolbarItems(),
			).App(true).Dark(true).Color("indigo"),
			VContent(
				innerPr.Schema.(h.HTMLComponent),
			),
		).Id("vt-app")

		pr.State = innerPr.State

		return
	}
}

func (b *Builder) initMux() {
	b.logger.Info("initializing mux for", zap.Reflect("models", modelNames(b.models)))
	mux := goji.NewMux()
	ub := b.builder

	mux.Handle(pat.Get("/assets/main.js"),
		ub.PacksHandler("text/javascript",
			branoverlay.JSComponentsPack(),
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
			branoverlay.CSSComponentsPack(),
			CSSComponentsPack(),
		),
	)

	for _, m := range b.models {
		muri := inflection.Plural(m.uriName)
		mux.Handle(
			pat.New(fmt.Sprintf("%s/%s", b.prefix, muri)),
			b.builder.Page(b.defaultLayout(m.listing.GetPageFunc())),
		)
		mux.Handle(
			pat.New(fmt.Sprintf("%s/%s/:id", b.prefix, muri)),
			b.builder.Page(b.defaultLayout(m.detailing.GetPageFunc())),
		)
		mux.Handle(
			pat.New(fmt.Sprintf("%s/%s/:id/edit", b.prefix, muri)),
			b.builder.Page(b.defaultLayout(m.editing.GetPageFunc())),
		)
		mux.Handle(
			pat.New(fmt.Sprintf("%s/%s/new", b.prefix, muri)),
			b.builder.Page(b.defaultLayout(m.editing.GetPageFunc())),
		)
	}

	b.mux = mux
}

func (b *Builder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if b.mux == nil {
		b.initMux()
	}
	b.mux.ServeHTTP(w, r)
}
