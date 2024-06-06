package vuetify

func (b *VRadioGroupBuilder) ErrorMessages(v ...string) (r *VRadioGroupBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VRadioGroupBuilder) Value(v interface{}) (r *VRadioGroupBuilder) {
	b.ModelValue(v)
	return b
}
