package users

import (
	"fmt"
	"mjm/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Create(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return fmt.Errorf("error binding user: %w", err)
	}

	verrs := user.Validate(tx)
	if verrs.HasAny() {
		return c.Render(422, r.JSON(verrs))
	}

	if err := tx.Create(user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return c.Render(200, r.JSON(user))
}
