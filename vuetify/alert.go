package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAlertBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAlert(children ...h.HTMLComponent) (r *VAlertBuilder) {
	r = &VAlertBuilder{
		tag: h.Tag("v-alert").Children(children...),
	}
	return
}

func (b *VAlertBuilder) Color(v string) (r *VAlertBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAlertBuilder) Dismissible(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":dismissible", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Icon(v string) (r *VAlertBuilder) {
	b.tag.Attr("icon", v)
	return b
}

func (b *VAlertBuilder) Mode(v string) (r *VAlertBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VAlertBuilder) Origin(v string) (r *VAlertBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VAlertBuilder) Outline(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":outline", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Transition(v string) (r *VAlertBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VAlertBuilder) Type(v string) (r *VAlertBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VAlertBuilder) Value(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Class(names ...string) (r *VAlertBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAlertBuilder) ClassIf(name string, add bool) (r *VAlertBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAlertBuilder) On(name string, value string) (r *VAlertBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAlertBuilder) Bind(name string, value string) (r *VAlertBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAlertBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
