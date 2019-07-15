package presets

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/sunfmin/bran"

	"github.com/qor/inflection"
	"goji.io/pat"

	goji "goji.io"
)

type PresetsBuilder struct {
	prefix  string
	models  []*ModelBuilder
	mux     *goji.Mux
	builder *bran.Builder
	logger  *zap.Logger
}

func New() *PresetsBuilder {
	l, _ := zap.NewDevelopment()
	return &PresetsBuilder{
		logger:  l,
		builder: bran.New(),
	}
}

func (b *PresetsBuilder) URIPrefix(v string) (r *PresetsBuilder) {
	b.prefix = v
	return b
}

func (b *PresetsBuilder) Builder(v *bran.Builder) (r *PresetsBuilder) {
	b.builder = v
	return b
}

func (b *PresetsBuilder) Logger(v *zap.Logger) (r *PresetsBuilder) {
	b.logger = v
	return b
}

func (b *PresetsBuilder) Model(v interface{}) (r *ModelBuilder) {
	r = &ModelBuilder{p: b}
	r.model = v
	b.models = append(b.models, r)
	r.listing = r.defaultListing()
	r.editing = r.defaultEditing()
	r.detailing = r.defaultDetailing()
	return r
}

func modelNames(ms []*ModelBuilder) (r []string) {
	for _, m := range ms {
		r = append(r, m.uriName)
	}
	return
}

func (b *PresetsBuilder) initMux() {
	b.logger.Info("initializing mux for", zap.Reflect("models", modelNames(b.models)))
	b.mux = goji.NewMux()
	for _, m := range b.models {
		muri := inflection.Plural(m.uriName)
		b.mux.Handle(
			pat.New(fmt.Sprintf("%s/%s", b.prefix, muri)),
			b.builder.Page(m.listingFunc),
		)
		b.mux.Handle(
			pat.New(fmt.Sprintf("%s/%s/:id", b.prefix, muri)),
			b.builder.Page(m.detailingFunc),
		)
		b.mux.Handle(
			pat.New(fmt.Sprintf("%s/%s/:id/edit", b.prefix, muri)),
			b.builder.Page(m.editingFunc),
		)
		b.mux.Handle(
			pat.New(fmt.Sprintf("%s/%s/new", b.prefix, muri)),
			b.builder.Page(m.editingFunc),
		)
	}
}

func (b *PresetsBuilder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if b.mux == nil {
		b.initMux()
	}
	b.mux.ServeHTTP(w, r)
}
