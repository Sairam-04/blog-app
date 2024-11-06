package utils

import (
	"fmt"
	"strings"

	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/go-playground/validator/v10"
)

func ValidationError(errs validator.ValidationErrors) types.GeneralResponse {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return types.GeneralResponse{
		Success: false,
		Message: "invalid payload data",
		Error:   strings.Join(errMsgs, ", "),
	}
}
