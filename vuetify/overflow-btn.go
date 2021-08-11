package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VOverflowBtnBuilder struct {
	tag *h.HTMLTagBuilder
}

func VOverflowBtn(children ...h.HTMLComponent) (r *VOverflowBtnBuilder) {
	r = &VOverflowBtnBuilder{
		tag: h.Tag("v-overflow-btn").Children(children...),
	}
	return
}

func (b *VOverflowBtnBuilder) AllowOverflow(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":allow-overflow", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) AppendIcon(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VOverflowBtnBuilder) AppendOuterIcon(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("append-outer-icon", v)
	return b
}

func (b *VOverflowBtnBuilder) Attach(v interface{}) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VOverflowBtnBuilder) AutoSelectFirst(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":auto-select-first", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Autofocus(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) BackgroundColor(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VOverflowBtnBuilder) CacheItems(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":cache-items", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Chips(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) ClearIcon(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VOverflowBtnBuilder) Clearable(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Color(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VOverflowBtnBuilder) Counter(v int) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) CounterValue(v interface{}) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VOverflowBtnBuilder) Dark(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) DeletableChips(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":deletable-chips", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Dense(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) DisableLookup(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":disable-lookup", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Disabled(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Eager(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Editable(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":editable", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Error(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) ErrorCount(v int) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) ErrorMessages(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("error-messages", v)
	return b
}

func (b *VOverflowBtnBuilder) Filled(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":filled", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Filter(v interface{}) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":filter", h.JSONString(v))
	return b
}

func (b *VOverflowBtnBuilder) Flat(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) FullWidth(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Height(v int) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) HideDetails(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) HideNoData(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) HideSelected(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":hide-selected", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Hint(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VOverflowBtnBuilder) Id(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VOverflowBtnBuilder) ItemColor(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("item-color", v)
	return b
}

func (b *VOverflowBtnBuilder) ItemDisabled(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("item-disabled", v)
	return b
}

func (b *VOverflowBtnBuilder) ItemText(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("item-text", v)
	return b
}

func (b *VOverflowBtnBuilder) ItemValue(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VOverflowBtnBuilder) Items(v interface{}) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VOverflowBtnBuilder) Label(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VOverflowBtnBuilder) Light(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) LoaderHeight(v int) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Loading(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) MenuProps(v interface{}) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":menu-props", h.JSONString(v))
	return b
}

func (b *VOverflowBtnBuilder) Messages(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VOverflowBtnBuilder) Multiple(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) NoDataText(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VOverflowBtnBuilder) NoFilter(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) OpenOnClear(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":open-on-clear", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Outlined(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) PersistentHint(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) PersistentPlaceholder(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Placeholder(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VOverflowBtnBuilder) Prefix(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VOverflowBtnBuilder) PrependIcon(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VOverflowBtnBuilder) PrependInnerIcon(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VOverflowBtnBuilder) Readonly(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) ReturnObject(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Reverse(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Rounded(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Rules(v interface{}) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VOverflowBtnBuilder) SearchInput(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("search-input", v)
	return b
}

func (b *VOverflowBtnBuilder) Segmented(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":segmented", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Shaped(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) SingleLine(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) SmallChips(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":small-chips", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Solo(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) SoloInverted(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Success(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) SuccessMessages(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VOverflowBtnBuilder) Suffix(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VOverflowBtnBuilder) Type(v string) (r *VOverflowBtnBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VOverflowBtnBuilder) ValidateOnBlur(v bool) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VOverflowBtnBuilder) Value(v interface{}) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VOverflowBtnBuilder) ValueComparator(v interface{}) (r *VOverflowBtnBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VOverflowBtnBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VOverflowBtnBuilder) Attr(vs ...interface{}) (r *VOverflowBtnBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VOverflowBtnBuilder) Children(children ...h.HTMLComponent) (r *VOverflowBtnBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VOverflowBtnBuilder) AppendChildren(children ...h.HTMLComponent) (r *VOverflowBtnBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VOverflowBtnBuilder) PrependChildren(children ...h.HTMLComponent) (r *VOverflowBtnBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VOverflowBtnBuilder) Class(names ...string) (r *VOverflowBtnBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VOverflowBtnBuilder) ClassIf(name string, add bool) (r *VOverflowBtnBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VOverflowBtnBuilder) On(name string, value string) (r *VOverflowBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VOverflowBtnBuilder) Bind(name string, value string) (r *VOverflowBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VOverflowBtnBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
