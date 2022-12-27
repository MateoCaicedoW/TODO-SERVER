package tasks

import (
	_ "embed"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

//go:embed task_for_id.sql
var taskForIDSQL string

type Task struct {
	ID          uuid.UUID `json:"task_id" db:"task_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      bool      `json:"status" db:"status"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	FullName    string    `json:"user_fullname" db:"user_fullname"`
	Email       string    `json:"user_email" db:"user_email"`
}

func TaskForID(tx *pop.Connection, ID uuid.UUID) (Task, error) {
	task := Task{}

	if err := tx.RawQuery(taskForIDSQL, ID).First(&task); err != nil {
		return task, err
	}

	return task, nil
}
