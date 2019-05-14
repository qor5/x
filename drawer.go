package branoverlay

import (
	"context"
	"fmt"

	. "github.com/theplant/htmlgo"
)

type DrawerBuilder struct {
	children []HTMLComponent

	trigger MutableAttrHTMLComponent
	tag     *HTMLTagBuilder
}

func Drawer(children ...HTMLComponent) (r *DrawerBuilder) {
	r = &DrawerBuilder{
		tag: Tag("bran-drawer"),
	}
	r.children = children
	return
}

func (b *DrawerBuilder) Trigger(v MutableAttrHTMLComponent) (r *DrawerBuilder) {
	b.trigger = v
	return b
}

func (b *DrawerBuilder) Width(v int) (r *DrawerBuilder) {
	b.tag.Attr("width", fmt.Sprint(v))
	return b
}

func (b *DrawerBuilder) Title(v string) (r *DrawerBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *DrawerBuilder) Visible(v bool) (r *DrawerBuilder) {
	b.tag.Attr("visible", fmt.Sprint(v))
	return b
}

func (b *DrawerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.trigger == nil {
		panic("Drawer().Trigger() required")
	}

	b.trigger.SetAttr("@click", "parent.show")

	b.tag.Children(
		Template(b.trigger).Attr("v-slot:trigger", "{ parent }"),
		Template(b.children...).Attr("v-slot:drawer", "{ parent }"),
	)
	return b.tag.MarshalHTML(ctx)
}
