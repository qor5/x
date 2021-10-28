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

type OverlayType string

const (
	Dialog OverlayType = "dialog"
	Drawer OverlayType = "drawer"
)

func FormEventFunc(eventFuncID string, overlayType OverlayType, params ...string) *web.VueEventTagBuilder {
	return web.Plaid().EventFunc(eventFuncID, append([]string{string(overlayType)}, params...)...)
}
