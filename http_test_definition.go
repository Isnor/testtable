package testtable

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// HTTPTest is a [testtable.Test] for HTTP handlers. The purpose of HTTPTest is to make testing
// APIs with several endpoints, inputs, and outputs simple, and allow rapid test-driven development.
//
// Example:
//
//	testtable.TestTable{
//		&testtable.HTTPTest[any, MockResponse]{
//			Name:    "no pet found",
//			Path:    "/foobar?id=foobar",
//			Handler: mockAPI.mockGetEndpoint,
//			Expectations: func(t testing.TB, resp *http.Response, output *MockResponse) {
//				expectResponse(t, resp, 204)
//			},
//		},
//		testtable.GetTest(
//			"pet found",
//			"/foobar?id=dog",
//			&Pet{},
//			mockAPI.mockGetEndpoint,
//			func(t testing.TB, resp *http.Response, output *MockResponse) {
//				expectResponse(t, resp, 200)
//				if output == nil {
//					t.Error("did not get a pet")
//					return
//				}
//				if output.Name != "dog" {
//					t.Error("did not get dog")
//				}
//			},
//		),
//	}.Run(t)
type HTTPTest[InputModel any, OutputModel any] struct {
	Name         string
	Path         string
	Method       string
	Body         *InputModel
	Handler      http.HandlerFunc
	Expectations func(t testing.TB, response *http.Response, output *OutputModel)
}

// Run is the [testtable.Run] implementation for HTTPTest. It uses the [httptest] package to invoke the [testtable.HTTPTest.Handler]
// of the test with its Body, if it has one.
func (test HTTPTest[I, O]) Run(t *testing.T) {
	// try to deserialize a response body
	t.Run(test.Name, func(t *testing.T) {
		// create a test request with the test input if there is one
		request := httptest.NewRequest(test.Method, test.Path, nil)
		if test.Body != nil {
			body, err := json.Marshal(test.Body)
			if err != nil {
				t.Errorf("error writing post request body %v %v\n", err, test.Body)
				return
			}
			request = httptest.NewRequest(test.Method, test.Path, bytes.NewBuffer(body))
		}

		// execute the endpoint
		recorder := httptest.NewRecorder()
		test.Handler(recorder, request)
		var resultOutput O
		httpResponse := recorder.Result()
		if err := json.NewDecoder(httpResponse.Body).Decode(&resultOutput); err != nil {
			t.Errorf("failed deserializing response body %v\n", err)
		}

		// run the assertions about the output
		test.Expectations(t, httpResponse, &resultOutput)
	})
}

/* convenience functions */
// Issues:
// We can do this for every method, but they all need to be updated whenever the HTTPTest struct is.

// PostTest is a convenience function to avoid writing Method: "POST" in every test.
// Golang supports infering function type parameters based on the arguments it's provided, but it can't do the same
// with structs, so this function also allows the programmer to be slightly less verbose when defining tests:
/*
	testtable.HTTPTest[Pet, MockResponse]{
		Name:   "post foobar",
		Path:   "/foobar",
		Method: "POST",
		Body: &Pet{
			Name:  "Mock",
			Items: []string{"foo", "bar"},
		},
		Handler: mockAPI.mockPostEndpoint,
		Expectations: func(t testing.TB, resp *http.Response, output *MockResponse) {...},
	}

vs.

	testtable.PostTest(
		"post foobar",
		"/foobar",
		&Pet{
			Name:  "Mock",
			Items: []string{"foo", "bar"},
		},
		mockAPI.mockPostEndpoint,
		func(t testing.TB, resp *http.Response, output *MockResponse) {...},
	)
*/
func PostTest[I any, O any](
	name string,
	path string,
	body *I,
	handler http.HandlerFunc,
	expectations func(t testing.TB, response *http.Response, output *O),
) HTTPTest[I, O] {
	return HTTPTest[I, O]{
		Name:         name,
		Path:         path,
		Method:       "POST",
		Body:         body,
		Handler:      handler,
		Expectations: expectations,
	}
}

// GetTest is similar to [PostTest], but for the GET method.
func GetTest[I any, O any](
	name string,
	path string,
	body *I,
	handler http.HandlerFunc,
	expectations func(t testing.TB, response *http.Response, output *O),
) HTTPTest[I, O] {
	return HTTPTest[I, O]{
		Name:         name,
		Path:         path,
		Method:       "GET",
		Body:         body,
		Handler:      handler,
		Expectations: expectations,
	}
}
