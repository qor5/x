package branoverlay

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type DrawerBuilder struct {
	children []h.HTMLComponent

	triggerElement h.MutableAttrHTMLComponent
	trigger        string
	tag            *h.HTMLTagBuilder
}

func Drawer(children ...h.HTMLComponent) (r *DrawerBuilder) {
	r = &DrawerBuilder{
		tag: h.Tag("bran-drawer"),
	}
	r.Placement("right")
	r.Trigger("click")
	r.children = children
	return
}

func (b *DrawerBuilder) TriggerElement(v h.MutableAttrHTMLComponent) (r *DrawerBuilder) {
	b.triggerElement = v
	return b
}

func (b *DrawerBuilder) Trigger(v string) (r *DrawerBuilder) {
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

	if b.triggerElement != nil {
		b.triggerElement.SetAttr(fmt.Sprintf("@%s", b.trigger), "parent.show")
	}

	b.tag.Children(
		h.If(b.triggerElement != nil, h.Template(b.triggerElement).Attr("v-slot:trigger", "{ parent }")),
		h.Template(b.children...).Attr("v-slot:drawer", "{ parent }"),
	)
	return b.tag.MarshalHTML(ctx)
}
