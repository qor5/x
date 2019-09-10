package bran

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type Builder struct {
	layoutFunc ui.LayoutFunc
}

func New() (b *Builder) {
	b = new(Builder)
	b.layoutFunc = defaultLayoutFunc
	return
}

func (b *Builder) LayoutFunc(mf ui.LayoutFunc) (r *Builder) {
	if mf == nil {
		panic("layout func is nil")
	}
	b.layoutFunc = mf
	return b
}

type ComponentsPack string

var startTime = time.Now()

func (b *Builder) PacksHandler(contentType string, packs ...ComponentsPack) http.Handler {
	var buf = bytes.NewBuffer(nil)
	for _, pk := range packs {
		// buf = append(buf, []byte(fmt.Sprintf("\n// pack %d\n", i+1))...)
		// buf = append(buf, []byte(fmt.Sprintf("\nconsole.log('pack %d, length %d');\n", i+1, len(pk)))...)
		buf.WriteString(string(pk))
		buf.WriteString("\n\n")
	}

	body := bytes.NewReader(buf.Bytes())

	return gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		http.ServeContent(w, r, "", startTime, body)
	}))
}

func NoopLayoutFunc(r *http.Request, injector *ui.PageInjector, body string) (output string, err error) {
	output = body
	return
}

func defaultLayoutFunc(r *http.Request, injector *ui.PageInjector, body string) (output string, err error) {

	root := h.HTML(
		h.Head(
			h.RawHTML(injector.GetHeadString()),
		),
		h.Body(
			h.Div(
				h.RawHTML(body),
			).Id("app").Attr("v-cloak", true),
			h.RawHTML(injector.GetTailString()),
		).Class("front"),
	)

	buf := bytes.NewBuffer(nil)
	ctx := new(ui.EventContext)
	ctx.R = r

	err = h.Fprint(buf, root, ui.WrapEventContext(context.TODO(), ctx))
	if err != nil {
		return
	}

	output = buf.String()
	return
}
