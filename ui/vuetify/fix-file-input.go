package vuetify

func (b *VFileInputBuilder) ErrorMessages(v ...string) (r *VFileInputBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
