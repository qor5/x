package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VMessagesBuilder struct {
	tag *h.HTMLTagBuilder
}

func VMessages(children ...h.HTMLComponent) (r *VMessagesBuilder) {
	r = &VMessagesBuilder{
		tag: h.Tag("v-messages").Children(children...),
	}
	return
}

func (b *VMessagesBuilder) Active(v bool) (r *VMessagesBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VMessagesBuilder) Color(v string) (r *VMessagesBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VMessagesBuilder) Messages(v interface{}) (r *VMessagesBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VMessagesBuilder) Transition(v interface{}) (r *VMessagesBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VMessagesBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VMessagesBuilder) Attr(vs ...interface{}) (r *VMessagesBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VMessagesBuilder) Children(children ...h.HTMLComponent) (r *VMessagesBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VMessagesBuilder) AppendChildren(children ...h.HTMLComponent) (r *VMessagesBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VMessagesBuilder) PrependChildren(children ...h.HTMLComponent) (r *VMessagesBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VMessagesBuilder) Class(names ...string) (r *VMessagesBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VMessagesBuilder) ClassIf(name string, add bool) (r *VMessagesBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VMessagesBuilder) On(name string, value string) (r *VMessagesBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VMessagesBuilder) Bind(name string, value string) (r *VMessagesBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VMessagesBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
