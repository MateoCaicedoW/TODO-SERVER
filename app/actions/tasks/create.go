package tasks

import (
	"fmt"
	"mjm/app/models"
	"mjm/internal/response"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	task := &models.Task{}
	res := response.Response{}
	if err := c.Bind(task); err != nil {
		return fmt.Errorf("error binding task: %w", err)
	}

	verrs := task.Validate()
	if verrs.HasAny() {

		res.Data = verrs.Errors
		res.Status = 422
		return c.Render(422, r.JSON(res))
	}

	if err := tx.Create(task); err != nil {
		return fmt.Errorf("error creating task: %w", err)
	}

	res.Data = task
	res.Status = 200

	return c.Render(200, r.JSON(res))
}
