package presets

type FilteringBuilder struct {
	fields []*FieldBuilder
}

func (b *ListingBuilder) Filtering(vs ...string) (r *FilteringBuilder) {
	r = b.filtering
	var newfields []*FieldBuilder
	for _, f := range vs {
		newfields = append(newfields, r.Filter(f))
	}
	r.fields = newfields
	return r
}

func (b *FilteringBuilder) Filter(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	r = &FieldBuilder{name: name}
	b.fields = append(b.fields, r)
	return
}
