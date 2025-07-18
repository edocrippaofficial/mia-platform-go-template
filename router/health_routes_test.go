package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/suite"
)

type HealthControllerTestSuite struct {
	suite.Suite
	router *Router
}

func TestHealthControllerSuite(t *testing.T) {
	suite.Run(t, new(HealthControllerTestSuite))
}

func (s *HealthControllerTestSuite) SetupTest() {
	logger, _ := test.NewNullLogger()
	logger.SetLevel(logrus.FatalLevel)
	s.router = NewRouter(logger)
}

func (s *HealthControllerTestSuite) TestHealthOkResponse() {
	// Make a request to the /-/healthz endpoint
	req := httptest.NewRequest(http.MethodGet, "/-/healthz", nil)
	w := httptest.NewRecorder()

	// Execute
	s.router.Handler.ServeHTTP(w, req)

	// Assertions
	s.Equal(http.StatusOK, w.Code, "Expected status code to be 200")

	var response HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Assert().NoError(err, "Should be able to unmarshal response")
	s.Assert().Equal("OK", response.Status)
}
