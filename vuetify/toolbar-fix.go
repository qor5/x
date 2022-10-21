package vuetify

func (b *VToolbarBuilder) AutoHeight(v bool) (r *VToolbarBuilder) {
	if v {
		b.tag.Attr(":height", `"auto"`)
	} else {
		b.tag.Attr(":height", ``)
	}
	return b
}
