package jsonvalidator

import (
	"fmt"

	"github.com/tidwall/gjson"
)

//ArrayConstraint Array validation constraint
type ArrayConstraint struct {
	constraints []Constraint
}

//Validate Validates array
func (ar *ArrayConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if !value.IsArray() {
		return
	}

	for index, item := range value.Array() {
		for _, rule := range ar.constraints {
			entryField := fmt.Sprintf("%s.%d", field, index)
			rule.Validate(entryField, &item, value, source, violations)
		}
	}
}

//Array Returns an instance of Array validation constraints
func Array(constraints []Constraint) *ArrayConstraint {
	return &ArrayConstraint{constraints: constraints}
}

//ArrayMinLengthConstraint Array min length constraint
type ArrayMinLengthConstraint struct {
	length  int
	message string
}

//Validate Validates array for minimum number of entries
func (arml *ArrayMinLengthConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if !value.IsArray() {
		return
	}

	if len(value.Array()) < arml.length {
		violations.Add(field, arml.message)
	}
}

//ArrayMinLength Creates a new ArrayMinLength constraint
func ArrayMinLength(length int, message string) *ArrayMinLengthConstraint {
	return &ArrayMinLengthConstraint{length: length, message: message}
}

//ArrayMaxLengthConstraint Array max length constraint
type ArrayMaxLengthConstraint struct {
	length  int
	message string
}

//Validate Validates array for minimum number of entries
func (arml *ArrayMaxLengthConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if !value.IsArray() {
		return
	}

	if len(value.Array()) > arml.length {
		violations.Add(field, arml.message)
	}
}

//ArrayMaxLength Creates a new ArrayMaxLength constraint
func ArrayMaxLength(length int, message string) *ArrayMaxLengthConstraint {
	return &ArrayMaxLengthConstraint{length: length, message: message}
}

//ArrayLengthConstraint Array length constraint
type ArrayLengthConstraint struct {
	length  int
	message string
}

//Validate Validates array for minimum number of entries
func (arml *ArrayLengthConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if !value.IsArray() {
		return
	}

	if len(value.Array()) != arml.length {
		violations.Add(field, arml.message)
	}
}

//ArrayLength Creates a new ArrayLength constraint
func ArrayLength(length int, message string) *ArrayLengthConstraint {
	return &ArrayLengthConstraint{length: length, message: message}
}
