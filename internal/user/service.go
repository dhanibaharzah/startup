package user

import (
	"errors"
)

func ValidateUser(user *User) error {
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if user.Role == "" {
		user.Role = "user" // Default role if not provided
	}
	return nil
}
