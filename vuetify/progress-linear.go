package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VProgressLinearBuilder struct {
	tag *h.HTMLTagBuilder
}

func VProgressLinear(children ...h.HTMLComponent) (r *VProgressLinearBuilder) {
	r = &VProgressLinearBuilder{
		tag: h.Tag("v-progress-linear").Children(children...),
	}
	return
}

func (b *VProgressLinearBuilder) Absolute(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
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

func (b *VProgressLinearBuilder) Bottom(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
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

func (b *VProgressLinearBuilder) Dark(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Fixed(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
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

func (b *VProgressLinearBuilder) Light(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Query(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":query", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Reverse(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Rounded(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Stream(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":stream", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Striped(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":striped", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Top(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":top", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Value(v int) (r *VProgressLinearBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VProgressLinearBuilder) Attr(vs ...interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VProgressLinearBuilder) Children(children ...h.HTMLComponent) (r *VProgressLinearBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VProgressLinearBuilder) AppendChildren(children ...h.HTMLComponent) (r *VProgressLinearBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VProgressLinearBuilder) PrependChildren(children ...h.HTMLComponent) (r *VProgressLinearBuilder) {
	b.tag.PrependChildren(children...)
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
