package postgresx_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/qor5/x/v3/gormx"
	"github.com/qor5/x/v3/gormx/postgresx"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var suite *gormx.TestSuite

func TestMain(m *testing.M) {
	suite = gormx.MustStartTestSuite(context.Background())
	defer suite.Stop(context.Background())
	m.Run()
}

func TestSavePointerDialectorInterface(t *testing.T) {
	// Test with PostgreSQL (which supports SavePointerDialectorInterface)
	d := postgres.Open(suite.DSN())
	_, ok := d.(gorm.SavePointerDialectorInterface)
	require.True(t, ok)

	d = postgresx.Open(suite.DSN())
	_, ok = d.(gorm.SavePointerDialectorInterface)
	require.True(t, ok)
}

func TestWithCauseVsWithoutCause(t *testing.T) {
	// Three scenarios to compare
	withoutTranslateDB, err := gorm.Open(
		postgres.Open(suite.DSN()),
		&gorm.Config{TranslateError: false},
	)
	require.NoError(t, err)

	withTranslateDB, err := gorm.Open(
		postgres.Open(suite.DSN()),
		&gorm.Config{TranslateError: true},
	)
	require.NoError(t, err)

	withCauseDB, err := gorm.Open(
		postgresx.Open(suite.DSN()),
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

func TestWithCause_NoTranslationNeeded(t *testing.T) {
	withCauseDB, err := gorm.Open(
		postgresx.Open(suite.DSN()),
		&gorm.Config{TranslateError: true},
	)
	require.NoError(t, err)

	// Create a table with a numeric column to cause a data type error
	type NumericTestEntity struct {
		ID    string `gorm:"primarykey"`
		Count int    `gorm:"not null"`
	}

	require.NoError(t, withCauseDB.AutoMigrate(&NumericTestEntity{}))
	t.Cleanup(func() {
		withCauseDB.Migrator().DropTable(&NumericTestEntity{})
	})

	// Execute raw SQL that causes a data type error (not in PostgreSQL translator's error code list)
	// This will cause error code 22P02 (invalid_text_representation) which is not translated by PostgreSQL driver
	rawErr := withCauseDB.Exec("INSERT INTO numeric_test_entities (id, count) VALUES ('test', 'not_a_number')").Error
	require.Error(t, rawErr)

	t.Logf("Raw SQL error (no translation): %s", rawErr.Error())

	// Verify this is a PostgreSQL error that was NOT translated
	var pgErr *pgconn.PgError
	require.True(t, errors.As(rawErr, &pgErr), "should be a PostgreSQL error")
	require.Equal(t, "22P02", pgErr.Code, "should be invalid_text_representation error")

	// Since this error code (22P02) is not in errCodes map, PostgreSQL driver returns the original error
	// WithCause should also return the same error without wrapping it
	require.False(t, strings.Contains(rawErr.Error(), "\n"), "error should not contain newlines indicating joined errors")
	require.False(t, errors.Is(rawErr, gorm.ErrDuplicatedKey), "should not be translated to any GORM error")
}
