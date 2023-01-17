package testtable_test

import (
	"encoding/json"
	"net/http"
	"testing"
	. "testtable"
)

type void *interface{}

func TestHTTPTest(t *testing.T) {
	api := &MockAPI{}
	TestTable{
		// NewHTTPTest[void](
		// 	"Test GET",
		// 	nil,
		// 	api.mockGetEndpoint,
		// 	func(r *http.Response) {
		// 		expectResponse(t, r, 200)
		// 	},
		// ),
		// NewHTTPTest(
		// 	"Test POST one",
		// 	&MockRequestPayload{},
		// 	api.mockPostEndpoint,
		// 	func(r *http.Response) {
		// 		expectResponse(t, r, 200)
		// 	},
		// ),
	}.Run(t)
}

func expectResponse(t *testing.T, r *http.Response, statusCode int) {
	if r.StatusCode != statusCode {
		t.Fail()
		t.Logf("expected response to have status code %d, had status code '%d'", statusCode, r.StatusCode)
	}
}

// ----- mock API -----

type MockAPI struct{}

func (m *MockAPI) mockGetEndpoint(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte("hello, mock!"))
}

type MockRequestPayload struct {
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

func (m *MockAPI) mockPostEndpoint(response http.ResponseWriter, request *http.Request) {

	var requestPayload MockRequestPayload
	if err := json.NewDecoder(request.Body).Decode(&requestPayload); err != nil {
		response.WriteHeader(400)
		response.Write([]byte("invalid request"))
		return
	}

	requestPayload.Items = append(requestPayload.Items, "foobar")

	response.WriteHeader(200)
	json.NewEncoder(response).Encode(requestPayload)
}

func (m *MockAPI) anotherMockPostEndpoint(response http.ResponseWriter, request *http.Request) {

	var requestPayload MockRequestPayload
	if err := json.NewDecoder(request.Body).Decode(&requestPayload); err != nil {
		response.WriteHeader(400)
		response.Write([]byte("invalid request"))
		return
	}

	if len(requestPayload.Name) == 0 {
		response.WriteHeader(400)
		response.Write([]byte("request payload must have a name"))
	}

	requestPayload.Items = append(requestPayload.Items, "raboof")

	response.WriteHeader(200)
	json.NewEncoder(response).Encode(requestPayload)
}
