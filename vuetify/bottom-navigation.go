package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBottomNavigationBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBottomNavigation(children ...h.HTMLComponent) (r *VBottomNavigationBuilder) {
	r = &VBottomNavigationBuilder{
		tag: h.Tag("v-bottom-navigation").Children(children...),
	}
	return
}

func (b *VBottomNavigationBuilder) Absolute(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) ActiveClass(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VBottomNavigationBuilder) App(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) BackgroundColor(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VBottomNavigationBuilder) Color(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBottomNavigationBuilder) Dark(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Fixed(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Grow(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":grow", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Height(v int) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) HideOnScroll(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":hide-on-scroll", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Horizontal(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":horizontal", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) InputValue(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":input-value", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Light(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Mandatory(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) MaxHeight(v int) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) MaxWidth(v int) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) MinHeight(v int) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) MinWidth(v int) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) ScrollTarget(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("scroll-target", v)
	return b
}

func (b *VBottomNavigationBuilder) ScrollThreshold(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("scroll-threshold", v)
	return b
}

func (b *VBottomNavigationBuilder) Shift(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":shift", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Tag(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBottomNavigationBuilder) Value(v interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VBottomNavigationBuilder) Width(v int) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBottomNavigationBuilder) Attr(vs ...interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBottomNavigationBuilder) Children(children ...h.HTMLComponent) (r *VBottomNavigationBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBottomNavigationBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBottomNavigationBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBottomNavigationBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBottomNavigationBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBottomNavigationBuilder) Class(names ...string) (r *VBottomNavigationBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBottomNavigationBuilder) ClassIf(name string, add bool) (r *VBottomNavigationBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBottomNavigationBuilder) On(name string, value string) (r *VBottomNavigationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBottomNavigationBuilder) Bind(name string, value string) (r *VBottomNavigationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBottomNavigationBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
