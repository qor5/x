package vuetify

func (b *VTextFieldBuilder) FieldName(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VTextareaBuilder) FieldName(v string) (r *VTextareaBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VCheckboxBuilder) FieldName(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VRadioGroupBuilder) FieldName(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VSwitchBuilder) FieldName(v string) (r *VSwitchBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VSliderBuilder) FieldName(v string) (r *VSliderBuilder) {
	b.tag.Attr("field-name", v)
	return b
}
