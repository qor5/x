package gormx_test

import (
	"context"
	"testing"

	"github.com/qor5/x/v3/gormx"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type TemplatePoolWidget struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

// seedWidget is the row the migrator writes into the template. Every fork must
// see it without re-running the migration.
const seedWidget = "seeded-in-template"

func migrateWidgets(_ context.Context, db *gorm.DB, dsn string) error {
	if dsn == "" {
		return gorm.ErrInvalidDB
	}
	if err := db.AutoMigrate(&TemplatePoolWidget{}); err != nil {
		return err
	}
	return db.Create(&TemplatePoolWidget{ID: seedWidget, Name: "template"}).Error
}

// Declared at package scope, exactly as a consumer would: nothing starts until
// the first Fork.
var widgetPool = gormx.NewTemplatePool(nil, migrateWidgets)

func TestTemplatePool(t *testing.T) {
	t.Cleanup(func() { require.NoError(t, widgetPool.Close(context.Background())) })

	// Reaching a fork at all proves the migrator's connection was released:
	// CREATE DATABASE ... TEMPLATE fails outright while any session is still
	// connected to the source.
	t.Run("fork inherits the migrated schema and seed", func(t *testing.T) {
		db := widgetPool.Fork(t)

		var got TemplatePoolWidget
		require.NoError(t, db.Where("id = ?", seedWidget).First(&got).Error)
		require.Equal(t, "template", got.Name)
	})

	t.Run("forks do not see each other's writes", func(t *testing.T) {
		a, b := widgetPool.Fork(t), widgetPool.Fork(t)

		require.NoError(t, a.Create(&TemplatePoolWidget{ID: "only-in-a", Name: "a"}).Error)

		var count int64
		require.NoError(t, b.Model(&TemplatePoolWidget{}).Where("id = ?", "only-in-a").Count(&count).Error)
		require.Zero(t, count, "a write in one fork must not be visible in another")

		// The seed is still there in both — isolation is not achieved by
		// handing one of them an empty database.
		require.NoError(t, b.Model(&TemplatePoolWidget{}).Where("id = ?", seedWidget).Count(&count).Error)
		require.Equal(t, int64(1), count)
	})

	// Dropping is not cosmetic: databases left behind make every later fork's
	// suite slower, because the autovacuum launcher round-robins over every
	// database that exists. This runs after the subtests above, so their three
	// forks must be gone and only this one may remain.
	t.Run("finished forks are dropped", func(t *testing.T) {
		db := widgetPool.Fork(t)

		var live []string
		require.NoError(t, db.Raw(
			`SELECT datname FROM pg_database WHERE datname LIKE 'fork_%' ORDER BY datname`,
		).Scan(&live).Error)
		require.Len(t, live, 1, "only the current fork should still exist, found %v", live)
	})

	t.Run("max_connections is raised above the stock 100", func(t *testing.T) {
		db := widgetPool.Fork(t)

		var maxConnections string
		require.NoError(t, db.Raw("SHOW max_connections").Scan(&maxConnections).Error)
		require.Equal(t, "300", maxConnections)
	})
}

// The pool raises max_connections because handing out many databases at once is
// its whole purpose, but a caller who names the setting must still win.
func TestTemplatePool_CallerArgsWin(t *testing.T) {
	conf := gormx.DefaultContainerConfig()
	conf.Args = []string{"-c", "max_connections=137"}

	pool := gormx.NewTemplatePool(conf, migrateWidgets)
	t.Cleanup(func() { require.NoError(t, pool.Close(context.Background())) })

	var maxConnections string
	require.NoError(t, pool.Fork(t).Raw("SHOW max_connections").Scan(&maxConnections).Error)
	require.Equal(t, "137", maxConnections)
}

func TestNewTemplatePool_RequiresMigrator(t *testing.T) {
	require.PanicsWithValue(t, "gormx: NewTemplatePool requires a TemplateMigrator", func() {
		gormx.NewTemplatePool(nil, nil)
	})
}
