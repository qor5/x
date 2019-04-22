package page

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type PageHeadBuilder struct {
	Nodes   []*html.Node
	Scripts []string
	Styles  []string
}

func (b *PageHeadBuilder) Title(title string) (r *PageHeadBuilder) {
	b.addNode(atom.Title, title)
	r = b
	return
}

func (b *PageHeadBuilder) HasTitle() (r bool) {
	for _, n := range b.Nodes {
		if n.Type == html.ElementNode && n.Data == "title" {
			return true
		}
	}
	return
}

func (b *PageHeadBuilder) MetaNameContent(name, content string) (r *PageHeadBuilder) {
	b.Meta("name", name, "content", content)
	r = b
	return
}

func (b *PageHeadBuilder) Meta(attrs ...string) (r *PageHeadBuilder) {
	b.addNode(atom.Meta, "", attrs...)
	r = b
	return
}

func (b *PageHeadBuilder) PutScript(script string) (r *PageHeadBuilder) {
	var exists bool
	for _, s := range b.Scripts {
		if s == script {
			exists = true
			break
		}
	}
	if !exists {
		b.Scripts = append(b.Scripts, script)
	}
	r = b
	return
}

func (b *PageHeadBuilder) PutStyle(style string) (r *PageHeadBuilder) {
	var exists bool
	for _, s := range b.Styles {
		if s == style {
			exists = true
			break
		}
	}
	if !exists {
		b.Styles = append(b.Styles, style)
	}
	r = b
	return
}

func (b *PageHeadBuilder) Clear() (r *PageHeadBuilder) {
	b.Nodes = []*html.Node{}
	r = b
	return
}

func (b *PageHeadBuilder) HTML(v string) (r *PageHeadBuilder) {
	n, err := html.Parse(strings.NewReader(v))
	if err != nil {
		panic(err)
	}
	// _ = n
	n = n.FirstChild.FirstChild.FirstChild
	for n != nil {
		b.Nodes = append(b.Nodes, n)
		n = n.NextSibling
	}
	r = b
	return
}

func haveAttr(key, val string, attrs []html.Attribute) (keyExists bool, keyValBothExists bool) {
	for _, attr := range attrs {
		if strings.ToLower(attr.Key) == strings.ToLower(key) {
			keyExists = true
			if strings.ToLower(attr.Val) == strings.ToLower(val) {
				keyValBothExists = true
			}
		}
	}
	return
}

func (b *PageHeadBuilder) addCharsetViewPortIfMissing() {
	var foundCharset, foundViewPort bool
	for _, n := range b.Nodes {
		if ok, _ := haveAttr("charset", "", n.Attr); ok {
			foundCharset = true
		}
		if _, both := haveAttr("name", "viewport", n.Attr); both {
			foundViewPort = true
		}
	}
	if !foundCharset {
		b.Meta("charset", "utf8")
	}
	if !foundViewPort {
		b.MetaNameContent("viewport", "width=device-width, initial-scale=1, shrink-to-fit=no")
	}
}

func (b *PageHeadBuilder) String() string {
	b.addCharsetViewPortIfMissing()
	buf := bytes.NewBuffer(nil)
	for _, n := range b.Nodes {
		html.Render(buf, n)
		buf.WriteString("\n")
	}
	return buf.String()
}

func (b *PageHeadBuilder) addNode(atom atom.Atom, body string, attrs ...string) {
	l := len(attrs)
	if l%2 != 0 {
		panic(fmt.Sprintf("attrs should be pairs: %+v, length: %d", attrs, l))
	}

	var htmlAttrs []html.Attribute
	for i := 0; i < l; i = i + 2 {
		htmlAttrs = append(htmlAttrs, html.Attribute{
			Key: attrs[i],
			Val: attrs[i+1],
		})
	}

	n := &html.Node{
		Type: html.ElementNode,
		Data: atom.String(),
		Attr: htmlAttrs,
	}

	if len(body) > 0 {
		n.AppendChild(&html.Node{
			Type: html.TextNode,
			Data: body,
		})
	}

	b.Nodes = append(b.Nodes, n)
}
