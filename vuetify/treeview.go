package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTreeviewBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTreeview(children ...h.HTMLComponent) (r *VTreeviewBuilder) {
	r = &VTreeviewBuilder{
		tag: h.Tag("v-treeview").Children(children...),
	}
	return
}

func (b *VTreeviewBuilder) Activatable(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":activatable", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Active(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":active", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) ActiveClass(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VTreeviewBuilder) Color(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTreeviewBuilder) Dark(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Dense(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) ExpandIcon(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("expand-icon", v)
	return b
}

func (b *VTreeviewBuilder) Filter(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":filter", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Hoverable(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":hoverable", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) IndeterminateIcon(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("indeterminate-icon", v)
	return b
}

func (b *VTreeviewBuilder) ItemChildren(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("item-children", v)
	return b
}

func (b *VTreeviewBuilder) ItemDisabled(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("item-disabled", v)
	return b
}

func (b *VTreeviewBuilder) ItemKey(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("item-key", v)
	return b
}

func (b *VTreeviewBuilder) ItemText(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("item-text", v)
	return b
}

func (b *VTreeviewBuilder) Items(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Light(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) LoadChildren(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":load-children", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) LoadingIcon(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("loading-icon", v)
	return b
}

func (b *VTreeviewBuilder) MultipleActive(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":multiple-active", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) OffIcon(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("off-icon", v)
	return b
}

func (b *VTreeviewBuilder) OnIcon(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("on-icon", v)
	return b
}

func (b *VTreeviewBuilder) Open(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":open", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) OpenAll(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":open-all", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) OpenOnClick(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) ReturnObject(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Rounded(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Search(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("search", v)
	return b
}

func (b *VTreeviewBuilder) Selectable(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":selectable", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) SelectedColor(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("selected-color", v)
	return b
}

func (b *VTreeviewBuilder) SelectionType(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("selection-type", v)
	return b
}

func (b *VTreeviewBuilder) Shaped(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Transition(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":transition", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Value(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTreeviewBuilder) Attr(vs ...interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTreeviewBuilder) Children(children ...h.HTMLComponent) (r *VTreeviewBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTreeviewBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTreeviewBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTreeviewBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTreeviewBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTreeviewBuilder) Class(names ...string) (r *VTreeviewBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTreeviewBuilder) ClassIf(name string, add bool) (r *VTreeviewBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTreeviewBuilder) On(name string, value string) (r *VTreeviewBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTreeviewBuilder) Bind(name string, value string) (r *VTreeviewBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTreeviewBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
