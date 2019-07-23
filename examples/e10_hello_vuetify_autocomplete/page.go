package e10_hello_vuetify_autocomplete

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/sunfmin/bran/ui"
	vt "github.com/sunfmin/bran/vuetify"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	Values1 []string
	Values2 []string
	Value3  string
}

type User struct {
	Login string
	Name  string
}

var selectedItems1 = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "charles", Name: "Charles"},
}

var options1 = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "john", Name: "John"},
	{Login: "charles", Name: "Charles"},
}

var selectedItems2 = []*User{
	{Login: "charles", Name: "Charles"},
}

var selectedItems3 = []*User{
	{Login: "charles", Name: "Charles"},
}

var options2 = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "john", Name: "John"},
	{Login: "charles", Name: "Charles"},
}

func HelloVuetifyAutocomplete(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("update", update)
	s := ctx.StateOrInit(&mystate{
		Values1: []string{
			"sam",
			"charles",
		},
		Values2: []string{
			"charles",
		},
		Value3: "charles",
	}).(*mystate)

	result := Ul()
	for _, v := range s.Values1 {
		result.AppendChildren(Li().Text(v))
	}
	result.AppendChildren(Li().Text("======"))
	for _, v := range s.Values2 {
		result.AppendChildren(Li().Text(v))
	}
	pr.Schema = vt.VApp(
		vt.VContent(
			vt.VContainer(
				H1("VAutocomplete"),
				vt.VAutocomplete().
					Items(options1).
					FieldName("Values1").
					ItemText("Name").
					ItemValue("Login").
					Label("Static Options"),

				vt.VAutocomplete().
					ItemsEventFunc(ctx.Hub, "users", users).
					ItemText("Name").
					ItemValue("Login").
					SelectedItems(selectedItems2).
					FieldName("Values2").
					Label("Load Options from Remote"),

				result,
				H1("VSelect"),
				vt.VSelect().
					Items(options1).
					ItemText("Name").
					ItemValue("Login").
					FieldName("Value3").
					Solo(true),
				Pre(s.Value3),
				vt.VBtn("Update").
					Color("success").
					OnClick("update"),
			),
		),
	)
	return
}

func users(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	us := []*User{}
	for _, u := range options1 {
		us = append(us, u)
	}
	if len(options2) <= 100 {
		for i := 0; i < 200; i++ {
			us = append(us, &User{
				Login: randomdata.Email(),
				Name:  randomdata.SillyName(),
			})
		}
		options2 = us
	}

	r.Data = options2
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	selectedItems1 = []*User{}
	for _, login := range s.Values1 {
		for _, u := range options1 {
			if u.Login == login {
				selectedItems1 = append(selectedItems1, u)
			}
		}
	}

	selectedItems2 = []*User{}
	for _, login := range s.Values2 {
		for _, u := range options2 {
			if u.Login == login {
				selectedItems2 = append(selectedItems2, u)
			}
		}
	}

	selectedItems3 = []*User{}
	for _, u := range options1 {
		if u.Login == s.Value3 {
			selectedItems3 = append(selectedItems3, u)
		}
	}
	r.Reload = true

	return
}
