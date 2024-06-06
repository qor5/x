package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSvgIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSvgIcon(children ...h.HTMLComponent) (r *VSvgIconBuilder) {
	r = &VSvgIconBuilder{
		tag: h.Tag("v-svg-icon").Children(children...),
	}
	return
}

func (b *VSvgIconBuilder) Icon(v interface{}) (r *VSvgIconBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VSvgIconBuilder) Tag(v string) (r *VSvgIconBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VSvgIconBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSvgIconBuilder) Attr(vs ...interface{}) (r *VSvgIconBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSvgIconBuilder) Children(children ...h.HTMLComponent) (r *VSvgIconBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSvgIconBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSvgIconBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSvgIconBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSvgIconBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSvgIconBuilder) Class(names ...string) (r *VSvgIconBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSvgIconBuilder) ClassIf(name string, add bool) (r *VSvgIconBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSvgIconBuilder) On(name string, value string) (r *VSvgIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSvgIconBuilder) Bind(name string, value string) (r *VSvgIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSvgIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
