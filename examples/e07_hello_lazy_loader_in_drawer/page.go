package e07_hello_lazy_loader_in_drawer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sunfmin/bran/ui"
	bo "github.com/sunfmin/branoverlay"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	drawerVisible bool
	Name          string
	NameError     string
}

func randStr(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s: %d", prefix, rand.Int31n(100))
}

func HelloLazyLoaderInDrawer(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{}).(*mystate)

	pr.Schema = Div(
		H1(s.Name),
		bo.Drawer(
			ui.LazyLoader(ctx.Hub, "editPage", editPage, "param1").LoadWhenParentVisible(),
			bo.Drawer(
				ui.LazyLoader(ctx.Hub, "editPage", editPage, "param2").LoadWhenParentVisible(),
			).Trigger(
				A().Text("New Drawer").Href("#"),
			),
		).Trigger(
			A().Text("Edit").Href("#"),
		).Width(500).DefaultOpen(s.drawerVisible),
	)
	return
}

func editPage(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	_ = s

	r.Schema = bo.Drawer(
		Button("Close").Attr("@click", "parent.close"),
		H1(ctx.Event.Params[0]),
		Div(Text(randStr("in editPage Drawer"))),
	).Trigger(
		A().Text("Open " + randStr("inner")).Href("#"),
	).Width(400)
	return
}
