package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type Executor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

var (
	_ Executor = (*sql.DB)(nil)
	_ Executor = (*sql.Tx)(nil)
)

type TransactionOption func(*transactionOptions)

type transactionOptions struct {
	txOptions *sql.TxOptions
}

func WithTxOptions(txOptions *sql.TxOptions) TransactionOption {
	return func(c *transactionOptions) {
		c.txOptions = txOptions
	}
}

func Transaction(ctx context.Context, exec Executor, fn func(ctx context.Context, tx *sql.Tx) error, opts ...TransactionOption) (xerr error) {
	if tx, ok := exec.(*sql.Tx); ok {
		return transactionWithSavepoint(ctx, tx, fn)
	}

	db, ok := exec.(*sql.DB)
	if !ok {
		return errors.New("executor must be *sql.DB or *sql.Tx")
	}

	var cfg transactionOptions
	for _, opt := range opts {
		opt(&cfg)
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
