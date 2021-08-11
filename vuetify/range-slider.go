package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VRangeSliderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VRangeSlider(children ...h.HTMLComponent) (r *VRangeSliderBuilder) {
	r = &VRangeSliderBuilder{
		tag: h.Tag("v-range-slider").Children(children...),
	}
	return
}

func (b *VRangeSliderBuilder) AppendIcon(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VRangeSliderBuilder) BackgroundColor(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VRangeSliderBuilder) Color(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VRangeSliderBuilder) Dark(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Dense(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Disabled(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Error(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) ErrorCount(v int) (r *VRangeSliderBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) ErrorMessages(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("error-messages", v)
	return b
}

func (b *VRangeSliderBuilder) Height(v int) (r *VRangeSliderBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) HideDetails(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Hint(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VRangeSliderBuilder) Id(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VRangeSliderBuilder) InverseLabel(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":inverse-label", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Label(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VRangeSliderBuilder) Light(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) LoaderHeight(v int) (r *VRangeSliderBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Loading(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Max(v int) (r *VRangeSliderBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Messages(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VRangeSliderBuilder) Min(v int) (r *VRangeSliderBuilder) {
	b.tag.Attr(":min", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) PersistentHint(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) PrependIcon(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VRangeSliderBuilder) Readonly(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Rules(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Step(v int) (r *VRangeSliderBuilder) {
	b.tag.Attr(":step", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Success(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) SuccessMessages(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VRangeSliderBuilder) ThumbColor(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("thumb-color", v)
	return b
}

func (b *VRangeSliderBuilder) ThumbLabel(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":thumb-label", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) ThumbSize(v int) (r *VRangeSliderBuilder) {
	b.tag.Attr(":thumb-size", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) TickLabels(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":tick-labels", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) TickSize(v int) (r *VRangeSliderBuilder) {
	b.tag.Attr(":tick-size", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Ticks(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":ticks", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) TrackColor(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("track-color", v)
	return b
}

func (b *VRangeSliderBuilder) TrackFillColor(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("track-fill-color", v)
	return b
}

func (b *VRangeSliderBuilder) ValidateOnBlur(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Value(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Vertical(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VRangeSliderBuilder) Attr(vs ...interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VRangeSliderBuilder) Children(children ...h.HTMLComponent) (r *VRangeSliderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VRangeSliderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VRangeSliderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VRangeSliderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VRangeSliderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VRangeSliderBuilder) Class(names ...string) (r *VRangeSliderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VRangeSliderBuilder) ClassIf(name string, add bool) (r *VRangeSliderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VRangeSliderBuilder) On(name string, value string) (r *VRangeSliderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VRangeSliderBuilder) Bind(name string, value string) (r *VRangeSliderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VRangeSliderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
