package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Fullname  string    `json:"fullname,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Passord   string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) Validate(state string) error {
	err := u.validateEmptyFields(state)
	if err != nil {
		return err
	}

	u.formatter()

	return nil
}

func (u *User) validateEmptyFields(state string) error {
	if u.Fullname == "" {
		return errors.New("O nome é obrigatório")
	}

	if u.Email == "" {
		return errors.New("O email é obrigatório")
	}

	if u.Nick == "" {
		return errors.New("O nick é obrigatório")
	}

	if state == "cadastro" && u.Passord == "" {
		return errors.New("A senha é obrigatória")
	}

	return nil
}

func (u *User) formatter() {
	u.Fullname = strings.TrimSpace(u.Fullname)
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)
}
