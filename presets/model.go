package presets

import (
	"reflect"
)

type ModelBuilder struct {
	p             *Builder
	model         interface{}
	uriName       string
	labels        []string
	placeholders  []string
	searchColumns []string
	listing       *ListingBuilder
	editing       *EditingBuilder
	detailing     *DetailingBuilder
}

func NewModelBuilder(p *Builder, model interface{}) (r *ModelBuilder) {
	r = &ModelBuilder{p: p, model: model}
	r.newListing()
	r.newDetailing()
	r.newEditing()
	r.inspectModel()
	return
}

func (b *ModelBuilder) inspectModel() {
	v := reflect.ValueOf(b.model)

	for v.Elem().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v = v.Elem()

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		//fmt.Println(f.Name, f.Type)
		ft := b.p.fieldTypeByType(f.Type)
		b.listing.Field(f.Name).ComponentFunc(ft.listingCompFunc)
		b.detailing.Field(f.Name).ComponentFunc(ft.detailingCompFunc)
		b.editing.Field(f.Name).ComponentFunc(ft.editingCompFunc)
	}
}

func (b *ModelBuilder) newListing() (r *ListingBuilder) {
	b.listing = &ListingBuilder{filtering: &FilteringBuilder{}, mb: b}
	return
}

func (b *ModelBuilder) newEditing() (r *EditingBuilder) {
	b.editing = &EditingBuilder{mb: b}
	return
}

func (b *ModelBuilder) newDetailing() (r *DetailingBuilder) {
	b.detailing = &DetailingBuilder{mb: b}
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

func (b *ModelBuilder) getLabel(field *FieldBuilder) (r string) {
	if len(field.label) > 0 {
		return field.label
	}

	for i := 0; i < len(b.labels)-1; i = i + 2 {
		if b.labels[i] == field.name {
			return b.labels[i+1]
		}
	}

	return field.name
}
