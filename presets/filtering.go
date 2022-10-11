package presets

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/vuetifyx"
)

func (b *ListingBuilder) FilterDataFunc(v FilterDataFunc) {
	b.filterDataFunc = func(ctx *web.EventContext) vuetifyx.FilterData {
		fd := v(ctx)
		for _, k := range fd {
			k.Key = "f_" + k.Key
		}
		return fd
	}
}

func (b *ListingBuilder) FilterTabsFunc(v FilterTabsFunc) {
	b.filterTabsFunc = v
}
