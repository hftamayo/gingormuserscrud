package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hftamayo/gingormrest/controller"
	"github.com/hftamayo/gingormrest/routes"
	"github.com/stretchr/testify/assert"
)

func TestUserRoute(t *testing.T) {
	// Create a router using the gin framework
	router := gin.Default()

	// Register routes using the UserRoute function
	routes.UserRoute(router)

	// Define a map to store registered routes
	registeredRoutes := map[string]string{
		"GET":    "/view-users",
		"POST":   "/add-user",
		"DELETE": "/delete-user/:id",
		"PUT":    "/update-user/:id",
	}

	// Check if all expected routes are registered
	for method, path := range registeredRoutes {
		routeInfo := router.Routes()

		// Search for the registered route
		var found bool
		for _, r := range routeInfo {
			if r.Path == path && r.Method == method {
				found = true
				break
			}
		}

		assert.Truef(t, found, "Route %s %s is not registered", method, path)
	}
}

// Mocking the controller functions for testing purposes
func init() {
	controller.GetUsers = func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GetUsers function"})
	}
	controller.CreateUser = func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "CreateUser function"})
	}
	controller.DeleteUser = func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "DeleteUser function"})
	}
	controller.UpdateUser = func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "UpdateUser function"})
	}
}

