package presets_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sunfmin/bran/ui"

	"github.com/theplant/htmlgo"

	"github.com/sunfmin/bran/presets"
)

type Company struct {
	Name      string
	FoundedAt time.Time
}

type User struct {
	ID      int
	Int1    int
	Float1  float32
	String1 string
	Bool1   bool
	Time1   time.Time
	Company *Company
}

func TestFields(t *testing.T) {

	vd := &presets.ValidationErrors{}
	_ = vd.FieldError("String1", "too small")

	ft := presets.NewFieldTypes(presets.WRITE).Exclude("ID")
	r := httptest.NewRequest("GET", "/hello", nil)

	ctx := &ui.EventContext{R: r}

	user := &User{
		ID:      1,
		Int1:    2,
		Float1:  23.1,
		String1: "hello",
		Bool1:   true,
		Time1:   time.Now(),
		Company: &Company{
			Name:      "Company1",
			FoundedAt: time.Now(),
		},
	}

	output := htmlgo.MustString(
		ft.InspectFields(&User{}).
			Only("Int1", "Float1", "String1", "Bool1", "Time1", "Company.Name", "Company.FoundedAt").
			ToComponent(
				user,
				vd,
				ctx),
		context.TODO(),
	)

	fmt.Println(output)

	output = htmlgo.MustString(
		ft.InspectFields(&User{}).
			Except("Company*").
			ToComponent(user, vd, ctx),
		context.TODO(),
	)

	fmt.Println(output)

	ftRead := presets.NewFieldTypes(presets.LIST)

	output = htmlgo.MustString(
		ftRead.InspectFields(&User{}).
			Except("Company*").ToComponent(user, vd, ctx),
		ui.WrapEventContext(context.TODO(), ctx),
	)

	fmt.Println(output)
}
