package branoverlay

import (
	"context"

	"github.com/sunfmin/bran/ui"
	. "github.com/theplant/htmlgo"
)

type LazyLoaderBuilder struct {
	loaderFunc *ui.EventFuncID
	tag        *HTMLTagBuilder
}

func LazyLoader(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *LazyLoaderBuilder) {
	r = &LazyLoaderBuilder{
		tag: Tag("bran-lazy-loader"),
		loaderFunc: &ui.EventFuncID{
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
