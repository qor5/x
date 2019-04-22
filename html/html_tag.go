package html

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	ui "github.com/sunfmin/page"
)

type HTMLTagBuilder struct {
	tag      string
	attrs    map[string]string
	text     string
	children []ui.HTMLComponent
}

func Tag(tag string) (r *HTMLTagBuilder) {
	r = &HTMLTagBuilder{}

	if r.attrs == nil {
		r.attrs = make(map[string]string)
	}

	r.Tag(tag)

	return
}

func (b *HTMLTagBuilder) Tag(v string) (r *HTMLTagBuilder) {
	b.tag = v
	r = b
	return
}

func (b *HTMLTagBuilder) Text(v string) (r *HTMLTagBuilder) {
	b.text = v
	r = b
	return
}

func (b *HTMLTagBuilder) Children(comps ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	b.children = comps
	r = b
	return
}

func (b *HTMLTagBuilder) Attr(k string, v string) (r *HTMLTagBuilder) {
	b.attrs[k] = v
	r = b
	return
}

func (b *HTMLTagBuilder) ClassNames(names ...string) (r *HTMLTagBuilder) {
	b.Attr("className", strings.TrimSpace(strings.Join(names, " ")))
	r = b
	return
}

func (b *HTMLTagBuilder) Style(v string) (r *HTMLTagBuilder) {
	b.Attr("style", v)
	r = b
	return
}

func (b *HTMLTagBuilder) OnClick(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *HTMLTagBuilder) {

	fid := &ui.EventFuncID{
		ID:     hub.RefEventFunc(eventFuncId, ef),
		Params: params,
	}

	jb, err := json.Marshal(fid)
	if err != nil {
		panic(err)
	}

	b.Attr("v-on:click", fmt.Sprintf("click(%s, $event)", string(jb)))
	r = b
	return
}

func (b *HTMLTagBuilder) AddChildren(c ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	b.children = append(b.children, c...)
	r = b
	return
}

func (b *HTMLTagBuilder) PrependChild(c ui.HTMLComponent) (r *HTMLTagBuilder) {
	b.children = append([]ui.HTMLComponent{c}, b.children...)
	r = b
	return
}

func (b *HTMLTagBuilder) MarshalHTML(phb *ui.PageHeadBuilder) (r []byte, err error) {
	// remove empty
	cs := []ui.HTMLComponent{}
	for _, c := range b.children {
		if c == nil {
			continue
		}
		cs = append(cs, c)
	}

	attrSegs := []string{}
	for k, v := range b.attrs {
		attrSegs = append(attrSegs, fmt.Sprintf("%s='%s'", k, v))
	}
	attrStr := ""
	if len(attrSegs) > 0 {
		attrStr = " " + strings.Join(attrSegs, " ")
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("<%s%s>\n", b.tag, attrStr))
	if len(cs) > 0 {
		for _, c := range cs {
			var child []byte
			child, err = c.MarshalHTML(phb)
			if err != nil {
				return
			}
			buf.Write(child)
		}
	} else if len(b.text) > 0 {
		buf.WriteString(b.text)
	}

	buf.WriteString(fmt.Sprintf("</%s>\n", b.tag))
	r = buf.Bytes()
	return
}
