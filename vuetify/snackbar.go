package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSnackbarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSnackbar(children ...h.HTMLComponent) (r *VSnackbarBuilder) {
	r = &VSnackbarBuilder{
		tag: h.Tag("v-snackbar").Children(children...),
	}
	return
}

func (b *VSnackbarBuilder) Absolute(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) App(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Bottom(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Centered(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":centered", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Color(v string) (r *VSnackbarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSnackbarBuilder) ContentClass(v string) (r *VSnackbarBuilder) {
	b.tag.Attr("content-class", v)
	return b
}

func (b *VSnackbarBuilder) Dark(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Elevation(v int) (r *VSnackbarBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Height(v int) (r *VSnackbarBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Left(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Light(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) MaxHeight(v int) (r *VSnackbarBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) MaxWidth(v int) (r *VSnackbarBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) MinHeight(v int) (r *VSnackbarBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) MinWidth(v int) (r *VSnackbarBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) MultiLine(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":multi-line", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Outlined(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Right(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Rounded(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Shaped(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Tag(v string) (r *VSnackbarBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VSnackbarBuilder) Text(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":text", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Tile(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Timeout(v int) (r *VSnackbarBuilder) {
	b.tag.Attr(":timeout", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Top(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":top", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Transition(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":transition", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Value(v interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSnackbarBuilder) Vertical(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Width(v int) (r *VSnackbarBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSnackbarBuilder) Attr(vs ...interface{}) (r *VSnackbarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSnackbarBuilder) Children(children ...h.HTMLComponent) (r *VSnackbarBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSnackbarBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSnackbarBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSnackbarBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSnackbarBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSnackbarBuilder) Class(names ...string) (r *VSnackbarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSnackbarBuilder) ClassIf(name string, add bool) (r *VSnackbarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSnackbarBuilder) On(name string, value string) (r *VSnackbarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSnackbarBuilder) Bind(name string, value string) (r *VSnackbarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSnackbarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
