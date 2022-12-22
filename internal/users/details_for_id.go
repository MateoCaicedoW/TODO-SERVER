package users

import (
	"fmt"
	"mjm/app/models"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func DetailsForID(tx *pop.Connection, ID uuid.UUID) (models.User, error) {
	user := models.User{}

	err := tx.Find(&user, ID)
	if err != nil {
		return user, fmt.Errorf("error finding user: %w", err)
	}

	return user, nil

}
