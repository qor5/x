package ui

import (
	"context"

	. "github.com/theplant/htmlgo"
)

type LazyLoaderBuilder struct {
	loaderFunc *EventFuncID
	tag        *HTMLTagBuilder
}

func LazyLoader(hub EventFuncHub, eventFuncId string, ef EventFunc, params ...string) (r *LazyLoaderBuilder) {
	r = &LazyLoaderBuilder{
		tag: Tag("bran-lazy-loader"),
		loaderFunc: &EventFuncID{
			ID:     hub.RefEventFunc(eventFuncId, ef),
			Params: params,
		},
	}
	return
}

func (b *LazyLoaderBuilder) Visible(v string) (r *LazyLoaderBuilder) {
	b.tag.Attr(":visible", v)
	return b
}

func (b *LazyLoaderBuilder) LoadWhenParentVisible() (r *LazyLoaderBuilder) {
	b.Visible("parent.isVisible")
	return b
}

func (b *LazyLoaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	b.tag.SetAttr(":loader-func", b.loaderFunc)
	return b.tag.MarshalHTML(ctx)
}
