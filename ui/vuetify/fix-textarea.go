package vuetify

func (b *VTextareaBuilder) ErrorMessages(v ...string) (r *VTextareaBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
