// controller_test.go

package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hftamayo/gingormrest/config"
	"github.com/hftamayo/gingormrest/controller"
	"github.com/hftamayo/gingormrest/models"
)

func TestCreateUser(t *testing.T) {
	// Set up a mock Gin context and request
	e := gin.Default()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name": "Alice", "email": "alice@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	c := e.NewContext(req, rec)

	// Create a mock user
	mockUser := models.User{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	// Create a mock database connection
	mockDB := config.DB.Begin()
	defer mockDB.Rollback()

	// Mock the behavior of the CreateUser function
	mockDB.On("Create", &mockUser).Return(&mockUser, nil)

	// Call the actual controller function
	controller.CreateUser(c)

	// Check the response status code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", rec.Code)
	}

	// Validate the response body (you can add more assertions here)
	expectedResponse := `{"name":"Alice","email":"alice@example.com"}`
	if rec.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rec.Body.String())
	}
}

