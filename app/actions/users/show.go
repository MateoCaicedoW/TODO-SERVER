package users

import (
	"mjm/internal/users"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func Show(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user, err := users.DetailsForID(tx, uuid.FromStringOrNil(c.Param("id")))
	if err != nil {
		return c.Render(404, r.JSON(map[string]string{"error": "user not found"}))
	}

	return c.Render(200, r.JSON(user))
}
