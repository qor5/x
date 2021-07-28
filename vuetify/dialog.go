package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDialogBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDialog(children ...h.HTMLComponent) (r *VDialogBuilder) {
	r = &VDialogBuilder{
		tag: h.Tag("v-dialog").Children(children...),
	}
	return
}

func (b *VDialogBuilder) Activator(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Attach(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) CloseDelay(v int) (r *VDialogBuilder) {
	b.tag.Attr(":close-delay", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) ContentClass(v string) (r *VDialogBuilder) {
	b.tag.Attr("content-class", v)
	return b
}

func (b *VDialogBuilder) Dark(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Disabled(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Eager(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Fullscreen(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":fullscreen", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) HideOverlay(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":hide-overlay", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) InternalActivator(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":internal-activator", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Light(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) MaxWidth(v string) (r *VDialogBuilder) {
	b.tag.Attr("max-width", v)
	return b
}

func (b *VDialogBuilder) NoClickAnimation(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":no-click-animation", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) OpenDelay(v int) (r *VDialogBuilder) {
	b.tag.Attr(":open-delay", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) OpenOnFocus(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) OpenOnHover(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Origin(v string) (r *VDialogBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VDialogBuilder) OverlayColor(v string) (r *VDialogBuilder) {
	b.tag.Attr("overlay-color", v)
	return b
}

func (b *VDialogBuilder) OverlayOpacity(v int) (r *VDialogBuilder) {
	b.tag.Attr(":overlay-opacity", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Persistent(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) RetainFocus(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":retain-focus", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) ReturnValue(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":return-value", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Scrollable(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Transition(v string) (r *VDialogBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VDialogBuilder) Value(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Width(v string) (r *VDialogBuilder) {
	b.tag.Attr("width", v)
	return b
}

func (b *VDialogBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDialogBuilder) Attr(vs ...interface{}) (r *VDialogBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDialogBuilder) Children(children ...h.HTMLComponent) (r *VDialogBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDialogBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDialogBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDialogBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDialogBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDialogBuilder) Class(names ...string) (r *VDialogBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDialogBuilder) ClassIf(name string, add bool) (r *VDialogBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDialogBuilder) On(name string, value string) (r *VDialogBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDialogBuilder) Bind(name string, value string) (r *VDialogBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDialogBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
