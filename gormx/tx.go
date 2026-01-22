package gormx

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Transaction(ctx context.Context, db *gorm.DB, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	err := db.WithContext(ctx).Transaction(fc, opts...)
	if err != nil {
		// Handle context cancellation race condition:
		// When ctx is cancelled, database/sql's internal awaitDone goroutine may rollback
		// the transaction before Commit() is called, resulting in ErrTxDone.
		// In this case, return the context error instead of ErrTxDone to reflect the true cause.
		// See: https://github.com/golang/go/issues/43507
		if errors.Is(err, sql.ErrTxDone) && ctx.Err() != nil {
			return errors.Wrap(ctx.Err(), "transaction failed")
		}
	}
	return err
}
