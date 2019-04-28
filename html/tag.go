package html

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	ui "github.com/sunfmin/page"
)

type HTMLTagBuilder struct {
	tag           string
	attrs         [][]string
	classNames    []string
	text          string
	children      []ui.HTMLComponent
	fieldName     *string
	onInputFuncID *ui.EventFuncID
}

func Tag(tag string) (r *HTMLTagBuilder) {
	r = &HTMLTagBuilder{}

	if r.attrs == nil {
		r.attrs = [][]string{}
	}

	r.Tag(tag)

	return
}

func (b *HTMLTagBuilder) Tag(v string) (r *HTMLTagBuilder) {
	b.tag = v
	return b
}

func (b *HTMLTagBuilder) Text(v string) (r *HTMLTagBuilder) {
	b.text = v
	return b
}

func (b *HTMLTagBuilder) Id(v string) (r *HTMLTagBuilder) {
	b.Attr("id", v)
	return b
}

func (b *HTMLTagBuilder) FieldName(v string) (r *HTMLTagBuilder) {
	if len(v) > 0 {
		b.fieldName = &v
	}

	return b
}

func (b *HTMLTagBuilder) Children(comps ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	b.children = comps
	return b
}

func (b *HTMLTagBuilder) Attr(k string, v string) (r *HTMLTagBuilder) {
	for _, at := range b.attrs {
		if at[0] == k {
			return b
		}
	}
	b.attrs = append(b.attrs, []string{k, v})
	return b
}

func (b *HTMLTagBuilder) ClassNames(names ...string) (r *HTMLTagBuilder) {
	b.classNames = []string{}
	for _, n := range names {
		b.classNames = append(b.classNames, strings.Split(n, " ")...)
	}
	return b
}

func (b *HTMLTagBuilder) Style(v string) (r *HTMLTagBuilder) {
	b.Attr("style", v)
	return b
}

func (b *HTMLTagBuilder) OnInput(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *HTMLTagBuilder) {

	b.onInputFuncID = &ui.EventFuncID{
		ID:     hub.RefEventFunc(eventFuncId, ef),
		Params: params,
	}

	return b
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

	b.Attr("v-on:click", fmt.Sprintf("onclick(%s, $event)", string(jb)))
	return b
}

func (b *HTMLTagBuilder) AddChildren(c ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	b.children = append(b.children, c...)
	return b
}

func (b *HTMLTagBuilder) PrependChild(c ui.HTMLComponent) (r *HTMLTagBuilder) {
	b.children = append([]ui.HTMLComponent{c}, b.children...)
	return b
}

func (b *HTMLTagBuilder) setupChange() {
	if b.fieldName == nil && b.onInputFuncID == nil {
		return
	}

	jb, err := json.Marshal(b.onInputFuncID)
	if err != nil {
		panic(err)
	}

	fieldName, err := json.Marshal(b.fieldName)
	if err != nil {
		panic(err)
	}

	b.Attr("v-on:input", fmt.Sprintf(`oninput(%s, %s, $event)`, string(jb), string(fieldName)))
}

func (b *HTMLTagBuilder) MarshalHTML(ctx *ui.EventContext) (r []byte, err error) {
	b.setupChange()
	class := strings.TrimSpace(strings.Join(b.classNames, " "))
	if len(class) > 0 {
		b.Attr("class", class)
	}

	// remove empty
	cs := []ui.HTMLComponent{}
	for _, c := range b.children {
		if c == nil {
			continue
		}
		cs = append(cs, c)
	}

	attrSegs := []string{}
	for _, at := range b.attrs {
		attrSegs = append(attrSegs, fmt.Sprintf("%s='%s'", at[0], at[1]))
	}

	attrStr := ""
	if len(attrSegs) > 0 {
		attrStr = " " + strings.Join(attrSegs, " ")
	}

	onlyText := false
	if len(b.text) > 0 {
		cs = append(cs, Text(b.text))
		if len(cs) == 1 {
			onlyText = true
		}
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("<%s%s>", b.tag, attrStr))
	if !onlyText {
		buf.WriteString("\n")
	}
	if len(cs) > 0 {
		for _, c := range cs {
			var child []byte
			child, err = c.MarshalHTML(ctx)
			if err != nil {
				return
			}
			buf.Write(child)
		}
	}
	buf.WriteString(fmt.Sprintf("</%s>\n", b.tag))
	r = buf.Bytes()
	return
}

type Styles struct {
	pairs [][]string
}

func (s *Styles) String() string {
	segs := []string{}
	for _, v := range s.pairs {
		segs = append(segs, fmt.Sprintf("%s:%s;", v[0], v[1]))
	}
	return strings.Join(segs, " ")
}

func (s *Styles) Put(name, value string) (r *Styles) {
	for _, el := range s.pairs {
		if el[0] == name {
			el[1] = value
			return s
		}
	}

	s.pairs = append(s.pairs, []string{name, value})
	return s
}
