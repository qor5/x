package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXSelectManyBuilder struct {
	tag             *h.HTMLTagBuilder
	selectedItems   interface{}
	items           interface{}
	searchItemsFunc string
}

func VXSelectMany(children ...h.HTMLComponent) (r *VXSelectManyBuilder) {
	r = &VXSelectManyBuilder{
		tag: h.Tag("vx-selectmany").Children(children...),
	}
	return
}

func (b *VXSelectManyBuilder) Items(v interface{}) (r *VXSelectManyBuilder) {
	b.items = v
	return b
}

func (b *VXSelectManyBuilder) SelectedItems(v interface{}) (r *VXSelectManyBuilder) {
	b.selectedItems = v
	return b
}

func (b *VXSelectManyBuilder) SearchItemsFunc(v string) (r *VXSelectManyBuilder) {
	b.searchItemsFunc = v
	return b
}

func (b *VXSelectManyBuilder) ItemText(v string) (r *VXSelectManyBuilder) {
	b.tag.Attr("item-text", v)
	return b
}

func (b *VXSelectManyBuilder) ItemValue(v string) (r *VXSelectManyBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VXSelectManyBuilder) Label(v string) (r *VXSelectManyBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VXSelectManyBuilder) AddItemLabel(v string) (r *VXSelectManyBuilder) {
	b.tag.Attr("add-item-label", v)
	return b
}

func (b *VXSelectManyBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.searchItemsFunc != "" {
		b.tag.Attr(":search-items-func", fmt.Sprintf(`function(val){return $plaid().eventFunc("%s").query("keyword", val).go()}`, b.searchItemsFunc))
	} else {
		b.tag.Attr(":items", b.items)
	}

	b.tag.Attr(":selected-items", b.selectedItems)
	return b.tag.MarshalHTML(ctx)
}
