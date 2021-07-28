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

func (b *VMenuBuilder) Activator(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) AllowOverflow(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":allow-overflow", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Attach(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
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

func (b *VMenuBuilder) ContentClass(v string) (r *VMenuBuilder) {
	b.tag.Attr("content-class", v)
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

func (b *VMenuBuilder) Eager(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Fixed(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) InternalActivator(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":internal-activator", fmt.Sprint(v))
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

func (b *VMenuBuilder) MaxHeight(v int) (r *VMenuBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
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

func (b *VMenuBuilder) OpenOnFocus(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
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
	b.tag.Attr(":return-value", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Right(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Rounded(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Tile(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
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
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) ZIndex(v int) (r *VMenuBuilder) {
	b.tag.Attr(":z-index", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VMenuBuilder) Attr(vs ...interface{}) (r *VMenuBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VMenuBuilder) Children(children ...h.HTMLComponent) (r *VMenuBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VMenuBuilder) AppendChildren(children ...h.HTMLComponent) (r *VMenuBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VMenuBuilder) PrependChildren(children ...h.HTMLComponent) (r *VMenuBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VMenuBuilder) Class(names ...string) (r *VMenuBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VMenuBuilder) ClassIf(name string, add bool) (r *VMenuBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VMenuBuilder) On(name string, value string) (r *VMenuBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VMenuBuilder) Bind(name string, value string) (r *VMenuBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VMenuBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
