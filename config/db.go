package config

import (
	"log"
	"os"
	"fmt"

	"github.com/hftamayo/gingormrest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	databaseName := os.Getenv("DATABASE_NAME")
	connectionString := fmt.Sprintf(databaseURL, databaseName)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open("postgres://gobank:gobank@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to the database.\n%v", err)
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	DB = db
	return db, nil
}
