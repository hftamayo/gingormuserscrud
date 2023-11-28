package config

import (
	"log"
	"os"
	"fmt"
	"strconv"
	"errors"

	"github.com/hftamayo/gingormrest/models"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	if !isTestEnviro() {
		if err := godotenv.Load(); err != nil {
			fmt.Println("Error reading context data")
		}
		user := os.Getenv("DATABASE_USER")
		password := os.Getenv("DATABASE_PASSWORD")
		databaseName := os.Getenv("DATABASE_NAME")
		host := os.Getenv("DATABASE_HOST")
		portStr := os.Getenv("DATABASE_PORT")

		port, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Println("Error converting value to integer:", err)
		}

		connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, databaseName)
		db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		//db, err := gorm.Open(postgres.Open("postgres://gobank:gobank@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
		if err != nil {
			log.Printf("Error connecting to the database.\n%v", err)
			return nil, err
		}
		db.AutoMigrate(&models.User{})
		DB = db
		return db, nil
	} else {
		return nil, errors.New("running in testing mode")
	}
}

func isTestEnviro() bool {
	envMode := os.Getenv("GOAPP_MODE")
	switch envMode {
	case "development":
		fmt.Println("running in development mode")
		return false
	case "testing":
		fmt.Println("running in testing mode")
		return true
	default:
		fmt.Println("running in production mode")
		return false
	}
}
