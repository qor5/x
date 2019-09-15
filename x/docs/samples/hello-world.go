package samples

//@snippet_begin(HelloWorldSample)
import (
	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func HelloWorld(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = H1("Hello World")
	return
}

//@snippet_end

const HelloWorldPath = "/samples/hello_world"
