package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTooltipBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTooltip(children ...h.HTMLComponent) (r *VTooltipBuilder) {
	r = &VTooltipBuilder{
		tag: h.Tag("v-tooltip").Children(children...),
	}
	return
}

func (b *VTooltipBuilder) Absolute(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Activator(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":activator", v)
	return b
}

func (b *VTooltipBuilder) AllowOverflow(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":allow-overflow", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Attach(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":attach", v)
	return b
}

func (b *VTooltipBuilder) Bottom(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) CloseDelay(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":close-delay", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Color(v string) (r *VTooltipBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTooltipBuilder) ContentClass(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":content-class", v)
	return b
}

func (b *VTooltipBuilder) Dark(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Debounce(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":debounce", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Disabled(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Fixed(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) InputActivator(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":input-activator", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Lazy(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":lazy", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Left(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Light(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) MaxWidth(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) MinWidth(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) NudgeBottom(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":nudge-bottom", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) NudgeLeft(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":nudge-left", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) NudgeRight(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":nudge-right", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) NudgeTop(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":nudge-top", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) NudgeWidth(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":nudge-width", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) OffsetOverflow(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":offset-overflow", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) OpenDelay(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":open-delay", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) PositionX(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":position-x", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) PositionY(v int) (r *VTooltipBuilder) {
	b.tag.Attr(":position-y", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Right(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Tag(v string) (r *VTooltipBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTooltipBuilder) Top(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":top", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Transition(v string) (r *VTooltipBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VTooltipBuilder) Value(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VTooltipBuilder) ZIndex(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":z-index", v)
	return b
}

func (b *VTooltipBuilder) Class(names ...string) (r *VTooltipBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTooltipBuilder) ClassIf(name string, add bool) (r *VTooltipBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTooltipBuilder) On(name string, value string) (r *VTooltipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTooltipBuilder) Bind(name string, value string) (r *VTooltipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTooltipBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
