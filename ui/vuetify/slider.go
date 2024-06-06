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

func (b *VSliderBuilder) Label(v string) (r *VSliderBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VSliderBuilder) Focused(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Reverse(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
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

func (b *VSliderBuilder) Readonly(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Max(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Min(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":min", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Step(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":step", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) ThumbColor(v string) (r *VSliderBuilder) {
	b.tag.Attr("thumb-color", v)
	return b
}

func (b *VSliderBuilder) ThumbLabel(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":thumb-label", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) ThumbSize(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":thumb-size", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) ShowTicks(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":show-ticks", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Ticks(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":ticks", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) TickSize(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":tick-size", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Color(v string) (r *VSliderBuilder) {
	b.tag.Attr("color", v)
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

func (b *VSliderBuilder) TrackSize(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":track-size", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Direction(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Rounded(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Tile(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Elevation(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Ripple(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Id(v string) (r *VSliderBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VSliderBuilder) AppendIcon(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) CenterAffix(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) PrependIcon(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) HideSpinButtons(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Hint(v string) (r *VSliderBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VSliderBuilder) PersistentHint(v bool) (r *VSliderBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VSliderBuilder) Messages(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Density(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) MaxWidth(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) MinWidth(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Width(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Theme(v string) (r *VSliderBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VSliderBuilder) MaxErrors(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) Name(v string) (r *VSliderBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VSliderBuilder) Rules(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) ModelValue(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) ValidateOn(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) ValidationValue(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VSliderBuilder) HideDetails(v interface{}) (r *VSliderBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
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
