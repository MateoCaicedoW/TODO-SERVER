package users

import (
	"mjm/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	users := &models.Users{}

	if err := tx.All(users); err != nil {
		return err
	}

	return c.Render(200, r.JSON(users))
}
