package vuetify

func (b *VListItemBuilder) Slot(v string) (r *VListItemBuilder) {
	b.tag.Attr("slot", v)
	return b
}
