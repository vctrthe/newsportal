package validatorLib

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	var errorMessages []string
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "email":
				errorMessages = append(errorMessages, "invalid email format")
			case "required":
				errorMessages = append(errorMessages, "field "+err.Field()+" wajib diisi")
			case "min":
				if err.Field() == "Password" {
					errorMessages = append(errorMessages, "password minimal 8 karakter")
				}
			case "eqfield":
				errorMessages = append(errorMessages, "field "+err.Field()+" dan "+err.Param()+" tidak sama")
			default:
				errorMessages = append(errorMessages, "field "+err.Field()+" tidak valid")
			}
		}

		return errors.New("validasi gagal: " + joinMessage(errorMessages))
	}

	return nil
}

func joinMessage(messages []string) string {
	result := ""
	for i, message := range messages {
		if i > 0 {
			result += ", "
		}
		result += message
	}

	return result
}
