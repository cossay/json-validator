package jsonvalidator

import "github.com/tidwall/gjson"

//ObjectConstraint Object validation constraint
type ObjectConstraint struct {
	constraints map[string][]Constraint
}

//Validate Validates a given object
func (oc *ObjectConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if !value.Exists() || violations.Has(field) {
		return
	}

	for subField, subRules := range oc.constraints {
		curField := field + "." + subField
		curValue := value.Get(subField)
		for _, subRule := range subRules {
			subRule.Validate(curField, &curValue, value, source, violations)
		}
	}

}

//Object Creates a new object validation constraint
func Object(fields map[string][]Constraint) *ObjectConstraint {
	return &ObjectConstraint{constraints: fields}
}

//IsObjectConstraint Constraint for checking a given value is an object
type IsObjectConstraint struct {
	message string
}

//Validate Validates a given value is an object
func (ioc *IsObjectConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {

	if !value.Exists() {
		return
	}

	if !value.IsObject() {
		violations.Add(field, ioc.message)
	}
}

//IsObject Creates a new constraints for checking that a given value is an ojbect
func IsObject(message string) *IsObjectConstraint {
	return &IsObjectConstraint{message: message}
}
