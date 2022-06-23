package testtable

import (
	"testing"
)

// Test is the base type
type Test interface {
	Run(*testing.T)
}

// TestTable is a slice of `Test`s. This allows us to define many different types of tests and test definitions in one slice
// and run them all without needing to duplicate the looping logic
type TestTable []Test

// TestTable's Run method loops over all of and runs the tests in it
func (tests TestTable) Run(t *testing.T) {
	for _, test := range tests {
		test.Run(t)
	}
}

// GenericTestDefinition is a generic description of a simple "given this input then I expect this function to do X"
type GenericTestDefinition[InputType any, OutputType any] struct {
	Name           string
	Input          InputType
	FunctionToTest func(InputType) OutputType
	Expectations   func(OutputType)
}

// The general Run method runs the `FunctionToTest` with the `Input` and `Expectations` of the test definition
func (test *GenericTestDefinition[I, O]) Run(parentTest *testing.T) {
	parentTest.Run(test.Name, func(*testing.T) {
		test.Expectations(test.FunctionToTest(test.Input))
	})
}
