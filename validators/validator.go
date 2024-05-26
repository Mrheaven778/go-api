package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUser(user interface{}) string {
	err := validate.Struct(user)
	if err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			switch tag {
			case "required":
				errorMessages = append(errorMessages, field+" es requerido")
			case "email":
				errorMessages = append(errorMessages, field+" debe ser un correo electrónico válido")
			case "password":
				errorMessages = append(errorMessages, field+" debe tener al menos 6 caracteres")
			default:
				errorMessages = append(errorMessages, "Error de validación en el campo "+field)
			}
		}
		return strings.Join(errorMessages, ", ")
	}
	return ""
}

func ValidateTask(task interface{}) string {
	err := validate.Struct(task)
	if err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			switch tag {
			case "required":
				errorMessages = append(errorMessages, field+" es requerido")
			default:
				errorMessages = append(errorMessages, "Error de validación en el campo "+field)
			}
		}
		return strings.Join(errorMessages, ", ")
	}
	return ""
}
