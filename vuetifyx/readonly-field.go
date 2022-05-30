package vuetifyx

import (
	"context"
	"fmt"

	"github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type VXReadonlyFieldBuilder struct {
	label     string
	value     interface{}
	valueComp h.HTMLComponent
	checkbox  bool
}

func VXReadonlyField() *VXReadonlyFieldBuilder {
	return &VXReadonlyFieldBuilder{}
}

func (b *VXReadonlyFieldBuilder) Label(v string) *VXReadonlyFieldBuilder {
	b.label = v
	return b
}

func (b *VXReadonlyFieldBuilder) Value(v interface{}) *VXReadonlyFieldBuilder {
	b.value = v
	return b
}

func (b *VXReadonlyFieldBuilder) ValueComponent(v h.HTMLComponent) *VXReadonlyFieldBuilder {
	b.valueComp = v
	return b
}

func (b *VXReadonlyFieldBuilder) Checkbox(v bool) *VXReadonlyFieldBuilder {
	b.checkbox = v
	return b
}

func (b *VXReadonlyFieldBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	var vComp h.HTMLComponent
	if b.valueComp != nil {
		vComp = b.valueComp
	} else {
		if b.checkbox {
			vComp = vuetify.VCheckbox().InputValue(b.value).
				Readonly(true).
				Ripple(false).
				HideDetails(true).
				Class("my-0 py-0")
		} else {
			vComp = h.Text(fmt.Sprint(b.value))
		}
	}

	return h.Div(
		h.Label(b.label).Class("v-label theme--light text-caption"),
		h.Div(vComp).Class("pt-1"),
	).Class("mb-4").MarshalHTML(ctx)
}
