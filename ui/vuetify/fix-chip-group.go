package vuetify

func (b *VChipGroupBuilder) Value(v interface{}) (r *VChipGroupBuilder) {
	b.ModelValue(v)
	return b
}
