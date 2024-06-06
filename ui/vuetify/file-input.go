package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFileInputBuilder struct {
	tag *h.HTMLTagBuilder
}

func VFileInput(children ...h.HTMLComponent) (r *VFileInputBuilder) {
	r = &VFileInputBuilder{
		tag: h.Tag("v-file-input").Children(children...),
	}
	return
}

func (b *VFileInputBuilder) Label(v string) (r *VFileInputBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VFileInputBuilder) Counter(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Flat(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Chips(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) CounterSizeString(v string) (r *VFileInputBuilder) {
	b.tag.Attr("counter-size-string", v)
	return b
}

func (b *VFileInputBuilder) CounterString(v string) (r *VFileInputBuilder) {
	b.tag.Attr("counter-string", v)
	return b
}

func (b *VFileInputBuilder) HideInput(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":hide-input", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Multiple(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) ShowSize(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":show-size", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Id(v string) (r *VFileInputBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VFileInputBuilder) AppendIcon(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) CenterAffix(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) PrependIcon(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) HideSpinButtons(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Hint(v string) (r *VFileInputBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VFileInputBuilder) PersistentHint(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Messages(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Direction(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Reverse(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Density(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) MaxWidth(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) MinWidth(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Width(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Theme(v string) (r *VFileInputBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VFileInputBuilder) Disabled(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Error(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) MaxErrors(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Name(v string) (r *VFileInputBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VFileInputBuilder) Readonly(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Rules(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) ModelValue(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) ValidateOn(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) ValidationValue(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Focused(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) HideDetails(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) AppendInnerIcon(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":append-inner-icon", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) BgColor(v string) (r *VFileInputBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VFileInputBuilder) Clearable(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) ClearIcon(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":clear-icon", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Active(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Color(v string) (r *VFileInputBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VFileInputBuilder) BaseColor(v string) (r *VFileInputBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VFileInputBuilder) Dirty(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":dirty", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) PersistentClear(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":persistent-clear", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) PrependInnerIcon(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":prepend-inner-icon", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) SingleLine(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Variant(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Loading(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Rounded(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Tile(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFileInputBuilder) Attr(vs ...interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFileInputBuilder) Children(children ...h.HTMLComponent) (r *VFileInputBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFileInputBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFileInputBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFileInputBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFileInputBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFileInputBuilder) Class(names ...string) (r *VFileInputBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFileInputBuilder) ClassIf(name string, add bool) (r *VFileInputBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFileInputBuilder) On(name string, value string) (r *VFileInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFileInputBuilder) Bind(name string, value string) (r *VFileInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFileInputBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
