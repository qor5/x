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

func (b *VRangeSliderBuilder) Label(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VRangeSliderBuilder) Focused(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Reverse(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Id(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VRangeSliderBuilder) AppendIcon(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) CenterAffix(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) PrependIcon(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) HideSpinButtons(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Hint(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VRangeSliderBuilder) PersistentHint(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Messages(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Direction(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Density(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) MaxWidth(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) MinWidth(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Width(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Theme(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("theme", v)
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

func (b *VRangeSliderBuilder) ErrorMessages(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":error-messages", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) MaxErrors(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Name(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("name", v)
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

func (b *VRangeSliderBuilder) ModelValue(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) ValidateOn(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) ValidationValue(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) HideDetails(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Max(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Min(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":min", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Step(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":step", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) ThumbColor(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("thumb-color", v)
	return b
}

func (b *VRangeSliderBuilder) ThumbLabel(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":thumb-label", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) ThumbSize(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":thumb-size", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) ShowTicks(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":show-ticks", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Ticks(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":ticks", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) TickSize(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":tick-size", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Color(v string) (r *VRangeSliderBuilder) {
	b.tag.Attr("color", v)
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

func (b *VRangeSliderBuilder) TrackSize(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":track-size", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Rounded(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Tile(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Elevation(v interface{}) (r *VRangeSliderBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VRangeSliderBuilder) Ripple(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VRangeSliderBuilder) Strict(v bool) (r *VRangeSliderBuilder) {
	b.tag.Attr(":strict", fmt.Sprint(v))
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
