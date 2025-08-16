package gormx

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var TimePrecision = time.Microsecond

type Model struct {
	ID        string         `gorm:"size:36;primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (m *Model) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}
