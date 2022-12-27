package tasks

import (
	"mjm/app/models"
	"mjm/internal/response"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Destroy(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	res := response.Response{}

	task := models.Task{}

	if err := tx.Find(&task, c.Param("id")); err != nil {
		return err
	}

	if err := tx.Destroy(&task); err != nil {
		return err
	}

	res.Status = http.StatusOK
	return c.Render(http.StatusOK, r.JSON(res))

}
