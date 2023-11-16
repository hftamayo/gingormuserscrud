package test

import (
	"testing"
	"github.com/hftamayo/gingormrest/models"
)

func TestValidateUser(t *testing.T) {
	cases := []struct {
		name   string
		user   models.User
		isValid bool
	}{
		{
			name:   "Valid user",
			user:   models.User{Name: "Herbert Tamayo", Email: "hftamayo@gmail.com", Password: "password123"},
			isValid: true,
		},
		{
			name:   "Invalid user with empty name",
			user:   models.User{Name: "", Email: "hftamayo@gmail.com", Password: "password123"},
			isValid: false,
		},
		{
			name:   "Invalid user with invalid email",
			user:   models.User{Name: "Herbert Tamayo", Email: "tamayoemail.com", Password: "password123"},
			isValid: false,
		},
		{
			name:   "Invalid user with empty password",
			user:   models.User{Name: "Herbert Tamayo", Email: "hftamayo@gmail.com", Password: ""},
			isValid: false,
		},
	}


	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.user.Validate()
			if c.isValid && err != nil {
				t.Errorf("Expected valid user, but got error: %v", err)
			} else if !c.isValid && err == nil {
				t.Errorf("Expected invalid user, but no error was returned")
			}
		})
	}
}

