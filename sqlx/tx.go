// Package sqlx provides transaction management utilities for database operations.
// It supports nested transactions using savepoints (optional) and context-based executor propagation.
package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

// Executor defines the common interface for database operations.
// Both *sql.DB and *sql.Tx implement this interface, allowing code to work
// with either a direct database connection or within a transaction context.
type Executor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

var (
	_ Executor = (*sql.DB)(nil)
	_ Executor = (*sql.Tx)(nil)
)

// TransactionOption configures transaction behavior.
type TransactionOption func(*transactionOptions)

type transactionOptions struct {
	txOptions        *sql.TxOptions
	disableSavepoint bool
}

// WithTxOptions sets the sql.TxOptions for the transaction.
// This allows configuring isolation level and read-only mode.
func WithTxOptions(txOptions *sql.TxOptions) TransactionOption {
	return func(c *transactionOptions) {
		c.txOptions = txOptions
	}
}

// WithDisableSavepoint disables savepoint usage for nested transactions.
// When disabled, if Transaction is called with an existing *sql.Tx, the function
// will be executed directly without creating a savepoint. This is useful for
// databases that do not support savepoints.
//
// Note: When savepoints are disabled, errors in nested transactions will affect
// the entire transaction, not just the nested portion.
func WithDisableSavepoint() TransactionOption {
	return func(c *transactionOptions) {
		c.disableSavepoint = true
	}
}

// Transaction executes fn within a database transaction.
//
// If exec is a *sql.DB, a new transaction is started. The transaction is
// automatically committed if fn returns nil, or rolled back if fn returns
// an error or panics.
//
// If exec is already a *sql.Tx (nested transaction), behavior depends on options:
//   - By default, a savepoint is created for partial rollback support.
//   - With WithDisableSavepoint(), fn is executed directly without savepoint.
//
// The function handles the context cancellation race condition where the
// database/sql's internal awaitDone goroutine may rollback the transaction
// before Commit() is called. In this case, the context error is returned
// instead of sql.ErrTxDone.
//
// Example:
//
//	err := sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
//	    _, err := tx.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", "Alice")
//	    if err != nil {
//	        return err
//	    }
//	    // Nested transaction using savepoint
//	    return sqlx.Transaction(ctx, tx, func(ctx context.Context, tx *sql.Tx) error {
//	        _, err := tx.ExecContext(ctx, "INSERT INTO logs (msg) VALUES (?)", "user created")
//	        return err
//	    })
//	})
func Transaction(ctx context.Context, exec Executor, fn func(ctx context.Context, tx *sql.Tx) error, opts ...TransactionOption) (xerr error) {
	var cfg transactionOptions
	for _, opt := range opts {
		opt(&cfg)
	}

	if tx, ok := exec.(*sql.Tx); ok {
		if cfg.disableSavepoint {
			return fn(ctx, tx)
		}
		return transactionWithSavepoint(ctx, tx, fn)
	}

	db, ok := exec.(*sql.DB)
	if !ok {
		return errors.New("executor must be *sql.DB or *sql.Tx")
	}

	tx, err := db.BeginTx(ctx, cfg.txOptions)
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	panicked := true
	defer func() {
		if panicked || xerr != nil {
			if err := tx.Rollback(); err != nil {
				if !errors.Is(err, sql.ErrTxDone) || ctx.Err() == nil {
					slog.ErrorContext(ctx, "failed to rollback transaction", "error", err)
				}
			}
		}
	}()

	err = fn(ctx, tx)
	panicked = false

	if err != nil {
		return errors.Wrap(err, "transaction failed")
	}

	// Handle context cancellation race condition:
	// When ctx is cancelled, database/sql's internal awaitDone goroutine may rollback
	// the transaction before Commit() is called, resulting in ErrTxDone.
	// In this case, return the context error instead of ErrTxDone to reflect the true cause.
	// See: https://github.com/golang/go/issues/43507
	if err := tx.Commit(); err != nil {
		if errors.Is(err, sql.ErrTxDone) && ctx.Err() != nil {
			return errors.Wrap(ctx.Err(), "failed to commit transaction")
		}
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

// transactionWithSavepoint executes fn within a savepoint for nested transactions.
// If fn returns an error or panics, only the savepoint is rolled back.
// This allows the outer transaction to continue or handle the error appropriately.
//
// Note: Savepoints are supported by PostgreSQL, MySQL, SQLite, and most modern databases.
// For databases that do not support savepoints, use WithDisableSavepoint() option.
func transactionWithSavepoint(ctx context.Context, tx *sql.Tx, fn func(ctx context.Context, tx *sql.Tx) error) (xerr error) {
	spID := xid.New().String()
	spName := fmt.Sprintf("sp_%s", spID)

	if _, err := tx.ExecContext(ctx, fmt.Sprintf("SAVEPOINT %s", spName)); err != nil {
		return errors.Wrap(err, "failed to create savepoint")
	}

	panicked := true
	defer func() {
		if panicked || xerr != nil {
			if _, err := tx.ExecContext(ctx, fmt.Sprintf("ROLLBACK TO SAVEPOINT %s", spName)); err != nil {
				slog.ErrorContext(ctx, "failed to rollback to savepoint", "savepoint", spName, "error", err)
			}
		}
	}()

	err := fn(ctx, tx)
	panicked = false

	if err != nil {
		return errors.Wrap(err, "transaction failed")
	}

	return nil
}
