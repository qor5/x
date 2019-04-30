package html

import (
	"html"

	"github.com/sunfmin/bran/ui"
)

func HTML(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("html").Children(children...)
}

func Head(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("head").Children(children...)
}

func Body(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("body").Children(children...)
}

func Div(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("div").Children(children...)
}

func Label(text string) (r *HTMLTagBuilder) {
	return Tag("label").Text(text)
}

func Span(text string) (r *HTMLTagBuilder) {
	return Tag("span").Text(text)
}

func Title(text string) (r *HTMLTagBuilder) {
	return Tag("title").Text(text)
}

func Strong(text string) (r *HTMLTagBuilder) {
	return Tag("strong").Text(text)
}

func Legend(text string) (r *HTMLTagBuilder) {
	return Tag("legend").Text(text)
}

func H1(text string) (r *HTMLTagBuilder) {
	return Tag("h1").Text(text)
}

func H2(text string) (r *HTMLTagBuilder) {
	return Tag("h2").Text(text)
}

func H3(text string) (r *HTMLTagBuilder) {
	return Tag("h3").Text(text)
}

func H4(text string) (r *HTMLTagBuilder) {
	return Tag("h4").Text(text)
}

func H5(text string) (r *HTMLTagBuilder) {
	return Tag("h5").Text(text)
}

func H6(text string) (r *HTMLTagBuilder) {
	return Tag("h6").Text(text)
}

func Table(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("table").Children(children...)
}

func Thead(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("thead").Children(children...)
}

func Tbody(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("tbody").Children(children...)
}

func Tfoot(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("tfoot").Children(children...)
}

func Th(text string) (r *HTMLTagBuilder) {
	return Tag("th").Text(text)
}

func Td(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("td").Children(children...)
}

func Details(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("details").Children(children...)
}

func Fieldset(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("fieldset").Children(children...)
}

func Nav(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("nav").Children(children...)
}

func P(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("p").Children(children...)
}

func I(text string) (r *HTMLTagBuilder) {
	return Tag("i").Text(text)
}

func Section(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("section").Children(children...)
}

func Pre(text string) (r *HTMLTagBuilder) {
	return Tag("pre").Text(text)
}

func UL(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ul").Children(children...)
}

func OL(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ol").Children(children...)
}

func Li(children ...ui.HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("li").Children(children...)
}

func Img(src string) (r *HTMLTagBuilder) {
	return Tag("img").Attr("src", src)
}

type HTMLLinkBuilder struct {
	HTMLTagBuilder
}

func A(text string) (r *HTMLLinkBuilder) {
	tag := Tag("a").Text(text)
	return &HTMLLinkBuilder{
		HTMLTagBuilder: *tag,
	}
}

func (b *HTMLLinkBuilder) Href(href string) (r *HTMLLinkBuilder) {
	b.Attr("href", href)
	return b
}

type HTMLFormBuilder struct {
	HTMLTagBuilder
}

func Form(children ...ui.HTMLComponent) (r *HTMLFormBuilder) {
	tag := Tag("form").Children(children...)
	return &HTMLFormBuilder{
		HTMLTagBuilder: *tag,
	}
}

func (b *HTMLFormBuilder) Action(v string) (r *HTMLFormBuilder) {
	b.Attr("action", v)
	return b
}

func (b *HTMLFormBuilder) Method(v string) (r *HTMLFormBuilder) {
	b.Attr("method", v)
	return b
}

type HTMLInputBuilder struct {
	HTMLTagBuilder
}

func Input(name string) (r *HTMLInputBuilder) {
	tag := Tag("input").Attr("name", name)
	return &HTMLInputBuilder{
		HTMLTagBuilder: *tag,
	}
}

func (b *HTMLInputBuilder) Type(v string) (r *HTMLInputBuilder) {
	b.Attr("type", v)
	return b
}

func (b *HTMLInputBuilder) Value(v string) (r *HTMLInputBuilder) {
	b.Attr("value", v)
	return b
}

func (b *HTMLInputBuilder) Placeholder(v string) (r *HTMLInputBuilder) {
	b.Attr("placeholder", v)
	return b
}

type HTMLMetaBuilder struct {
	HTMLTagBuilder
}

func Meta(name string) (r *HTMLMetaBuilder) {
	tag := Tag("meta").Attr("name", name)
	return &HTMLMetaBuilder{
		HTMLTagBuilder: *tag,
	}
}

func (b *HTMLMetaBuilder) Content(content string) (r *HTMLMetaBuilder) {
	b.Attr("content", content)
	return b
}

func Text(text string) (r ui.HTMLComponent) {
	return ui.RawHTML(html.EscapeString(text))
}

func Script(script string) (r ui.HTMLComponent) {
	return Tag("script").
		Attr("type", "text/javascript").
		Children(ui.RawHTML(script))
}

type ButtonBuilder struct {
	HTMLTagBuilder
}

func Button(label string) (r *ButtonBuilder) {
	tag := Tag("button").Text(label)
	r = &ButtonBuilder{
		HTMLTagBuilder: *tag,
	}
	return
}

func (b *ButtonBuilder) Type(v string) (r *ButtonBuilder) {
	b.Attr("type", v)
	return b
}
