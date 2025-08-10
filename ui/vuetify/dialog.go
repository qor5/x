package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDialogBuilder struct {
	tag          *h.HTMLTagBuilder
	rounded      bool
	roundedClass string
}

func VDialog(children ...h.HTMLComponent) (r *VDialogBuilder) {
	r = &VDialogBuilder{
		tag:          h.Tag("v-dialog").Children(children...),
		rounded:      true,      // default to rounded
		roundedClass: "rounded", // default rounded class
	}
	return
}

func (b *VDialogBuilder) Activator(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Fullscreen(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":fullscreen", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) RetainFocus(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":retain-focus", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Scrollable(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Absolute(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) CloseOnBack(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":close-on-back", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Contained(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":contained", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) ContentClass(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) ContentProps(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":content-props", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Disabled(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Opacity(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) NoClickAnimation(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":no-click-animation", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) ModelValue(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Persistent(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) Scrim(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":scrim", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) ZIndex(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":z-index", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Target(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":target", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) ActivatorProps(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":activator-props", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) OpenOnClick(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) OpenOnHover(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) OpenOnFocus(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) CloseOnContentClick(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":close-on-content-click", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) CloseDelay(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) OpenDelay(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Height(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) MaxHeight(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) MaxWidth(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) MinHeight(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) MinWidth(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Width(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Eager(v bool) (r *VDialogBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VDialogBuilder) LocationStrategy(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":location-strategy", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Location(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Origin(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":origin", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Offset(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":offset", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) ScrollStrategy(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":scroll-strategy", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Theme(v string) (r *VDialogBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VDialogBuilder) Transition(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Attach(v interface{}) (r *VDialogBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VDialogBuilder) Rounded(v interface{}) (r *VDialogBuilder) {
	// Accept bool or size string (e.g., "lg", "xl") and store to be applied on content container
	switch vv := v.(type) {
	case bool:
		b.rounded = vv
		if vv {
			b.roundedClass = "rounded"
		}
	case string:
		switch vv { // prefer tagged switch for clarity and linter friendliness
		case "true":
			b.rounded = true
			b.roundedClass = "rounded"
		case "false":
			b.rounded = false
		default:
			b.rounded = true
			b.roundedClass = fmt.Sprintf("rounded-%s", vv)
		}
	default:
		b.rounded = true
		b.roundedClass = "rounded"
	}
	return b
}

func (b *VDialogBuilder) NoRounded() (r *VDialogBuilder) {
	b.rounded = false
	return b
}

func (b *VDialogBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDialogBuilder) Attr(vs ...interface{}) (r *VDialogBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDialogBuilder) Children(children ...h.HTMLComponent) (r *VDialogBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDialogBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDialogBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDialogBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDialogBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDialogBuilder) Class(names ...string) (r *VDialogBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDialogBuilder) ClassIf(name string, add bool) (r *VDialogBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDialogBuilder) On(name string, value string) (r *VDialogBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDialogBuilder) Bind(name string, value string) (r *VDialogBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDialogBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	// Apply rounded style on overlay content instead of root tag for better visual correctness
	if b.rounded {
		b.tag.Attr(":content-class", fmt.Sprintf("'%s'", b.roundedClass))
	} else {
		b.tag.Attr(":content-class", "'rounded-0'")
	}
	return b.tag.MarshalHTML(ctx)
}
