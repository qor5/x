package gormx

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type omitAssociationsPlugin struct{}

// OmitAssociationsPlugin automatically omits associations for Create, Update and Delete operations
var OmitAssociationsPlugin gorm.Plugin = &omitAssociationsPlugin{}

func (p *omitAssociationsPlugin) Name() string {
	return "omit-associations-plugin"
}

func (p *omitAssociationsPlugin) Initialize(db *gorm.DB) error {
	omitAssociations := func(db *gorm.DB) {
		db.Statement.Omit(clause.Associations)
	}

	if err := db.Callback().Create().Before("gorm:before_create").Register("omit_associations", omitAssociations); err != nil {
		return errors.Wrap(err, "failed to register omit associations callback for Create operation")
	}

	if err := db.Callback().Delete().Before("gorm:before_delete").Register("omit_associations", omitAssociations); err != nil {
		return errors.Wrap(err, "failed to register omit associations callback for Delete operation")
	}

	if err := db.Callback().Update().Before("gorm:before_update").Register("omit_associations", omitAssociations); err != nil {
		return errors.Wrap(err, "failed to register omit associations callback for Update operation")
	}

	return nil
}
