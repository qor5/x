package html

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"strings"

	ui "github.com/sunfmin/page"
)

type HTMLTagBuilder struct {
	tag           string
	attrs         map[string]string
	classNames    []string
	text          string
	children      []ui.HTMLComponent
	fieldName     *string
	onInputFuncID *ui.EventFuncID
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

func (b *HTMLTagBuilder) FieldName(v string) (r *HTMLTagBuilder) {
	if len(v) > 0 {
		b.fieldName = &v
	}

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
	b.classNames = names
	r = b
	return
}

func (b *HTMLTagBuilder) Style(v string) (r *HTMLTagBuilder) {
	b.Attr("style", v)
	r = b
	return
}

func (b *HTMLTagBuilder) OnInput(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *HTMLTagBuilder) {

	b.onInputFuncID = &ui.EventFuncID{
		ID:     hub.RefEventFunc(eventFuncId, ef),
		Params: params,
	}

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

	b.Attr("v-on:click", fmt.Sprintf("onclick(%s, $event)", string(jb)))
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
	b.Attr("class", strings.TrimSpace(strings.Join(b.classNames, " ")))

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
			child, err = c.MarshalHTML(ctx)
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

func Div(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("div").Children(children...)
}

func Text(text string) (r ui.HTMLComponent) {
	r = ui.StringHTMLComponent(html.EscapeString(text))
	return
}

type Styles map[string]string

func (s Styles) String() string {
	segs := []string{}
	for k, v := range s {
		segs = append(segs, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(segs, "; ")
}

func (s Styles) Put(name, value string) (r Styles) {
	s[name] = value
	return
}
