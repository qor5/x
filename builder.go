//go:generate gorazor templates templates
package pagui

import (
	"bytes"
	"net/http"
	"time"

	"github.com/sunfmin/pagui/templates"
)

type Builder struct {
	layoutMiddleFunc LayoutMiddleFn
	frontDev         bool
	prefix           string
}

type LayoutFn func(r *http.Request, body string) (output string, err error)

type LayoutMiddleFn func(in LayoutFn, head *PageHeadBuilder) (out LayoutFn)

func New() (b *Builder) {
	b = new(Builder)
	return
}

func (b *Builder) LayoutMiddleFn(mf LayoutMiddleFn) (r *Builder) {
	b.layoutMiddleFunc = mf
	r = b
	return
}

func (b *Builder) FrontDev(v bool) (r *Builder) {
	b.frontDev = v
	r = b
	return
}

type ComponentsPack string

func (b *Builder) Prefix(prefix string) (r *Builder) {
	b.prefix = prefix
	r = b
	return
}

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

func (b *Builder) defaultLayoutMiddleFunc(in LayoutFn, head *PageHeadBuilder) (out LayoutFn) {
	return func(r *http.Request, body string) (output string, err error) {
		output = templates.App(b.frontDev, b.prefix, head.String(), body)
		return
	}
}

func (b *Builder) GetLayoutMiddleFunc() (lm LayoutMiddleFn) {
	if b.layoutMiddleFunc != nil {
		return b.layoutMiddleFunc
	}
	return b.defaultLayoutMiddleFunc
}
