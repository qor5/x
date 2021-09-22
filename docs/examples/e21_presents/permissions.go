package e21_presents

import (
	"net/http"

	"github.com/goplaid/x/perm"
	"github.com/goplaid/x/presets"
	"gorm.io/gorm"
)

// @snippet_begin(PresetsPermissionsSample)
type User struct {
	ID       uint
	Username string
}

type Group struct {
	ID   uint
	Name string
}

func PresetsPermissions(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, dp, db = PresetsDetailPageCards(b)
	b.URIPrefix(PresetsPermissionsPath)
	perm.Verbose = true
	b.Permission(perm.New().
		Policies(
			perm.PolicyFor("editor").WhoAre(perm.Allowed).ToDo(perm.Anything).On(perm.Anything),
			perm.PolicyFor("editor").WhoAre(perm.Denied).ToDo(presets.PermRead...).On("*user_management*"),
			perm.PolicyFor("editor").WhoAre(perm.Denied).
				ToDo(presets.PermCreate, presets.PermDelete).On("*customers*"),
			perm.PolicyFor("editor").WhoAre(perm.Denied).
				ToDo(presets.PermCreate, presets.PermUpdate).On("*companies*"),
			perm.PolicyFor("editor").WhoAre(perm.Denied).
				ToDo(presets.PermGet).On("*customers:*:company_id*"),
		).
		SubjectsFunc(func(r *http.Request) []string {
			return []string{"editor"}
		}))

	err := db.AutoMigrate(&User{}, &Group{})
	if err != nil {
		panic(err)
	}

	b.Model(&User{}).MenuGroup("User Management")
	b.Model(&Group{}).MenuGroup("User Management")
	return
}

const PresetsPermissionsPath = "/samples/presets-permissions"

// @snippet_end
