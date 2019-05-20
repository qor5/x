package e08_hello_popover

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sunfmin/bran/ui"
	bo "github.com/sunfmin/branoverlay"
	. "github.com/theplant/htmlgo"
)

type popoverState struct {
	Visible     bool
	EditingName string
	NameError   string
}

type mystate struct {
	Popover1 *popoverState
	Popover2 *popoverState
	Popover3 *popoverState
}

var globalState = &struct {
	Name string
}{}

func randStr(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s: %d", prefix, rand.Int31n(100))
}

func overlay(s *popoverState, ctx *ui.EventContext, subState string) HTMLComponent {
	return Div(
		ui.Bind(Input("").Type("text").Value(s.EditingName)).FieldName(fmt.Sprintf("%s.EditingName", subState)),
		Label(s.NameError).Style("color:red"),
		ui.Bind(Button("Update")).OnClick(ctx.Hub, "update", update, subState),
	).Style("padding: 20px; background-color: white;")
}

func HelloPopover(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{
		Popover1: &popoverState{
			EditingName: globalState.Name,
		},
		Popover2: &popoverState{
			EditingName: globalState.Name,
		},
		Popover3: &popoverState{
			EditingName: globalState.Name,
		},
	}).(*mystate)

	pr.Schema = Div(
		H1(globalState.Name),

		H2("Default"),
		bo.Popover(
			overlay(s.Popover1, ctx, "Popover1"),
		).TriggerElement(
			A().Text("Edit").Href("#"),
		).DefaultVisible(s.Popover1.Visible).Placement("right"),

		H2("Load from remote"),
		bo.Popover(
			ui.LazyLoader(ctx.Hub, "remoteOverlay", remoteOverlay).
				Visible("true").
				ParentForceUpdateAfterLoaded(),
		).TriggerElement(
			A().Text("Remote Loader").Href("#"),
		).DefaultVisible(s.Popover2.Visible).Placement("right"),

		H2("Load when mouse over"),
		bo.Popover(
			overlay(s.Popover3, ctx, "Popover3"),
		).TriggerElement(
			A().Text("Mouseover").Href("#"),
		).DefaultVisible(s.Popover3.Visible).
			Trigger("hover").Placement("bottom"),
	)
	return
}

func remoteOverlay(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	r.Schema = overlay(s.Popover2, ctx, "Popover2")
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	subState := ctx.Event.Params[0]
	s := ctx.SubStateOrInit(subState, &popoverState{}).(*popoverState)
	if len(s.EditingName) < 10 {
		s.NameError = "name is too short"
		s.Visible = true
	} else {
		globalState.Name = s.EditingName
		s.NameError = ""
		s.Visible = false
	}
	return
}
