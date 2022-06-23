package testtable

import (
	"testing"
)

// Test is the base type
type Test interface {
	Run()
}

type TestTable []Test

func (tests TestTable) Run() {
	for _, test := range tests {
		test.Run()
	}
}

// GenericTestDefinition is a generic description of a simple "given this input then I expect this function to do X"
type GenericTestDefinition[InputType any, OutputType any] struct {
	Name           string
	Parent         *testing.T
	Input          InputType
	FunctionToTest func(InputType) OutputType
	Expectations   func(OutputType)
}

// The general Run method runs the `FunctionToTest` with the `Input` and `Expectations` of the test definition
func (test *GenericTestDefinition[I, O]) Run() {
	test.Parent.Run(test.Name, func(t *testing.T) {
		test.Expectations(test.FunctionToTest(test.Input))
	})
}
