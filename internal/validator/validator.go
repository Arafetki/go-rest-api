package validator

import (
	"slices"
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) HasErrors() bool {
	return len(v.Errors) != 0
}

func (v *Validator) AddError(key, msg string) {

	if _, exist := v.Errors[key]; !exist {
		v.Errors[key] = msg
	}
}

func (v *Validator) Check(ok bool, key string, msg string) {
	if !ok {
		v.AddError(key, msg)
	}
}

func PermittedValues[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}

func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}
