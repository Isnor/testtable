package testtable_test

import (
	"testing"
	. "testtable"
)

func TestStringLength(t *testing.T) {

	TestTable{
		// test nil string
		&GenericTestDefinition[*string, int]{
			Name:           "Given a nil string, StringLength should return -1",
			FunctionToTest: StringLength,
			Input:          nil,
			Expectations: func(result int) {
				if result != -1 {
					t.Error("result should have been -1")
				}
			},
		},
		// test empty string
		&GenericTestDefinition[*string, int]{
			Name:           "Given an empty string, StringLength should return -1",
			FunctionToTest: StringLength,
			Input:          func(s string) *string {return &s}(""),
			Expectations: func(result int) {
				if result != -1 {
					t.Error("result should have been -1")
				}
			},
		},
		// test a non-empty string
		&GenericTestDefinition[*string, int]{
			Name:           "Given a non-empty string, StringLength should return the length of the string",
			FunctionToTest: StringLength,
			Input:          func(s string) *string {return &s}("hello, world!"),
			Expectations: func(result int) {
				if result <= 0 {
					t.Error("result should have been > 0")
				}
			},
		},
	}.Run(t)
}

// StringLength returns the length of a string, or -1 if an empty string or nil is provided
func StringLength(input *string) int {
	if input == nil {
		return -1
	}
	if length := len(*input); length == 0 {
		return -1
	} else {
		return length
	}
}