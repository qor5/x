package presets

import (
	"fmt"

	"github.com/sunfmin/bran/ui"
)

type Messages struct {
	SuccessfullyUpdated string
	New                 string
	Update              string
	EditingObjectTitle  func(label string) string
	ListingObjectTitle  func(label string) string
}

var Messages_en_US = Messages{
	SuccessfullyUpdated: "Successfully Updated",
	EditingObjectTitle: func(label string) string {
		return fmt.Sprintf("Editing %s", label)
	},
	ListingObjectTitle: func(label string) string {
		return fmt.Sprintf("Listing %s", label)
	},
	New:    "New",
	Update: "Update",
}

func defaultMessageFunc(ctx *ui.EventContext) *Messages {
	msg := Messages_en_US
	return &msg
}
