package ui

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type LazyLoaderBuilder struct {
	loaderFunc *EventFuncID
	tag        *h.HTMLTagBuilder
}

func LazyLoader(eventFuncId string, params ...string) (r *LazyLoaderBuilder) {
	r = &LazyLoaderBuilder{
		tag: h.Tag("bran-lazy-loader"),
		loaderFunc: &EventFuncID{
			ID:     eventFuncId,
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

func (b *LazyLoaderBuilder) ParentForceUpdateAfterLoaded() (r *LazyLoaderBuilder) {
	b.tag.Attr(":after-loaded", "parent.forceUpdate")
	return b
}

func (b *LazyLoaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	b.tag.SetAttr(":loader-func", b.loaderFunc)
	return b.tag.MarshalHTML(ctx)
}
