package tasks

import (
	"fmt"
	"mjm/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	task := &models.Task{}

	if err := c.Bind(task); err != nil {
		return fmt.Errorf("error binding task: %w", err)
	}

	verrs := task.Validate()
	if verrs.HasAny() {
		return c.Render(422, r.JSON(verrs))
	}

	if err := tx.Create(task); err != nil {
		return fmt.Errorf("error creating task: %w", err)
	}

	return c.Render(200, r.JSON(task))
}
