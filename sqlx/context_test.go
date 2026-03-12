package sqlx_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/qor5/x/v3/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromContext_NoExecutorInContext(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	// When no executor in context, should return the fallback
	exec := sqlx.FromContext(ctx, db)
	assert.Equal(t, db, exec, "FromContext should return fallback when no executor in context")
}

func TestFromContext_WithExecutorInContext(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	// Start a transaction and put it in context
	tx, err := db.BeginTx(ctx, nil)
	require.NoError(t, err, "Failed to begin transaction")
	defer func() { _ = tx.Rollback() }()

	ctxWithTx := sqlx.NewContext(ctx, tx)

	// FromContext should return the tx from context, not the fallback db
	exec := sqlx.FromContext(ctxWithTx, db)
	assert.Equal(t, tx, exec, "FromContext should return executor from context")
	assert.NotEqual(t, db, exec, "FromContext should not return fallback when executor exists in context")
}

func TestFromContext_WithDBInContext(t *testing.T) {
	db1 := setupTestDB(t)
	defer db1.Close()

	db2, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err, "Failed to open second database")
	defer db2.Close()

	ctx := context.Background()

	// Put db1 in context
	ctxWithDB := sqlx.NewContext(ctx, db1)

	// FromContext should return db1 from context, not the fallback db2
	exec := sqlx.FromContext(ctxWithDB, db2)
	assert.Equal(t, db1, exec, "FromContext should return executor from context")
	assert.NotEqual(t, db2, exec, "FromContext should not return fallback when executor exists in context")
}

func TestNewContext_OverwritesPreviousExecutor(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	// Start two transactions
	tx1, err := db.BeginTx(ctx, nil)
	require.NoError(t, err, "Failed to begin transaction 1")
	defer func() { _ = tx1.Rollback() }()

	tx2, err := db.BeginTx(ctx, nil)
	require.NoError(t, err, "Failed to begin transaction 2")
	defer func() { _ = tx2.Rollback() }()

	// Put tx1 in context, then overwrite with tx2
	ctxWithTx1 := sqlx.NewContext(ctx, tx1)
	ctxWithTx2 := sqlx.NewContext(ctxWithTx1, tx2)

	// FromContext should return tx2 (the most recent one)
	exec := sqlx.FromContext(ctxWithTx2, db)
	assert.Equal(t, tx2, exec, "FromContext should return the most recently set executor")
	assert.NotEqual(t, tx1, exec, "FromContext should not return the overwritten executor")
}

func TestFromContext_NilFallback(t *testing.T) {
	ctx := context.Background()

	// When no executor in context and fallback is nil, should return nil
	exec := sqlx.FromContext(ctx, nil)
	assert.Nil(t, exec, "FromContext should return nil when no executor in context and fallback is nil")
}

func TestNewContext_NilExecutor(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	// Put nil executor in context
	ctxWithNil := sqlx.NewContext(ctx, nil)

	// FromContext should return the fallback since nil is stored
	exec := sqlx.FromContext(ctxWithNil, db)
	assert.Equal(t, db, exec, "FromContext should return fallback when nil executor is stored in context")
}
