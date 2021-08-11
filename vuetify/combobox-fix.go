package vuetify

func (b *VComboboxBuilder) FieldName(v string) (r *VComboboxBuilder) {
	b.tag.Attr("field-name", v)
	return b
}
