package branoverlay

import (
	"context"
	"fmt"

	. "github.com/theplant/htmlgo"
)

type PopoverBuilder struct {
	children []HTMLComponent

	overlay []HTMLComponent
	tag     *HTMLTagBuilder
}

func Popover(children ...HTMLComponent) (r *PopoverBuilder) {
	r = &PopoverBuilder{
		tag: Tag("bran-popover"),
	}
	r.Placement("top")
	r.Trigger("click")
	r.children = children
	return
}

func (b *PopoverBuilder) Overlay(vs ...HTMLComponent) (r *PopoverBuilder) {
	b.overlay = vs
	return b
}

func (b *PopoverBuilder) DefaultVisible(v bool) (r *PopoverBuilder) {
	if v {
		b.tag.Attr(":default-visible", fmt.Sprint(v))
	}
	return b
}

func (b *PopoverBuilder) PrefixClass(v string) (r *PopoverBuilder) {
	b.tag.Attr(":prefix-cls", v)
	return b
}

func (b *PopoverBuilder) OverlayClassName(v string) (r *PopoverBuilder) {
	b.tag.Attr("overlay-class-name", v)
	return b
}

func (b *PopoverBuilder) Placement(v string) (r *PopoverBuilder) {
	b.tag.Attr("placement", v)
	return b
}

func (b *PopoverBuilder) Trigger(v string) (r *PopoverBuilder) {
	b.tag.Attr("trigger", v)
	return b
}

func (b *PopoverBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {

	b.tag.Children(b.children...)
	b.tag.AppendChildren(
		Template(b.overlay...).Attr("v-slot:overlay", "{ parent }"),
	)

	return b.tag.MarshalHTML(ctx)
}
