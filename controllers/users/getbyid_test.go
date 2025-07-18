package users

import (
	"mia_template_service_name_placeholder/middlewares"
	userService "mia_template_service_name_placeholder/services/users"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetByID(id string, name string) (*userService.User, error) {
	args := m.Called(id, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userService.User), args.Error(1)
}
func (m *MockUserService) Create(user *userService.User) (*userService.User, error) { return nil, nil }

type UserControllerTestSuite struct {
	suite.Suite
	mockService *MockUserService
	controller  *UserController
	echo        *echo.Echo
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockService = new(MockUserService)
	suite.controller = &UserController{
		userService: suite.mockService,
	}
	suite.echo = echo.New()
	suite.echo.Validator = middlewares.NewValidator()
}

func (suite *UserControllerTestSuite) TearDownTest() {
	suite.mockService.AssertExpectations(suite.T())
}

func TestUserController_GetByIDHandler_TestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (suite *UserControllerTestSuite) TestSuccessfulRequest() {
	suite.mockService.On("GetByID", "123", "John").Return(&userService.User{
		ID:    "123",
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/123?name=John", nil)
	req.Header.Set("x-api-key", "valid-api-key")

	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("123")

	err := middlewares.Bind[GetByIDRequest](suite.controller.GetByIDHandler)(c)

	suite.NoError(err)
	suite.Equal(http.StatusOK, rec.Code)

	var response GetByIDResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("123", response.ID)
	suite.Equal("John Doe", response.Name)
	suite.Equal("john@example.com", response.Email)
}

func (suite *UserControllerTestSuite) TestUserNotFound() {
	suite.mockService.On("GetByID", "999", "").Return(nil, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
	req.Header.Set("x-api-key", "valid-api-key")

	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	err := middlewares.Bind[GetByIDRequest](suite.controller.GetByIDHandler)(c)

	suite.NoError(err)
	suite.Equal(http.StatusNotFound, rec.Code)

	var errorResponse map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &errorResponse)
	suite.NoError(err)
	suite.Contains(errorResponse, "error")
	suite.NotEmpty(errorResponse["error"])
}

func (suite *UserControllerTestSuite) TestMissingApiKey() {
	req := httptest.NewRequest(http.MethodGet, "/users/123?name=John", nil)

	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("123")

	err := middlewares.Bind[GetByIDRequest](suite.controller.GetByIDHandler)(c)

	suite.NoError(err)
	suite.Equal(http.StatusBadRequest, rec.Code)

	var errorResponse map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &errorResponse)
	suite.NoError(err)
	suite.Contains(errorResponse, "error")
	suite.Equal("Invalid input", errorResponse["error"])
	suite.Contains(errorResponse, "validation")
}
