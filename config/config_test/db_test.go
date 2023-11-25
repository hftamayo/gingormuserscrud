package config_test

import (
	"testing"

	"github.com/hftamayo/gingormrest/config"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnect(t *testing.T) {
	// Create a mock database object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock: %s", err)
	}

	defer db.Close()

	//set expectations for the mock db object
	mock.ExpectPing()

	//replace the global DB variable with the mock database object
	config.DB = db

	config.Connect()

	//ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

