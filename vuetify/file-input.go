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

func (b *VFileInputBuilder) AppendIcon(v string) (r *VFileInputBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VFileInputBuilder) AppendOuterIcon(v string) (r *VFileInputBuilder) {
	b.tag.Attr("append-outer-icon", v)
	return b
}

func (b *VFileInputBuilder) Autofocus(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) BackgroundColor(v string) (r *VFileInputBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VFileInputBuilder) Chips(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) ClearIcon(v string) (r *VFileInputBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VFileInputBuilder) Clearable(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Color(v string) (r *VFileInputBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VFileInputBuilder) Counter(v int) (r *VFileInputBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
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

func (b *VFileInputBuilder) CounterValue(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Dark(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Dense(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
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

func (b *VFileInputBuilder) ErrorCount(v int) (r *VFileInputBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Filled(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":filled", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Flat(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) FullWidth(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Height(v int) (r *VFileInputBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) HideDetails(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) HideInput(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":hide-input", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Hint(v string) (r *VFileInputBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VFileInputBuilder) Id(v string) (r *VFileInputBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VFileInputBuilder) Label(v string) (r *VFileInputBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VFileInputBuilder) Light(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) LoaderHeight(v int) (r *VFileInputBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Loading(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Messages(v string) (r *VFileInputBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VFileInputBuilder) Multiple(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Outlined(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) PersistentHint(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) PersistentPlaceholder(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Placeholder(v string) (r *VFileInputBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VFileInputBuilder) Prefix(v string) (r *VFileInputBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VFileInputBuilder) PrependIcon(v string) (r *VFileInputBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VFileInputBuilder) PrependInnerIcon(v string) (r *VFileInputBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VFileInputBuilder) Reverse(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Rounded(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Rules(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VFileInputBuilder) Shaped(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) ShowSize(v int) (r *VFileInputBuilder) {
	b.tag.Attr(":show-size", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) SingleLine(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) SmallChips(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":small-chips", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Solo(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) SoloInverted(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Success(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) SuccessMessages(v string) (r *VFileInputBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VFileInputBuilder) Suffix(v string) (r *VFileInputBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VFileInputBuilder) TruncateLength(v int) (r *VFileInputBuilder) {
	b.tag.Attr(":truncate-length", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Type(v string) (r *VFileInputBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VFileInputBuilder) ValidateOnBlur(v bool) (r *VFileInputBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VFileInputBuilder) Value(v interface{}) (r *VFileInputBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
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
