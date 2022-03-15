package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSelectBuilder struct {
	tag *h.HTMLTagBuilder
}

func (b *VSelectBuilder) AppendIcon(v string) (r *VSelectBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VSelectBuilder) AppendOuterIcon(v string) (r *VSelectBuilder) {
	b.tag.Attr("append-outer-icon", v)
	return b
}

func (b *VSelectBuilder) Attach(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Autofocus(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) BackgroundColor(v string) (r *VSelectBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VSelectBuilder) CacheItems(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":cache-items", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Chips(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) ClearIcon(v string) (r *VSelectBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VSelectBuilder) Clearable(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Color(v string) (r *VSelectBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSelectBuilder) Counter(v int) (r *VSelectBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) CounterValue(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Dark(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) DeletableChips(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":deletable-chips", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Dense(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) DisableLookup(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":disable-lookup", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Disabled(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Eager(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Error(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) ErrorCount(v int) (r *VSelectBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Filled(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":filled", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Flat(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) FullWidth(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Height(v int) (r *VSelectBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) HideDetails(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) HideSelected(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":hide-selected", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Hint(v string) (r *VSelectBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VSelectBuilder) Id(v string) (r *VSelectBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VSelectBuilder) ItemColor(v string) (r *VSelectBuilder) {
	b.tag.Attr("item-color", v)
	return b
}

func (b *VSelectBuilder) ItemDisabled(v string) (r *VSelectBuilder) {
	b.tag.Attr("item-disabled", v)
	return b
}

func (b *VSelectBuilder) ItemText(v string) (r *VSelectBuilder) {
	b.tag.Attr("item-text", v)
	return b
}

func (b *VSelectBuilder) ItemValue(v string) (r *VSelectBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VSelectBuilder) Items(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Label(v string) (r *VSelectBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VSelectBuilder) Light(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) LoaderHeight(v int) (r *VSelectBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Loading(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) MenuProps(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":menu-props", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Messages(v string) (r *VSelectBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VSelectBuilder) Multiple(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) NoDataText(v string) (r *VSelectBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VSelectBuilder) OpenOnClear(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":open-on-clear", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Outlined(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) PersistentHint(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) PersistentPlaceholder(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Placeholder(v string) (r *VSelectBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VSelectBuilder) Prefix(v string) (r *VSelectBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VSelectBuilder) PrependIcon(v string) (r *VSelectBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VSelectBuilder) PrependInnerIcon(v string) (r *VSelectBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VSelectBuilder) Readonly(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) ReturnObject(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Reverse(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Rounded(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Rules(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Shaped(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) SingleLine(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) SmallChips(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":small-chips", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Solo(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) SoloInverted(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Success(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) SuccessMessages(v string) (r *VSelectBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VSelectBuilder) Suffix(v string) (r *VSelectBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VSelectBuilder) Type(v string) (r *VSelectBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VSelectBuilder) ValidateOnBlur(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Value(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) ValueComparator(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSelectBuilder) Attr(vs ...interface{}) (r *VSelectBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSelectBuilder) Children(children ...h.HTMLComponent) (r *VSelectBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSelectBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSelectBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSelectBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSelectBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSelectBuilder) Class(names ...string) (r *VSelectBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSelectBuilder) ClassIf(name string, add bool) (r *VSelectBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSelectBuilder) On(name string, value string) (r *VSelectBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSelectBuilder) Bind(name string, value string) (r *VSelectBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSelectBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
