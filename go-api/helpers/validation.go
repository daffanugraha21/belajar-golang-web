package helpers

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationErrors(err error) []ValidationError{
	var errors []ValidationError

	for _, e := range err.(validator.ValidationErrors) {
		var msg string

		switch e.Tag() {
			case "required":
				msg = "field wajib diisi"
			case "min":
				msg = "minimal" + e.Param() + "karakter"
			default:
				msg = "input tidak valid"
		}
		errors = append(errors, ValidationError{
			Field: e.Field(),
			Message: msg,
		})
	}
	return errors
}