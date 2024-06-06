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

func (b *VBottomNavigationBuilder) BaseColor(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VBottomNavigationBuilder) BgColor(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VBottomNavigationBuilder) Color(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBottomNavigationBuilder) Grow(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":grow", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Mode(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VBottomNavigationBuilder) Height(v interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VBottomNavigationBuilder) Active(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Border(v interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VBottomNavigationBuilder) Density(v interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VBottomNavigationBuilder) Elevation(v interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VBottomNavigationBuilder) Rounded(v interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VBottomNavigationBuilder) Tile(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Name(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VBottomNavigationBuilder) Order(v interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":order", h.JSONString(v))
	return b
}

func (b *VBottomNavigationBuilder) Absolute(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Tag(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBottomNavigationBuilder) ModelValue(v interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VBottomNavigationBuilder) Multiple(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Max(v int) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) SelectedClass(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VBottomNavigationBuilder) Disabled(v bool) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBottomNavigationBuilder) Mandatory(v interface{}) (r *VBottomNavigationBuilder) {
	b.tag.Attr(":mandatory", h.JSONString(v))
	return b
}

func (b *VBottomNavigationBuilder) Theme(v string) (r *VBottomNavigationBuilder) {
	b.tag.Attr("theme", v)
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
