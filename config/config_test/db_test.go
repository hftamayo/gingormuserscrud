package config_test

import (
	"testing"

	"github.com/hftamayo/gingormrest/config"
	sqlmock "github.com/stretchr/testify/mock"
)

func TestConnect(t *testing.T) {
	// Create a mock database object
	dbMock := sqlmock.Open(`postgres`)

	// Set expectations for the mock database object
	dbMock.Expect().Ping().WillReturn(nil)

	// Replace the global DB variable with the mock database object
	config.DB = dbMock

	// Call the Connect function
	config.Connect()

	// Verify that the mock database object was called
	dbMock.AssertExpectations(t)
}
