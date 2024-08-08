package vuetifyx

import (
	"fmt"

	h "github.com/theplant/htmlgo"
)

func (b *VXAutocompleteBuilder) AllowOverflow(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":allow-overflow", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) AppendIcon(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VXAutocompleteBuilder) AppendOuterIcon(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("append-outer-icon", v)
	return b
}

func (b *VXAutocompleteBuilder) Attach(v interface{}) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VXAutocompleteBuilder) AutoSelectFirst(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":auto-select-first", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Autofocus(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) BackgroundColor(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VXAutocompleteBuilder) CacheItems(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":cache-items", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Chips(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) ClearIcon(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VXAutocompleteBuilder) Clearable(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Color(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXAutocompleteBuilder) Counter(v int) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) CounterValue(v interface{}) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VXAutocompleteBuilder) Dark(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) DeletableChips(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":deletable-chips", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Dense(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) DisableLookup(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":disable-lookup", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Disabled(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Eager(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Error(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) ErrorCount(v int) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Filled(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":filled", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Filter(v interface{}) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":filter", h.JSONString(v))
	return b
}

func (b *VXAutocompleteBuilder) Flat(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) FullWidth(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Height(v int) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) HideDetails(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) HideNoData(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) HideSelected(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":hide-selected", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Hint(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VXAutocompleteBuilder) Id(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VXAutocompleteBuilder) ItemColor(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("item-color", v)
	return b
}

func (b *VXAutocompleteBuilder) ItemDisabled(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("item-disabled", v)
	return b
}

func (b *VXAutocompleteBuilder) ItemTitle(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("item-title", v)
	return b
}

func (b *VXAutocompleteBuilder) ItemValue(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VXAutocompleteBuilder) Label(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VXAutocompleteBuilder) Light(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) LoaderHeight(v int) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Loading(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) MenuProps(v interface{}) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":menu-props", h.JSONString(v))
	return b
}

func (b *VXAutocompleteBuilder) Messages(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VXAutocompleteBuilder) Multiple(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) NoDataText(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VXAutocompleteBuilder) NoFilter(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) OpenOnClear(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":open-on-clear", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Outlined(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) PersistentHint(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) PersistentPlaceholder(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Placeholder(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VXAutocompleteBuilder) Prefix(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VXAutocompleteBuilder) PrependIcon(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VXAutocompleteBuilder) PrependInnerIcon(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VXAutocompleteBuilder) Readonly(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) ReturnObject(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Reverse(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Rounded(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Rules(v interface{}) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VXAutocompleteBuilder) SearchInput(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("search-input", v)
	return b
}

func (b *VXAutocompleteBuilder) Shaped(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) SingleLine(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) SmallChips(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":small-chips", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Solo(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) SoloInverted(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Success(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) SuccessMessages(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VXAutocompleteBuilder) Suffix(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VXAutocompleteBuilder) Type(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VXAutocompleteBuilder) ValidateOnBlur(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VXAutocompleteBuilder) Value(v interface{}) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXAutocompleteBuilder) ValueComparator(v interface{}) (r *VXAutocompleteBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VXAutocompleteBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXAutocompleteBuilder) Attr(vs ...interface{}) (r *VXAutocompleteBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXAutocompleteBuilder) Children(children ...h.HTMLComponent) (r *VXAutocompleteBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXAutocompleteBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXAutocompleteBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXAutocompleteBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXAutocompleteBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXAutocompleteBuilder) Class(names ...string) (r *VXAutocompleteBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXAutocompleteBuilder) ClassIf(name string, add bool) (r *VXAutocompleteBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXAutocompleteBuilder) On(name string, value string) (r *VXAutocompleteBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXAutocompleteBuilder) Bind(name string, value string) (r *VXAutocompleteBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}
