package xeno

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if client.httpClient != http.DefaultClient {
		t.Errorf("Expected default HTTP client, got different client")
	}

	if client.baseURL.String() != APIEndpointBase {
		t.Errorf("Expected baseURL %s, got %s", APIEndpointBase, client.baseURL.String())
	}
}

func TestClient_Get(t *testing.T) {
	testCases := []struct {
		name           string
		query          string
		expectedResult *Response
		mockResponse   string
		mockStatusCode int
	}{
		{
			name:           "Successful response",
			query:          "test",
			expectedResult: &Response{Page: 1},
			mockResponse:   `{"page":1}`,
			mockStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.mockStatusCode)
				w.Write([]byte(tc.mockResponse))
			}))
			defer mockServer.Close()

			client, _ := NewClient(WithBaseURL(mockServer.URL + "?query="))

			response, err := client.Get(context.Background(), tc.query)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(response, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, response)
			}
		})
	}
}
