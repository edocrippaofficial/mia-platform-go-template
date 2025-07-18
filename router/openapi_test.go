package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestSwaggerUI(t *testing.T) {
	log, _ := test.NewNullLogger()
	router := NewRouter(log)

	// Check if Swagger UI is exposed
	assert.NotNil(t, router.OpenAPI, "OpenAPI should not be nil")
	assert.NotEmpty(t, router.OpenAPI.Info.Title, "OpenAPI title should not be empty")
	assert.NotEmpty(t, router.OpenAPI.Info.Version, "OpenAPI version should not be empty")

	// Make a request to the Swagger UI endpoint
	req := httptest.NewRequest(http.MethodGet, "/documentation.json", nil)
	w := httptest.NewRecorder()
	router.Handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Expected status code to be 200")

	// Check if the response is a valid JSON
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")
	assert.NotEmpty(t, response, "Response should not be empty")
}
