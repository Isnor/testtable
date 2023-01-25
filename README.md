# Test Table

A small library to help write table-defined tests.

## Motivation

Table-defined tests are a [well-defined pattern](https://github.com/golang/go/wiki/TableDrivenTests) in Golang, but it often results in a lot of duplicated boiler plate code for defining the test definition and running the tests. I've found myself using the same patterns for testing various microservices and with the introduction of generics in go 1.18, I decided to try to capture them as a general pattern.

## Goals

The goal is to remove the need to write boiler plate code to run tests. 

```golang
func TestFoobarEndpoint(t *testing.T) {

  // define a table of tests - put any implementation of `testtable.Test` here
  TestTable{
    // This is a generic "given this input, I expect this output"
    &testtable.PostTest[any, MockResponse]{
			Name:    "post foobar with empty payload",
			Path:    "/foobar",
			Handler: mockAPI.mockPostEndpoint,
			Expectations: func(t testing.TB, resp *http.Response, output *MockResponse) {
				expectResponse(t, resp, 400)

				if output == nil {
					t.Error("output should not be null")
					return
				}

				if output.Status != "no pet provided" {
					t.Errorf("unexpected status \"%s\"\n", output.Status)
				}
			},
		},
    AnotherTestDefinitionType{...},
  }.Run()
}
```

The hope is that this is very easy to use for test-driven development and eventually, behaviour-driven tests. Ultimately, the hope for this library is that the programmer can just _write the tests_ and not need to worry about the boilerplate of defining, running, and maintaining them.

- [x] Test Table
- [x] Generic Test Definition
- [x] HTTP Test Definition
- [ ] RPC Test Definition

## Design

`testtable` is meant to be extremely small and often won't be used directly by the programmer writing tests, but instead a programmer creating a testing library. The "useful" parts of `testtable` will be the modules built on top of it. The first of these modules and the initial motivation for this project will be testing API endpoints.

## Use-cases

Test multiple endpoints of a mock service without defining several test structures
Quickly write unit tests