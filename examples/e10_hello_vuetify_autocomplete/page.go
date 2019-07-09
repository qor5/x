package e10_hello_vuetify_autocomplete

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/sunfmin/bran/ui"
	vt "github.com/sunfmin/bran/vuetify"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	MyValues   []string
	YourValues []string
}

type User struct {
	Login string
	Name  string
}

var selectedItems = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "charles", Name: "Charles"},
}

var options = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "john", Name: "John"},
	{Login: "charles", Name: "Charles"},
}

var yourSelectedItems = []*User{
	{Login: "charles", Name: "Charles"},
}

var yourOptions = []*User{}

func HelloVuetifyAutocomplete(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{
		MyValues: []string{
			"sam",
			"charles",
		},
		YourValues: []string{
			"charles",
		},
	}).(*mystate)

	result := Ul()
	for _, v := range s.MyValues {
		result.AppendChildren(Li().Text(v))
	}
	result.AppendChildren(Li().Text("======"))
	for _, v := range s.YourValues {
		result.AppendChildren(Li().Text(v))
	}

	pr.Schema = vt.VApp(
		vt.VContent(
			vt.VContainer(
				vt.VAutoComplete().Items(options).SelectedItems(selectedItems).
					FieldName("MyValues").
					ItemText("Name").
					ItemValue("Login").
					Label("Static Options"),

				vt.VAutoComplete().
					ItemsEventFunc(ctx.Hub, "users", users).
					ItemText("Name").
					ItemValue("Login").
					SelectedItems(yourSelectedItems).
					FieldName("YourValues").
					Label("Load Options from Remote"),

				result,
				vt.VBtn("Update").
					Color("success").
					OnClick(ctx.Hub, "update", update),
			),
		),
	)
	return
}

func users(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	us := []*User{}
	for _, u := range options {
		us = append(us, u)
	}
	if len(yourOptions) <= 100 {
		for i := 0; i < 200; i++ {
			us = append(us, &User{
				Login: randomdata.Email(),
				Name:  randomdata.SillyName(),
			})
		}
		yourOptions = us
	}

	r.Data = yourOptions
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	selectedItems = []*User{}
	for _, login := range s.MyValues {
		for _, u := range options {
			if u.Login == login {
				selectedItems = append(selectedItems, u)
			}
		}
	}

	yourSelectedItems = []*User{}
	for _, login := range s.YourValues {
		for _, u := range yourOptions {
			if u.Login == login {
				yourSelectedItems = append(yourSelectedItems, u)
			}
		}
	}
	r.Reload = true

	return
}
