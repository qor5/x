package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXLinkageSelectRemotetBuilber struct {
	tag *h.HTMLTagBuilder
}

func VXLinkageSelectRemote() *VXLinkageSelectRemotetBuilber {
	b := &VXLinkageSelectRemotetBuilber{
		tag: h.Tag("vx-linkageselect-remote"),
	}
	return b
}

func (b *VXLinkageSelectRemotetBuilber) Labels(vs ...string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":labels", vs)
	return b
}

func (b *VXLinkageSelectRemotetBuilber) ErrorMessages(vs ...[]string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":error-messages", vs)
	return b
}

func (b *VXLinkageSelectRemotetBuilber) Disabled(v bool) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":disabled", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectRemotetBuilber) SelectOutOfOrder(v bool) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":select-out-of-order", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectRemotetBuilber) Chips(v bool) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":chips", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectRemotetBuilber) Row(v bool) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":row", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) HideDetails(v bool) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectRemotetBuilber) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXLinkageSelectRemotetBuilber) Attr(vs ...interface{}) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXLinkageSelectRemotetBuilber) MarshalHTML(ctx context.Context) ([]byte, error) {
	return b.tag.MarshalHTML(ctx)
}
