package models

import (
	"github.com/gofrs/uuid"
)

type Users []User
type User struct {
	ID           uuid.UUID `db:"id" `
	FirstName    string    `db:"first_name" `
	LastName     string    `db:"last_name" `
	DNI          string    `db:"dni" `
	EmailAddress string    `db:"email_address" `
	PhoneNumber  string    `db:"phone_number" `
}
