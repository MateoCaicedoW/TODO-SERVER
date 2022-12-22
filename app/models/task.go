package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Task struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title" fako:"job_title" `
	Description string    `db:"description" fako:"sentence" `
	CreatedAt   time.Time `db:"created_at" `
	UpdatedAt   time.Time `db:"updated_at" `
	UserID      uuid.UUID `db:"user_id" `
	Must        time.Time `db:"must" `
	Status      bool      `db:"status" `
	Complete    time.Time `db:"completeby" `
	User        *User     `belongs_to:"users"`
}
