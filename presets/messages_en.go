package presets

import (
	"fmt"
	"net/http"
)

type Messages struct {
	SuccessfullyUpdated    string
	New                    string
	Update                 string
	Delete                 string
	Edit                   string
	OK                     string
	Cancel                 string
	Create                 string
	DeleteConfirmationText func(id string) string
	EditingObjectTitle     func(label string) string
	CreatingObjectTitle    func(label string) string
	ListingObjectTitle     func(label string) string
	DetailingObjectTitle   func(label string, name string) string
}

var Messages_en_US = Messages{
	SuccessfullyUpdated: "Successfully Updated",
	EditingObjectTitle: func(label string) string {
		return fmt.Sprintf("Editing %s", label)
	},
	CreatingObjectTitle: func(label string) string {
		return fmt.Sprintf("New %s", label)
	},
	ListingObjectTitle: func(label string) string {
		return fmt.Sprintf("Listing %s", label)
	},
	DetailingObjectTitle: func(label string, name string) string {
		return fmt.Sprintf("%s %s", label, name)
	},
	DeleteConfirmationText: func(id string) string {
		return fmt.Sprintf("Are you sure you want to delete object with id: %s?", id)
	},

	New:    "New",
	Update: "Update",
	Delete: "Delete",
	Edit:   "Edit",
	OK:     "OK",
	Cancel: "Cancel",
	Create: "Create",
}

func defaultMessageFunc(r *http.Request) *Messages {
	msg := Messages_en_US
	return &msg
}
