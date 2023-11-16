package models

import (
	"errors"
	"strings"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id int `json:"ID" gorm:"primary_key"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (user *User) Validate() error {
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}

	if !strings.Contains(user.Email, "@") {
		return errors.New("invalid email address")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	return nil
}
