package utils

import (
	"fmt"

	"github.com/shurcooL/sanitized_anchor_name"
	. "github.com/theplant/htmlgo"
)

func Anchor(h *HTMLTagBuilder, text string) HTMLComponent {
	anchorName := sanitized_anchor_name.Create(text)
	return h.Children(
		Text(text),
		A().Class("anchor").Href(fmt.Sprintf("#%s", anchorName)),
	).Id(anchorName)
}

func Demo(title string, path string) HTMLComponent {
	return Div(
		A().Text(title).Href(path).Target("_blank"),
	).Class("demo")
}
