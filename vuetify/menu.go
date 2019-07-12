package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VMenuBuilder struct {
	tag *h.HTMLTagBuilder
}

func VMenu(children ...h.HTMLComponent) (r *VMenuBuilder) {
	r = &VMenuBuilder{
		tag: h.Tag("v-menu").Children(children...),
	}
	return
}
func (b *VMenuBuilder) Absolute(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Activator(v string) (r *VMenuBuilder) {
	b.tag.Attr(":activator", v)
	return b
}

func (b *VMenuBuilder) AllowOverflow(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":allow-overflow", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Attach(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":attach", v)
	return b
}

func (b *VMenuBuilder) Auto(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":auto", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Bottom(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) CloseDelay(v int) (r *VMenuBuilder) {
	b.tag.Attr(":close-delay", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) CloseOnClick(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":close-on-click", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) CloseOnContentClick(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":close-on-content-click", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) ContentClass(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":content-class", v)
	return b
}

func (b *VMenuBuilder) Dark(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) DisableKeys(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":disable-keys", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Disabled(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Fixed(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) FullWidth(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) InputActivator(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":input-activator", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Lazy(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":lazy", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Left(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Light(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) MaxHeight(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":max-height", v)
	return b
}

func (b *VMenuBuilder) MaxWidth(v int) (r *VMenuBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) MinWidth(v int) (r *VMenuBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) NudgeBottom(v int) (r *VMenuBuilder) {
	b.tag.Attr(":nudge-bottom", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) NudgeLeft(v int) (r *VMenuBuilder) {
	b.tag.Attr(":nudge-left", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) NudgeRight(v int) (r *VMenuBuilder) {
	b.tag.Attr(":nudge-right", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) NudgeTop(v int) (r *VMenuBuilder) {
	b.tag.Attr(":nudge-top", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) NudgeWidth(v int) (r *VMenuBuilder) {
	b.tag.Attr(":nudge-width", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) OffsetOverflow(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":offset-overflow", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) OffsetX(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":offset-x", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) OffsetY(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":offset-y", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) OpenDelay(v int) (r *VMenuBuilder) {
	b.tag.Attr(":open-delay", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) OpenOnClick(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) OpenOnHover(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Origin(v string) (r *VMenuBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VMenuBuilder) PositionX(v int) (r *VMenuBuilder) {
	b.tag.Attr(":position-x", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) PositionY(v int) (r *VMenuBuilder) {
	b.tag.Attr(":position-y", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) ReturnValue(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":return-value", v)
	return b
}

func (b *VMenuBuilder) Right(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Top(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":top", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Transition(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":transition", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Value(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VMenuBuilder) ZIndex(v int) (r *VMenuBuilder) {
	b.tag.Attr(":z-index", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
