package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSliderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSlider(children ...h.HTMLComponent) (r *VSliderBuilder) {
	r = &VSliderBuilder{
		tag: h.Tag("v-slider").Children(children...),
	}
	return
}

func (b *VSliderBuilder) AppendIcon(v string) (r *VSliderBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VSliderBuilder) BackgroundColor(v string) (r *VSliderBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VSliderBuilder) Color(v string) (r *VSliderBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSliderBuilder) Dark(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Dense(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Disabled(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Error(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) ErrorCount(v int) (r *VSliderBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Height(v int) (r *VSliderBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) HideDetails(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Hint(v string) (r *VSliderBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VSliderBuilder) Id(v string) (r *VSliderBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VSliderBuilder) InverseLabel(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":inverse-label", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Label(v string) (r *VSliderBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VSliderBuilder) Light(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) LoaderHeight(v int) (r *VSliderBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Loading(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Max(v int) (r *VSliderBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Messages(v string) (r *VSliderBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VSliderBuilder) Min(v int) (r *VSliderBuilder) {
	b.tag.Attr(":min", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) PersistentHint(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) PrependIcon(v string) (r *VSliderBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VSliderBuilder) Readonly(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Rules(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Step(v int) (r *VSliderBuilder) {
	b.tag.Attr(":step", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Success(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) SuccessMessages(v string) (r *VSliderBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VSliderBuilder) ThumbColor(v string) (r *VSliderBuilder) {
	b.tag.Attr("thumb-color", v)
	return b
}

func (b *VSliderBuilder) ThumbLabel(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":thumb-label", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) ThumbSize(v int) (r *VSliderBuilder) {
	b.tag.Attr(":thumb-size", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) TickLabels(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":tick-labels", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) TickSize(v int) (r *VSliderBuilder) {
	b.tag.Attr(":tick-size", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Ticks(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":ticks", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) TrackColor(v string) (r *VSliderBuilder) {
	b.tag.Attr("track-color", v)
	return b
}

func (b *VSliderBuilder) TrackFillColor(v string) (r *VSliderBuilder) {
	b.tag.Attr("track-fill-color", v)
	return b
}

func (b *VSliderBuilder) ValidateOnBlur(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Value(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Vertical(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSliderBuilder) Attr(vs ...interface{}) (r *VSliderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSliderBuilder) Children(children ...h.HTMLComponent) (r *VSliderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSliderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSliderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSliderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSliderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSliderBuilder) Class(names ...string) (r *VSliderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSliderBuilder) ClassIf(name string, add bool) (r *VSliderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSliderBuilder) On(name string, value string) (r *VSliderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSliderBuilder) Bind(name string, value string) (r *VSliderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSliderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
