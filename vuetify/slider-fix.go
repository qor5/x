package vuetify

import h "github.com/theplant/htmlgo"

func VSlider(children ...h.HTMLComponent) (r *VSliderBuilder) {
	r = &VSliderBuilder{
		tag: h.Tag("vw-slider").Children(children...),
	}
	return
}

func (b *VSliderBuilder) ErrorMessages(v ...string) (r *VSliderBuilder) {
	setErrorMessages(b.tag, v)
	return b
}

func (b *VSliderBuilder) FieldName(v string) (r *VSliderBuilder) {
	b.tag.Attr("field-name", v)
	return b
}
