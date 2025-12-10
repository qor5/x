package gormx

import (
	"cmp"
	"context"
	"database/sql"
	"database/sql/driver"
	"io"

	"github.com/pkg/errors"
	"github.com/theplant/appkit/logtracing"
	"gorm.io/gorm"

	kitlog "github.com/theplant/appkit/log"
)

const (
	keyTracingSpanName       = "keyTracingSpanName"
	keyTracingMaxQueryLength = "keyTracingMaxQueryLength"
	keyTracingDisabled       = "keyTracingDisabled"
)

const defaultMaxQueryLength = 4096

func WithSpanName(spanName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Set(keyTracingSpanName, spanName)
	}
}

func WithMaxQueryLength(maxLength int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Set(keyTracingMaxQueryLength, maxLength)
	}
}

// WithoutTracing returns a scope function that disables tracing for the current query.
// This is useful when you want to temporarily disable tracing for specific queries.
func WithoutTracing() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Set(keyTracingDisabled, true)
	}
}

type TracingConfig struct {
	ExcludeQuery     bool                      `confx:"excludeQuery" usage:"Exclude query"`
	ExcludeQueryVars bool                      `confx:"excludeQueryVars" usage:"Exclude query vars"`
	MaxQueryLength   int                       `confx:"maxQueryLength" usage:"Maximum query length for tracing, 0 uses default (4096)"`
	QueryFormatter   func(query string) string `confx:"-" json:"-"`
	Logger           *kitlog.Logger            `confx:"-" json:"-" inject:"optional"`
}

func NewTracingPlugin(conf *TracingConfig) gorm.Plugin {
	return &logtracingPlugin{
		TracingConfig: conf,
	}
}

type logtracingPlugin struct {
	*TracingConfig
}

func (p *logtracingPlugin) Name() string {
	return "logtracing"
}

type gormHookFunc func(tx *gorm.DB)

type gormRegister interface {
	Register(name string, fn func(*gorm.DB)) error
}

func (p *logtracingPlugin) Initialize(db *gorm.DB) (err error) {
	cb := db.Callback()
	hooks := []struct {
		callback gormRegister
		hook     gormHookFunc
		name     string
	}{
		{
			name:     "before:create",
			callback: cb.Create().Before("gorm:create"),
			hook:     p.before("gorm.Create"),
		},
		{
			name:     "after:create",
			callback: cb.Create().After("gorm:create"),
			hook:     p.after(),
		},

		{
			name:     "before:select",
			callback: cb.Query().Before("gorm:query"),
			hook:     p.before("gorm.Query"),
		},
		{
			name:     "after:select",
			callback: cb.Query().After("gorm:query"),
			hook:     p.after(),
		},

		{
			name:     "before:delete",
			callback: cb.Delete().Before("gorm:delete"),
			hook:     p.before("gorm.Delete"),
		},
		{
			name:     "after:delete",
			callback: cb.Delete().After("gorm:delete"),
			hook:     p.after(),
		},

		{
			name:     "before:update",
			callback: cb.Update().Before("gorm:update"),
			hook:     p.before("gorm.Update"),
		},
		{
			name:     "after:update",
			callback: cb.Update().After("gorm:update"),
			hook:     p.after(),
		},

		{
			name:     "before:row",
			callback: cb.Row().Before("gorm:row"),
			hook:     p.before("gorm.Row"),
		},
		{
			name:     "after:row",
			callback: cb.Row().After("gorm:row"),
			hook:     p.after(),
		},

		{
			callback: cb.Raw().Before("gorm:raw"),
			hook:     p.before("gorm.Raw"),
			name:     "before:raw",
		},
		{
			callback: cb.Raw().After("gorm:raw"),
			hook:     p.after(),
			name:     "after:raw",
		},
	}

	var firstErr error
	for _, h := range hooks {
		if err := h.callback.Register("logtracing:"+h.name, h.hook); err != nil && firstErr == nil {
			firstErr = errors.Wrapf(err, "callback register %s failed", h.name)
		}
	}

	return firstErr
}

type contextWrapper struct {
	context.Context
	parent context.Context
	table  string
}

func parseTable(tx *gorm.DB) string {
	var table string
	stmt := tx.Statement
	if stmt.Table != "" {
		table = stmt.Table
	} else {
		model := cmp.Or(stmt.Model, stmt.Dest)
		if model != nil {
			s, err := ParseSchema(tx, model)
			if err == nil {
				table = s.Table
			}
		}
	}
	return table
}

func isTracingDisabled(tx *gorm.DB) bool {
	if val, ok := tx.Get(keyTracingDisabled); ok {
		if disabled, ok := val.(bool); ok && disabled {
			return true
		}
	}
	return false
}

func (p *logtracingPlugin) before(spanPrefix string) gormHookFunc {
	return func(tx *gorm.DB) {
		if isTracingDisabled(tx) {
			return
		}

		var spanName string
		val, ok := tx.Get(keyTracingSpanName)
		if ok {
			sn, ok := val.(string)
			if ok && sn != "" {
				spanName = "gorm:" + sn
			}
		}

		table := parseTable(tx)

		if spanName == "" {
			spanName = spanPrefix + ":" + table
		}

		parentCtx := tx.Statement.Context
		ctx, _ := logtracing.StartSpan(tx.Statement.Context, spanName)
		tx.Statement.Context = &contextWrapper{ctx, parentCtx, table}
	}
}

func (p *logtracingPlugin) after() gormHookFunc {
	return func(tx *gorm.DB) {
		if isTracingDisabled(tx) {
			return
		}

		cw, _ := tx.Statement.Context.(*contextWrapper)
		defer func() {
			if cw != nil {
				// recover previous context
				tx.Statement.Context = cw.parent
			}
		}()

		span := logtracing.SpanFromContext(tx.Statement.Context)
		if span == nil || !span.IsRecording() {
			return
		}

		var xerr error
		defer func() {
			ctx := tx.Statement.Context
			if p.Logger != nil {
				if _, ok := kitlog.FromContext(ctx); !ok {
					ctx = kitlog.Context(ctx, *p.Logger)
				}
			}
			logtracing.EndSpan(ctx, xerr)
		}()

		span.AppendKVs(
			"span.type", "sql",
			"span.role", "client",
		)

		if cw != nil && cw.table != "" {
			span.AppendKVs("sql.table", cw.table)
		}

		if tx.Statement.RowsAffected != -1 {
			span.AppendKVs("sql.rows_affected", tx.Statement.RowsAffected)
		}

		if !p.ExcludeQuery {
			vars := tx.Statement.Vars
			var query string
			if p.ExcludeQueryVars {
				query = tx.Statement.SQL.String()
			} else {
				query = tx.Dialector.Explain(tx.Statement.SQL.String(), vars...)
			}
			span.AppendKVs("sql.query", p.formatQuery(tx, query))
		}

		if tx.Error != nil &&
			!(errors.Is(tx.Error, gorm.ErrRecordNotFound)) &&
			!(errors.Is(tx.Error, driver.ErrSkip)) &&
			!(errors.Is(tx.Error, io.EOF)) &&
			!(errors.Is(tx.Error, sql.ErrNoRows)) {
			xerr = errors.WithStack(tx.Error)
		}
	}
}

func (p *logtracingPlugin) formatQuery(tx *gorm.DB, query string) string {
	if p.QueryFormatter != nil {
		query = p.QueryFormatter(query)
	}

	// Priority: request-level > config-level > default
	maxLen := p.MaxQueryLength
	if val, ok := tx.Get(keyTracingMaxQueryLength); ok {
		if reqMaxLen, ok := val.(int); ok && reqMaxLen > 0 {
			maxLen = reqMaxLen
		}
	}

	if maxLen <= 0 {
		maxLen = defaultMaxQueryLength
	}

	if len(query) > maxLen {
		return query[:maxLen] + "... (truncated)"
	}

	return query
}
