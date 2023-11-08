# Test Table

A library to provide scaffolding for table-driven testing.

## Motivation

Table-driven tests are a [well-defined pattern](https://github.com/golang/go/wiki/TableDrivenTests) in Golang, but it often results in a lot of duplicated boiler plate code for defining the test definition and running the tests. I've found myself using the same patterns for testing various microservices and with the introduction of generics in go 1.18, I decided to try to capture them as a general pattern.

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

Ultimately, the hope for this library is that the programmer can just _write the tests_: "this is how I expect my function to behave given this input".

- [x] Test Table
- [x] HTTP Test Definition
- [ ] RPC Test Definition
- [ ] Generate tests given a spec / based on swagger annotations (?)

## Design

`testtable` is a very small package intended to make it easier to quickly produce and write tests that can be easily extended and modified later. It provides very little functionality and is intended to be built upon by implementing the `testtable.Run` function.

## Use-cases

Test multiple endpoints of a mock service without defining several test structures
Quickly write unit tests