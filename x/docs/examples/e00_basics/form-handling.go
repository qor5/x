package e00_basics

//@snippet_begin(FormHandlingSample)
import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/goplaid/x/docs/utils"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

type MyData struct {
	Text1          string
	Checkbox1      string
	Color1         string
	Email1         string
	Radio1         string
	Range1         int
	Url1           string
	Tel1           string
	Month1         string
	Time1          string
	Week1          string
	DatetimeLocal1 string
	File1          []*multipart.FileHeader
}

func FormHandlingPage(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("checkvalue", checkvalue)

	var fv MyData
	err = ctx.UnmarshalForm(&fv)
	if err != nil {
		panic(err)
	}

	pr.Body = Div(
		H1("Form Handling"),
		H3("Form Content"),
		utils.PrettyFormAsJSON(ctx),
		H3("File1 Content"),
		Pre(fv.File1Bytes()).Style("width: 400px; white-space: pre-wrap;"),
		Div(
			Label("Text1"),
			web.Bind(Input("").Type("text").Value(fv.Text1)).FieldName("Text1"),
		),
		Div(
			Label("Checkbox1"),
			web.Bind(Input("").Type("checkbox").Value("1").Checked(fv.Checkbox1 == "1")).FieldName("Checkbox1"),
		),
		Div(
			Label("Color1"),
			web.Bind(Input("").Type("color").Value(fv.Color1)).FieldName("Color1"),
		),
		Div(
			Label("Email1"),
			web.Bind(Input("").Type("email").Value(fv.Email1)).FieldName("Email1"),
		),
		Div(
			Fieldset(
				Legend("Radio"),
				Label("Radio Value 1"),
				web.Bind(Input("Radio1").Type("radio").
					Value("1").Checked(fv.Radio1 == "1")).FieldName("Radio1"),
				Label("Radio Value 2"),
				web.Bind(Input("Radio1").Type("radio").
					Value("2").Checked(fv.Radio1 == "2")).FieldName("Radio1"),
			),
		),
		Div(
			Label("Range1"),
			web.Bind(Input("").Type("range").Value(fmt.Sprint(fv.Range1))).FieldName("Range1"),
		),
		Div(
			Label("Url1"),
			web.Bind(Input("").Type("url").Value(fv.Url1)).FieldName("Url1"),
		),
		Div(
			Label("Tel1"),
			web.Bind(Input("").Type("tel").Value(fv.Tel1)).FieldName("Tel1"),
		),
		Div(
			Label("Month1"),
			web.Bind(Input("").Type("month").Value(fv.Month1)).FieldName("Month1"),
		),
		Div(
			Label("Time1"),
			web.Bind(Input("").Type("time").Value(fv.Time1)).FieldName("Time1"),
		),
		Div(
			Label("Week1"),
			web.Bind(Input("").Type("week").Value(fv.Week1)).FieldName("Week1"),
		),
		Div(
			Label("DatetimeLocal1"),
			web.Bind(Input("").Type("datetime-local").Value(fv.DatetimeLocal1)).FieldName("DatetimeLocal1"),
		),
		Div(
			Label("File1"),
			web.Bind(Input("").Type("file").Value("")).FieldName("File1"),
		),
		Div(
			web.Bind(
				Button("Submit"),
			).OnClick("checkvalue"),
		),
	)
	return
}

func checkvalue(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Reload = true
	return
}

func (m *MyData) File1Bytes() string {
	if m.File1 == nil || len(m.File1) == 0 {
		return ""
	}
	f, err := m.File1[0].Open()
	if err != nil {
		panic(err)
	}
	var b = make([]byte, 200)
	_, err = io.ReadFull(f, b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%+v ...", b)
}

//@snippet_end

const FormHandlingPagePath = "/samples/form_handling"
