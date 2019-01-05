package jsonvalidator

import "github.com/tidwall/gjson"

func isString(value *gjson.Result) bool {

	_, ok := value.Value().(string)

	return ok
}

//IsStringConstraint Constraint for validating that a given value is a string
type IsStringConstraint struct {
	message string
}

//Validate Validates that a given value is a string
func (sc *IsStringConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if isEmpty(value) || isString(value) {
		return
	}

	violations.Add(field, sc.message)
}

//String Returns a new instance of String validation constraint
func String(message string) *IsStringConstraint {
	return &IsStringConstraint{message: message}
}
