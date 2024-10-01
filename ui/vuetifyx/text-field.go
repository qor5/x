package vuetifyx

import (
	"context"
	"fmt"

	"github.com/qor5/web/v3"

	h "github.com/theplant/htmlgo"
)

type VXTextFieldBuilder struct {
	label     string
	readOnly  bool
	dense     string
	vField    vField
	text      string
	class     string
	valueType string
	suffix    string
}

type vField struct {
	formKey string
	value   interface{}
}

func VXTextField() *VXTextFieldBuilder {
	return &VXTextFieldBuilder{}
}

func (b *VXTextFieldBuilder) Label(label string) *VXTextFieldBuilder {
	b.label = label
	return b
}

func (b *VXTextFieldBuilder) ReadOnly(readOnly bool) *VXTextFieldBuilder {
	b.readOnly = readOnly
	return b
}

func (b *VXTextFieldBuilder) Dense(dense string) *VXTextFieldBuilder {
	b.dense = dense
	return b
}

func (b *VXTextFieldBuilder) Text(value string) *VXTextFieldBuilder {
	b.text = value
	b.readOnly = true
	return b
}

func (b *VXTextFieldBuilder) VField(formKey string, value interface{}) *VXTextFieldBuilder {
	b.vField.formKey = formKey
	b.vField.value = value
	return b
}

func (b *VXTextFieldBuilder) Class(class string) *VXTextFieldBuilder {
	b.class = class
	return b
}

func (b *VXTextFieldBuilder) Type(valueType string) *VXTextFieldBuilder {
	b.valueType = valueType
	return b
}

func (b *VXTextFieldBuilder) Suffix(suffix string) *VXTextFieldBuilder {
	b.suffix = suffix
	return b
}

func (b *VXTextFieldBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	// var labelStyle string = "font-size:16px; font-weight:500;"
	var label h.HTMLComponent
	if b.label != "" {
		label = VXLabel(h.Span(b.label)).Class("mb-2")
	}
	if b.readOnly {
		div := h.Div().Class(b.class)
		if b.label != "" {
			div.AppendChildren(label)
		}
		if b.suffix != "" {
			b.text = fmt.Sprintf("%s %s", b.text, b.suffix)
		}
		div.AppendChildren(
			h.Div(h.Span(b.text)),
		)
		return div.MarshalHTML(ctx)
	}

	var valueType string = "text"
	if b.valueType != "" {
		valueType = b.valueType
	}
	content := VXField().Type(valueType).
		Attr("suffix", b.suffix).
		Attr(web.VField(b.vField.formKey, b.vField.value)...)
	return h.Div(
		label,
		content,
	).Class(b.class).MarshalHTML(ctx)
}
