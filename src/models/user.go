package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
}

func (user *User) Validate(step string) error {
	if err := user.format(step); err != nil {
		return err
	}
	if err := user.checkValidate(step); err != nil {
		return err
	}
	return nil
}

func (user *User) checkValidate(step string) error {
	if user.Name == "" {
		return errors.New("name is required.")
	}
	if user.Nick == "" {
		return errors.New("nick is required.")
	}
	if user.Email == "" {
		return errors.New("email is required.")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email invalid!")
	}
	if step == "create" && user.Password == "" {
		return errors.New("password is required.")
	}
	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if step == "create" {
		passworwithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passworwithHash)
	}
	return nil
}
