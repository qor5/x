package sqlx_test

import (
	"context"
	"database/sql"
	"strings"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/qor5/x/v3/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, balance INTEGER)`)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestTransaction_Basic(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "Alice", 100)
		return err
	})

	assert.NoError(t, err)

	var count int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestTransaction_Rollback(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "Bob", 200)
		if err != nil {
			return err
		}
		return errors.New("intentional error")
	})

	assert.Error(t, err)

	var count int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}

func TestTransaction_Nested(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "Charlie", 300)
		if err != nil {
			return err
		}

		err = sqlx.Transaction(ctx, tx, func(ctx context.Context, tx *sql.Tx) error {
			_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "David", 400)
			if err != nil {
				return err
			}
			return errors.New("nested transaction error")
		})

		if err == nil {
			t.Fatal("expected nested transaction to fail")
		}

		_, err = tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "Eve", 500)
		return err
	})

	assert.NoError(t, err)

	var count int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 2, count)

	var names []string
	rows, err := db.QueryContext(ctx, "SELECT name FROM users ORDER BY name")
	assert.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			t.Fatal(err)
		}
		names = append(names, name)
	}

	assert.Equal(t, []string{"Charlie", "Eve"}, names)
}

func TestTransaction_NestedSuccess(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "User1", 100)
		if err != nil {
			return err
		}

		err = sqlx.Transaction(ctx, tx, func(ctx context.Context, tx *sql.Tx) error {
			_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "User2", 200)
			return err
		})
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "User3", 300)
		return err
	})

	assert.NoError(t, err)

	var count int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 3, count)
}

func TestTransaction_ContextCancellation(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())

	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "CancelUser", 100)
		if err != nil {
			return err
		}

		// Cancel context and wait for awaitDone goroutine to rollback the transaction.
		// This simulates the race condition where context cancellation triggers
		// internal rollback before Commit() is called.
		// See: https://github.com/golang/go/issues/43507
		cancel()
		time.Sleep(50 * time.Millisecond)
		return nil
	})

	// Should return context.Canceled instead of sql.ErrTxDone
	assert.Error(t, err)
	assert.True(t, errors.Is(err, context.Canceled), "expected context.Canceled, got: %v", err)
	assert.True(t, strings.Contains(err.Error(), "failed to commit transaction"), "expected 'failed to commit transaction' in error, got: %v", err)
}

func TestTransaction_WithContextExecutor(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	// Test that Transaction uses executor from context when available
	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		// Put tx in context
		ctx = sqlx.NewContext(ctx, tx)

		// Insert using the executor from context
		exec := sqlx.FromContext(ctx, db)
		_, err := exec.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "ContextUser", 100)
		return err
	})

	require.NoError(t, err, "Transaction should not error")

	// Verify the insert was committed
	var count int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE name = ?", "ContextUser").Scan(&count)
	require.NoError(t, err, "Query should not error")
	assert.Equal(t, 1, count, "Should have 1 user inserted")
}

func TestTransaction_NestedWithContextPropagation(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	// Simulate a service function that uses FromContext
	insertUser := func(ctx context.Context, db *sql.DB, name string, balance int) error {
		exec := sqlx.FromContext(ctx, db)
		_, err := exec.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", name, balance)
		return err
	}

	// Outer transaction
	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		ctx = sqlx.NewContext(ctx, tx)

		// This should use the tx from context
		if err := insertUser(ctx, db, "User1", 100); err != nil {
			return err
		}

		// Nested transaction (savepoint)
		err := sqlx.Transaction(ctx, tx, func(ctx context.Context, tx *sql.Tx) error {
			ctx = sqlx.NewContext(ctx, tx)

			// This should use the inner tx from context
			if err := insertUser(ctx, db, "User2", 200); err != nil {
				return err
			}

			// Fail the nested transaction
			return errors.New("nested transaction failed")
		})

		// Nested transaction failed, but outer should continue
		if err == nil {
			t.Fatal("expected nested transaction to fail")
		}

		// This should still use the outer tx
		return insertUser(ctx, db, "User3", 300)
	})

	require.NoError(t, err, "Outer transaction should not error")

	// Verify: User1 and User3 should exist, User2 should not (rolled back with savepoint)
	var names []string
	rows, err := db.QueryContext(ctx, "SELECT name FROM users ORDER BY name")
	require.NoError(t, err, "Query should not error")
	defer rows.Close()

	for rows.Next() {
		var name string
		require.NoError(t, rows.Scan(&name), "Scan should not error")
		names = append(names, name)
	}

	assert.Equal(t, []string{"User1", "User3"}, names, "Only User1 and User3 should exist")
}

func TestTransaction_WithDisableSavepoint(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	// With savepoint disabled, nested transaction errors affect the entire transaction
	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "User1", 100)
		if err != nil {
			return err
		}

		// Nested transaction without savepoint
		err = sqlx.Transaction(ctx, tx, func(ctx context.Context, tx *sql.Tx) error {
			_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "User2", 200)
			return err
		}, sqlx.WithDisableSavepoint())
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "User3", 300)
		return err
	})

	require.NoError(t, err, "Transaction should not error")

	// All three users should exist since no error occurred
	var count int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	require.NoError(t, err, "Query should not error")
	assert.Equal(t, 3, count, "Should have 3 users inserted")
}

func TestTransaction_WithDisableSavepoint_NestedError(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	// With savepoint disabled, nested transaction error will propagate up
	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "User1", 100)
		if err != nil {
			return err
		}

		// Nested transaction without savepoint - error will propagate
		err = sqlx.Transaction(ctx, tx, func(ctx context.Context, tx *sql.Tx) error {
			_, err := tx.ExecContext(ctx, "INSERT INTO users (name, balance) VALUES (?, ?)", "User2", 200)
			if err != nil {
				return err
			}
			return errors.New("nested error")
		}, sqlx.WithDisableSavepoint())

		// This error propagates directly without savepoint rollback
		return err
	})

	assert.Error(t, err, "Transaction should error")

	// No users should exist since the entire transaction was rolled back
	var count int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	require.NoError(t, err, "Query should not error")
	assert.Equal(t, 0, count, "Should have 0 users since transaction was rolled back")
}
