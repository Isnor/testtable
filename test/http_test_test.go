package testtable_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Isnor/testtable"
)

// simple nonsense "API"
var mockAPI = &MockAPI{
	Store: map[string]Pet{
		"dog": {
			Name:  "dog",
			Items: []string{"head", "legs"},
		},
	},
}

func TestGet(t *testing.T) {
	testtable.TestTable{
		&testtable.HTTPTest[any, MockResponse]{
			Name:    "no pet found",
			Path:    "/foobar?id=foobar",
			Handler: mockAPI.mockGetEndpoint,
			Expectations: func(t testing.TB, resp *http.Response, output *MockResponse) {
				expectResponse(t, resp, 204)
			},
		},
		testtable.GetTest(
			"pet found",
			"/foobar?id=dog",
			&Pet{},
			mockAPI.mockGetEndpoint,
			func(t testing.TB, resp *http.Response, output *MockResponse) {
				expectResponse(t, resp, 200)
				if output == nil {
					t.Error("did not get a pet")
					return
				}
				if output.Name != "dog" {
					t.Error("did not get dog")
				}
			},
		),
	}.Run(t)
}

func TestPost(t *testing.T) {

	testtable.TestTable{
		&testtable.HTTPTest[any, MockResponse]{
			Name:    "post foobar with empty payload",
			Path:    "/foobar",
			Method:  "POST",
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
		testtable.PostTest(
			"post raboof",
			"/raboof",
			&Pet{
				Name:  "Mock",
				Items: []string{"foo", "bar"},
			},
			mockAPI.mockPostEndpoint,
			func(t testing.TB, resp *http.Response, output *MockResponse) {
				expectResponse(t, resp, 200)
			},
		),
		testtable.PostTest(
			"post raboof with invalid payload",
			"/raboof",
			&Pet{},
			mockAPI.mockPostEndpoint,
			func(t testing.TB, resp *http.Response, output *MockResponse) {
				expectResponse(t, resp, 400)
			},
		),
	}.Run(t)
}

func expectResponse(t testing.TB, r *http.Response, statusCode int) {
	if r.StatusCode != statusCode {
		t.Errorf("expected response to have status code %d, had status code '%d'\n", statusCode, r.StatusCode)
	}
}

// ----- mock API -----

type MockAPI struct {
	Store map[string]Pet
}

func (m *MockAPI) mockGetEndpoint(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(MockResponse{
			Status: "no id provided",
		})
		return
	}
	pet, found := m.Store[id]
	if !found {
		w.WriteHeader(204)
		json.NewEncoder(w).Encode(MockResponse{
			Status: "no pet found",
		})
		return
	}
	json.NewEncoder(w).Encode(pet)
}

type Pet struct {
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

type MockResponse struct {
	Status string `json:"status"`
	*Pet
}

func (m *MockAPI) mockPostEndpoint(response http.ResponseWriter, request *http.Request) {

	var requestPayload Pet
	if err := json.NewDecoder(request.Body).Decode(&requestPayload); err != nil {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(MockResponse{
			Status: "no pet provided",
		})
		return
	}

	if len(requestPayload.Name) == 0 {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(MockResponse{
			Status: "invalid pet provided",
		})
		return
	}

	m.Store[requestPayload.Name] = requestPayload
	json.NewEncoder(response).Encode(MockResponse{Status: "pet added", Pet: &requestPayload})
}
