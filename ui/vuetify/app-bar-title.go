package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAppBarTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAppBarTitle(children ...h.HTMLComponent) (r *VAppBarTitleBuilder) {
	r = &VAppBarTitleBuilder{
		tag: h.Tag("v-app-bar-title").Children(children...),
	}
	return
}

func (b *VAppBarTitleBuilder) Text(v string) (r *VAppBarTitleBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VAppBarTitleBuilder) Tag(v string) (r *VAppBarTitleBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VAppBarTitleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAppBarTitleBuilder) Attr(vs ...interface{}) (r *VAppBarTitleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAppBarTitleBuilder) Children(children ...h.HTMLComponent) (r *VAppBarTitleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAppBarTitleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAppBarTitleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAppBarTitleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAppBarTitleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAppBarTitleBuilder) Class(names ...string) (r *VAppBarTitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAppBarTitleBuilder) ClassIf(name string, add bool) (r *VAppBarTitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAppBarTitleBuilder) On(name string, value string) (r *VAppBarTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAppBarTitleBuilder) Bind(name string, value string) (r *VAppBarTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAppBarTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
