package e21_presents

import (
	"github.com/goplaid/x/presets"
	"gorm.io/gorm"
)

// @snippet_begin(PresetsPermissionsSample)

func PresetsPermissions(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, dp, db = PresetsDetailPageCards(b)
	b.URIPrefix(PresetsPermissionsPath)

	return
}

const PresetsPermissionsPath = "/samples/presets-permissions"

// @snippet_end
