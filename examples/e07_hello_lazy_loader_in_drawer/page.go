package e07_hello_lazy_loader_in_drawer

import (
	"fmt"
	"math/rand"
	"time"

	bo "github.com/sunfmin/bran/overlay"
	"github.com/sunfmin/bran/ui"
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
	ctx.Hub.RegisterEventFunc("editPage", editPage)

	s := ctx.StateOrInit(&mystate{}).(*mystate)

	pr.Schema = Div(
		H1(s.Name),
		bo.Drawer(
			ui.LazyLoader("editPage", "param1").LoadWhenParentVisible(),
			bo.Drawer(
				ui.LazyLoader("editPage", "param2").LoadWhenParentVisible(),
			).TriggerElement(
				A().Text("New Drawer").Href("#"),
			),
		).TriggerElement(
			A().Text("Edit").Href("#"),
		).Width(500).DefaultOpen(s.drawerVisible, false),
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
	).TriggerElement(
		A().Text("Open " + randStr("inner")).Href("#"),
	).Width(400)
	return
}
