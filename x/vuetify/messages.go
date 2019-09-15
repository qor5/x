package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VMessagesBuilder struct {
	tag *h.HTMLTagBuilder
}

func VMessages() (r *VMessagesBuilder) {
	r = &VMessagesBuilder{
		tag: h.Tag("v-messages"),
	}
	return
}

func (b *VMessagesBuilder) Color(v string) (r *VMessagesBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VMessagesBuilder) Dark(v bool) (r *VMessagesBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VMessagesBuilder) Light(v bool) (r *VMessagesBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VMessagesBuilder) Value(v []string) (r *VMessagesBuilder) {
	b.tag.Attr(":value", v)
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
