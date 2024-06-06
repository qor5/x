package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSnackbarQueueBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSnackbarQueue(children ...h.HTMLComponent) (r *VSnackbarQueueBuilder) {
	r = &VSnackbarQueueBuilder{
		tag: h.Tag("v-snackbar-queue").Children(children...),
	}
	return
}

func (b *VSnackbarQueueBuilder) Activator(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Text(v string) (r *VSnackbarQueueBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VSnackbarQueueBuilder) MultiLine(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":multi-line", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) Timer(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":timer", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Timeout(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":timeout", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Vertical(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) Location(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Position(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Absolute(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) Rounded(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Tile(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) Color(v string) (r *VSnackbarQueueBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSnackbarQueueBuilder) Variant(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Theme(v string) (r *VSnackbarQueueBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VSnackbarQueueBuilder) CloseOnBack(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":close-on-back", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) Contained(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":contained", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) ContentClass(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) ContentProps(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":content-props", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Disabled(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) Opacity(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) ModelValue(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) ZIndex(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":z-index", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Target(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":target", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) ActivatorProps(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":activator-props", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) OpenOnClick(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) OpenOnHover(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) OpenOnFocus(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) CloseOnContentClick(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":close-on-content-click", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) CloseDelay(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) OpenDelay(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Height(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) MaxHeight(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) MaxWidth(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) MinHeight(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) MinWidth(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Width(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Eager(v bool) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VSnackbarQueueBuilder) LocationStrategy(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":location-strategy", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Origin(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":origin", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Offset(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":offset", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Transition(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Attach(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) Closable(v interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(":closable", h.JSONString(v))
	return b
}

func (b *VSnackbarQueueBuilder) CloseText(v string) (r *VSnackbarQueueBuilder) {
	b.tag.Attr("close-text", v)
	return b
}

func (b *VSnackbarQueueBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSnackbarQueueBuilder) Attr(vs ...interface{}) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSnackbarQueueBuilder) Children(children ...h.HTMLComponent) (r *VSnackbarQueueBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSnackbarQueueBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSnackbarQueueBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSnackbarQueueBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSnackbarQueueBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSnackbarQueueBuilder) Class(names ...string) (r *VSnackbarQueueBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSnackbarQueueBuilder) ClassIf(name string, add bool) (r *VSnackbarQueueBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSnackbarQueueBuilder) On(name string, value string) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSnackbarQueueBuilder) Bind(name string, value string) (r *VSnackbarQueueBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSnackbarQueueBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
