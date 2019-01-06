package jsonvalidator

import (
	"strings"

	"github.com/tidwall/gjson"
)

//IsEmpty Checks if a field is considered empty
func IsEmpty(value *gjson.Result) bool {
	if !value.Exists() || value.Type == gjson.Null || strings.TrimSpace(value.String()) == "" {
		return true
	}
	return false
}

//Required Required validation constraint
func Required(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
		if IsEmpty(value) {
			violations.Add(field, message)
		}
	})
}
