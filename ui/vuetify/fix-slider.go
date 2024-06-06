package vuetify

func (b *VSliderBuilder) ErrorMessages(v ...string) (r *VSliderBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
