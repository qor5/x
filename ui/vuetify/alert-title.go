package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAlertTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAlertTitle(children ...h.HTMLComponent) (r *VAlertTitleBuilder) {
	r = &VAlertTitleBuilder{
		tag: h.Tag("v-alert-title").Children(children...),
	}
	return
}

func (b *VAlertTitleBuilder) Tag(v string) (r *VAlertTitleBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VAlertTitleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAlertTitleBuilder) Attr(vs ...interface{}) (r *VAlertTitleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAlertTitleBuilder) Children(children ...h.HTMLComponent) (r *VAlertTitleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAlertTitleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAlertTitleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAlertTitleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAlertTitleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAlertTitleBuilder) Class(names ...string) (r *VAlertTitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAlertTitleBuilder) ClassIf(name string, add bool) (r *VAlertTitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAlertTitleBuilder) On(name string, value string) (r *VAlertTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAlertTitleBuilder) Bind(name string, value string) (r *VAlertTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAlertTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
