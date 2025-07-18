package users

import (
	"mia_template_service_name_placeholder/middlewares"
	userService "mia_template_service_name_placeholder/services/users"
	
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserServiceForCreate struct {
	mock.Mock
}

func (m *MockUserServiceForCreate) GetByID(id string, name string) (*userService.User, error) {
	args := m.Called(id, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userService.User), args.Error(1)
}

func (m *MockUserServiceForCreate) Create(user *userService.User) (*userService.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userService.User), args.Error(1)
}

type CreateUserControllerTestSuite struct {
	suite.Suite
	mockService *MockUserServiceForCreate
	controller  *UserController
	echo        *echo.Echo
}

func (suite *CreateUserControllerTestSuite) SetupTest() {
	suite.mockService = new(MockUserServiceForCreate)
	suite.controller = &UserController{
		userService: suite.mockService,
	}
	suite.echo = echo.New()
	suite.echo.Validator = middlewares.NewValidator()
}

func (suite *CreateUserControllerTestSuite) TearDownTest() {
	suite.mockService.AssertExpectations(suite.T())
}

func TestUserController_CreateHandler_TestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserControllerTestSuite))
}

func (suite *CreateUserControllerTestSuite) TestSuccessfulCreateWithoutMetadata() {
	expectedUser := &userService.User{
		Name: "John Doe",
	}

	createdUser := &userService.User{
		ID:   "123",
		Name: "John Doe",
	}

	suite.mockService.On("Create", expectedUser).Return(createdUser, nil)

	requestBody := CreateRequest{
		Name: "John Doe",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)

	err := middlewares.Bind[CreateRequest](suite.controller.CreateHandler)(c)

	suite.NoError(err)
	suite.Equal(http.StatusCreated, rec.Code)

	var response CreateResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("123", response.ID)
}

func (suite *CreateUserControllerTestSuite) TestSuccessfulCreateWithMetadata() {
	expectedUser := &userService.User{
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	createdUser := &userService.User{
		ID:    "456",
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	suite.mockService.On("Create", expectedUser).Return(createdUser, nil)

	requestBody := CreateRequest{
		Name: "Jane Doe",
		Metadata: &struct {
			Age   int    `json:"age"`
			Email string `json:"email" validate:"required,email"`
		}{
			Age:   25,
			Email: "jane@example.com",
		},
	}

	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)

	err := middlewares.Bind[CreateRequest](suite.controller.CreateHandler)(c)

	suite.NoError(err)
	suite.Equal(http.StatusCreated, rec.Code)

	var response CreateResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("456", response.ID)
}

func (suite *CreateUserControllerTestSuite) TestMissingRequiredName() {
	requestBody := CreateRequest{}

	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)

	err := middlewares.Bind[CreateRequest](suite.controller.CreateHandler)(c)

	suite.NoError(err)
	suite.Equal(http.StatusBadRequest, rec.Code)

	var errorResponse map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &errorResponse)
	suite.NoError(err)
	suite.Contains(errorResponse, "error")
	suite.Equal("Invalid input", errorResponse["error"])
	suite.Contains(errorResponse, "validation")
}

func (suite *CreateUserControllerTestSuite) TestInvalidEmailInMetadata() {
	requestBody := CreateRequest{
		Name: "John Doe",
		Metadata: &struct {
			Age   int    `json:"age"`
			Email string `json:"email" validate:"required,email"`
		}{
			Age:   30,
			Email: "invalid-email",
		},
	}

	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)

	err := middlewares.Bind[CreateRequest](suite.controller.CreateHandler)(c)

	suite.NoError(err)
	suite.Equal(http.StatusBadRequest, rec.Code)

	var errorResponse map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &errorResponse)
	suite.NoError(err)
	suite.Contains(errorResponse, "error")
	suite.Equal("Invalid input", errorResponse["error"])
	suite.Contains(errorResponse, "validation")
}

func (suite *CreateUserControllerTestSuite) TestMissingEmailInMetadata() {
	requestBody := CreateRequest{
		Name: "John Doe",
		Metadata: &struct {
			Age   int    `json:"age"`
			Email string `json:"email" validate:"required,email"`
		}{
			Age: 30,
		},
	}

	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)

	err := middlewares.Bind[CreateRequest](suite.controller.CreateHandler)(c)

	suite.NoError(err)
	suite.Equal(http.StatusBadRequest, rec.Code)

	var errorResponse map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &errorResponse)
	suite.NoError(err)
	suite.Contains(errorResponse, "error")
	suite.Equal("Invalid input", errorResponse["error"])
	suite.Contains(errorResponse, "validation")
}

func (suite *CreateUserControllerTestSuite) TestServiceError() {
	expectedUser := &userService.User{
		Name: "John Doe",
	}

	suite.mockService.On("Create", expectedUser).Return(nil, assert.AnError)

	requestBody := CreateRequest{
		Name: "John Doe",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)

	err := middlewares.Bind[CreateRequest](suite.controller.CreateHandler)(c)

	suite.NoError(err)
	suite.Equal(http.StatusInternalServerError, rec.Code)

	var errorResponse map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &errorResponse)
	suite.NoError(err)
	suite.Contains(errorResponse, "error")
	suite.NotEmpty(errorResponse["error"])
}
