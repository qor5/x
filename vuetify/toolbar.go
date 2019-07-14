package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VToolbarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VToolbar(children ...h.HTMLComponent) (r *VToolbarBuilder) {
	r = &VToolbarBuilder{
		tag: h.Tag("v-toolbar").Children(children...),
	}
	return
}

func (b *VToolbarBuilder) Absolute(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) App(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Card(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":card", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) ClippedLeft(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":clipped-left", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) ClippedRight(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":clipped-right", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Color(v string) (r *VToolbarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VToolbarBuilder) Dark(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Dense(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Extended(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":extended", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) ExtensionHeight(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":extension-height", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Fixed(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Flat(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Floating(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":floating", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Height(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) InvertedScroll(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":inverted-scroll", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Light(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) ManualScroll(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":manual-scroll", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Prominent(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":prominent", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) ScrollOffScreen(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":scroll-off-screen", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) ScrollTarget(v string) (r *VToolbarBuilder) {
	b.tag.Attr("scroll-target", v)
	return b
}

func (b *VToolbarBuilder) ScrollThreshold(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":scroll-threshold", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) ScrollToolbarOffScreen(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":scroll-toolbar-off-screen", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Tabs(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":tabs", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Class(names ...string) (r *VToolbarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VToolbarBuilder) ClassIf(name string, add bool) (r *VToolbarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VToolbarBuilder) On(name string, value string) (r *VToolbarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VToolbarBuilder) Bind(name string, value string) (r *VToolbarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VToolbarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
