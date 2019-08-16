package ui

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type LazyPortalBuilder struct {
	loaderFunc *EventFuncID
	tag        *h.HTMLTagBuilder
}

func LazyPortal(children ...h.HTMLComponent) (r *LazyPortalBuilder) {
	r = &LazyPortalBuilder{
		tag: h.Tag("bran-lazy-portal").Children(children...),
	}
	r.Visible("true")
	return
}

func (b *LazyPortalBuilder) EventFunc(eventFuncId string, params ...string) (r *LazyPortalBuilder) {
	b.loaderFunc = &EventFuncID{
		ID:     eventFuncId,
		Params: params,
	}
	return b
}

func (b *LazyPortalBuilder) Visible(v string) (r *LazyPortalBuilder) {
	b.tag.Attr(":visible", v)
	return b
}

func (b *LazyPortalBuilder) Name(v string) (r *LazyPortalBuilder) {
	b.tag.Attr("portal-name", v)
	return b
}

func (b *LazyPortalBuilder) Children(comps ...h.HTMLComponent) (r *LazyPortalBuilder) {
	b.tag.Children(comps...)
	return b
}

func (b *LazyPortalBuilder) LoadWhenParentVisible() (r *LazyPortalBuilder) {
	b.Visible("parent.isVisible")
	return b
}

func (b *LazyPortalBuilder) ParentForceUpdateAfterLoaded() (r *LazyPortalBuilder) {
	b.tag.Attr(":after-loaded", "parent.forceUpdate")
	return b
}

func (b *LazyPortalBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	b.tag.SetAttr(":loader-func", b.loaderFunc)
	return b.tag.MarshalHTML(ctx)
}
