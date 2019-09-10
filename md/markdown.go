package md

import (
	"context"

	"github.com/shurcooL/github_flavored_markdown"
	. "github.com/theplant/htmlgo"
)

func Markdown(body string) HTMLComponent {
	return ComponentFunc(func(c context.Context) (r []byte, err error) {
		root := RawHTML(github_flavored_markdown.Markdown([]byte(body)))
		return root.MarshalHTML(c)
	})
}
