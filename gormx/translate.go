package gormx

import (
	"errors"

	"gorm.io/gorm"
)

type withCause struct {
	gorm.Dialector
}

func WithCause(d gorm.Dialector) gorm.Dialector {
	if _, ok := d.(gorm.ErrorTranslator); !ok {
		return d
	}
	return &withCause{
		Dialector: d,
	}
}

func (d *withCause) Translate(cause error) error {
	translatedErr := d.Dialector.(gorm.ErrorTranslator).Translate(cause)
	if translatedErr == cause {
		return cause
	}
	return errors.Join(translatedErr, cause)
}
