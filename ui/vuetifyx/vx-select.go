package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXSelectBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXSelect(children ...h.HTMLComponent) (r *VXSelectBuilder) {
	r = &VXSelectBuilder{
		tag: h.Tag("vx-select").Children(children...),
	}
	return
}

func (b *VXSelectBuilder) Class(names ...string) (r *VXSelectBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXSelectBuilder) Type(v string) (r *VXSelectBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VXSelectBuilder) Label(v string) (r *VXSelectBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VXSelectBuilder) Placeholder(v string) (r *VXSelectBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VXSelectBuilder) Chips(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) ClosableChips(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":closable-chips", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) HideNoData(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) HideDetails(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":hideDetails", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Required(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":required", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Multiple(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Clearable(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Items(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(":items", v)
	return b
}

func (b *VXSelectBuilder) Disabled(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) ItemTitle(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr("item-title", v)
	return b
}

func (b *VXSelectBuilder) ErrorMessages(errMsgs ...string) (r *VXSelectBuilder) {
	b.tag.Attr(":error-messages", errMsgs)
	return b
}

func (b *VXSelectBuilder) ItemValue(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VXSelectBuilder) Tips(v string) (r *VXSelectBuilder) {
	b.tag.Attr(":tips", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Attr(vs ...interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(vs...)
	return b
}
func (b *VXSelectBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXSelectBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
