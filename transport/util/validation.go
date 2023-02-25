package util

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func validateRequest(r *http.Request, b any) error {
	validate = validator.New()

	return nil
}
