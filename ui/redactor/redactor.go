package redactor

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type Builder struct {
	tag *h.HTMLTagBuilder
}

type Config struct {
	Plugins []string `json:"plugins"`
}

func New() (r *Builder) {
	r = &Builder{
		tag: h.Tag("redactor"),
	}
	return
}

func (b *Builder) Value(v string) (r *Builder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *Builder) Placeholder(v string) (r *Builder) {
	b.tag.Attr(":placeholder", h.JSONString(v))
	return b
}

func (b *Builder) Config(v Config) (r *Builder) {
	b.tag.Attr(":config", h.JSONString(v))
	return b
}

func (b *Builder) RawConfig(v interface{}) (r *Builder) {
	b.tag.Attr(":config", h.JSONString(v))
	return b
}

func (b *Builder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *Builder) Attr(vs ...interface{}) (r *Builder) {
	b.tag.Attr(vs...)
	return b
}

func (b *Builder) Children(children ...h.HTMLComponent) (r *Builder) {
	b.tag.Children(children...)
	return b
}

func (b *Builder) AppendChildren(children ...h.HTMLComponent) (r *Builder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *Builder) PrependChildren(children ...h.HTMLComponent) (r *Builder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *Builder) Class(names ...string) (r *Builder) {
	b.tag.Class(names...)
	return b
}

func (b *Builder) ClassIf(name string, add bool) (r *Builder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *Builder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
