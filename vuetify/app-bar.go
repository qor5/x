package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAppBarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAppBar(children ...h.HTMLComponent) (r *VAppBarBuilder) {
	r = &VAppBarBuilder{
		tag: h.Tag("v-app-bar").Children(children...),
	}
	return
}

func (b *VAppBarBuilder) Absolute(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) App(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Bottom(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) ClippedLeft(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":clipped-left", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) ClippedRight(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":clipped-right", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Collapse(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":collapse", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) CollapseOnScroll(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":collapse-on-scroll", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Color(v string) (r *VAppBarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAppBarBuilder) Dark(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Dense(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) ElevateOnScroll(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":elevate-on-scroll", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Elevation(v int) (r *VAppBarBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Extended(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":extended", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) ExtensionHeight(v int) (r *VAppBarBuilder) {
	b.tag.Attr(":extension-height", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) FadeImgOnScroll(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":fade-img-on-scroll", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Fixed(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Flat(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Floating(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":floating", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Height(v int) (r *VAppBarBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) HideOnScroll(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":hide-on-scroll", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) InvertedScroll(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":inverted-scroll", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Light(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) MaxHeight(v int) (r *VAppBarBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) MaxWidth(v int) (r *VAppBarBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) MinHeight(v int) (r *VAppBarBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) MinWidth(v int) (r *VAppBarBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Outlined(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Prominent(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":prominent", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Rounded(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) ScrollOffScreen(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":scroll-off-screen", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) ScrollTarget(v string) (r *VAppBarBuilder) {
	b.tag.Attr("scroll-target", v)
	return b
}

func (b *VAppBarBuilder) ScrollThreshold(v string) (r *VAppBarBuilder) {
	b.tag.Attr("scroll-threshold", v)
	return b
}

func (b *VAppBarBuilder) Shaped(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Short(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":short", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) ShrinkOnScroll(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":shrink-on-scroll", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Src(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":src", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) Tag(v string) (r *VAppBarBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VAppBarBuilder) Tile(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Value(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Width(v int) (r *VAppBarBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAppBarBuilder) Attr(vs ...interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAppBarBuilder) Children(children ...h.HTMLComponent) (r *VAppBarBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAppBarBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAppBarBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAppBarBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAppBarBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAppBarBuilder) Class(names ...string) (r *VAppBarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAppBarBuilder) ClassIf(name string, add bool) (r *VAppBarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAppBarBuilder) On(name string, value string) (r *VAppBarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAppBarBuilder) Bind(name string, value string) (r *VAppBarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAppBarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
