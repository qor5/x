package gormx_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/qor5/x/v3/gormx"
	"github.com/qor5/x/v3/gormx/postgresx"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type SoftDeleteUpdatedAtModel struct {
	gormx.Model
	Name string
}

type SoftDeleteWithoutUpdatedAtModel struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

type SoftDeleteCustomFieldNameModel struct {
	ID         string         `gorm:"primaryKey"`
	CreatedAt  time.Time      `gorm:"not null"`
	ModifiedAt time.Time      `gorm:"not null;autoUpdateTime"`
	RemovedAt  gorm.DeletedAt `gorm:"index"`
	Name       string
}

type SoftDeleteMultiAutoUpdateModel struct {
	ID         string         `gorm:"primaryKey"`
	CreatedAt  time.Time      `gorm:"not null"`
	UpdatedAt  time.Time      `gorm:"not null"`
	ModifiedAt time.Time      `gorm:"not null;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Name       string
}

func TestSoftDeleteUpdatedAtPlugin(t *testing.T) {
	ctx := context.Background()

	db, err := gorm.Open(postgresx.Open(suite.DSN()))
	require.NoError(t, err)
	require.NoError(t, db.Use(gormx.SoftDeleteUpdatedAtPlugin))

	t.Cleanup(func() {
		require.NoError(t, db.Migrator().DropTable(
			&SoftDeleteUpdatedAtModel{},
			&SoftDeleteWithoutUpdatedAtModel{},
			&SoftDeleteCustomFieldNameModel{},
			&SoftDeleteMultiAutoUpdateModel{},
		))
	})

	t.Run("soft delete SQL includes updated_at", func(t *testing.T) {
		require.NoError(t, suite.ResetDB(ctx, &SoftDeleteUpdatedAtModel{}))

		record := SoftDeleteUpdatedAtModel{Name: "soft-delete-sql"}
		require.NoError(t, db.WithContext(ctx).Create(&record).Error)

		dryRunDB := db.Session(&gorm.Session{DryRun: true}).WithContext(ctx)
		stmt := dryRunDB.Delete(&record).Statement
		sql := stmt.SQL.String()

		require.True(t, strings.Contains(sql, "\"updated_at\""), sql)
		require.True(t, strings.Contains(sql, "\"deleted_at\""), sql)
	})

	t.Run("soft delete updates UpdatedAt", func(t *testing.T) {
		require.NoError(t, suite.ResetDB(ctx, &SoftDeleteUpdatedAtModel{}))

		record := SoftDeleteUpdatedAtModel{Name: "soft-delete"}
		require.NoError(t, db.WithContext(ctx).Create(&record).Error)

		require.NoError(t, db.WithContext(ctx).Delete(&record).Error)

		var deletedRecord SoftDeleteUpdatedAtModel
		require.NoError(t, db.WithContext(ctx).Unscoped().Where("id = ?", record.ID).First(&deletedRecord).Error)
		require.False(t, deletedRecord.DeletedAt.Time.IsZero())
		require.True(t, deletedRecord.UpdatedAt.Equal(deletedRecord.DeletedAt.Time))
	})

	t.Run("hard delete does not try to update UpdatedAt", func(t *testing.T) {
		require.NoError(t, suite.ResetDB(ctx, &SoftDeleteUpdatedAtModel{}))

		record := SoftDeleteUpdatedAtModel{Name: "hard-delete"}
		require.NoError(t, db.WithContext(ctx).Create(&record).Error)

		require.NoError(t, db.WithContext(ctx).Unscoped().Delete(&record).Error)

		var count int64
		require.NoError(t, db.WithContext(ctx).Model(&SoftDeleteUpdatedAtModel{}).Unscoped().Where("id = ?", record.ID).Count(&count).Error)
		require.Equal(t, int64(0), count)
	})

	t.Run("model without UpdatedAt still deletes successfully", func(t *testing.T) {
		require.NoError(t, suite.ResetDB(ctx, &SoftDeleteWithoutUpdatedAtModel{}))

		record := SoftDeleteWithoutUpdatedAtModel{ID: uuid.NewString(), Name: "without-updated-at"}
		require.NoError(t, db.WithContext(ctx).Create(&record).Error)

		require.NoError(t, db.WithContext(ctx).Delete(&record).Error)

		var deletedRecord SoftDeleteWithoutUpdatedAtModel
		require.NoError(t, db.WithContext(ctx).Unscoped().Where("id = ?", record.ID).First(&deletedRecord).Error)
		require.False(t, deletedRecord.DeletedAt.Time.IsZero())
	})

	t.Run("custom field names with autoUpdateTime tag", func(t *testing.T) {
		require.NoError(t, suite.ResetDB(ctx, &SoftDeleteCustomFieldNameModel{}))

		record := SoftDeleteCustomFieldNameModel{ID: uuid.NewString(), Name: "custom-fields"}
		require.NoError(t, db.WithContext(ctx).Create(&record).Error)

		require.NoError(t, db.WithContext(ctx).Delete(&record).Error)

		var deletedRecord SoftDeleteCustomFieldNameModel
		require.NoError(t, db.WithContext(ctx).Unscoped().Where("id = ?", record.ID).First(&deletedRecord).Error)
		require.False(t, deletedRecord.RemovedAt.Time.IsZero())
		require.True(t, deletedRecord.ModifiedAt.Equal(deletedRecord.RemovedAt.Time))
	})

	t.Run("custom field names dry run SQL", func(t *testing.T) {
		require.NoError(t, suite.ResetDB(ctx, &SoftDeleteCustomFieldNameModel{}))

		record := SoftDeleteCustomFieldNameModel{ID: uuid.NewString(), Name: "custom-fields-sql"}
		require.NoError(t, db.WithContext(ctx).Create(&record).Error)

		dryRunDB := db.Session(&gorm.Session{DryRun: true}).WithContext(ctx)
		stmt := dryRunDB.Delete(&record).Statement
		sql := stmt.SQL.String()

		require.True(t, strings.Contains(sql, "\"modified_at\""), sql)
		require.True(t, strings.Contains(sql, "\"removed_at\""), sql)
	})

	t.Run("multiple AutoUpdateTime fields", func(t *testing.T) {
		require.NoError(t, suite.ResetDB(ctx, &SoftDeleteMultiAutoUpdateModel{}))

		record := SoftDeleteMultiAutoUpdateModel{ID: uuid.NewString(), Name: "multi-auto-update"}
		require.NoError(t, db.WithContext(ctx).Create(&record).Error)

		require.NoError(t, db.WithContext(ctx).Delete(&record).Error)

		var deletedRecord SoftDeleteMultiAutoUpdateModel
		require.NoError(t, db.WithContext(ctx).Unscoped().Where("id = ?", record.ID).First(&deletedRecord).Error)
		require.False(t, deletedRecord.DeletedAt.Time.IsZero())
		require.True(t, deletedRecord.UpdatedAt.Equal(deletedRecord.DeletedAt.Time))
		require.True(t, deletedRecord.ModifiedAt.Equal(deletedRecord.DeletedAt.Time))
	})

	t.Run("multiple AutoUpdateTime fields dry run SQL", func(t *testing.T) {
		require.NoError(t, suite.ResetDB(ctx, &SoftDeleteMultiAutoUpdateModel{}))

		record := SoftDeleteMultiAutoUpdateModel{ID: uuid.NewString(), Name: "multi-sql"}
		require.NoError(t, db.WithContext(ctx).Create(&record).Error)

		dryRunDB := db.Session(&gorm.Session{DryRun: true}).WithContext(ctx)
		stmt := dryRunDB.Delete(&record).Statement
		sql := stmt.SQL.String()

		require.True(t, strings.Contains(sql, "\"updated_at\""), sql)
		require.True(t, strings.Contains(sql, "\"modified_at\""), sql)
		require.True(t, strings.Contains(sql, "\"deleted_at\""), sql)
	})

	t.Run("concurrent soft deletes", func(t *testing.T) {
		require.NoError(t, suite.ResetDB(ctx, &SoftDeleteUpdatedAtModel{}))

		const n = 20
		records := make([]SoftDeleteUpdatedAtModel, n)
		for i := range records {
			records[i] = SoftDeleteUpdatedAtModel{Name: "concurrent"}
			require.NoError(t, db.WithContext(ctx).Create(&records[i]).Error)
		}

		// Use a fresh DB+plugin so the schema hasn't been patched yet,
		// maximizing the chance of concurrent patch attempts.
		freshDB, err := gorm.Open(postgresx.Open(suite.DSN()))
		require.NoError(t, err)
		require.NoError(t, freshDB.Use(gormx.SoftDeleteUpdatedAtPlugin))

		errs := make(chan error, n)
		for i := range records {
			go func(r SoftDeleteUpdatedAtModel) {
				errs <- freshDB.WithContext(ctx).Delete(&r).Error
			}(records[i])
		}
		for range records {
			require.NoError(t, <-errs)
		}

		for _, r := range records {
			var deletedRecord SoftDeleteUpdatedAtModel
			require.NoError(t, db.WithContext(ctx).Unscoped().Where("id = ?", r.ID).First(&deletedRecord).Error)
			require.False(t, deletedRecord.DeletedAt.Time.IsZero())
			require.True(t, deletedRecord.UpdatedAt.Equal(deletedRecord.DeletedAt.Time))
		}
	})
}
