package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTextareaBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTextarea(children ...h.HTMLComponent) (r *VTextareaBuilder) {
	r = &VTextareaBuilder{
		tag: h.Tag("vw-textarea").Children(children...),
	}
	return
}

func (b *VTextareaBuilder) AppendIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VTextareaBuilder) AppendOuterIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("append-outer-icon", v)
	return b
}

func (b *VTextareaBuilder) AutoGrow(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":auto-grow", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Autofocus(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) BackgroundColor(v string) (r *VTextareaBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VTextareaBuilder) Box(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":box", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) BrowserAutocomplete(v string) (r *VTextareaBuilder) {
	b.tag.Attr("browser-autocomplete", v)
	return b
}

func (b *VTextareaBuilder) ClearIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VTextareaBuilder) Clearable(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Color(v string) (r *VTextareaBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTextareaBuilder) Counter(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Dark(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Disabled(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) DontFillMaskBlanks(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":dont-fill-mask-blanks", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Error(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) ErrorCount(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) ErrorMessages(v string) (r *VTextareaBuilder) {
	b.tag.Attr("error-messages", v)
	return b
}

func (b *VTextareaBuilder) Flat(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) FullWidth(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Height(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) HideDetails(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Hint(v string) (r *VTextareaBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VTextareaBuilder) Label(v string) (r *VTextareaBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VTextareaBuilder) Light(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Loading(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Mask(v string) (r *VTextareaBuilder) {
	b.tag.Attr("mask", v)
	return b
}

func (b *VTextareaBuilder) Messages(v string) (r *VTextareaBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VTextareaBuilder) NoResize(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":no-resize", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Outline(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":outline", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) PersistentHint(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Placeholder(v string) (r *VTextareaBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VTextareaBuilder) Prefix(v string) (r *VTextareaBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VTextareaBuilder) PrependIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VTextareaBuilder) PrependInnerIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VTextareaBuilder) Readonly(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) ReturnMaskedValue(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":return-masked-value", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Reverse(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) RowHeight(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":row-height", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Rows(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":rows", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Rules(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr("rules", v)
	return b
}

func (b *VTextareaBuilder) SingleLine(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Solo(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) SoloInverted(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Success(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) SuccessMessages(v string) (r *VTextareaBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VTextareaBuilder) Suffix(v string) (r *VTextareaBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VTextareaBuilder) Type(v string) (r *VTextareaBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VTextareaBuilder) ValidateOnBlur(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}

func (b *VTextareaBuilder) FieldName(v string) (r *VTextareaBuilder) {
	b.tag.Attr("field-name", v)
	return b
}
