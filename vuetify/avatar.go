package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAvatarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAvatar(children ...h.HTMLComponent) (r *VAvatarBuilder) {
	r = &VAvatarBuilder{
		tag: h.Tag("v-avatar").Children(children...),
	}
	return
}

func (b *VAvatarBuilder) Color(v string) (r *VAvatarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAvatarBuilder) Size(v int) (r *VAvatarBuilder) {
	b.tag.Attr(":size", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Tile(v bool) (r *VAvatarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Class(names ...string) (r *VAvatarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAvatarBuilder) ClassIf(name string, add bool) (r *VAvatarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAvatarBuilder) On(name string, value string) (r *VAvatarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAvatarBuilder) Bind(name string, value string) (r *VAvatarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAvatarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
