package config

import (
	"github.com/hftamayo/gingormrest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://gobank:gobank@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
