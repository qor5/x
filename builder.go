package bran

import (
	"bytes"
	"net/http"
	"time"

	"github.com/sunfmin/bran/ui"

	h "github.com/sunfmin/bran/html"
)

type Builder struct {
	layoutMiddleFunc ui.LayoutMiddleFunc
}

func New() (b *Builder) {
	b = new(Builder)
	return
}

func (b *Builder) LayoutMiddleFunc(mf ui.LayoutMiddleFunc) (r *Builder) {
	b.layoutMiddleFunc = mf
	return b
}

type ComponentsPack string

var startTime = time.Now()

func (b *Builder) PacksHandler(contentType string, packs ...ComponentsPack) http.HandlerFunc {
	var buf = bytes.NewBuffer(nil)
	for _, pk := range packs {
		// buf = append(buf, []byte(fmt.Sprintf("\n// pack %d\n", i+1))...)
		// buf = append(buf, []byte(fmt.Sprintf("\nconsole.log('pack %d, length %d');\n", i+1, len(pk)))...)
		buf.WriteString(string(pk))
		buf.WriteString("\n\n")
	}

	body := bytes.NewReader(buf.Bytes())

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		http.ServeContent(w, r, "", startTime, body)
	}
}

func (b *Builder) defaultLayoutMiddleFunc(in ui.LayoutFunc, head ui.PageInjector) (out ui.LayoutFunc) {
	return func(r *http.Request, body string) (output string, err error) {

		root := h.HTML(
			h.Head(
				ui.RawHTML(head.HeadString()),
			),
			h.Body(
				ui.RawHTML(body),
			).Class("front"),
		)

		buf := bytes.NewBuffer(nil)
		buf.WriteString("<!DOCTYPE html>\n")

		var b []byte
		ctx := new(ui.EventContext)
		ctx.R = r
		b, err = root.MarshalHTML(ctx)
		if err != nil {
			return
		}
		buf.Write(b)

		output = buf.String()
		return
	}
}

func (b *Builder) getLayoutMiddleFunc() (lm ui.LayoutMiddleFunc) {
	if b.layoutMiddleFunc != nil {
		return b.layoutMiddleFunc
	}
	return b.defaultLayoutMiddleFunc
}
