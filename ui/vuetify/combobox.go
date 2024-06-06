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

func (b *VComboboxBuilder) Label(v string) (r *VComboboxBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VComboboxBuilder) AutoSelectFirst(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":auto-select-first", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) ClearOnSelect(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":clear-on-select", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Type(v string) (r *VComboboxBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VComboboxBuilder) FilterMode(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":filter-mode", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) NoFilter(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) CustomFilter(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":custom-filter", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Reverse(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Flat(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) CustomKeyFilter(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":custom-key-filter", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) FilterKeys(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":filter-keys", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Chips(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) ClosableChips(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":closable-chips", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) CloseText(v string) (r *VComboboxBuilder) {
	b.tag.Attr("close-text", v)
	return b
}

func (b *VComboboxBuilder) OpenText(v string) (r *VComboboxBuilder) {
	b.tag.Attr("open-text", v)
	return b
}

func (b *VComboboxBuilder) Eager(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
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

func (b *VComboboxBuilder) ListProps(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":list-props", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) BaseColor(v string) (r *VComboboxBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VComboboxBuilder) BgColor(v string) (r *VComboboxBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VComboboxBuilder) Disabled(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Multiple(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Density(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) MaxWidth(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) MinWidth(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Width(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Items(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) ItemTitle(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":item-title", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) ItemValue(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) ItemChildren(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":item-children", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) ItemProps(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":item-props", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) ReturnObject(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) ValueComparator(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Rounded(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Tile(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Theme(v string) (r *VComboboxBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VComboboxBuilder) Color(v string) (r *VComboboxBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VComboboxBuilder) Variant(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Name(v string) (r *VComboboxBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VComboboxBuilder) Menu(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":menu", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) MenuIcon(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":menu-icon", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) MenuProps(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":menu-props", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Id(v string) (r *VComboboxBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VComboboxBuilder) ModelValue(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Transition(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) NoDataText(v string) (r *VComboboxBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VComboboxBuilder) OpenOnClear(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":open-on-clear", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) ItemColor(v string) (r *VComboboxBuilder) {
	b.tag.Attr("item-color", v)
	return b
}

func (b *VComboboxBuilder) Autofocus(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Counter(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":counter", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Prefix(v string) (r *VComboboxBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VComboboxBuilder) Placeholder(v string) (r *VComboboxBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VComboboxBuilder) PersistentPlaceholder(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) PersistentCounter(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":persistent-counter", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Suffix(v string) (r *VComboboxBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VComboboxBuilder) Role(v string) (r *VComboboxBuilder) {
	b.tag.Attr("role", v)
	return b
}

func (b *VComboboxBuilder) AppendIcon(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) CenterAffix(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) PrependIcon(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) HideSpinButtons(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Hint(v string) (r *VComboboxBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VComboboxBuilder) PersistentHint(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Messages(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Direction(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Error(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) ErrorMessages(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":error-messages", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) MaxErrors(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Readonly(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Rules(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) ValidateOn(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Focused(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) HideDetails(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Clearable(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) ClearIcon(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":clear-icon", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Active(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) PersistentClear(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":persistent-clear", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) PrependInnerIcon(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":prepend-inner-icon", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) SingleLine(v bool) (r *VComboboxBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VComboboxBuilder) Loading(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) CounterValue(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) ModelModifiers(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":model-modifiers", h.JSONString(v))
	return b
}

func (b *VComboboxBuilder) Delimiters(v interface{}) (r *VComboboxBuilder) {
	b.tag.Attr(":delimiters", h.JSONString(v))
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
