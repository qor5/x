package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VNavigationDrawerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VNavigationDrawer(children ...h.HTMLComponent) (r *VNavigationDrawerBuilder) {
	r = &VNavigationDrawerBuilder{
		tag: h.Tag("v-navigation-drawer").Children(children...),
	}
	return
}

func (b *VNavigationDrawerBuilder) Image(v string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr("image", v)
	return b
}

func (b *VNavigationDrawerBuilder) Color(v string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VNavigationDrawerBuilder) DisableResizeWatcher(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":disable-resize-watcher", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) DisableRouteWatcher(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":disable-route-watcher", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) ExpandOnHover(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":expand-on-hover", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Floating(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":floating", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) ModelValue(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Permanent(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":permanent", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Rail(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":rail", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) RailWidth(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":rail-width", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Scrim(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":scrim", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Temporary(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":temporary", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Persistent(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Touchless(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":touchless", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Width(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Location(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Sticky(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":sticky", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Border(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) CloseDelay(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) OpenDelay(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Mobile(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) MobileBreakpoint(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Elevation(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Name(v string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VNavigationDrawerBuilder) Order(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":order", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Absolute(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Rounded(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Tile(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Tag(v string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VNavigationDrawerBuilder) Theme(v string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VNavigationDrawerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VNavigationDrawerBuilder) Attr(vs ...interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VNavigationDrawerBuilder) Children(children ...h.HTMLComponent) (r *VNavigationDrawerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VNavigationDrawerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VNavigationDrawerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VNavigationDrawerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VNavigationDrawerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VNavigationDrawerBuilder) Class(names ...string) (r *VNavigationDrawerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VNavigationDrawerBuilder) ClassIf(name string, add bool) (r *VNavigationDrawerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VNavigationDrawerBuilder) On(name string, value string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VNavigationDrawerBuilder) Bind(name string, value string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VNavigationDrawerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
