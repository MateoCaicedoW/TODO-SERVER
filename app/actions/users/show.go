package users

import (
	"mjm/internal/response"
	"mjm/internal/users"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func Show(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	res := response.Response{}

	user, err := users.DetailsForID(tx, uuid.FromStringOrNil(c.Param("id")))
	if err != nil {
		return c.Render(404, r.JSON(map[string]string{"error": "user not found"}))
	}

	res.Data = user
	res.Status = http.StatusOK

	return c.Render(http.StatusOK, r.JSON(res))
}
