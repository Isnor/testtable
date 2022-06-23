# Test Table

A small library to help write table-defined tests.

## Motivation

Table-defined tests are a [well-defined pattern](https://go.dev/doc/code#Testing) in Golang, but it often results in a lot of duplicated boiler plate code for defining the test definition and running the tests. I've found myself using the same patterns for testing various microservices and with the introduction of generics in go 1.18, I decided to try to capture them as a general pattern.

## Goals

The goal is for the user to write tests like:

```golang
func TestTestTable(t *testing.T) {

  // define a table of tests - put any implementation of `testtable.Test` here
  TestTable{
    // This is a generic "given this input, I expect this output"
    TestDefinition[FooInput, FooOutput]{
      Name: "Test the `Foobar` function",
      FunctionToTest: Foobar, // Foobar(input FooInput) FooOutput
      Input: FooInput{...},
      Parent: t,
      Expectations: func(result FooOutput) {
        // make assertions about the result, based on the Input and FunctionToTest
        if result == nil {
          t.Error("result should not be nil")
        }
      }
    },
    AnotherTestDefinitionType{...},
  }.Run()
```

The hope is that this is very easy to use for test-driven development and eventually, behaviour-driven tests. Ultimately, the hope for this library is that the programmer can just _write the tests_ and not need to worry about the boilerplate of defining, running, and maintaining them.

- [x] Test Table
- [ ] Generic Test Definition
- [ ] HTTP extension / module
- [ ] BDD Extension / module

## Use-cases

Test multiple endpoints of a mock service without defining several test structures
Quickly write unit tests