package presets

const (
	PermModule = "presets"
	PermList   = "list"
	PermGet    = "get"
	PermCreate = "create"
	PermUpdate = "update"
	PermDelete = "delete"
)

var (
	PermRead = []string{PermList, PermGet}
)
