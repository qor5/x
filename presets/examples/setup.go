package examples

import (
	"fmt"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sunfmin/bran/presets"
	"github.com/sunfmin/bran/presets/gormop"
	s "github.com/sunfmin/bran/stripeui"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type Thumb struct {
	Name string
}

type Customer struct {
	ID              int
	Name            string
	Email           string
	Description     string
	Thumb1          *Thumb
	CompanyID       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ApprovedAt      *time.Time
	ApprovalComment string
	LanguageCode    string
}

type Note struct {
	ID         int
	SourceType string
	SourceID   int
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreditCard struct {
	ID              int
	CustomerID      int
	Number          string
	ExpireYearMonth string
	Name            string
	Type            string
	Phone           string
	Email           string
}

type Payment struct {
	ID                   int
	CustomerID           int
	CurrencyCode         string
	Amount               int
	PaymentMethodID      int
	StatementDescription string
	Description          string
	AuthorizeOnly        bool
	CreatedAt            time.Time
}

type Event struct {
	ID          int
	SourceType  string // Payment, Customer
	SourceID    int
	CreatedAt   time.Time
	Type        string
	Description string
}

type Language struct {
	Code string
	Name string
}

type Company struct {
	ID   int
	Name string
}

type Product struct {
	ID        int
	Name      string
	OwnerName string
}

func Preset1(db *gorm.DB) (r *presets.Builder) {
	err := db.AutoMigrate(
		&Customer{},
		&Note{},
		&CreditCard{},
		&Payment{},
		&Event{},
		&Company{},
		&Product{},
		&Language{},
	).Error
	if err != nil {
		panic(err)
	}

	p := presets.New().URIPrefix("/admin").PrimaryColor("cyan darken-3")

	p.BrandFunc(func(ctx *ui.EventContext) h.HTMLComponent {
		return h.Components(
			//h.Img("https://material.io/tools/icons/static/ic_material_192px_light.svg").Style("height: 32px"),
			VIcon("directions_boat").Class("pr-2"),
			VToolbarTitle("My Admin"),
		)
	}).BrandTitle("The Plant")

	p.FieldType(&Thumb{}).
		ComponentFunc(presets.LISTING, func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
			i, err := reflectutils.Get(obj, field.Name)
			if err != nil {
				panic(err)
			}
			return h.Text(i.(*Thumb).Name)
		}).
		ComponentFunc(presets.DETAILING, func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
			i, err := reflectutils.Get(obj, field.Name)
			if err != nil {
				panic(err)
			}
			return h.Text(i.(*Thumb).Name)
		}).
		ComponentFunc(presets.EDITING, func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
			i, err := reflectutils.Get(obj, field.Name)
			if err != nil {
				panic(err)
			}
			return h.Text(i.(*Thumb).Name)
		})

	p.DataOperator(gormop.DataOperator(db))

	p.MenuGroup("Customer Management").Icon("group")
	mp := p.Model(&Product{}).MenuIcon("laptop")
	mp.Listing().PerPage(3)

	m := p.Model(&Customer{}).URIName("customer").MenuGroup("Customer Management")
	p.Model(&Company{}).MenuGroup("Customer Management")
	m.Labels(
		"Name", "名字",
		"Bool1", "性别",
		"Float1", "体重",
		"CompanyID", "公司",
	).Placeholders(
		"Name", "请输入你的名字",
	)

	l := m.Listing("Name", "CompanyID", "ApprovalComment").SearchColumns("name", "job_title").PerPage(5)
	l.Field("Name").Label("列表的名字").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*Customer)
		return h.Td(ui.Bind(h.A().Text(u.Name)).PushStateURL(fmt.Sprintf("/admin/customers/%d/edit", u.ID)))
	})

	l.Field("CompanyID").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*Customer)
		var comp Company
		err := db.Find(&comp, u.CompanyID).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			panic(err)
		}
		return h.Td(ui.Bind(
			h.A().Text(comp.Name)).
			URL("/admin/companies").
			EventFunc("formDrawerEdit", fmt.Sprint(comp.ID)))
	})

	l.BulkAction("Approve").Label("Approve").UpdateFunc(func(selectedIds []string, form *multipart.Form, ctx *ui.EventContext) (err error) {
		comment := ctx.R.FormValue("ApprovalComment")
		if len(comment) < 10 {
			ctx.Flash = "comment should larger than 10"
			return
		}
		err = db.Model(&Customer{}).
			Where("id IN (?)", selectedIds).
			Updates(map[string]interface{}{"approved_at": time.Now(), "approval_comment": comment}).Error
		if err != nil {
			ctx.Flash = err.Error()
		}
		return
	}).ComponentFunc(func(selectedIds []string, ctx *ui.EventContext) h.HTMLComponent {
		comment := ctx.R.FormValue("ApprovalComment")
		errorMessage := ""
		if ctx.Flash != nil {
			errorMessage = ctx.Flash.(string)
		}
		return VTextField().
			FieldName("ApprovalComment").
			Value(comment).
			Label("Comment").
			Error(len(errorMessage) > 0).
			ErrorMessages(errorMessage)
	})

	l.BulkAction("Delete").Label("Delete").UpdateFunc(func(selectedIds []string, form *multipart.Form, ctx *ui.EventContext) (err error) {
		err = db.Where("id IN (?)", selectedIds).Delete(&Customer{}).Error
		return
	}).ComponentFunc(func(selectedIds []string, ctx *ui.EventContext) h.HTMLComponent {
		return h.Div().Text(fmt.Sprintf("Are you sure you want to delete %s ?", selectedIds)).Class("title deep-orange--text")
	})

	l.FilterDataFunc(func(ctx *ui.EventContext) FilterData {
		var companyOptions []*SelectItem
		err := db.Model(&Company{}).Select("name as text, id as value").Scan(&companyOptions).Error
		if err != nil {
			panic(err)
		}

		return []*FilterItem{
			{
				Key:          "created",
				Label:        "Created",
				ItemType:     ItemTypeDate,
				SQLCondition: `extract(epoch from created_at) %s ?`,
			},
			{
				Key:          "approved",
				Label:        "Approved",
				ItemType:     ItemTypeDate,
				SQLCondition: `extract(epoch from approved_at) %s ?`,
			},
			{
				Key:          "name",
				Label:        "Name",
				ItemType:     ItemTypeString,
				SQLCondition: `name %s ?`,
			},
			{
				Key:          "company",
				Label:        "Company",
				ItemType:     ItemTypeSelect,
				SQLCondition: `company_id %s ?`,
				Options:      companyOptions,
			},
		}
	})

	l.FilterTabsFunc(func(ctx *ui.EventContext) []*presets.FilterTab {
		var c Company
		db.First(&c)
		return []*presets.FilterTab{
			{
				Label: "Felix",
				Query: url.Values{"name.ilike": []string{"felix"}},
			},
			{
				Label: "The Plant",
				Query: url.Values{"company": []string{fmt.Sprint(c.ID)}},
			},
			{
				Label: "Approved",
				Query: url.Values{"approved.gt": []string{fmt.Sprint(1)}},
			},
			{
				Label: "All",
				Query: url.Values{"all": []string{"1"}},
			},
		}
	})

	ef := m.Editing("Name", "CompanyID", "LanguageCode")
	ef.Field("LanguageCode").Label("语言").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*Customer)
		var langs []Language
		err := db.Find(&langs).Error
		if err != nil {
			panic(err)
		}
		return VAutocomplete().
			FieldName(field.Name).
			Label(field.Label).
			Items(langs).
			ItemText("Name").
			ItemValue("Code").
			Multiple(false).
			Value(u.LanguageCode)
	})

	ef.Field("CompanyID").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*Customer)
		var companies []*Company
		err := db.Find(&companies).Error
		if err != nil {
			panic(err)
		}
		return VSelect().
			FieldName("CompanyID").
			Label(field.Label).
			Items(companies).
			ItemText("Name").
			ItemValue("ID").
			Multiple(false).
			Value(u.CompanyID)
	})

	dp := m.Detailing("MainInfo", "Details", "Cards", "Payments", "Events")

	dp.Field("MainInfo").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {

		cu := obj.(*Customer)

		title := cu.Name
		if len(title) == 0 {
			title = cu.Description
		}

		var notes []*Note
		err := db.Where("source_type = 'Customer' AND source_id = ?", cu.ID).
			Order("id DESC").
			Find(&notes).Error
		if err != nil {
			panic(err)
		}

		dt := s.DataTable(notes).WithoutHeader(true).LoadMoreAt(2, "Show More")

		dt.Column("Content").CellComponentFunc(func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent {
			n := obj.(*Note)
			return h.Td(h.Div(
				h.Div(
					VIcon("comment").Color("blue").Small(true).Class("pr-2"),
					h.Text(n.Content),
				).Class("body-1"),
				h.Div(
					h.Text(n.CreatedAt.Format("Jan 02,15:04 PM")),
					h.Text(" by Felix Sun"),
				).Class("grey--text pl-7 body-2"),
			).Class("my-3"))
		})

		cusID := fmt.Sprint(cu.ID)
		dt.RowMenuItemsFunc(func(obj interface{}, id string, ctx *ui.EventContext) []h.HTMLComponent {
			return []h.HTMLComponent{
				ui.Bind(VListItem(
					VListItemIcon(VIcon("edit")),
					VListItemTitle(h.Text("Edit")),
				)).OnClick("formDrawerEdit", id, "Customer", cusID).URL("/admin/notes"),

				ui.Bind(VListItem(
					VListItemIcon(VIcon("delete")),
					VListItemTitle(h.Text("Delete")),
				)).OnClick("deleteConfirmation", cusID).URL("/admin/notes"),
			}
		})

		return s.Card(
			dt,
		).HeaderTitle(title).
			Actions(
				ui.Bind(VBtn("Add Note").
					Depressed(true)).OnClick(
					"formDrawerNew",
					"",
					"Customer",
					cusID,
				).URL("/admin/notes"),
			).Class("mb-4")
	})

	dp.Field("Details").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		cu := obj.(*Customer)
		cusID := fmt.Sprint(cu.ID)

		var lang Language
		db.Where("code = ?", cu.LanguageCode).First(&lang)

		detail := s.DetailInfo(
			s.DetailColumn(
				s.DetailField(s.OptionalText(cu.Name).ZeroLabel("No Name")).Label("Name"),
				s.DetailField(s.OptionalText(cu.Email).ZeroLabel("No Email")).Label("Email"),
				s.DetailField(s.OptionalText(cu.Description).ZeroLabel("No Description")).Label("Description"),
				s.DetailField(s.OptionalText(cusID).ZeroLabel("No ID")).Label("ID"),
				s.DetailField(s.OptionalText(cu.CreatedAt.Format("Jan 02,15:04 PM")).ZeroLabel("")).Label("Created"),
				s.DetailField(s.OptionalText(lang.Name).ZeroLabel("No Language")).Label("Language"),
			).Header("ACCOUNT INFORMATION"),
			s.DetailColumn().Header("BILLING INFORMATION"),
		)

		return s.Card(detail).HeaderTitle("Details").
			Actions(
				ui.Bind(VBtn("Update details").
					Depressed(true)).OnClick(
					"formDrawerEdit",
					cusID,
				).URL("/admin/customers"),
			)
	})

	p.Model(&Note{}).
		InMenu(false).
		Editing("Content").
		SetterFunc(func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) {
			note := obj.(*Note)
			note.SourceID = ctx.Event.ParamAsInt(2)
			note.SourceType = ctx.Event.Params[1]
		})

	p.Model(&Language{}).PrimaryField("Code")

	ie := dp.Field("Bool1").InplaceEdit()
	ie.ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		//u := obj.(*Customer)
		return VCheckbox().FieldName("Bool1")
	}).UpdateFunc(func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) (err error) {
		u := obj.(*Customer)
		err = db.Model(&Customer{}).UpdateColumn("Name", u.Name).Error
		return
	})

	dp.Action("Approve").UpdateFunc(func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) (err error) {
		u := obj.(*Customer)
		err = db.Model(&Customer{}).Where("id = ?", u.ID).UpdateColumn("approved_at = ?", time.Now()).Error
		return
	}).ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		return VBtn("Approve")
	})
	return p
}
