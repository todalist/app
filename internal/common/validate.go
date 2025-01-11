package common

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var (
	onceValidate sync.Once
	_validator   *validator.Validate
)

func lazyInit() {
	onceValidate.Do(func() {
		_validator = validator.New()
	})
}

func Valid(i interface{}) error {
	lazyInit()
	return _validator.Struct(i)
}

func ValidField(i interface{}, tag string) error {
	lazyInit()
	return _validator.Var(i, tag)
}

type StructValidator struct {
	V *validator.Validate
}

func (v *StructValidator) Validate(i interface{}) error {
	return v.V.Struct(i)
}
