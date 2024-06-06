package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VRatingBuilder struct {
	tag *h.HTMLTagBuilder
}

func VRating(children ...h.HTMLComponent) (r *VRatingBuilder) {
	r = &VRatingBuilder{
		tag: h.Tag("v-rating").Children(children...),
	}
	return
}

func (b *VRatingBuilder) Length(v interface{}) (r *VRatingBuilder) {
	b.tag.Attr(":length", h.JSONString(v))
	return b
}

func (b *VRatingBuilder) Name(v string) (r *VRatingBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VRatingBuilder) ItemAriaLabel(v string) (r *VRatingBuilder) {
	b.tag.Attr("item-aria-label", v)
	return b
}

func (b *VRatingBuilder) ActiveColor(v string) (r *VRatingBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VRatingBuilder) Color(v string) (r *VRatingBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VRatingBuilder) Clearable(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Disabled(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) EmptyIcon(v interface{}) (r *VRatingBuilder) {
	b.tag.Attr(":empty-icon", h.JSONString(v))
	return b
}

func (b *VRatingBuilder) FullIcon(v interface{}) (r *VRatingBuilder) {
	b.tag.Attr(":full-icon", h.JSONString(v))
	return b
}

func (b *VRatingBuilder) HalfIncrements(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":half-increments", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Hover(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":hover", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Readonly(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) ModelValue(v interface{}) (r *VRatingBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VRatingBuilder) ItemLabelPosition(v string) (r *VRatingBuilder) {
	b.tag.Attr("item-label-position", v)
	return b
}

func (b *VRatingBuilder) Ripple(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Density(v interface{}) (r *VRatingBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VRatingBuilder) Size(v interface{}) (r *VRatingBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VRatingBuilder) Tag(v string) (r *VRatingBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VRatingBuilder) Theme(v string) (r *VRatingBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VRatingBuilder) ItemLabels(v interface{}) (r *VRatingBuilder) {
	b.tag.Attr(":item-labels", h.JSONString(v))
	return b
}

func (b *VRatingBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VRatingBuilder) Attr(vs ...interface{}) (r *VRatingBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VRatingBuilder) Children(children ...h.HTMLComponent) (r *VRatingBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VRatingBuilder) AppendChildren(children ...h.HTMLComponent) (r *VRatingBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VRatingBuilder) PrependChildren(children ...h.HTMLComponent) (r *VRatingBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VRatingBuilder) Class(names ...string) (r *VRatingBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VRatingBuilder) ClassIf(name string, add bool) (r *VRatingBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VRatingBuilder) On(name string, value string) (r *VRatingBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VRatingBuilder) Bind(name string, value string) (r *VRatingBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VRatingBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
