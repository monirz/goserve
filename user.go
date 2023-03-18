package goserve

import (
	"errors"
	"regexp"
)

type UserService interface {
	CreateUser(*User) (int, error)
	FindUserByID(uuid string) (*User, error)
	GetUsers() (*[]User, error)
}

type User struct {
	ID       int64  `json:"id,omitempty"`
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	if u.Password == "" {
		return errors.New("password cannot be empty")
	}

	// validate email format
	match, _ := regexp.MatchString(`\b[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Z|a-z]{2,}\b`, u.Email)
	if !match {
		return errors.New("invalid email format")
	}

	// validate password strength
	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}
