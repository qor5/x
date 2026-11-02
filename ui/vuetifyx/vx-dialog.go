package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXDialogBuilder struct {
	tag *h.HTMLTagBuilder
}

type (
	VXDialogType string
	VXDialogSize string
)

const (
	DialogDefault VXDialogType = "default"
	DialogInfo    VXDialogType = "info"
	DialogSuccess VXDialogType = "success"
	DialogWarn    VXDialogType = "warn"
	DialogError   VXDialogType = "error"
)

const (
	DialogSizeDefault VXDialogSize = "default"
	DialogSizeLarge   VXDialogSize = "large"
)

/*
@doc https://qor5.github.io/x/Components/VXDialog/
*/
func VXDialog(children ...h.HTMLComponent) (r *VXDialogBuilder) {
	r = &VXDialogBuilder{
		tag: h.Tag("vx-dialog").Children(children...),
	}
	return
}

/*
Set the title of VXDialog
*/
func (b *VXDialogBuilder) Title(v string) (r *VXDialogBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VXDialogBuilder) Type(v VXDialogType) (r *VXDialogBuilder) {
	b.tag.Attr("type", string(v))
	return b
}

func (b *VXDialogBuilder) Size(v VXDialogSize) (r *VXDialogBuilder) {
	b.tag.Attr("size", string(v))
	return b
}

func (b *VXDialogBuilder) Text(v string) (r *VXDialogBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VXDialogBuilder) NoClickAnimation(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":no-click-animation", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) HideCancel(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":hide-cancel", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) DisableOk(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":disable-ok", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) LoadingOk(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":loading-ok", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) HideOk(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":hide-ok", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) HideClose(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":hide-close", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) HideFooter(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":hide-footer", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) ModelValue(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) OkText(v string) (r *VXDialogBuilder) {
	b.tag.Attr("ok-text", v)
	return b
}

func (b *VXDialogBuilder) CancelText(v string) (r *VXDialogBuilder) {
	b.tag.Attr("cancel-text", v)
	return b
}

func (b *VXDialogBuilder) Persistent(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) ContentPadding(v string) (r *VXDialogBuilder) {
	b.tag.Attr("content-padding", v)
	return b
}

func (b *VXDialogBuilder) ContentOnlyMode(v bool) (r *VXDialogBuilder) {
	b.tag.Attr(":content-only-mode", fmt.Sprint(v))
	return b
}

func (b *VXDialogBuilder) ContentHeight(v int) (r *VXDialogBuilder) {
	b.tag.Attr("content-height", h.JSONString(v))
	return b
}

func (b *VXDialogBuilder) Width(v int) (r *VXDialogBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VXDialogBuilder) Height(v int) (r *VXDialogBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VXDialogBuilder) MaxWidth(v int) (r *VXDialogBuilder) {
	b.tag.Attr("max-width", h.JSONString(v))
	return b
}

func (b *VXDialogBuilder) Attr(vs ...interface{}) (r *VXDialogBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXDialogBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXDialogBuilder) Children(children ...h.HTMLComponent) (r *VXDialogBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXDialogBuilder) Class(names ...string) (r *VXDialogBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXDialogBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
