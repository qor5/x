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

func (b *VSnackbarBuilder) AutoHeight(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":auto-height", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Bottom(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Color(v string) (r *VSnackbarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSnackbarBuilder) Left(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) MultiLine(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":multi-line", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Right(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
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

func (b *VSnackbarBuilder) Value(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VSnackbarBuilder) Vertical(v bool) (r *VSnackbarBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
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
