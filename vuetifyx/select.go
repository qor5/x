package vuetifyx

import (
	"fmt"

	h "github.com/theplant/htmlgo"
)

func (b *VXSelectBuilder) AppendIcon(v string) (r *VXSelectBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VXSelectBuilder) AppendOuterIcon(v string) (r *VXSelectBuilder) {
	b.tag.Attr("append-outer-icon", v)
	return b
}

func (b *VXSelectBuilder) Attach(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VXSelectBuilder) Autofocus(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) BackgroundColor(v string) (r *VXSelectBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VXSelectBuilder) CacheItems(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":cache-items", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Chips(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) ClearIcon(v string) (r *VXSelectBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VXSelectBuilder) Clearable(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Color(v string) (r *VXSelectBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXSelectBuilder) Counter(v int) (r *VXSelectBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) CounterValue(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VXSelectBuilder) Dark(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) DeletableChips(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":deletable-chips", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Dense(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) DisableLookup(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":disable-lookup", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Disabled(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Eager(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Error(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) ErrorCount(v int) (r *VXSelectBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Filled(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":filled", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Flat(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) FullWidth(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Height(v int) (r *VXSelectBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) HideDetails(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) HideSelected(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":hide-selected", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Hint(v string) (r *VXSelectBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VXSelectBuilder) Id(v string) (r *VXSelectBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VXSelectBuilder) ItemColor(v string) (r *VXSelectBuilder) {
	b.tag.Attr("item-color", v)
	return b
}

func (b *VXSelectBuilder) ItemDisabled(v string) (r *VXSelectBuilder) {
	b.tag.Attr("item-disabled", v)
	return b
}

func (b *VXSelectBuilder) ItemText(v string) (r *VXSelectBuilder) {
	b.tag.Attr("item-text", v)
	return b
}

func (b *VXSelectBuilder) ItemValue(v string) (r *VXSelectBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VXSelectBuilder) Label(v string) (r *VXSelectBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VXSelectBuilder) Light(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) LoaderHeight(v int) (r *VXSelectBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Loading(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) MenuProps(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(":menu-props", h.JSONString(v))
	return b
}

func (b *VXSelectBuilder) Messages(v string) (r *VXSelectBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VXSelectBuilder) Multiple(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) NoDataText(v string) (r *VXSelectBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VXSelectBuilder) OpenOnClear(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":open-on-clear", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Outlined(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) PersistentHint(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) PersistentPlaceholder(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Placeholder(v string) (r *VXSelectBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VXSelectBuilder) Prefix(v string) (r *VXSelectBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VXSelectBuilder) PrependIcon(v string) (r *VXSelectBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VXSelectBuilder) PrependInnerIcon(v string) (r *VXSelectBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VXSelectBuilder) Readonly(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) ReturnObject(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Reverse(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Rounded(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Rules(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VXSelectBuilder) Shaped(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) SingleLine(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) SmallChips(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":small-chips", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Solo(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) SoloInverted(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Success(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) SuccessMessages(v string) (r *VXSelectBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VXSelectBuilder) Suffix(v string) (r *VXSelectBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VXSelectBuilder) Type(v string) (r *VXSelectBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VXSelectBuilder) ValidateOnBlur(v bool) (r *VXSelectBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VXSelectBuilder) Value(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXSelectBuilder) ValueComparator(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VXSelectBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXSelectBuilder) Attr(vs ...interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXSelectBuilder) Children(children ...h.HTMLComponent) (r *VXSelectBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXSelectBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXSelectBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXSelectBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXSelectBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXSelectBuilder) Class(names ...string) (r *VXSelectBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXSelectBuilder) ClassIf(name string, add bool) (r *VXSelectBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXSelectBuilder) On(name string, value string) (r *VXSelectBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXSelectBuilder) Bind(name string, value string) (r *VXSelectBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}
