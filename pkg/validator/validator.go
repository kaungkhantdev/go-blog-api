package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func ValidateStruct(input interface{}) error {
	err := validate.Struct(input)
	if err != nil {
		formattedErrors := formatValidationErrors(err)
		return fmt.Errorf("%v", formattedErrors)
	}
	return nil
}

func formatValidationErrors(err error) string {
	var messages []string
	for _, err := range err.(validator.ValidationErrors) {
		var message string
		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("The field '%s' is required.", err.Field())
		case "min":
			message = fmt.Sprintf("The field '%s' must be at least %s characters long.", err.Field(), err.Param())
		case "max":
			message = fmt.Sprintf("The field '%s' must be at most %s characters long.", err.Field(), err.Param())
		case "email":
			message = fmt.Sprintf("The field '%s' must be a valid email address.", err.Field())
		default:
			message = fmt.Sprintf("The field '%s' is invalid due to '%s' validation.", err.Field(), err.Tag())
		}
		messages = append(messages, message)
	}
	return strings.Join(messages, " ")
}
