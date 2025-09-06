package gormx

import (
	"errors"
	"fmt"

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
	et, ok := d.Dialector.(gorm.ErrorTranslator)
	if !ok {
		panic(fmt.Sprintf("dialector %T does not implement gorm.ErrorTranslator", d.Dialector))
	}
	translatedErr := et.Translate(cause)
	if translatedErr == cause {
		return translatedErr
	}
	return errors.Join(translatedErr, cause)
}
