package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidarFormato(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^[A-Za-z]{3}[0-9]{3}$`)
	return regex.MatchString(fl.Field().String())
}

func ValidarFormatoMaritima(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^[A-Za-z]{3}[0-9]{4}[A-Za-z]$`)
	return regex.MatchString(fl.Field().String())
}
