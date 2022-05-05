package e21_presents

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/actions"
	"github.com/goplaid/x/stripeui"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

// @snippet_begin(PresetsDetailPageTopNotesSample)

type Note struct {
	ID         int
	SourceType string
	SourceID   int
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func PresetsDetailPageTopNotes(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsEditingCustomizationValidation(b)
	b.URIPrefix(PresetsDetailPageTopNotesPath)
	err := db.AutoMigrate(&Note{})
	if err != nil {
		panic(err)
	}

	dp = cust.Detailing("TopNotes", "Details", "Cards")

	dp.Field("TopNotes").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		mi := field.ModelInfo
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

		dt := stripeui.DataTable(notes).WithoutHeader(true).LoadMoreAt(2, "Show More")

		dt.Column("Content").CellComponentFunc(func(obj interface{}, fieldName string, ctx *web.EventContext) h.HTMLComponent {
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
		dt.RowMenuItemFuncs(presets.EditDeleteRowMenuItemFuncs(mi, mi.PresetsPrefix()+"/notes", url.Values{"model": []string{"Customer"}, "model_id": []string{cusID}})...)

		return stripeui.Card(
			dt,
		).HeaderTitle(title).
			Actions(
				VBtn("Add Note").
					Depressed(true).
					Attr("@click",
						web.POST().EventFunc(actions.New).
							Query("model", "Customer").
							Query("model_id", cusID).
							URL(mi.PresetsPrefix()+"/notes").
							Go(),
					),
			).Class("mb-4")
	})

	b.Model(&Note{}).
		InMenu(false).
		Editing("Content").
		SetterFunc(func(obj interface{}, ctx *web.EventContext) {
			note := obj.(*Note)
			note.SourceID = ctx.QueryAsInt("model_id")
			note.SourceType = ctx.R.FormValue("model")
		})
	return
}

const PresetsDetailPageTopNotesPath = "/samples/presets-detail-page-top-notes"

// @snippet_end

// @snippet_begin(PresetsDetailPageDetailsSample)

func PresetsDetailPageDetails(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, dp, db = PresetsDetailPageTopNotes(b)
	b.URIPrefix(PresetsDetailPageDetailsPath)
	err := db.AutoMigrate(&CreditCard{})
	if err != nil {
		panic(err)
	}

	dp.Field("Details").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		mi := field.ModelInfo
		cu := obj.(*Customer)
		cusID := fmt.Sprint(cu.ID)

		var termAgreed string
		if cu.TermAgreedAt != nil {
			termAgreed = cu.TermAgreedAt.Format("Jan 02,15:04 PM")
		}

		detail := stripeui.DetailInfo(
			stripeui.DetailColumn(
				stripeui.DetailField(stripeui.OptionalText(cu.Name).ZeroLabel("No Name")).Label("Name"),
				stripeui.DetailField(stripeui.OptionalText(cu.Email).ZeroLabel("No Email")).Label("Email"),
				stripeui.DetailField(stripeui.OptionalText(cusID).ZeroLabel("No ID")).Label("ID"),
				stripeui.DetailField(stripeui.OptionalText(cu.CreatedAt.Format("Jan 02,15:04 PM")).ZeroLabel("")).Label("Created"),
				stripeui.DetailField(stripeui.OptionalText(termAgreed).ZeroLabel("Not Agreed Yet")).Label("Terms Agreed"),
			).Header("ACCOUNT INFORMATION"),
			stripeui.DetailColumn(
				stripeui.DetailField(h.RawHTML(cu.Description)).Label("Description"),
			).Header("DETAILS"),
		)

		return stripeui.Card(detail).HeaderTitle("Details").
			Actions(
				VBtn("Agree Terms").
					Depressed(true).Class("mr-2").
					Attr("@click", web.POST().
						EventFunc(actions.Action).
						Query(presets.ParamAction, "AgreeTerms").
						Query(presets.ParamID, cusID).
						Go(),
					),

				VBtn("Update details").
					Depressed(true).
					Attr("@click", web.POST().
						EventFunc(actions.Edit).
						Query(presets.ParamOverlay, actions.Dialog).
						Query(presets.ParamID, cusID).
						URL(mi.PresetsPrefix()+"/customers").
						Go(),
					),
			).Class("mb-4")
	})

	dp.Action("AgreeTerms").UpdateFunc(func(selectedIds []string, ctx *web.EventContext) (err error) {
		if ctx.R.FormValue("Agree") != "true" {
			ve := &web.ValidationErrors{}
			ve.GlobalError("You must agree the terms")
			err = ve
			return
		}

		err = db.Model(&Customer{}).Where("id = ?", selectedIds[0]).
			Updates(map[string]interface{}{"term_agreed_at": time.Now()}).Error

		return
	}).ComponentFunc(func(selectedIds []string, ctx *web.EventContext) h.HTMLComponent {
		var alert h.HTMLComponent

		if ve, ok := ctx.Flash.(*web.ValidationErrors); ok {
			alert = VAlert(h.Text(ve.GetGlobalError())).Border("left").
				Type("error").
				Elevation(2).
				ColoredBorder(true)
		}

		return h.Components(
			alert,
			VCheckbox().FieldName("Agree").Value(ctx.R.FormValue("Agree")).Label("Agree the terms"),
		)
	})
	return
}

const PresetsDetailPageDetailsPath = "/samples/presets-detail-page-details"

// @snippet_end

// @snippet_begin(PresetsDetailPageCardsSample)

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

func PresetsDetailPageCards(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, dp, db = PresetsDetailPageDetails(b)
	b.URIPrefix(PresetsDetailPageCardsPath)
	err := db.AutoMigrate(&CreditCard{})
	if err != nil {
		panic(err)
	}

	dp.Field("Cards").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		mi := field.ModelInfo
		cu := obj.(*Customer)
		cusID := fmt.Sprint(cu.ID)

		var cards []*CreditCard
		err := db.Where("customer_id = ?", cu.ID).Order("id ASC").Find(&cards).Error
		if err != nil {
			panic(err)
		}

		dt := stripeui.DataTable(cards).
			WithoutHeader(true).
			RowExpandFunc(func(obj interface{}, ctx *web.EventContext) h.HTMLComponent {
				card := obj.(*CreditCard)
				return stripeui.DetailInfo(
					stripeui.DetailColumn(
						stripeui.DetailField(stripeui.OptionalText(card.Name).ZeroLabel("No Name")).Label("Name"),
						stripeui.DetailField(stripeui.OptionalText(card.Number).ZeroLabel("No Number")).Label("Number"),
						stripeui.DetailField(stripeui.OptionalText(card.ExpireYearMonth).ZeroLabel("No Expires")).Label("Expires"),
						stripeui.DetailField(stripeui.OptionalText(card.Type).ZeroLabel("No Type")).Label("Type"),
						stripeui.DetailField(stripeui.OptionalText(card.Phone).ZeroLabel("No phone provided")).Label("Phone"),
						stripeui.DetailField(stripeui.OptionalText(card.Email).ZeroLabel("No email provided")).Label("Email"),
					),
				)
			}).RowMenuItemFuncs(presets.EditDeleteRowMenuItemFuncs(mi, mi.PresetsPrefix()+"/credit-cards", url.Values{"customerID": []string{cusID}})...)

		dt.Column("Type")
		dt.Column("Number")
		dt.Column("ExpireYearMonth")

		return stripeui.Card(dt).HeaderTitle("Cards").
			Actions(
				VBtn("Add Card").
					Depressed(true).
					Attr("@click",
						web.POST().
							EventFunc(actions.New).
							Query("customerID", cusID).
							URL(mi.PresetsPrefix()+"/credit-cards").
							Go(),
					).Class("mb-4"),
			)
	})

	cc := b.Model(&CreditCard{}).
		InMenu(false)

	ccedit := cc.Editing("ExpireYearMonth", "Phone", "Email").
		SetterFunc(func(obj interface{}, ctx *web.EventContext) {
			card := obj.(*CreditCard)
			card.CustomerID = ctx.QueryAsInt("customerID")
		})

	ccedit.Creating("Number")
	return
}

const PresetsDetailPageCardsPath = "/samples/presets-detail-page-cards"

// @snippet_end
