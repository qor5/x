package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VInfiniteScrollBuilder struct {
	tag *h.HTMLTagBuilder
}

func VInfiniteScroll(children ...h.HTMLComponent) (r *VInfiniteScrollBuilder) {
	r = &VInfiniteScrollBuilder{
		tag: h.Tag("v-infinite-scroll").Children(children...),
	}
	return
}

func (b *VInfiniteScrollBuilder) Color(v string) (r *VInfiniteScrollBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VInfiniteScrollBuilder) Direction(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) Side(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":side", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) Mode(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":mode", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) Margin(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":margin", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) LoadMoreText(v string) (r *VInfiniteScrollBuilder) {
	b.tag.Attr("load-more-text", v)
	return b
}

func (b *VInfiniteScrollBuilder) EmptyText(v string) (r *VInfiniteScrollBuilder) {
	b.tag.Attr("empty-text", v)
	return b
}

func (b *VInfiniteScrollBuilder) Height(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) MaxHeight(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) MaxWidth(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) MinHeight(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) MinWidth(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) Width(v interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VInfiniteScrollBuilder) Tag(v string) (r *VInfiniteScrollBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VInfiniteScrollBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VInfiniteScrollBuilder) Attr(vs ...interface{}) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VInfiniteScrollBuilder) Children(children ...h.HTMLComponent) (r *VInfiniteScrollBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VInfiniteScrollBuilder) AppendChildren(children ...h.HTMLComponent) (r *VInfiniteScrollBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VInfiniteScrollBuilder) PrependChildren(children ...h.HTMLComponent) (r *VInfiniteScrollBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VInfiniteScrollBuilder) Class(names ...string) (r *VInfiniteScrollBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VInfiniteScrollBuilder) ClassIf(name string, add bool) (r *VInfiniteScrollBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VInfiniteScrollBuilder) On(name string, value string) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VInfiniteScrollBuilder) Bind(name string, value string) (r *VInfiniteScrollBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VInfiniteScrollBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
