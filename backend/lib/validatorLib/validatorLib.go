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
				errorMessages = append(errorMessages, "field "+err.Field()+" cannot be empty")
			case "min":
				if err.Field() == "Password" {
					errorMessages = append(errorMessages, "minimum password length is 6 characters")
				}
			case "eqfield":
				errorMessages = append(errorMessages, "field "+err.Field()+" and "+err.Param()+" not match")
			default:
				errorMessages = append(errorMessages, "field "+err.Field()+" invalid")
			}
		}

		return errors.New("validation failed: " + joinMessage(errorMessages))
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
