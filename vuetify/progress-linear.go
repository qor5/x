package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VProgressLinearBuilder struct {
	tag *h.HTMLTagBuilder
}

func VProgressLinear() (r *VProgressLinearBuilder) {
	r = &VProgressLinearBuilder{
		tag: h.Tag("v-progress-linear"),
	}
	return
}

func (b *VProgressLinearBuilder) Active(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) BackgroundColor(v string) (r *VProgressLinearBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VProgressLinearBuilder) BackgroundOpacity(v int) (r *VProgressLinearBuilder) {
	b.tag.Attr(":background-opacity", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) BufferValue(v int) (r *VProgressLinearBuilder) {
	b.tag.Attr(":buffer-value", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Color(v string) (r *VProgressLinearBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VProgressLinearBuilder) Height(v int) (r *VProgressLinearBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Indeterminate(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Query(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":query", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Value(v int) (r *VProgressLinearBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Class(names ...string) (r *VProgressLinearBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VProgressLinearBuilder) ClassIf(name string, add bool) (r *VProgressLinearBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VProgressLinearBuilder) On(name string, value string) (r *VProgressLinearBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VProgressLinearBuilder) Bind(name string, value string) (r *VProgressLinearBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VProgressLinearBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
