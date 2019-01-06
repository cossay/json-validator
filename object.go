package jsonvalidator

import "github.com/tidwall/gjson"

//Object Creates a new object validation constraint
func Object(fields map[string][]Constraint) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || violations.Has(field) {
			return
		}

		for subField, subRules := range fields {
			curField := field + "." + subField
			curValue := value.Get(subField)
			for _, subRule := range subRules {
				subRule.Validate(curField, &curValue, value, source, violations, validator)
			}
		}
	})
}

//TypeObject Creates a new constraints for checking that a given value is an ojbect
func TypeObject(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !value.IsObject() {
			violations.Add(field, message)
		}
	})
}
