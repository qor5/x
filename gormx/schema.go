package gormx

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ParseSchema(db *gorm.DB, model any) (*schema.Schema, error) {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(model); err != nil {
		return nil, errors.Wrap(err, "parse statement")
	}
	return stmt.Schema, nil
}
