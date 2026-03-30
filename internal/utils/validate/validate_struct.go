package validate

import "github.com/go-playground/validator/v10"

func ValidateStruct(request any) map[string]string {
	validateData := validator.New()
	err := validateData.Struct(request)
	if err != nil {
		errors := make(map[string]string)

		for _, e := range err.(validator.ValidationErrors) {
			// Meýdan we ýalňyşlygy goşmak
			errors[e.Field()] = e.Field() + " is " + e.Tag()
		}
		return errors
	}
	return nil
}
