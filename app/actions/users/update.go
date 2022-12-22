package users

import (
	"fmt"
	"mjm/internal/users"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user, err := users.DetailsForID(tx, uuid.FromStringOrNil(c.Param("id")))
	if err != nil {
		return c.Render(404, r.JSON(map[string]string{"error": "user not found"}))
	}

	if err := c.Bind(&user); err != nil {
		return fmt.Errorf("error binding user: %w", err)
	}

	verrs := user.Validate(tx)
	if verrs.HasAny() {
		return c.Render(422, r.JSON(verrs))
	}

	if err := tx.Update(&user); err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": "error updating user"}))
	}

	return c.Render(200, r.JSON(user))
}
