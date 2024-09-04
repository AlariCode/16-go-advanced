package req

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
