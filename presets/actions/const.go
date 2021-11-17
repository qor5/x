package actions

import (
	"encoding/json"
)

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

type OverlayOptionsBuilder struct {
	Type       string
	NextScript string
}

func (opts *OverlayOptionsBuilder) String() string {
	r, _ := json.Marshal(opts)
	return string(r)
}

func ParamAsOptions(param string) (r OverlayOptionsBuilder) {
	if param == "" {
		return
	}

	if param[0] == '{' {
		_ = json.Unmarshal([]byte(param), &r)
		return
	}

	r.Type = param
	return
}

func OptionType(v string) (r *OverlayOptionsBuilder) {
	r = &OverlayOptionsBuilder{Type: v}
	return
}

func (b *OverlayOptionsBuilder) SetNextScript(v string) (r *OverlayOptionsBuilder) {
	b.NextScript = v
	return b
}
