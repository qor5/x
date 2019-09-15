package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCounterBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCounter() (r *VCounterBuilder) {
	r = &VCounterBuilder{
		tag: h.Tag("v-counter"),
	}
	return
}

func (b *VCounterBuilder) Dark(v bool) (r *VCounterBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VCounterBuilder) Light(v bool) (r *VCounterBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VCounterBuilder) Max(v int) (r *VCounterBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VCounterBuilder) Value(v int) (r *VCounterBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VCounterBuilder) Class(names ...string) (r *VCounterBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCounterBuilder) ClassIf(name string, add bool) (r *VCounterBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCounterBuilder) On(name string, value string) (r *VCounterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCounterBuilder) Bind(name string, value string) (r *VCounterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCounterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
