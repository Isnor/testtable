// Package testtable is a testing library to quickly produce [table defined tests].
//
// [table defined tests]: https://github.com/golang/go/wiki/TableDrivenTests
package testtable

import "testing"

// Test describes something that can run within a testing context.
type Test interface {
	Run(*testing.T)
}

// TestTable is a slice of `Test`s.
type TestTable []Test

// Run simply loops over the tests in the TestTable and invokes their run functions.
func (tests TestTable) Run(t *testing.T) {
	for _, test := range tests {
		test.Run(t)
	}
}
