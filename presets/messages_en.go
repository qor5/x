package presets

import (
	"fmt"

	"github.com/sunfmin/bran/ui"
)

type Messages struct {
	SuccessfullyUpdated string
	EditingObjectTitle  func(label string) string
}

type MessagesFunc func(ctx *ui.EventContext) *Messages

var Messages_en_US = Messages{
	SuccessfullyUpdated: "Successfully Updated",
	EditingObjectTitle: func(label string) string {
		return fmt.Sprintf("Editing %s", label)
	},
}

func defaultMessageFunc(ctx *ui.EventContext) *Messages {
	msg := Messages_en_US
	return &msg
}
