package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VToolbarItemsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VToolbarItems() (r *VToolbarItemsBuilder) {
	r = &VToolbarItemsBuilder{
		tag: h.Tag("v-toolbar-items"),
	}
	return
}

func (b *VToolbarItemsBuilder) Class(names ...string) (r *VToolbarItemsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VToolbarItemsBuilder) ClassIf(name string, add bool) (r *VToolbarItemsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VToolbarItemsBuilder) On(name string, value string) (r *VToolbarItemsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VToolbarItemsBuilder) Bind(name string, value string) (r *VToolbarItemsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VToolbarItemsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
