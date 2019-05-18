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
	r.Placement("right")
	r.children = children
	return
}

func (b *DrawerBuilder) Trigger(v MutableAttrHTMLComponent) (r *DrawerBuilder) {
	b.trigger = v
	return b
}

func (b *DrawerBuilder) Width(v int) (r *DrawerBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *DrawerBuilder) Height(v int) (r *DrawerBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *DrawerBuilder) DefaultOpen(open bool, animation bool) (r *DrawerBuilder) {
	if open {
		b.tag.Attr(":default-open", fmt.Sprint(open), ":first-enter", fmt.Sprint(!animation))
	}
	return b
}

func (b *DrawerBuilder) ClassName(v string) (r *DrawerBuilder) {
	b.tag.Attr(":class-name", v)
	return b
}

func (b *DrawerBuilder) Level(v []string) (r *DrawerBuilder) {
	b.tag.Attr(":level", v)
	return b
}

func (b *DrawerBuilder) Placement(v string) (r *DrawerBuilder) {
	b.tag.Attr("placement", v)
	return b
}

func (b *DrawerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {

	if b.trigger != nil {
		b.trigger.SetAttr("@click", "parent.show")
	}

	b.tag.Children(
		If(b.trigger != nil, Template(b.trigger).Attr("v-slot:trigger", "{ parent }")),
		Template(b.children...).Attr("v-slot:drawer", "{ parent }"),
	)
	return b.tag.MarshalHTML(ctx)
}
