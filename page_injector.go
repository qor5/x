package pagui

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type DefaultPageInjector struct {
	headNodes []*html.Node
	scripts   []string
	styles    []string
	tailHtmls []string
}

func (b *DefaultPageInjector) Title(title string) {
	b.addNode(atom.Title, title)
	return
}

func (b *DefaultPageInjector) HasTitle() (r bool) {
	for _, n := range b.headNodes {
		if n.Type == html.ElementNode && n.Data == "title" {
			return true
		}
	}
	return
}

func (b *DefaultPageInjector) MetaNameContent(name, content string) {
	b.Meta("name", name, "content", content)
	return
}

func (b *DefaultPageInjector) Meta(attrs ...string) {
	b.addNode(atom.Meta, "", attrs...)
	return
}

func (b *DefaultPageInjector) PutScript(script string) {
	var exists bool
	for _, s := range b.scripts {
		if s == script {
			exists = true
			break
		}
	}
	if !exists {
		b.scripts = append(b.scripts, script)
	}
	return
}

func (b *DefaultPageInjector) PutStyle(style string) {
	for _, s := range b.styles {
		if s == style {
			return
		}
	}

	b.styles = append(b.styles, style)
	return
}

func (b *DefaultPageInjector) PutTailHTML(v string) {
	for _, s := range b.tailHtmls {
		if s == v {
			return
		}
	}

	b.tailHtmls = append(b.tailHtmls, v)
	return
}

func (b *DefaultPageInjector) MainStyles(htmlTag bool) (r string) {

	if len(b.styles) == 0 {
		return
	}

	body := bytes.NewBuffer(nil)
	if htmlTag {
		body.WriteString(`<style id="main_styles" type="text/css">`)
		body.WriteString("\n")
	}
	body.WriteString(strings.Join(b.styles, "\n\n"))
	if htmlTag {
		body.WriteString("\n</style>\n")
	}

	return body.String()
}

func (b *DefaultPageInjector) MainScripts(htmlTag bool) (r string) {
	if len(b.scripts) == 0 {
		return
	}

	body := bytes.NewBuffer(nil)
	if htmlTag {
		body.WriteString("<script id=\"main_scripts\">\n")
	}
	body.WriteString(strings.Join(b.scripts, "\n\n"))
	if htmlTag {
		body.WriteString("\n</script>\n")
	}
	return body.String()
}

func (b *DefaultPageInjector) TailHTML() (r string) {
	return strings.Join(b.tailHtmls, "\n")
}

func (b *DefaultPageInjector) Clear() (r *DefaultPageInjector) {
	b.headNodes = []*html.Node{}
	return b
}

func (b *DefaultPageInjector) PutHeadHTML(v string) {
	n, err := html.Parse(strings.NewReader(v))
	if err != nil {
		panic(err)
	}
	// _ = n
	n = n.FirstChild.FirstChild.FirstChild
	for n != nil {
		b.headNodes = append(b.headNodes, n)
		n = n.NextSibling
	}
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

func (b *DefaultPageInjector) addCharsetViewPortIfMissing() {
	var foundCharset, foundViewPort bool
	for _, n := range b.headNodes {
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

func (b *DefaultPageInjector) HeadString() string {
	b.addCharsetViewPortIfMissing()
	buf := bytes.NewBuffer(nil)
	for _, n := range b.headNodes {
		html.Render(buf, n)
		buf.WriteString("\n")
	}
	return buf.String()
}

func (b *DefaultPageInjector) addNode(atom atom.Atom, body string, attrs ...string) {
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

	b.headNodes = append(b.headNodes, n)
}
