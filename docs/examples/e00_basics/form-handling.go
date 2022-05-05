package e00_basics

// @snippet_begin(FormHandlingSample)
import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/goplaid/web"
	"github.com/goplaid/x/docs/utils"
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
	HiddenValue1   string
}

func FormHandlingPage(ctx *web.EventContext) (pr web.PageResponse, err error) {

	var fv MyData
	err = ctx.UnmarshalForm(&fv)
	if fv.Text1 == "" {
		fv.Text1 = `Hello '1
World`
	}

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
			Input("").Type("text").Value(fv.Text1).Attr(web.VFieldName("Text1")...),
		),
		Div(
			Label("Checkbox1"),
			Input("").Type("checkbox").Value("1").Checked(fv.Checkbox1 == "1").Attr(web.VFieldName("Checkbox1")...),
		),

		web.Scope(
			Fieldset(
				Legend("Nested Form"),

				Div(
					Label("Color1"),
					Input("").Type("color").
						Value(fv.Color1).
						Attr(web.VFieldName("Color1")...),
				),
				Div(
					Label("Email1"),
					Input("").Type("email").Value(fv.Email1).Attr(web.VFieldName("Email1")...),
				),

				Input("").Type("checkbox").
					Attr("v-model", "locals.checked").
					Attr(web.VFieldName("Checked123")...),

				Button("Uncheck it").Attr("@click", "locals.checked = false"),
				Hr(),
				Button("Send").Attr("@click", web.POST().
					EventFunc("checkvalue").
					Query("id", 123).
					FieldValue("name", "azuma").
					Go()),
			),
		).VSlot("{ plaidForm, locals }").Init("{checked: true}"),
		web.Scope(
			Fieldset(
				Legend("Nested Form 2"),

				Div(
					Label("Email1"),
					Input("").Type("email").Value(fv.Email1).Attr(web.VFieldName("Email1")...),
				),

				Button("Send").Attr("@click", web.POST().
					EventFunc("checkvalue").
					Go()),
			),
		).VSlot("{ plaidForm, locals }").Init("{checked: true}"),
		Div(
			Fieldset(
				Legend("Radio"),
				Label("Radio Value 1"),
				Input("Radio1").Type("radio").
					Value("1").Checked(fv.Radio1 == "1").Attr(web.VFieldName("Radio1")...),
				Label("Radio Value 2"),
				Input("Radio1").Type("radio").
					Value("2").Checked(fv.Radio1 == "2").Attr(web.VFieldName("Radio1")...),
			),
		),
		Div(
			Label("Range1"),
			Input("").Type("range").Value(fmt.Sprint(fv.Range1)).Attr(web.VFieldName("Range1")...),
		),

		web.Scope(
			Div(
				Label("Url1"),
				Input("").Type("url").Value(fv.Url1).Attr(web.VFieldName("Url1")...),
			),
			Div(
				Label("Tel1"),
				Input("").Type("tel").Value(fv.Tel1).Attr(web.VFieldName("Tel1")...),
			),
			Div(
				Label("Month1"),
				Input("").Type("month").Value(fv.Month1).Attr(web.VFieldName("Month1")...),
			),
		).VSlot("{ locals }"),

		Div(
			Label("Time1"),
			Input("").Type("time").Value(fv.Time1).Attr(web.VFieldName("Time1")...),
		),
		Div(
			Label("Week1"),
			Input("").Type("week").Value(fv.Week1).Attr(web.VFieldName("Week1")...),
		),
		Div(
			Label("DatetimeLocal1"),
			Input("").Type("datetime-local").Value(fv.DatetimeLocal1).Attr(web.VFieldName("DatetimeLocal1")...),
		),
		Div(
			Label("File1"),
			Input("").Type("file").Value("").Attr(web.VFieldName("File1")...),
		),
		Div(
			Label("Hidden values with default"),
			Input("").Type("hidden").Value(`hidden value 
'123`).Attr(web.VFieldName("HiddenValue1")...),
		),
		Div(
			Button("Submit").Attr("@click", web.POST().EventFunc("checkvalue").Go()),
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

var FormHandlingPagePB = web.Page(FormHandlingPage).
	EventFunc("checkvalue", checkvalue)

const FormHandlingPagePath = "/samples/form_handling"

// @snippet_end
