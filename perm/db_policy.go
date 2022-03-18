package perm

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type DefaultDBPolicy struct {
	gorm.Model

	ReferID   string
	Subject   string
	Effect    string
	Actions   pq.StringArray `gorm:"type:text[]"`
	Resources pq.StringArray `gorm:"type:text[]"`
}

func (p DefaultDBPolicy) LoadDBPolicies(db *gorm.DB, startFrom *time.Time) (toUpdate []*PolicyBuilder, toDelete []*PolicyBuilder) {
	var ps []DefaultDBPolicy
	if startFrom == nil || startFrom.IsZero() {
		db.Find(&ps)
	} else {
		db.Unscoped().Where("updated_at >= ? or deleted_at >= ?", startFrom, startFrom).Find(&ps)
	}

	for _, p := range ps {
		if p.DeletedAt.Valid {
			toDelete = append(toDelete, p.ToPolicy())
		} else {
			toUpdate = append(toUpdate, p.ToPolicy())
		}
	}
	return
}

func (p DefaultDBPolicy) ToPolicy() *PolicyBuilder {
	res := strings.Split(strings.Join(p.Resources, ","), ",")
	return PolicyFor(p.Subject).WhoAre(p.Effect).ToDo(p.Actions...).On(res...).ID(strconv.Itoa(int(p.ID)))
}
