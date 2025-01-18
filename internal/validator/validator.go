package validator

import (
	"fmt"
	"strings"
)

type ValidationErrors struct {
	Errors map[string]string
}

func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		Errors: make(map[string]string),
	}
}

func (v *ValidationErrors) Add(field, message string) {
	v.Errors[field] = message
}

func (v *ValidationErrors) Error() string {
	if len(v.Errors) == 0 {
		return "Validation passed."
	}

	var builder strings.Builder

	for field, message := range v.Errors {
		builder.WriteString(fmt.Sprintf("%s: %s; ", field, message))
	}

	return strings.TrimSpace(builder.String())
}
