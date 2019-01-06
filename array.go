package jsonvalidator

import (
	"fmt"

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
			counts := 0

			for _, target := range x {
				if cmp.Equal(item.Value(), target.Value()) {
					counts++
				}
			}

			if counts > 1 {
				violations.Add(field, message)
				break
			}
		}
	})
}
