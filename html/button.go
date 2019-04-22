package html

import (
	"encoding/json"
	"fmt"

	ui "github.com/sunfmin/page"
)

type ButtonBuilder struct {
	onMouseDownFuncID *ui.EventFuncID
}

func Button() (r *ButtonBuilder) {
	r = &ButtonBuilder{}
	return
}

func (b *ButtonBuilder) OnMouseDown(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *ButtonBuilder) {

	b.onMouseDownFuncID = &ui.EventFuncID{
		ID:     hub.RefEventFunc(eventFuncId, ef),
		Params: params,
	}

	r = b
	return
}

func (b *ButtonBuilder) MarshalHTML(phb *ui.PageHeadBuilder) (r []byte, err error) {
	mid, _ := json.Marshal(b.onMouseDownFuncID)
	r = []byte(fmt.Sprintf(`<button v-on:click='click(%s, $event)'>Hello</button>`, string(mid)))
	return
}
