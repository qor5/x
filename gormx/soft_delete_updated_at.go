package gormx

import (
	"database/sql"
	"reflect"
	"sync"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type softDeleteUpdatedAtPlugin struct{}

// SoftDeleteUpdatedAtPlugin updates UpdatedAt together with DeletedAt during
// soft delete by patching the cached schema's DeleteClauses on first access.
var SoftDeleteUpdatedAtPlugin gorm.Plugin = &softDeleteUpdatedAtPlugin{}

func (p *softDeleteUpdatedAtPlugin) Name() string {
	return "gormx:soft_delete_updated_at"
}

func (p *softDeleteUpdatedAtPlugin) Initialize(db *gorm.DB) error {
	patcher := &deleteClausesPatcher{}
	if err := db.Callback().Delete().Before("gorm:delete").Register("gormx:patch_delete_clauses", patcher.patch); err != nil {
		return errors.Wrap(err, "failed to register soft delete updated at callback")
	}
	return nil
}

type deleteClausesPatcher struct {
	mu      sync.Mutex
	patched sync.Map
}

func (p *deleteClausesPatcher) patch(tx *gorm.DB) {
	if tx.Statement == nil || tx.Statement.Schema == nil {
		return
	}

	s := tx.Statement.Schema
	if _, ok := p.patched.Load(s); ok {
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	// Double-check after acquiring the lock.
	if _, ok := p.patched.Load(s); ok {
		return
	}
	defer func() { p.patched.Store(s, true) }()

	var updatedAtFields []*schema.Field
	for _, f := range s.Fields {
		if f.AutoUpdateTime > 0 {
			updatedAtFields = append(updatedAtFields, f)
		}
	}
	if len(updatedAtFields) == 0 {
		return
	}

	clauses := make([]clause.Interface, 0, len(s.DeleteClauses))
	replaced := false
	for _, c := range s.DeleteClauses {
		if sd, ok := c.(gorm.SoftDeleteDeleteClause); ok {
			clauses = append(clauses, softDeleteWithUpdatedAtClause{
				ZeroValue:       sd.ZeroValue,
				Field:           sd.Field,
				UpdatedAtFields: updatedAtFields,
			})
			replaced = true
			continue
		}
		clauses = append(clauses, c)
	}

	if replaced {
		s.DeleteClauses = clauses
	}
}

// softDeleteWithUpdatedAtClause is a drop-in replacement for
// gorm.SoftDeleteDeleteClause that additionally sets UpdatedAt.
//
// The implementation mirrors gorm.SoftDeleteDeleteClause.ModifyStatement
// (gorm@v1.31.1/soft_delete.go) with the only addition being the
// updated_at column in the SET clause and SetColumn call.
type softDeleteWithUpdatedAtClause struct {
	ZeroValue       sql.NullString
	Field           *schema.Field
	UpdatedAtFields []*schema.Field
}

func (sd softDeleteWithUpdatedAtClause) Name() string               { return "" }
func (sd softDeleteWithUpdatedAtClause) Build(clause.Builder)       {}
func (sd softDeleteWithUpdatedAtClause) MergeClause(*clause.Clause) {}

func (sd softDeleteWithUpdatedAtClause) ModifyStatement(stmt *gorm.Statement) {
	if stmt.SQL.Len() == 0 && !stmt.Statement.Unscoped {
		curTime := stmt.DB.NowFunc()
		set := clause.Set{{Column: clause.Column{Name: sd.Field.DBName}, Value: curTime}}
		for _, f := range sd.UpdatedAtFields {
			set = append(set, clause.Assignment{Column: clause.Column{Name: f.DBName}, Value: curTime})
		}
		stmt.AddClause(set)
		stmt.SetColumn(sd.Field.DBName, curTime, true)
		for _, f := range sd.UpdatedAtFields {
			stmt.SetColumn(f.DBName, curTime, true)
		}

		if stmt.Schema != nil {
			_, queryValues := schema.GetIdentityFieldValuesMap(stmt.Context, stmt.ReflectValue, stmt.Schema.PrimaryFields)
			column, values := schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)

			if len(values) > 0 {
				stmt.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
			}

			if stmt.ReflectValue.CanAddr() && stmt.Dest != stmt.Model && stmt.Model != nil {
				_, queryValues = schema.GetIdentityFieldValuesMap(stmt.Context, reflect.ValueOf(stmt.Model), stmt.Schema.PrimaryFields)
				column, values = schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)

				if len(values) > 0 {
					stmt.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
				}
			}
		}

		gorm.SoftDeleteQueryClause{Field: sd.Field, ZeroValue: sd.ZeroValue}.ModifyStatement(stmt)
		stmt.AddClauseIfNotExists(clause.Update{})
		stmt.Build(stmt.DB.Callback().Update().Clauses...)
	}
}
