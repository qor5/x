package sqlx

import (
	"context"
)

// ctxKeyExecutor is the context key for storing an Executor.
type ctxKeyExecutor struct{}

// FromContext retrieves an Executor from the context.
// If no Executor is found in the context, it returns the provided fallback.
// This is useful for propagating transactions through the call stack.
//
// Example:
//
//	func doSomething(ctx context.Context, db *sql.DB) error {
//	    exec := sqlx.FromContext(ctx, db)
//	    _, err := exec.ExecContext(ctx, "INSERT INTO ...")
//	    return err
//	}
func FromContext(ctx context.Context, fallback Executor) Executor {
	if executor, ok := ctx.Value(ctxKeyExecutor{}).(Executor); ok {
		return executor
	}
	return fallback
}

// NewContext returns a new context with the given Executor attached.
// Use this to propagate a transaction through the call stack, allowing
// nested functions to participate in the same transaction.
//
// Example:
//
//	sqlx.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
//	    ctx = sqlx.NewContext(ctx, tx)
//	    return doSomething(ctx, db) // doSomething will use tx via FromContext
//	})
func NewContext(ctx context.Context, executor Executor) context.Context {
	return context.WithValue(ctx, ctxKeyExecutor{}, executor)
}
