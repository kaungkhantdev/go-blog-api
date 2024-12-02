package validator

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func ValidateStruct(input interface{}) error {
	return validate.Struct(input)
}
