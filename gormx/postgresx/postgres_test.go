package postgresx_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/qor5/x/v3/gormx"
	"github.com/qor5/x/v3/gormx/postgresx"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var suite *gormx.TestSuite

func TestMain(m *testing.M) {
	ctx := context.Background()
	suite = gormx.MustStartTestSuite(ctx)
	defer func() { _ = suite.Stop(ctx) }()
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
		_ = withTranslateDB.Migrator().DropTable(&TestEntity{})
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
		_ = withCauseDB.Migrator().DropTable(&NumericTestEntity{})
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

func TestConfigureTimezone(t *testing.T) {
	// Test that ConfigureTimezone correctly configures timestamp scanning with timezone
	// PostgreSQL stores timestamp without timezone as-is, and we want to scan it with a specific location

	t.Run("with timezone configured", func(t *testing.T) {
		// Parse DSN with timezone parameter
		dsnWithTZ := suite.DSN() + "&timezone=Asia/Shanghai"
		conf, err := pgx.ParseConfig(dsnWithTZ)
		require.NoError(t, err)

		// Apply timezone configuration
		tzOpts := postgresx.ConfigureTimezone(conf)
		require.NotNil(t, tzOpts, "should return options when timezone is set")

		// Open connection with timezone options
		conn := stdlib.OpenDB(*conf, tzOpts...)
		db, err := gorm.Open(postgresx.New(postgres.Config{Conn: conn}), &gorm.Config{})
		require.NoError(t, err)

		// Query current timestamp from PostgreSQL
		var result time.Time
		err = db.Raw("SELECT '2024-01-15 10:30:00'::timestamp").Scan(&result).Error
		require.NoError(t, err)

		// The scanned time should be in Asia/Shanghai timezone
		shanghai, _ := time.LoadLocation("Asia/Shanghai")
		require.Equal(t, shanghai.String(), result.Location().String())
		t.Logf("Scanned time with Asia/Shanghai: %v", result)
	})

	t.Run("without timezone configured", func(t *testing.T) {
		// Parse DSN without timezone parameter
		conf, err := pgx.ParseConfig(suite.DSN())
		require.NoError(t, err)

		// ConfigureTimezone should return nil when no timezone is set
		tzOpts := postgresx.ConfigureTimezone(conf)
		require.Nil(t, tzOpts, "should return nil when timezone is not set")
	})

	t.Run("compare with and without timezone", func(t *testing.T) {
		dsnWithTZ := suite.DSN() + "&timezone=Asia/Tokyo"

		// Without timezone - tz will be ignored
		confWithoutTZ, err := pgx.ParseConfig(dsnWithTZ)
		require.NoError(t, err)
		connWithoutTZ := stdlib.OpenDB(*confWithoutTZ)
		dbWithoutTZ, err := gorm.Open(postgresx.New(postgres.Config{Conn: connWithoutTZ}), &gorm.Config{})
		require.NoError(t, err)

		// With timezone - uses Asia/Tokyo
		confWithTZ, err := pgx.ParseConfig(dsnWithTZ)
		require.NoError(t, err)
		tzOpts := postgresx.ConfigureTimezone(confWithTZ)
		connWithTZ := stdlib.OpenDB(*confWithTZ, tzOpts...)
		dbWithTZ, err := gorm.Open(postgresx.New(postgres.Config{Conn: connWithTZ}), &gorm.Config{})
		require.NoError(t, err)

		// Query the same timestamp
		var resultWithoutTZ, resultWithTZ time.Time
		err = dbWithoutTZ.Raw("SELECT '2024-01-15 10:30:00'::timestamp").Scan(&resultWithoutTZ).Error
		require.NoError(t, err)
		err = dbWithTZ.Raw("SELECT '2024-01-15 10:30:00'::timestamp").Scan(&resultWithTZ).Error
		require.NoError(t, err)

		t.Logf("Without timezone: %v (location: %s)", resultWithoutTZ, resultWithoutTZ.Location())
		t.Logf("With Asia/Tokyo: %v (location: %s)", resultWithTZ, resultWithTZ.Location())

		// Verify different locations
		tokyo, _ := time.LoadLocation("Asia/Tokyo")
		require.Equal(t, tokyo.String(), resultWithTZ.Location().String())
		require.NotEqual(t, resultWithoutTZ.Location().String(), resultWithTZ.Location().String())
	})
}
