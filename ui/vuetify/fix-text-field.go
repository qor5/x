package vuetify

func (b *VTextFieldBuilder) ErrorMessages(v ...string) (r *VTextFieldBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
