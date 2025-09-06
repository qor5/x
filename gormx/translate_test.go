package gormx

import (
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestWithCauseVsWithoutCause(t *testing.T) {
	dsn := db.Config.Dialector.(*postgres.Dialector).Config.DSN

	// Three scenarios to compare
	withoutTranslateDB, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{TranslateError: false},
	)
	require.NoError(t, err)

	withTranslateDB, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{TranslateError: true},
	)
	require.NoError(t, err)

	withCauseDB, err := gorm.Open(
		WithCause(postgres.Open(dsn)),
		&gorm.Config{TranslateError: true},
	)
	require.NoError(t, err)

	type TestEntity struct {
		ID   string `gorm:"primarykey"`
		Name string `gorm:"unique"`
	}

	require.NoError(t, withTranslateDB.AutoMigrate(&TestEntity{}))
	t.Cleanup(func() {
		withTranslateDB.Migrator().DropTable(&TestEntity{})
	})

	entity1 := &TestEntity{ID: "1", Name: "unique_name"}
	require.NoError(t, withTranslateDB.Create(entity1).Error)

	// Test all three scenarios
	entityWithoutTranslate := &TestEntity{ID: "2", Name: "unique_name"}
	withoutTranslateErr := withoutTranslateDB.Create(entityWithoutTranslate).Error
	require.Error(t, withoutTranslateErr)

	entityWithTranslate := &TestEntity{ID: "3", Name: "unique_name"}
	withTranslateErr := withTranslateDB.Create(entityWithTranslate).Error
	require.Error(t, withTranslateErr)

	entityWithCause := &TestEntity{ID: "4", Name: "unique_name"}
	withCauseErr := withCauseDB.Create(entityWithCause).Error
	require.Error(t, withCauseErr)

	t.Logf("WithoutTranslate error: %s", withoutTranslateErr.Error())
	t.Logf("WithTranslate error: %s", withTranslateErr.Error())
	t.Logf("WithCause error: %s", withCauseErr.Error())

	// WithoutTranslate: should be raw PostgreSQL error
	var withoutTranslatePgErr *pgconn.PgError
	require.True(t, errors.As(withoutTranslateErr, &withoutTranslatePgErr))
	require.Equal(t, "23505", withoutTranslatePgErr.Code)
	require.False(t, errors.Is(withoutTranslateErr, gorm.ErrDuplicatedKey), "withoutTranslate should not be GORM error")

	// WithTranslate: should be translated GORM error only
	require.True(t, errors.Is(withTranslateErr, gorm.ErrDuplicatedKey))
	var withTranslatePgErr *pgconn.PgError
	require.False(t, errors.As(withTranslateErr, &withTranslatePgErr), "withTranslate error does not contain original PostgreSQL error")

	// WithCause: should contain both GORM error and original PostgreSQL error
	require.True(t, errors.Is(withCauseErr, gorm.ErrDuplicatedKey))
	var withCausePgErr *pgconn.PgError
	require.True(t, errors.As(withCauseErr, &withCausePgErr), "withCause error should contain original PostgreSQL error")
	require.Equal(t, "23505", withCausePgErr.Code)

	t.Logf("Comparison summary:")
	t.Logf("  WithoutTranslate: Raw PostgreSQL error only")
	t.Logf("  WithTranslate: GORM error only (loses details)")
	t.Logf("  WithCause: GORM error + PostgreSQL details (best of both)")
}
