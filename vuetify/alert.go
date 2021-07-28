package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAlertBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAlert(children ...h.HTMLComponent) (r *VAlertBuilder) {
	r = &VAlertBuilder{
		tag: h.Tag("v-alert").Children(children...),
	}
	return
}

func (b *VAlertBuilder) Border(v string) (r *VAlertBuilder) {
	b.tag.Attr("border", v)
	return b
}

func (b *VAlertBuilder) CloseIcon(v string) (r *VAlertBuilder) {
	b.tag.Attr("close-icon", v)
	return b
}

func (b *VAlertBuilder) CloseLabel(v string) (r *VAlertBuilder) {
	b.tag.Attr("close-label", v)
	return b
}

func (b *VAlertBuilder) Color(v string) (r *VAlertBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAlertBuilder) ColoredBorder(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":colored-border", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Dark(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Dense(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Dismissible(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":dismissible", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Elevation(v int) (r *VAlertBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Height(v int) (r *VAlertBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Icon(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":icon", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Light(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) MaxHeight(v int) (r *VAlertBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) MaxWidth(v int) (r *VAlertBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) MinHeight(v int) (r *VAlertBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) MinWidth(v int) (r *VAlertBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Mode(v string) (r *VAlertBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VAlertBuilder) Origin(v string) (r *VAlertBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VAlertBuilder) Outlined(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Prominent(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":prominent", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Rounded(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Shaped(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Tag(v string) (r *VAlertBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VAlertBuilder) Text(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":text", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Tile(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Transition(v string) (r *VAlertBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VAlertBuilder) Type(v string) (r *VAlertBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VAlertBuilder) Value(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Width(v int) (r *VAlertBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAlertBuilder) Attr(vs ...interface{}) (r *VAlertBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAlertBuilder) Children(children ...h.HTMLComponent) (r *VAlertBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAlertBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAlertBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAlertBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAlertBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAlertBuilder) Class(names ...string) (r *VAlertBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAlertBuilder) ClassIf(name string, add bool) (r *VAlertBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAlertBuilder) On(name string, value string) (r *VAlertBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAlertBuilder) Bind(name string, value string) (r *VAlertBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAlertBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
