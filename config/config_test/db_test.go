package config_test

import (
	"testing"
	"github.com/hftamayo/gingormrest/config"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnect(t *testing.T) {
	// Create a mock database object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock: %s", err)
	}

	defer db.Close()

	//wrap mock database in gorm.DB object
	dialect := postgres.New(postgres.Config{Conn: db})
	gdb, err := gorm.Open(dialect, &gorm.Config{SkipDefaultConfig: true})
	if err != nil {
		t.Errorf("Failed to connect to the database: %v", err)
	}

	//replace the global DB variable with the mock database object
	config.DB = gdb

	config.Connect()

	//ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

