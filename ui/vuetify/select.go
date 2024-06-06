package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSelectBuilder struct {
	tag *h.HTMLTagBuilder
}

func (b *VSelectBuilder) Label(v string) (r *VSelectBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VSelectBuilder) Chips(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) ClosableChips(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":closable-chips", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) CloseText(v string) (r *VSelectBuilder) {
	b.tag.Attr("close-text", v)
	return b
}

func (b *VSelectBuilder) Type(v string) (r *VSelectBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VSelectBuilder) OpenText(v string) (r *VSelectBuilder) {
	b.tag.Attr("open-text", v)
	return b
}

func (b *VSelectBuilder) Eager(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) HideNoData(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) HideSelected(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":hide-selected", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) ListProps(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":list-props", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) BaseColor(v string) (r *VSelectBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VSelectBuilder) BgColor(v string) (r *VSelectBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VSelectBuilder) Disabled(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Multiple(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Reverse(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Flat(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Density(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) MaxWidth(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) MinWidth(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Width(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Items(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) ItemTitle(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":item-title", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) ItemValue(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) ItemChildren(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":item-children", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) ItemProps(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":item-props", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) ReturnObject(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) ValueComparator(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Rounded(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Tile(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Theme(v string) (r *VSelectBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VSelectBuilder) Color(v string) (r *VSelectBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSelectBuilder) Variant(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Name(v string) (r *VSelectBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VSelectBuilder) Menu(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":menu", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) MenuIcon(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":menu-icon", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) MenuProps(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":menu-props", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Id(v string) (r *VSelectBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VSelectBuilder) ModelValue(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Transition(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
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

func (b *VSelectBuilder) ItemColor(v string) (r *VSelectBuilder) {
	b.tag.Attr("item-color", v)
	return b
}

func (b *VSelectBuilder) Autofocus(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Counter(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":counter", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Prefix(v string) (r *VSelectBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VSelectBuilder) Placeholder(v string) (r *VSelectBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VSelectBuilder) PersistentPlaceholder(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) PersistentCounter(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":persistent-counter", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Suffix(v string) (r *VSelectBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VSelectBuilder) Role(v string) (r *VSelectBuilder) {
	b.tag.Attr("role", v)
	return b
}

func (b *VSelectBuilder) AppendIcon(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) CenterAffix(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) PrependIcon(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) HideSpinButtons(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Hint(v string) (r *VSelectBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VSelectBuilder) PersistentHint(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Messages(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Direction(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Error(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) MaxErrors(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Readonly(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Rules(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) ValidateOn(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Focused(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) HideDetails(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Clearable(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) ClearIcon(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":clear-icon", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) Active(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) PersistentClear(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":persistent-clear", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) PrependInnerIcon(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":prepend-inner-icon", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) SingleLine(v bool) (r *VSelectBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VSelectBuilder) Loading(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) CounterValue(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VSelectBuilder) ModelModifiers(v interface{}) (r *VSelectBuilder) {
	b.tag.Attr(":model-modifiers", h.JSONString(v))
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
