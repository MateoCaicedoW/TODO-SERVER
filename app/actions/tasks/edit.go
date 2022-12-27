package tasks

import (
	"mjm/app/models"
	"mjm/internal/response"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	task := models.Task{}
	response := response.Response{}
	if err := tx.Find(&task, c.Param("id")); err != nil {
		return err
	}

	if err := c.Bind(&task); err != nil {
		return err
	}

	if err := tx.Update(&task); err != nil {
		return err
	}

	response.Data = task
	response.Status = http.StatusOK

	return c.Render(http.StatusOK, r.JSON(response))
}


