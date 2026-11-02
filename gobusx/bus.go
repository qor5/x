package gobusx

import (
	"context"

	"github.com/pkg/errors"
	"github.com/qor5/go-bus"
	"github.com/qor5/go-bus/pgbus"
	"github.com/qor5/x/v3/gormx"
	"github.com/theplant/inject/lifecycle"
	"gorm.io/gorm"
)

type Config struct {
	Database gormx.DatabaseConfig `confx:"database"`
}

var SetupBus = []any{
	SetupDatabase,
	CreateBus,
}

type DB *gorm.DB

var SetupDatabase = SetupDatabaseFactory("bus-database")

func SetupDatabaseFactory(name string) func(ctx context.Context, lc *lifecycle.Lifecycle, conf *Config) (DB, error) {
	return func(ctx context.Context, lc *lifecycle.Lifecycle, conf *Config) (DB, error) {
		return gormx.SetupDatabaseFactory(name)(ctx, lc, &conf.Database)
	}
}

func CreateBus(db DB) (bus.Bus, error) {
	sqlDB, err := (*gorm.DB)(db).DB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get database connection")
	}
	b, err := pgbus.New(sqlDB, bus.WithMigrate(false))
	if err != nil {
		return nil, err
	}
	return b, nil
}
