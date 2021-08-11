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

func (b *VRatingBuilder) BackgroundColor(v string) (r *VRatingBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VRatingBuilder) Clearable(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) CloseDelay(v int) (r *VRatingBuilder) {
	b.tag.Attr(":close-delay", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Color(v string) (r *VRatingBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VRatingBuilder) Dark(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Dense(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) EmptyIcon(v string) (r *VRatingBuilder) {
	b.tag.Attr("empty-icon", v)
	return b
}

func (b *VRatingBuilder) FullIcon(v string) (r *VRatingBuilder) {
	b.tag.Attr("full-icon", v)
	return b
}

func (b *VRatingBuilder) HalfIcon(v string) (r *VRatingBuilder) {
	b.tag.Attr("half-icon", v)
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

func (b *VRatingBuilder) IconLabel(v string) (r *VRatingBuilder) {
	b.tag.Attr("icon-label", v)
	return b
}

func (b *VRatingBuilder) Large(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":large", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Length(v int) (r *VRatingBuilder) {
	b.tag.Attr(":length", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Light(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) OpenDelay(v int) (r *VRatingBuilder) {
	b.tag.Attr(":open-delay", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Readonly(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Ripple(v interface{}) (r *VRatingBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VRatingBuilder) Size(v int) (r *VRatingBuilder) {
	b.tag.Attr(":size", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Small(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":small", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) Value(v int) (r *VRatingBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) XLarge(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":x-large", fmt.Sprint(v))
	return b
}

func (b *VRatingBuilder) XSmall(v bool) (r *VRatingBuilder) {
	b.tag.Attr(":x-small", fmt.Sprint(v))
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
