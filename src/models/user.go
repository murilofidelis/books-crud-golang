package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
}

func (user *User) Validate() error {
	user.format()
	if err := user.checkValidate(); err != nil {
		return err
	}
	return nil
}

func (user *User) checkValidate() error {
	if user.Name == "" {
		return errors.New("name is required.")
	}
	if user.Nick == "" {
		return errors.New("nick is required.")
	}
	if user.Email == "" {
		return errors.New("email is required.")
	}
	if user.Password == "" {
		return errors.New("password is required.")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
