package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBottomNavBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBottomNav(children ...h.HTMLComponent) (r *VBottomNavBuilder) {
	r = &VBottomNavBuilder{
		tag: h.Tag("v-bottom-nav").Children(children...),
	}
	return
}

func (b *VBottomNavBuilder) Absolute(v bool) (r *VBottomNavBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VBottomNavBuilder) Active(v int) (r *VBottomNavBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VBottomNavBuilder) App(v bool) (r *VBottomNavBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VBottomNavBuilder) Color(v string) (r *VBottomNavBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBottomNavBuilder) Dark(v bool) (r *VBottomNavBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VBottomNavBuilder) Fixed(v bool) (r *VBottomNavBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VBottomNavBuilder) Height(v int) (r *VBottomNavBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VBottomNavBuilder) Light(v bool) (r *VBottomNavBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VBottomNavBuilder) Mandatory(v bool) (r *VBottomNavBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VBottomNavBuilder) Shift(v bool) (r *VBottomNavBuilder) {
	b.tag.Attr(":shift", fmt.Sprint(v))
	return b
}

func (b *VBottomNavBuilder) Value(v interface{}) (r *VBottomNavBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VBottomNavBuilder) Class(names ...string) (r *VBottomNavBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBottomNavBuilder) ClassIf(name string, add bool) (r *VBottomNavBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBottomNavBuilder) On(name string, value string) (r *VBottomNavBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBottomNavBuilder) Bind(name string, value string) (r *VBottomNavBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBottomNavBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
