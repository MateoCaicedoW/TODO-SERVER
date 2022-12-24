package users

import (
	"fmt"
	"mjm/app/models"
	"mjm/internal/response"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Create(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	user := &models.User{}
	res := response.Response{}

	if err := c.Bind(user); err != nil {
		return fmt.Errorf("error binding user: %w", err)
	}

	verrs := user.Validate(tx)
	if verrs.HasAny() {

		res.Data = verrs.Errors
		res.Status = 422

		return c.Render(422, r.JSON(res))
	}

	if err := tx.Create(user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	res.Data = user
	res.Status = 200

	return c.Render(200, r.JSON(res))
}
