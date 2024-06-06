package vuetify

func (b *VSwitchBuilder) ErrorMessages(v ...string) (r *VSwitchBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
