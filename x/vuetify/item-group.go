package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VItemGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VItemGroup(children ...h.HTMLComponent) (r *VItemGroupBuilder) {
	r = &VItemGroupBuilder{
		tag: h.Tag("v-item-group").Children(children...),
	}
	return
}

func (b *VItemGroupBuilder) ActiveClass(v string) (r *VItemGroupBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VItemGroupBuilder) Dark(v bool) (r *VItemGroupBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VItemGroupBuilder) Light(v bool) (r *VItemGroupBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VItemGroupBuilder) Mandatory(v bool) (r *VItemGroupBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VItemGroupBuilder) Max(v int) (r *VItemGroupBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VItemGroupBuilder) Multiple(v bool) (r *VItemGroupBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VItemGroupBuilder) Value(v interface{}) (r *VItemGroupBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VItemGroupBuilder) Class(names ...string) (r *VItemGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VItemGroupBuilder) ClassIf(name string, add bool) (r *VItemGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VItemGroupBuilder) On(name string, value string) (r *VItemGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VItemGroupBuilder) Bind(name string, value string) (r *VItemGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VItemGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
