package jsonvalidator

import (
	"strings"

	"github.com/tidwall/gjson"
)

//RequiredConstraint Required validation constraint
type RequiredConstraint struct {
	message string
}

//Validate Validates a given value
func (rc *RequiredConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if !value.Exists() || value.Type == gjson.Null || strings.TrimSpace(value.String()) == "" {
		violations.Add(field, rc.message)
	}
}

//Required Required validation constraint
func Required(message string) *RequiredConstraint {
	return &RequiredConstraint{message: message}
}
