package helper

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/kekasicoid/kekasigohelper"
)

func ValidatorInit() *validator.Validate {
	validate := validator.New()
	validate.RegisterAlias("req-email", "required,email")
	validate.RegisterAlias("req-alphanum", "required,alphanum")
	validate.RegisterAlias("req-alpha", "required,alpha")
	validate.RegisterAlias("req-numeric", "required,numeric")
	validate.RegisterAlias("req-dive", "required,dive")
	validate.RegisterValidation("alphanum-space", AlphaNumSpace)
	validate.RegisterValidation("null-numeric", NullOrNumeric)
	return validate
}

func AlphaNumSpace(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if !regexp.MustCompile("^[a-zA-Z0-9 ]*$").MatchString(value) {
			return false
		}
	}
	return true
}

func NullOrNumeric(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value == "" {
			return true
		}
		if !regexp.MustCompile("^[0-9]*$").MatchString(value) {
			kekasigohelper.LoggerWarning("regex")
			return false
		}
	}
	return true
}
