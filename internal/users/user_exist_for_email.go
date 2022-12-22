package users

import (
	_ "embed"
	"fmt"

	"github.com/gobuffalo/pop/v6"
)

//go:embed user_exist_for_email.sql
var userExistForEmailSQL string

func UserExistForEmail(tx *pop.Connection, email string) (bool, error) {
	var exist bool

	err := tx.RawQuery(userExistForEmailSQL, email).First(&exist)
	if err != nil {
		return exist, fmt.Errorf("error finding user: %w", err)
	}

	return exist, nil
}
