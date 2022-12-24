package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

type User struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	Email                string    `json:"email" db:"email"`
	FirstName            string    `json:"first_name" db:"first_name"`
	LastName             string    `json:"last_name" db:"last_name"`
	Password             string    `json:"password" db:"-"`
	PasswordHash         string    `db:"password_hash" `
	PasswordConfirmation string    `json:"password_confirmation" db:"-"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
	Rol                  string    `json:"role" db:"rol"`
	Tasks                []Task    `has_many:"tasks"`
}

type Users []User

func (u *User) Validate(tx *pop.Connection) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName"},
		&validators.StringIsPresent{Field: u.LastName, Name: "LastName"},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.FirstName != "" && !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(u.FirstName) {
					return false
				}
				return true
			},
			Name:    "First Name",
			Message: "%s First Name must be letters only.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.LastName != "" && !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(u.LastName) {
					return false
				}
				return true
			},
			Name:    "Last Name",
			Message: "%s Last Name must be letters only.",
		},
		&validators.FuncValidator{

			Fn: func() bool {
				if u.FirstName != "" && len(u.FirstName) > 50 && regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.FirstName) {
					return false
				}
				return true
			},
			Name:    "First Name",
			Message: "%s First Name must be less than 50 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.LastName != "" && len(u.LastName) > 50 && regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.LastName) {
					return false
				}
				return true
			},
			Name:    "Last Name",
			Message: "%s Last Name must be less than 50 characters.",
		},
		&validators.FuncValidator{

			Name:    "Email",
			Message: "%s Email is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("email = ?", u.Email)
				if u.ID != uuid.Nil {
					q = q.Where("id != ?", u.ID)
				}
				b, err := q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},
		&validators.FuncValidator{
			Fn: func() bool {
				ex := `^[\w\.]+@([\w-]+\.)+[\w-]{2,4}$`
				if u.Email != "" && !regexp.MustCompile(ex).MatchString(u.Email) {
					return false
				}
				return true
			},
			Name:    "Email",
			Message: "%s Email is invalid",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.Email != "" && strings.Contains(u.Email, "@") {
					local := strings.Split(u.Email, "@")
					str := local[0]
					if len(str) > 64 {
						return false
					}
				}
				return true
			},
			Name:    "Email",
			Message: "%s Before @ Email must be less or equal than 64 characters ",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.Email != "" && strings.Contains(u.Email, "@") {
					local := strings.Split(u.Email, "@")
					str := local[1]
					if len(str) > 255 {
						return false
					}
				}
				return true
			},
			Name:    "Email",
			Message: "%s After @ Email must be less or equal than 255 characters ",
		},
	)
}
