package vuetify

import (
	"context"
	"fmt"
	h "github.com/theplant/htmlgo"
)

type VAppBuilder struct {
	tag *h.HTMLTagBuilder
}

func VApp(children ...h.HTMLComponent) (r *VAppBuilder) {
	r = &VAppBuilder{
		tag: h.Tag("v-app").Children(children...),
	}
	return
}

func (b *VAppBuilder) Dark(v bool) (r *VAppBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VAppBuilder) Id(v string) (r *VAppBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VAppBuilder) Light(v bool) (r *VAppBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VAppBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
