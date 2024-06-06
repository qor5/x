package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDefaultsProviderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDefaultsProvider(children ...h.HTMLComponent) (r *VDefaultsProviderBuilder) {
	r = &VDefaultsProviderBuilder{
		tag: h.Tag("v-defaults-provider").Children(children...),
	}
	return
}

func (b *VDefaultsProviderBuilder) Disabled(v bool) (r *VDefaultsProviderBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDefaultsProviderBuilder) Reset(v interface{}) (r *VDefaultsProviderBuilder) {
	b.tag.Attr(":reset", h.JSONString(v))
	return b
}

func (b *VDefaultsProviderBuilder) Root(v interface{}) (r *VDefaultsProviderBuilder) {
	b.tag.Attr(":root", h.JSONString(v))
	return b
}

func (b *VDefaultsProviderBuilder) Scoped(v bool) (r *VDefaultsProviderBuilder) {
	b.tag.Attr(":scoped", fmt.Sprint(v))
	return b
}

func (b *VDefaultsProviderBuilder) Defaults(v interface{}) (r *VDefaultsProviderBuilder) {
	b.tag.Attr(":defaults", h.JSONString(v))
	return b
}

func (b *VDefaultsProviderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDefaultsProviderBuilder) Attr(vs ...interface{}) (r *VDefaultsProviderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDefaultsProviderBuilder) Children(children ...h.HTMLComponent) (r *VDefaultsProviderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDefaultsProviderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDefaultsProviderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDefaultsProviderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDefaultsProviderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDefaultsProviderBuilder) Class(names ...string) (r *VDefaultsProviderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDefaultsProviderBuilder) ClassIf(name string, add bool) (r *VDefaultsProviderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDefaultsProviderBuilder) On(name string, value string) (r *VDefaultsProviderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDefaultsProviderBuilder) Bind(name string, value string) (r *VDefaultsProviderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDefaultsProviderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
