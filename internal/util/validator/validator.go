package validator

import (
	v10 "github.com/go-playground/validator/v10"
)

var validate *v10.Validate

func GetValidator() *v10.Validate {
	if validate == nil {
		validate = v10.New()
	}

	return validate
}

func Validate(s interface{}) error {
	vr := GetValidator()
	return vr.Struct(s)
}
