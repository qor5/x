package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXFieldBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXField(children ...h.HTMLComponent) (r *VXFieldBuilder) {
	r = &VXFieldBuilder{
		tag: h.Tag("vx-field").Children(children...),
	}
	return
}

func (b *VXFieldBuilder) Label(v string) (r *VXFieldBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VXFieldBuilder) Type(v string) (r *VXFieldBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VXFieldBuilder) Name(v string) (r *VXFieldBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VXFieldBuilder) Id(v string) (r *VXFieldBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VXFieldBuilder) Placeholder(v string) (r *VXFieldBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VXFieldBuilder) Readonly(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Width(v int) (r *VXFieldBuilder) {
	b.tag.Attr("width", h.JSONString(v))
	return b
}

func (b *VXFieldBuilder) PasswordVisibleToggle(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":password-visible-toggle", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) PasswordVisibleDefault(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":password-visible-default", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Disabled(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Required(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":required", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Autofocus(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) HideDetails(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}
func (b *VXFieldBuilder) Clearable(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) ModelValue(v interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VXFieldBuilder) Attr(vs ...interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXFieldBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXFieldBuilder) Children(children ...h.HTMLComponent) (r *VXFieldBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXFieldBuilder) Class(names ...string) (r *VXFieldBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXFieldBuilder) Tips(v string) (r *VXFieldBuilder) {
	b.tag.Attr("tips", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) On(name string, value string) (r *VXFieldBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXFieldBuilder) Bind(name string, value string) (r *VXFieldBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXFieldBuilder) ErrorMessages(errMsgs ...string) (r *VXFieldBuilder) {
	b.tag.Attr(":error-messages", errMsgs)
	return b
}

func (b *VXFieldBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
