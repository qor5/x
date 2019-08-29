package presets_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sunfmin/bran/presets"
	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
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

	ft := presets.NewFieldDefaults(presets.WRITE).Exclude("ID")
	ft.FieldType(time.Time{}).ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		return h.Div().Class("time-control").Text(field.Value(obj).(time.Time).Format("2006-01-02"))
	})

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

	output := h.MustString(
		ft.InspectFields(&User{}).
			Labels("Int1", "整数1", "Company.Name", "公司名").
			Only("Int1", "Float1", "String1", "Bool1", "Time1", "Company.Name", "Company.FoundedAt").
			ToComponent(
				user,
				vd,
				ctx),
		context.TODO(),
	)

	fmt.Println(output)

	output = h.MustString(
		ft.InspectFields(&User{}).
			Except("Company*").
			ToComponent(user, vd, ctx),
		context.TODO(),
	)

	fmt.Println(output)

	ftRead := presets.NewFieldDefaults(presets.LIST)

	output = h.MustString(
		ftRead.InspectFields(&User{}).
			Except("Company*").ToComponent(user, vd, ctx),
		ui.WrapEventContext(context.TODO(), ctx),
	)

	fmt.Println(output)
}
