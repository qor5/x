package html

import (
	"github.com/sunfmin/bran/ui"
)

func If(v bool, yes ui.HTMLComponent, no ui.HTMLComponent) (r ui.HTMLComponent) {
	return ui.ComponentFunc(func(ctx *ui.EventContext) (r []byte, err error) {
		if v && yes != nil {
			return yes.MarshalHTML(ctx)
		}
		if no != nil {
			return no.MarshalHTML(ctx)
		}
		return
	})
}
