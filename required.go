package jsonvalidator

import (
	"fmt"
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
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			violations.Add(field, message)
		}
	})
}

//NotExpected Unpected field validation constraint
func NotExpected(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if value.Exists() {
			violations.Add(field, message)
		}
	})
}

//NotExpectedDefault Same as NotExpected but with hard coded message
func NotExpectedDefault() *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if value.Exists() {
			violations.Add(field, fmt.Sprintf("Field '%s' was not expected", field))
		}
	})
}
