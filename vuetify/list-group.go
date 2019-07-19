package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListGroup(children ...h.HTMLComponent) (r *VListGroupBuilder) {
	r = &VListGroupBuilder{
		tag: h.Tag("v-list-group").Children(children...),
	}
	return
}
func (b *VListGroupBuilder) ActiveClass(v string) (r *VListGroupBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VListGroupBuilder) AppendIcon(v string) (r *VListGroupBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VListGroupBuilder) Disabled(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) Group(v string) (r *VListGroupBuilder) {
	b.tag.Attr("group", v)
	return b
}

func (b *VListGroupBuilder) Lazy(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":lazy", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) NoAction(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":no-action", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) PrependIcon(v string) (r *VListGroupBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VListGroupBuilder) SubGroup(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":sub-group", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) Value(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
