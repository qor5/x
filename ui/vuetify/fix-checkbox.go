package vuetify

func (b *VCheckboxBuilder) ErrorMessages(v ...string) (r *VCheckboxBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
