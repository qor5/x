package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTab(children ...h.HTMLComponent) (r *VTabBuilder) {
	r = &VTabBuilder{
		tag: h.Tag("v-tab").Children(children...),
	}
	return
}

func (b *VTabBuilder) ActiveClass(v string) (r *VTabBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VTabBuilder) Append(v bool) (r *VTabBuilder) {
	b.tag.Attr(":append", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Dark(v bool) (r *VTabBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Disabled(v bool) (r *VTabBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Exact(v bool) (r *VTabBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) ExactActiveClass(v string) (r *VTabBuilder) {
	b.tag.Attr("exact-active-class", v)
	return b
}

func (b *VTabBuilder) ExactPath(v bool) (r *VTabBuilder) {
	b.tag.Attr(":exact-path", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Href(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":href", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Light(v bool) (r *VTabBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Link(v bool) (r *VTabBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Nuxt(v bool) (r *VTabBuilder) {
	b.tag.Attr(":nuxt", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Replace(v bool) (r *VTabBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Ripple(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Tag(v string) (r *VTabBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTabBuilder) Target(v string) (r *VTabBuilder) {
	b.tag.Attr("target", v)
	return b
}

func (b *VTabBuilder) To(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VTabBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTabBuilder) Attr(vs ...interface{}) (r *VTabBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTabBuilder) Children(children ...h.HTMLComponent) (r *VTabBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTabBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTabBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTabBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTabBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTabBuilder) Class(names ...string) (r *VTabBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabBuilder) ClassIf(name string, add bool) (r *VTabBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabBuilder) On(name string, value string) (r *VTabBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabBuilder) Bind(name string, value string) (r *VTabBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
