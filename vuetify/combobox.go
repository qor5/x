package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VComboboxBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCombobox(children ...h.HTMLComponent) (r *VComboboxBuilder) {
	r = &VComboboxBuilder{
		tag: h.Tag("v-combobox").Children(children...),
	}
	return
}

func (b *VComboboxBuilder) AllowOverflow(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":allow-overflow", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) AppendIcon(v string) (r *VComboboxBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VComboboxBuilder) AppendOuterIcon(v string) (r *VComboboxBuilder) {
	b.tag.Attr("append-outer-icon", v)
	return b
}

func (b *VComboboxBuilder) Attach(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) AutoSelectFirst(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":auto-select-first", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Autofocus(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) BackgroundColor(v string) (r *VComboboxBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VComboboxBuilder) CacheItems(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":cache-items", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Chips(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) ClearIcon(v string) (r *VComboboxBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VComboboxBuilder) Clearable(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Color(v string) (r *VComboboxBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VComboboxBuilder) Counter(v int) (r *VComboboxBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) CounterValue(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Dark(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) DeletableChips(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":deletable-chips", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Delimiters(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":delimiters", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Dense(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) DisableLookup(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":disable-lookup", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Disabled(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Eager(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Error(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) ErrorCount(v int) (r *VComboboxBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) ErrorMessages(v string) (r *VComboboxBuilder) {
	b.tag.Attr("error-messages", v)
	return b
}

func (b *VComboboxBuilder) Filled(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":filled", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Filter(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":filter", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Flat(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) FullWidth(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Height(v int) (r *VComboboxBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) HideDetails(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) HideNoData(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) HideSelected(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":hide-selected", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Hint(v string) (r *VComboboxBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VComboboxBuilder) Id(v string) (r *VComboboxBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VComboboxBuilder) ItemColor(v string) (r *VComboboxBuilder) {
	b.tag.Attr("item-color", v)
	return b
}

func (b *VComboboxBuilder) ItemDisabled(v string) (r *VComboboxBuilder) {
	b.tag.Attr("item-disabled", v)
	return b
}

func (b *VComboboxBuilder) ItemText(v string) (r *VComboboxBuilder) {
	b.tag.Attr("item-text", v)
	return b
}

func (b *VComboboxBuilder) ItemValue(v string) (r *VComboboxBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VComboboxBuilder) Items(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Label(v string) (r *VComboboxBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VComboboxBuilder) Light(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) LoaderHeight(v int) (r *VComboboxBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Loading(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) MenuProps(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":menu-props", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Messages(v string) (r *VComboboxBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VComboboxBuilder) Multiple(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) NoDataText(v string) (r *VComboboxBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VComboboxBuilder) NoFilter(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) OpenOnClear(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":open-on-clear", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Outlined(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) PersistentHint(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) PersistentPlaceholder(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Placeholder(v string) (r *VComboboxBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VComboboxBuilder) Prefix(v string) (r *VComboboxBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VComboboxBuilder) PrependIcon(v string) (r *VComboboxBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VComboboxBuilder) PrependInnerIcon(v string) (r *VComboboxBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VComboboxBuilder) Readonly(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) ReturnObject(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Reverse(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Rounded(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Rules(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) SearchInput(v string) (r *VComboboxBuilder) {
	b.tag.Attr("search-input", v)
	return b
}

func (b *VComboboxBuilder) Shaped(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) SingleLine(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) SmallChips(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":small-chips", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Solo(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) SoloInverted(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Success(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) SuccessMessages(v string) (r *VComboboxBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VComboboxBuilder) Suffix(v string) (r *VComboboxBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VComboboxBuilder) Type(v string) (r *VComboboxBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VComboboxBuilder) ValidateOnBlur(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Value(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) ValueComparator(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VComboboxBuilder) Attr(vs ...interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VComboboxBuilder) Children(children ...h.HTMLComponent) (r *VComboboxBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VComboboxBuilder) AppendChildren(children ...h.HTMLComponent) (r *VComboboxBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VComboboxBuilder) PrependChildren(children ...h.HTMLComponent) (r *VComboboxBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VComboboxBuilder) Class(names ...string) (r *VComboboxBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VComboboxBuilder) ClassIf(name string, add bool) (r *VComboboxBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VComboboxBuilder) On(name string, value string) (r *VComboboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VComboboxBuilder) Bind(name string, value string) (r *VComboboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VComboboxBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
