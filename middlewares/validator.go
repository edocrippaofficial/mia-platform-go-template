package middlewares

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.Validator.Struct(i)
}

func validateErrorMapping(err error) map[string]string {
	if err == nil {
		return nil
	}

	validationErrors := make(map[string]string)
	if _, ok := err.(*validator.InvalidValidationError); ok {
		validationErrors["error"] = "Invalid validation error"
		return validationErrors
	}

	for _, fieldErr := range err.(validator.ValidationErrors) {
		validationErrors[fieldErr.Field()] = fieldErr.Tag()
	}
	return validationErrors
}
