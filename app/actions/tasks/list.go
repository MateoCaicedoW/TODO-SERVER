package tasks

import (
	"fmt"
	"mjm/internal/tasks"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	tasks, err := tasks.All(tx)
	if err != nil {
		return fmt.Errorf("error listing tasks: %w", err)
	}

	return c.Render(200, r.JSON(tasks))
}
