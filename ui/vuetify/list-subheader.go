package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListSubheaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListSubheader(children ...h.HTMLComponent) (r *VListSubheaderBuilder) {
	r = &VListSubheaderBuilder{
		tag: h.Tag("v-list-subheader").Children(children...),
	}
	return
}

func (b *VListSubheaderBuilder) Color(v string) (r *VListSubheaderBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListSubheaderBuilder) Inset(v bool) (r *VListSubheaderBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VListSubheaderBuilder) Sticky(v bool) (r *VListSubheaderBuilder) {
	b.tag.Attr(":sticky", fmt.Sprint(v))
	return b
}

func (b *VListSubheaderBuilder) Title(v string) (r *VListSubheaderBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VListSubheaderBuilder) Tag(v string) (r *VListSubheaderBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListSubheaderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListSubheaderBuilder) Attr(vs ...interface{}) (r *VListSubheaderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListSubheaderBuilder) Children(children ...h.HTMLComponent) (r *VListSubheaderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListSubheaderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListSubheaderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListSubheaderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListSubheaderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListSubheaderBuilder) Class(names ...string) (r *VListSubheaderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListSubheaderBuilder) ClassIf(name string, add bool) (r *VListSubheaderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListSubheaderBuilder) On(name string, value string) (r *VListSubheaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListSubheaderBuilder) Bind(name string, value string) (r *VListSubheaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListSubheaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
