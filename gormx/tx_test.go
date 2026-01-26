package gormx_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/qor5/x/v3/gormx"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type TxTestUser struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	Balance int
}

func TestTransaction_ContextCancellation(t *testing.T) {
	ctx := context.Background()
	db := suite.DB()

	err := suite.ResetDB(ctx, &TxTestUser{})
	assert.NoError(t, err)

	cancelCtx, cancel := context.WithCancel(ctx)

	err = gormx.Transaction(db.WithContext(cancelCtx), func(tx *gorm.DB) error {
		err := tx.Create(&TxTestUser{ID: "1", Name: "CancelUser", Balance: 100}).Error
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
	assert.True(t, strings.Contains(err.Error(), "transaction failed"), "expected 'transaction failed' in error, got: %v", err)
}
