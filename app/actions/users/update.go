package users

import (
	"fmt"
	"mjm/internal/response"
	"mjm/internal/users"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	res := response.Response{}
	user, err := users.DetailsForID(tx, uuid.FromStringOrNil(c.Param("id")))
	if err != nil {
		return c.Render(404, r.JSON(map[string]string{"error": "user not found"}))
	}

	if err := c.Bind(&user); err != nil {
		return fmt.Errorf("error binding user: %w", err)
	}

	verrs := user.Validate(tx)
	if verrs.HasAny() {
		res.Data = verrs.Errors
		res.Status = http.StatusUnprocessableEntity
		return c.Render(http.StatusUnprocessableEntity, r.JSON(res))
	}

	if err := tx.Update(&user); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": "error updating user"}))
	}

	res.Data = user
	res.Status = http.StatusOK

	return c.Render(http.StatusOK, r.JSON(res))
}
