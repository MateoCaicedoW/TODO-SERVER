package models

import (
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

type User struct {
	ID                   uuid.UUID `db:"id" `
	Email                string    `db:"email"`
	FirstName            string    `db:"first_name" `
	LastName             string    `db:"last_name" `
	Password             string    `db:"-" `
	PasswordHash         string    `db:"password_hash" `
	PasswordConfirmation string    `db:"-"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
	Rol                  string    `db:"rol"`
	Tasks                []Task    `has_many:"tasks"`
}

type Users []User

func (u *User) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName"},
		&validators.StringIsPresent{Field: u.LastName, Name: "LastName"},
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringIsPresent{Field: u.PasswordConfirmation, Name: "PasswordConfirmation"},
	)
}
