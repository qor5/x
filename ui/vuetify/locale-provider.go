package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VLocaleProviderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VLocaleProvider(children ...h.HTMLComponent) (r *VLocaleProviderBuilder) {
	r = &VLocaleProviderBuilder{
		tag: h.Tag("v-locale-provider").Children(children...),
	}
	return
}

func (b *VLocaleProviderBuilder) Locale(v string) (r *VLocaleProviderBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VLocaleProviderBuilder) FallbackLocale(v string) (r *VLocaleProviderBuilder) {
	b.tag.Attr("fallback-locale", v)
	return b
}

func (b *VLocaleProviderBuilder) Messages(v interface{}) (r *VLocaleProviderBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VLocaleProviderBuilder) Rtl(v bool) (r *VLocaleProviderBuilder) {
	b.tag.Attr(":rtl", fmt.Sprint(v))
	return b
}

func (b *VLocaleProviderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VLocaleProviderBuilder) Attr(vs ...interface{}) (r *VLocaleProviderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VLocaleProviderBuilder) Children(children ...h.HTMLComponent) (r *VLocaleProviderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VLocaleProviderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VLocaleProviderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VLocaleProviderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VLocaleProviderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VLocaleProviderBuilder) Class(names ...string) (r *VLocaleProviderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VLocaleProviderBuilder) ClassIf(name string, add bool) (r *VLocaleProviderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VLocaleProviderBuilder) On(name string, value string) (r *VLocaleProviderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VLocaleProviderBuilder) Bind(name string, value string) (r *VLocaleProviderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VLocaleProviderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
