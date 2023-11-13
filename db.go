package config

import (
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *form.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
