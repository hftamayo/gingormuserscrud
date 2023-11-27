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
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, databaseName)
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
