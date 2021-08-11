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

func (b *VNavigationDrawerBuilder) Absolute(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) App(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Bottom(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Clipped(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":clipped", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Color(v string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VNavigationDrawerBuilder) Dark(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
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

func (b *VNavigationDrawerBuilder) Fixed(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Floating(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":floating", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Height(v int) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) HideOverlay(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":hide-overlay", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Light(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) MiniVariant(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":mini-variant", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) MiniVariantWidth(v int) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":mini-variant-width", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) MobileBreakpoint(v int) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":mobile-breakpoint", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) OverlayColor(v string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr("overlay-color", v)
	return b
}

func (b *VNavigationDrawerBuilder) OverlayOpacity(v int) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":overlay-opacity", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Permanent(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":permanent", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Right(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Src(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":src", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Stateless(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":stateless", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Tag(v string) (r *VNavigationDrawerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VNavigationDrawerBuilder) Temporary(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":temporary", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Touchless(v bool) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":touchless", fmt.Sprint(v))
	return b
}

func (b *VNavigationDrawerBuilder) Value(v interface{}) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VNavigationDrawerBuilder) Width(v int) (r *VNavigationDrawerBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
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
