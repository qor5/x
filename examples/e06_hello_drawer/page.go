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

func randStr(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s: %d", prefix, rand.Int31n(100))
}

func HelloDrawer(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{}).(*mystate)

	pr.Schema = Drawer(
		ui.Bind(Button("Hello")).OnClick(ctx.Hub, "reload", reload),
		Button("Close").Attr("@click", "parent.close"),
		Div(Text(randStr("homeDrawer"))),
		LazyLoader(ctx.Hub, "editPage", editPage, "param1").ParentVisible(),
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

func editPage(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	_ = s

	r.Schema = Drawer(
		Button("Close").Attr("@click", "parent.close"),
		H1(ctx.Event.Params[0]),
		Div(Text(randStr("in editPage Drawer"))),
	).Trigger(
		A().Text("Open Inner").Href("#"),
	)
	return
}
