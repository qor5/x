package presets

import (
	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type ModelBuilder struct {
	p             *PresetsBuilder
	model         interface{}
	uriName       string
	labels        []string
	placeholders  []string
	searchColumns []string
	listing       *ListingBuilder
	editing       *EditingBuilder
	detailing     *DetailingBuilder
}

func (b *ModelBuilder) defaultListing() (r *ListingBuilder) {
	r = &ListingBuilder{filtering: &FilteringBuilder{}}
	return
}

func (b *ModelBuilder) defaultEditing() (r *EditingBuilder) {
	r = &EditingBuilder{}
	return
}

func (b *ModelBuilder) defaultDetailing() (r *DetailingBuilder) {
	r = &DetailingBuilder{}
	return
}

func (b *ModelBuilder) URIName(v string) (r *ModelBuilder) {
	b.uriName = v
	return b
}

func (b *ModelBuilder) Labels(vs ...string) (r *ModelBuilder) {
	b.labels = append(b.labels, vs...)
	return b
}

func (b *ModelBuilder) Placeholders(vs ...string) (r *ModelBuilder) {
	b.placeholders = append(b.placeholders, vs...)
	return b
}

func (b *ModelBuilder) SearchColumns(vs ...string) (r *ModelBuilder) {
	b.searchColumns = vs
	return b
}

func (b *ModelBuilder) listingFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	r.Schema = h.Div(
		h.Text("Hello"),
	)
	return
}

func (b *ModelBuilder) detailingFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	return
}

func (b *ModelBuilder) editingFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	return
}
