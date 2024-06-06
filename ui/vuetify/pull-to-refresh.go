package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VPullToRefreshBuilder struct {
	tag *h.HTMLTagBuilder
}

func VPullToRefresh(children ...h.HTMLComponent) (r *VPullToRefreshBuilder) {
	r = &VPullToRefreshBuilder{
		tag: h.Tag("v-pull-to-refresh").Children(children...),
	}
	return
}

func (b *VPullToRefreshBuilder) PullDownThreshold(v int) (r *VPullToRefreshBuilder) {
	b.tag.Attr(":pull-down-threshold", fmt.Sprint(v))
	return b
}

func (b *VPullToRefreshBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VPullToRefreshBuilder) Attr(vs ...interface{}) (r *VPullToRefreshBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VPullToRefreshBuilder) Children(children ...h.HTMLComponent) (r *VPullToRefreshBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VPullToRefreshBuilder) AppendChildren(children ...h.HTMLComponent) (r *VPullToRefreshBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VPullToRefreshBuilder) PrependChildren(children ...h.HTMLComponent) (r *VPullToRefreshBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VPullToRefreshBuilder) Class(names ...string) (r *VPullToRefreshBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VPullToRefreshBuilder) ClassIf(name string, add bool) (r *VPullToRefreshBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VPullToRefreshBuilder) On(name string, value string) (r *VPullToRefreshBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VPullToRefreshBuilder) Bind(name string, value string) (r *VPullToRefreshBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VPullToRefreshBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
