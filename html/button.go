package html

import (
	ui "github.com/sunfmin/page"
)

type ButtonBuilder struct {
}

func Button() (r *ButtonBuilder) {
	r = &ButtonBuilder{}
	return
}

func (b *ButtonBuilder) MarshalHTML(phb *ui.PageHeadBuilder) (r []byte, err error) {
	r = []byte(`<button v-on:click='click({id: "hello"}, $event)'>Hello</button>`)
	return
}
