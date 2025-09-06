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
	if et, ok := d.Dialector.(gorm.ErrorTranslator); ok {
		translatedErr := et.Translate(cause)
		if translatedErr == cause {
			return translatedErr
		}
		return errors.Join(translatedErr, cause)
	}
	return cause
}
