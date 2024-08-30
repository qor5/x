package vuetifyx

import (
	"context"
	"fmt"

	"github.com/qor5/web/v3"
	"github.com/qor5/x/v3/ui/vuetify"
	h "github.com/theplant/htmlgo"
)

type VXReadonlyFieldBuilder struct {
	label    string
	value    interface{}
	children h.HTMLComponents
	checkbox bool
}

func VXReadonlyField(children ...h.HTMLComponent) *VXReadonlyFieldBuilder {
	b := &VXReadonlyFieldBuilder{}
	if len(children) > 0 {
		b.children = children
	}
	return b
}

func (b *VXReadonlyFieldBuilder) Label(v string) *VXReadonlyFieldBuilder {
	b.label = v
	return b
}

func (b *VXReadonlyFieldBuilder) Value(v interface{}) *VXReadonlyFieldBuilder {
	b.value = v
	return b
}

func (b *VXReadonlyFieldBuilder) Children(children ...h.HTMLComponent) *VXReadonlyFieldBuilder {
	b.children = children
	return b
}

func (b *VXReadonlyFieldBuilder) Checkbox(v bool) *VXReadonlyFieldBuilder {
	b.checkbox = v
	return b
}

func (b *VXReadonlyFieldBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	var vComp h.HTMLComponent
	if b.children != nil {
		vComp = b.children
	} else {
		if b.checkbox {
			vComp = vuetify.VCheckbox().
				Attr(web.VField(b.label, b.value)...).
				Readonly(true).
				Disabled(true).
				Ripple(false).
				HideDetails(true).
				Class("my-0 py-0 text-black-lighten-9")
		} else {
			vComp = h.Text(fmt.Sprint(b.value))
		}
	}

	return h.Div(
		h.Label(b.label).Class("v-label theme--light text-caption"),
		h.Div(vComp).Class("pt-1"),
	).Class("mb-4").MarshalHTML(ctx)
}
