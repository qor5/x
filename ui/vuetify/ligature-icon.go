package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VLigatureIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VLigatureIcon(children ...h.HTMLComponent) (r *VLigatureIconBuilder) {
	r = &VLigatureIconBuilder{
		tag: h.Tag("v-ligature-icon").Children(children...),
	}
	return
}

func (b *VLigatureIconBuilder) Icon(v interface{}) (r *VLigatureIconBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VLigatureIconBuilder) Tag(v string) (r *VLigatureIconBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VLigatureIconBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VLigatureIconBuilder) Attr(vs ...interface{}) (r *VLigatureIconBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VLigatureIconBuilder) Children(children ...h.HTMLComponent) (r *VLigatureIconBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VLigatureIconBuilder) AppendChildren(children ...h.HTMLComponent) (r *VLigatureIconBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VLigatureIconBuilder) PrependChildren(children ...h.HTMLComponent) (r *VLigatureIconBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VLigatureIconBuilder) Class(names ...string) (r *VLigatureIconBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VLigatureIconBuilder) ClassIf(name string, add bool) (r *VLigatureIconBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VLigatureIconBuilder) On(name string, value string) (r *VLigatureIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VLigatureIconBuilder) Bind(name string, value string) (r *VLigatureIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VLigatureIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
