package models

import (
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

type Task struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `json:"title" db:"title" fako:"job_title" `
	Description string    `json:"description" db:"description" fako:"sentence" `
	CreatedAt   time.Time `db:"created_at" `
	UpdatedAt   time.Time `db:"updated_at" `
	UserID      uuid.UUID `json:"user_id" db:"user_id" `
	Status      bool      `json:"status" db:"status" `
	Complete    time.Time `db:"completeby" `
	User        *User     `belongs_to:"users"`
}

type Tasks []Task

func (t *Task) Validate() *validate.Errors {

	return validate.Validate(
		&validators.StringIsPresent{Field: t.Title, Name: "Title"},
		&validators.StringIsPresent{Field: t.Description, Name: "Description"},
		&validators.FuncValidator{
			Fn: func() bool {
				if t.UserID == uuid.Nil {
					return false
				}
				return true
			},
			Field:   "",
			Name:    "UserID",
			Message: "%s User can't be blank.",
		},

		&validators.FuncValidator{
			Fn: func() bool {
				if len(t.Title) > 50 {
					return false
				}
				return true
			},
			Field:   "",
			Name:    "Title",
			Message: "%s Title must be less than 50 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if len(t.Description) > 450 {
					return false
				}
				return true
			},
			Field:   "",
			Name:    "Description",
			Message: "%s Description must be less than 450 characters.",
		},
	)

}
