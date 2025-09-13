package gobusx

import (
	"context"

	"github.com/pkg/errors"
	"github.com/qor5/go-bus/pgbus"
	"gorm.io/gorm"
)

type migrator struct{}

func Migrate(ctx context.Context, db DB) (*migrator, error) {
	sqlDB, err := (*gorm.DB)(db).DB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get database connection")
	}
	err = pgbus.Migrate(ctx, sqlDB)
	if err != nil {
		return nil, err
	}
	return &migrator{}, nil
}
