package presets

import (
	"strings"
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
	CreatingObjectTitle    func(label string) string
	EditingObjectTitle     func(label string, name string) string
	ListingObjectTitle     func(label string) string
	DetailingObjectTitle   func(label string, name string) string
}

var Messages_en_US = &Messages{
	SuccessfullyUpdated: "Successfully Updated",
	EditingObjectTitle: func(label string, name string) string {
		return strings.NewReplacer("{id}", name, "{modelName}", label).
			Replace("Editing {modelName} {id}")
	},
	CreatingObjectTitle: func(label string) string {
		return strings.NewReplacer("{modelName}", label).
			Replace("New {modelName}")
	},
	ListingObjectTitle: func(label string) string {
		return strings.NewReplacer("{modelName}", label).
			Replace("Listing {modelName}")
	},
	DetailingObjectTitle: func(label string, name string) string {
		return strings.NewReplacer("{id}", name, "{modelName}", label).
			Replace("{modelName} {id}")
	},
	DeleteConfirmationText: func(id string) string {
		return strings.NewReplacer("{id}", id).
			Replace("Are you sure you want to delete object with id: {id}?")
	},

	New:    "New",
	Update: "Update",
	Delete: "Delete",
	Edit:   "Edit",
	OK:     "OK",
	Cancel: "Cancel",
	Create: "Create",
}
