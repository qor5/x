package vuetify

import (
	"context"
	"fmt"
	h "github.com/theplant/htmlgo"
)

type VListTileBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListTile(children ...h.HTMLComponent) (r *VListTileBuilder) {
	r = &VListTileBuilder{
		tag: h.Tag("v-list-tile").Children(children...),
	}
	return
}
func (b *VListTileBuilder) ActiveClass(v string) (r *VListTileBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VListTileBuilder) Append(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":append", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) Avatar(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":avatar", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) Color(v string) (r *VListTileBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListTileBuilder) Dark(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) Disabled(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) Exact(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) ExactActiveClass(v string) (r *VListTileBuilder) {
	b.tag.Attr("exact-active-class", v)
	return b
}

func (b *VListTileBuilder) Href(v string) (r *VListTileBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VListTileBuilder) Inactive(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":inactive", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) Light(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) Nuxt(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":nuxt", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) Replace(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) Ripple(v bool) (r *VListTileBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VListTileBuilder) Tag(v string) (r *VListTileBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListTileBuilder) Target(v string) (r *VListTileBuilder) {
	b.tag.Attr("target", v)
	return b
}

func (b *VListTileBuilder) To(v string) (r *VListTileBuilder) {
	b.tag.Attr("to", v)
	return b
}

func (b *VListTileBuilder) Value(v interface{}) (r *VListTileBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VListTileBuilder) Class(names ...string) (r *VListTileBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListTileBuilder) ClassIf(name string, add bool) (r *VListTileBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListTileBuilder) On(name string, value string) (r *VListTileBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListTileBuilder) Bind(name string, value string) (r *VListTileBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListTileBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
