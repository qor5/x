package actions

import "github.com/goplaid/web"

const (
	New                = "presets_New"
	Edit               = "presets_Edit"
	Action             = "presets_Action"
	DeleteConfirmation = "presets_DeleteConfirmation"
	Update             = "presets_Update"
	DoAction           = "presets_DoAction"
	DoDelete           = "presets_DoDelete"
	DoBulkAction       = "presets_DoBulkAction"
)

const (
	Dialog = "dialog"
	Drawer = "drawer"
)

func FormEventFunc(eventFuncID string, overlayType string, params ...string) *web.VueEventTagBuilder {
	return web.Plaid().EventFunc(eventFuncID, append([]string{overlayType}, params...)...)
}
