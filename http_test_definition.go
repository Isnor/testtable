package testtable

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
)

// HTTPTest is a special type of GenericTestDefinition whose purpose is to make testing many `http.HandlerFunc`s
// with different types of payloads easier. Because the signature of `http.HandlerFunc` and the test functions `func(input) output`,
// the easiest way to use this is to call NewHTTPTest which does the wrapping and input marshaling
type HTTPTest[requestType any] GenericTestDefinition[requestType, *http.Response]

// NewHTTPTest returns
func NewHTTPTest[requestType any](name string, input *requestType, functionToTest http.HandlerFunc, expectations func(*http.Response)) HTTPTest[*requestType] {
	return HTTPTest[*requestType]{
		Name:           name,
		Input:          input,
		FunctionToTest: handlerFuncToTestFunc[requestType](functionToTest),
		Expectations:   expectations,
	}
}

// TODO:  ah, too bad: the type alias doesn't allow us to inherit the functions defined for the type we alias. I guess that's why they
// call their functions "receivers". In any case, given this restriction is there a benefit in using the generic test definition at all?
//
// func (test *HTTPTest[requestType]) Run(t *testing.T) {

// }

// turn an http.HandlerFunc into a function from http.Request -> http.Response
func handlerFuncToTestFunc[requestInputType any](h http.HandlerFunc) func(*requestInputType) *http.Response {
	return func(input *requestInputType) *http.Response {
		request := httptest.NewRequest("", "", nil)
		if input != nil {
			body, err := json.Marshal(input)
			// TODO: handle error better
			if err != nil {
				log.Fatalf("you've made me very sad with your non-JSON input. i can't believe this is happening %v %v", err, input)
				return nil
			}
			request = httptest.NewRequest("", "", bytes.NewBuffer(body))
		}
		recorder := httptest.NewRecorder()
		h(recorder, request)
		return recorder.Result()
	}
}
