package users

import (
	"mjm/app/models"
	"mjm/internal/response"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	users := &models.Users{}

	response := response.Response{}

	if err := tx.All(users); err != nil {
		return err
	}

	response.Data = users
	response.Status = 200
	return c.Render(200, r.JSON(response))
}
