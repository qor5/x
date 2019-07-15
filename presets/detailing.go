package presets

type DetailingBuilder struct {
	fieldNames []string
	fields     []*FieldBuilder
	actions    []*ActionBuilder
}

func (b *ModelBuilder) Detailing(vs ...string) (r *DetailingBuilder) {
	r = b.detailing
	r.fieldNames = vs
	var newfields []*FieldBuilder
	for _, f := range vs {
		newfields = append(newfields, r.Field(f))
	}
	r.fields = newfields
	return r
}

func (b *DetailingBuilder) Field(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	r = &FieldBuilder{name: name}
	b.fields = append(b.fields, r)
	return
}
