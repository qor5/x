package vuetifyx

import (
	"context"

	"github.com/goplaid/web"
	h "github.com/theplant/htmlgo"
)

type VXLinkageSelectBuilber struct {
	tag *h.HTMLTagBuilder
}

func VXLinkageSelect() *VXLinkageSelectBuilber {
	b := &VXLinkageSelectBuilber{
		tag: h.Tag("vx-linkageselect"),
	}
	return b
}

func (b *VXLinkageSelectBuilber) FieldName(v string) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}

type LinkageSelectData struct {
	Label string
	Items []*LinkageSelectItem

	SelectedID string
}

type LinkageSelectItem struct {
	ID          string
	Name        string
	ChildrenIDs []string
}

// order by level
func (b *VXLinkageSelectBuilber) Data(v []*LinkageSelectData) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":data", v)
	return b
}

func (b *VXLinkageSelectBuilber) SelectOutOfOrder(v bool) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":select-out-of-order", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectBuilber) MarshalHTML(ctx context.Context) ([]byte, error) {
	return b.tag.MarshalHTML(ctx)
}
