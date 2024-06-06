package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VThemeProviderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VThemeProvider(children ...h.HTMLComponent) (r *VThemeProviderBuilder) {
	r = &VThemeProviderBuilder{
		tag: h.Tag("v-theme-provider").Children(children...),
	}
	return
}

func (b *VThemeProviderBuilder) WithBackground(v bool) (r *VThemeProviderBuilder) {
	b.tag.Attr(":with-background", fmt.Sprint(v))
	return b
}

func (b *VThemeProviderBuilder) Theme(v string) (r *VThemeProviderBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VThemeProviderBuilder) Tag(v string) (r *VThemeProviderBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VThemeProviderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VThemeProviderBuilder) Attr(vs ...interface{}) (r *VThemeProviderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VThemeProviderBuilder) Children(children ...h.HTMLComponent) (r *VThemeProviderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VThemeProviderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VThemeProviderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VThemeProviderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VThemeProviderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VThemeProviderBuilder) Class(names ...string) (r *VThemeProviderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VThemeProviderBuilder) ClassIf(name string, add bool) (r *VThemeProviderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VThemeProviderBuilder) On(name string, value string) (r *VThemeProviderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VThemeProviderBuilder) Bind(name string, value string) (r *VThemeProviderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VThemeProviderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
