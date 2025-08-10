package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXDialogBuilder struct {
	tag          *h.HTMLTagBuilder
	rounded      bool
	roundedClass string
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
		tag:          h.Tag("vx-dialog").Children(children...),
		rounded:      true,      // default to rounded
		roundedClass: "rounded", // default rounded class
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

func (b *VXDialogBuilder) Rounded(v interface{}) (r *VXDialogBuilder) {
	// Accept bool or size string (e.g., "lg", "xl") and store to be applied on content container
	switch vv := v.(type) {
	case bool:
		b.rounded = vv
		if vv {
			b.roundedClass = "rounded"
		}
	case string:
		switch vv { // prefer tagged switch for clarity and linter friendliness
		case "true":
			b.rounded = true
			b.roundedClass = "rounded"
		case "false":
			b.rounded = false
		default:
			b.rounded = true
			b.roundedClass = fmt.Sprintf("rounded-%s", vv)
		}
	default:
		b.rounded = true
		b.roundedClass = "rounded"
	}
	return b
}

func (b *VXDialogBuilder) NoRounded() (r *VXDialogBuilder) {
	b.rounded = false
	return b
}

func (b *VXDialogBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	// Apply rounded style on overlay content via content-class for v-dialog inside VXDialog
	if b.rounded {
		b.tag.Attr(":content-class", fmt.Sprintf("'%s'", b.roundedClass))
	} else {
		b.tag.Attr(":content-class", "'rounded-0'")
	}
	return b.tag.MarshalHTML(ctx)
}
