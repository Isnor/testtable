# Test Table

A library to help accelerate and maintain testing

## Motivation

Table-driven tests are a [well-defined pattern](https://github.com/golang/go/wiki/TableDrivenTests) in Golang, but it often results in a lot of duplicated boiler plate code for defining the test definition and running the tests. 

## Goals

Ultimately, the hope for this library is that the programmer can just _write the tests_: "this is how I expect my function to behave given this input".

## Design

`testtable` is a very small package intended to make it easier to quickly produce and write tests that can be easily refactored. It doesn't provide functionality - it's intended to be built upon by implementing the `testtable.Run` function.