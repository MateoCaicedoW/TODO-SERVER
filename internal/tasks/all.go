package tasks

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

//go:embed all_query.sql
var allSQL string

type Task struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Completed   time.Time `json:"complete" db:"completed"`
	Status      bool      `json:"status" db:"status"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	FullName    string    `json:"user_fullname" db:"user_fullname"`
	Email       string    `json:"user_email" db:"user_email"`
	UserRole    string    `json:"user_role" db:"user_role"`
}

func All(tx *pop.Connection) ([]Task, error) {
	tasks := []Task{}

	err := tx.RawQuery(allSQL).All(&tasks)
	if err != nil {
		return nil, fmt.Errorf("error listing tasks: %w", err)
	}

	return tasks, err
}
