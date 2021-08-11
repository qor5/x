package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBtnBuilder struct {
	tag *h.HTMLTagBuilder
}

func (b *VBtnBuilder) Absolute(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) ActiveClass(v string) (r *VBtnBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VBtnBuilder) Append(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":append", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Block(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":block", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Bottom(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Color(v string) (r *VBtnBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBtnBuilder) Dark(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Depressed(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":depressed", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Disabled(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Elevation(v int) (r *VBtnBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Exact(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) ExactActiveClass(v string) (r *VBtnBuilder) {
	b.tag.Attr("exact-active-class", v)
	return b
}

func (b *VBtnBuilder) ExactPath(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":exact-path", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Fab(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":fab", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Fixed(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Height(v int) (r *VBtnBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Href(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":href", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Icon(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":icon", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) InputValue(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":input-value", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Large(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":large", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Left(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Light(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Link(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Loading(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) MaxHeight(v int) (r *VBtnBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) MaxWidth(v int) (r *VBtnBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) MinHeight(v int) (r *VBtnBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) MinWidth(v int) (r *VBtnBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Nuxt(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":nuxt", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Outlined(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Plain(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":plain", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Replace(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) RetainFocusOnClick(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":retain-focus-on-click", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Right(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Ripple(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Rounded(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Shaped(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Small(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":small", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Tag(v string) (r *VBtnBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBtnBuilder) Target(v string) (r *VBtnBuilder) {
	b.tag.Attr("target", v)
	return b
}

func (b *VBtnBuilder) Text(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":text", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Tile(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) To(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Top(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":top", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Type(v string) (r *VBtnBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VBtnBuilder) Value(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Width(v int) (r *VBtnBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) XLarge(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":x-large", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) XSmall(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":x-small", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBtnBuilder) Attr(vs ...interface{}) (r *VBtnBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBtnBuilder) Children(children ...h.HTMLComponent) (r *VBtnBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBtnBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBtnBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBtnBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBtnBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBtnBuilder) Class(names ...string) (r *VBtnBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBtnBuilder) ClassIf(name string, add bool) (r *VBtnBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBtnBuilder) On(name string, value string) (r *VBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBtnBuilder) Bind(name string, value string) (r *VBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBtnBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
