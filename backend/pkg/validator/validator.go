package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	// TODO: register custom validators here
	// validate.RegisterValidation("custom_tag", customValidatorFunc)
}

func Validate(data interface{}) error {
	return validate.Struct(data)
}

func ValidateVar(field interface{}, tag string) error {
	return validate.Var(field, tag)
}