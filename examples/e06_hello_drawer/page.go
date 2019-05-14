package e06_hello_drawer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/branoverlay"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	drawerVisible bool
}

func HelloDrawer(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{}).(*mystate)
	rand.Seed(time.Now().UnixNano())
	v := rand.Int31n(100)

	pr.Schema = Drawer(
		ui.Bind(Button("Hello")).OnClick(ctx.Hub, "reload", reload),
		Button("Close").Attr("@click", "parent.close"),
		Div(Text(fmt.Sprint(v))),
	).Trigger(
		A().Text("Open").Href("#"),
	).Width(500).Visible(s.drawerVisible)
	return
}

func reload(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	s := ctx.State.(*mystate)
	s.drawerVisible = true
	return
}
