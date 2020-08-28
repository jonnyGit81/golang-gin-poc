package validators

import (
	"github.com/go-playground/validator"
	"strings"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}
