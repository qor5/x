package presets

const (
	PermModule = "presets"
	PermList   = "presets:list"
	PermGet    = "presets:get"
	PermCreate = "presets:create"
	PermUpdate = "presets:update"
	PermDelete = "presets:delete"
)

var (
	PermRead = []string{PermList, PermGet}
)

// params
const (
	ParamID             = "id"
	ParamAction         = "action"
	ParamOverlay        = "overlay"
	ParamBulkActionName = "bulk_action"
	ParamSelectedIds    = "selected_ids"
)
