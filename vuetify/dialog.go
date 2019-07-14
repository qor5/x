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

func (b *VDialogBuilder) Attach(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":attach", v)
	return b
}

func (b *VDialogBuilder) ContentClass(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":content-class", v)
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

func (b *VDialogBuilder) FullWidth(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
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

func (b *VDialogBuilder) Lazy(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":lazy", fmt.Sprint(v))
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

func (b *VDialogBuilder) Origin(v string) (r *VDialogBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VDialogBuilder) Persistent(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) ReturnValue(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":return-value", v)
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
	b.tag.Attr(":value", v)
	return b
}

func (b *VDialogBuilder) Width(v string) (r *VDialogBuilder) {
	b.tag.Attr("width", v)
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
