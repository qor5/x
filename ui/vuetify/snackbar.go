package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSnackbarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSnackbar(children ...h.HTMLComponent) (r *VSnackbarBuilder) {
	r = &VSnackbarBuilder{
		tag: h.Tag("v-snackbar").Children(children...),
	}
	return
}

func (b *VSnackbarBuilder) Activator(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Text(v string) (r *VSnackbarBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VSnackbarBuilder) MultiLine(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":multi-line", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Timer(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":timer", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Timeout(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":timeout", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Vertical(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Location(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Position(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Absolute(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Rounded(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Tile(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Color(v string) (r *VSnackbarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSnackbarBuilder) Variant(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Theme(v string) (r *VSnackbarBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VSnackbarBuilder) CloseOnBack(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":close-on-back", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Contained(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":contained", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) ContentClass(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) ContentProps(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":content-props", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Disabled(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Opacity(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) ModelValue(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) ZIndex(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":z-index", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Target(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":target", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) ActivatorProps(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":activator-props", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) OpenOnClick(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) OpenOnHover(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) OpenOnFocus(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) CloseOnContentClick(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":close-on-content-click", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) CloseDelay(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) OpenDelay(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Height(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) MaxHeight(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) MaxWidth(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) MinHeight(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) MinWidth(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Width(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Eager(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) LocationStrategy(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":location-strategy", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Origin(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":origin", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Offset(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":offset", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Transition(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Attach(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSnackbarBuilder) Attr(vs ...interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSnackbarBuilder) Children(children ...h.HTMLComponent) (r *VSnackbarBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSnackbarBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSnackbarBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSnackbarBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSnackbarBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSnackbarBuilder) Class(names ...string) (r *VSnackbarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSnackbarBuilder) ClassIf(name string, add bool) (r *VSnackbarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSnackbarBuilder) On(name string, value string) (r *VSnackbarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSnackbarBuilder) Bind(name string, value string) (r *VSnackbarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSnackbarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
