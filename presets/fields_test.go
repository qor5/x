package presets_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/presets"
	h "github.com/theplant/htmlgo"
	"github.com/theplant/testingutils"
)

type Company struct {
	Name      string
	FoundedAt time.Time
}

type Media string

type User struct {
	ID      int
	Int1    int
	Float1  float32
	String1 string
	Bool1   bool
	Time1   time.Time
	Company *Company
	Media1  Media
}

func TestFields(t *testing.T) {

	vd := &web.ValidationErrors{}
	vd.FieldError("String1", "too small")

	ft := NewFieldDefaults(WRITE).Exclude("ID")
	ft.FieldType(time.Time{}).ComponentFunc(func(obj interface{}, field *FieldContext, ctx *web.EventContext) h.HTMLComponent {
		return h.Div().Class("time-control").Text(field.Value(obj).(time.Time).Format("2006-01-02")).Attr("field-name", field.Name)
	})

	ft.FieldType(Media("")).ComponentFunc(func(obj interface{}, field *FieldContext, ctx *web.EventContext) h.HTMLComponent {
		if field.ContextValue("a") == nil {
			return h.Text("")
		}
		return h.Text(field.ContextValue("a").(string) + ", " + field.ContextValue("b").(string))
	})

	r := httptest.NewRequest("GET", "/hello", nil)

	ctx := &web.EventContext{R: r}

	user := &User{
		ID:      1,
		Int1:    2,
		Float1:  23.1,
		String1: "hello",
		Bool1:   true,
		Time1:   time.Unix(1567048169, 0),
		Company: &Company{
			Name:      "Company1",
			FoundedAt: time.Unix(1567048169, 0),
		},
	}

	ftRead := NewFieldDefaults(LIST)

	var cases = []struct {
		name           string
		toComponentFun func() h.HTMLComponent
		expect         string
	}{
		{
			name: "Only with additional nested object",
			toComponentFun: func() h.HTMLComponent {
				return ft.InspectFields(&User{}).
					Labels("Int1", "整数1", "Company.Name", "公司名").
					Only("Int1", "Float1", "String1", "Bool1", "Time1", "Company.Name", "Company.FoundedAt").
					ToComponent(
						user,
						vd,
						ctx)
			},
			expect: `
<vw-text-field type='number' field-name='Int1' label='整数1' :value='"2"'></vw-text-field>

<vw-text-field type='number' field-name='Float1' label='Float1' :value='"23.1"'></vw-text-field>

<vw-text-field type='text' field-name='String1' label='String1' :value='"hello"' :error-messages='["too small"]'></vw-text-field>

<vw-checkbox field-name='Bool1' label='Bool1' :input-value='true'></vw-checkbox>

<div field-name='Time1' class='time-control'>2019-08-29</div>

<vw-text-field type='text' field-name='Company.Name' label='公司名' :value='"Company1"'></vw-text-field>

<div field-name='Company.FoundedAt' class='time-control'>2019-08-29</div>
`,
		},

		{
			name: "Except with file glob pattern",
			toComponentFun: func() h.HTMLComponent {
				return ft.InspectFields(&User{}).
					Except("Bool*").
					ToComponent(user, vd, ctx)
			},
			expect: `
<vw-text-field type='number' field-name='Int1' label='Int1' :value='"2"'></vw-text-field>

<vw-text-field type='number' field-name='Float1' label='Float1' :value='"23.1"'></vw-text-field>

<vw-text-field type='text' field-name='String1' label='String1' :value='"hello"' :error-messages='["too small"]'></vw-text-field>

<div field-name='Time1' class='time-control'>2019-08-29</div>
`,
		},

		{
			name: "Read Except with file glob pattern",
			toComponentFun: func() h.HTMLComponent {
				return ftRead.InspectFields(&User{}).
					Except("Float*").ToComponent(user, vd, ctx)
			},
			expect: `
<td>1</td>

<td>2</td>

<td>hello</td>

<td>true</td>
`,
		},

		{
			name: "Read for a time field",
			toComponentFun: func() h.HTMLComponent {
				return ftRead.InspectFields(&User{}).
					Only("Time1", "Int1").ToComponent(user, vd, ctx)
			},
			expect: `
<td>2019-08-29 11:09:29 +0800 CST</td>

<td>2</td>
`,
		},

		{
			name: "pass in context",
			toComponentFun: func() h.HTMLComponent {
				fb := ft.InspectFields(&User{}).
					Only("Media1")
				fb.Field("Media1").
					WithContextValue("a", "context value1").
					WithContextValue("b", "context value2")
				return fb.ToComponent(user, vd, ctx)
			},
			expect: `context value1, context value2`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := h.MustString(c.toComponentFun(), web.WrapEventContext(context.TODO(), ctx))
			diff := testingutils.PrettyJsonDiff(c.expect, output)
			if len(diff) > 0 {
				t.Error(c.name, diff)
			}
		})
	}

}
