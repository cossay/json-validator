package jsonvalidator

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/tidwall/gjson"
)

//TypeArray Validates that a given value is an array
func TypeArray(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !value.IsArray() {
			violations.Add(field, message)
		}
	})
}

//Array Returns an instance of Array validation constraints
func Array(constraints []Constraint) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if !value.IsArray() {
			return
		}

		for index, item := range value.Array() {
			for _, rule := range constraints {
				entryField := fmt.Sprintf("%s.%d", field, index)
				rule.Validate(entryField, &item, value, source, violations, validator)
			}
		}
	})
}

//ArrayMinLength Creates a new ArrayMinLength constraint
func ArrayMinLength(length int, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if !value.IsArray() {
			return
		}

		if len(value.Array()) < length {
			violations.Add(field, message)
		}
	})
}

//ArrayMaxLength Creates a new ArrayMaxLength constraint
func ArrayMaxLength(length int, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if !value.IsArray() {
			return
		}

		if len(value.Array()) > length {
			violations.Add(field, message)
		}
	})
}

//ArrayLength Creates a new ArrayLength constraint
func ArrayLength(length int, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if !value.IsArray() {
			return
		}

		if len(value.Array()) != length {
			violations.Add(field, message)
		}
	})
}

//ArrayUnique Creates a new constraint for ensuring that elements of an array are unique
func ArrayUnique(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if !value.IsArray() {
			return
		}

		x := value.Array()

		for _, item := range value.Array() {
			duplicates := 0

			for _, target := range x {
				if cmp.Equal(item.Value(), target.Value()) {
					duplicates++
				}
			}

			if duplicates > 1 {
				violations.Add(field, message)
				break
			}
		}
	})
}

//ArrayUniqueObjectStringField Creates a new constraint for ensuring that elements of an array of objects has unique values for a given field
func ArrayUniqueObjectStringField(fieldName, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if !value.IsArray() {
			return
		}

		matchedValues := value.Get("#." + fieldName).Array()

		for position, curMatch := range matchedValues {
			dups := 0
			val := strings.ToLower(CollapseWhiteSpaces(curMatch.String()))

			for _, curVal := range matchedValues {
				if strings.ToLower(CollapseWhiteSpaces(curVal.String())) == val {
					dups++
				}
			}

			if dups > 1 {
				key := fmt.Sprintf("%s.%d.%s", field, position, fieldName)
				violations.Add(key, message)
				break
			}
		}

	})
}

//ArrayUniqueObjectNumberField Creates a new constraint for ensuring that elements of an array of objects has unique values for a given field
func ArrayUniqueObjectNumberField(fieldName, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if !value.IsArray() {
			return
		}

		matchedValues := value.Get("#." + fieldName).Array()

		for position, curMatch := range matchedValues {
			dups := 0
			val := curMatch.Float()

			for _, curVal := range matchedValues {
				if curVal.Float() == val {
					dups++
				}
			}

			if dups > 1 {
				key := fmt.Sprintf("%s.%d.%s", field, position, fieldName)
				violations.Add(key, message)
				break
			}
		}
	})
}
