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

type LinkageSelectItem struct {
	ID          string
	Name        string
	ChildrenIDs []string
}

func (b *VXLinkageSelectBuilber) Items(vs ...[]*LinkageSelectItem) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":items", vs)
	return b
}

func (b *VXLinkageSelectBuilber) Labels(vs ...string) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":labels", vs)
	return b
}

func (b *VXLinkageSelectBuilber) SelectedIDs(vs ...string) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":value", vs)
	return b
}

func (b *VXLinkageSelectBuilber) ErrorMessages(vs ...[]string) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":error-messages", vs)
	return b
}

func (b *VXLinkageSelectBuilber) Disabled(v bool) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":disabled", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectBuilber) SelectOutOfOrder(v bool) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":select-out-of-order", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectBuilber) Chips(v bool) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":chips", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectBuilber) Row(v bool) (r *VXLinkageSelectBuilber) {
	b.tag.Attr(":row", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectBuilber) MarshalHTML(ctx context.Context) ([]byte, error) {
	return b.tag.MarshalHTML(ctx)
}
