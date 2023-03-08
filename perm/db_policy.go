package perm

import (
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type DBPolicyBuilder struct {
	db            *gorm.DB
	model         DBPolicy
	loadFrequency time.Duration
}

func NewDBPolicy(db *gorm.DB) *DBPolicyBuilder {
	return &DBPolicyBuilder{
		db:            db,
		model:         DefaultDBPolicy{},
		loadFrequency: time.Minute,
	}
}

func (dpb *DBPolicyBuilder) Model(m DBPolicy) *DBPolicyBuilder {
	dpb.model = m
	return dpb
}

func (dpb *DBPolicyBuilder) LoadFrequency(d time.Duration) *DBPolicyBuilder {
	dpb.loadFrequency = d
	return dpb
}

type DefaultDBPolicy struct {
	gorm.Model

	ReferID   string
	Subject   string
	Effect    string
	Actions   pq.StringArray `gorm:"type:text[]"`
	Resources pq.StringArray `gorm:"type:text[]"`
}

func (p DefaultDBPolicy) LoadDBPolicies(db *gorm.DB, startFrom *time.Time) (toUpdateOrCreate []*PolicyBuilder, toDelete []*PolicyBuilder) {
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
			toUpdateOrCreate = append(toUpdateOrCreate, p.ToPolicy())
		}
	}
	return
}

func (p DefaultDBPolicy) ToPolicy() *PolicyBuilder {
	res := strings.Split(strings.Join(p.Resources, ","), ",")
	return PolicyFor(p.Subject).WhoAre(p.Effect).ToDo(p.Actions...).On(res...).ID(strconv.Itoa(int(p.ID)))
}
