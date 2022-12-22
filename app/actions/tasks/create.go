package tasks

import (
	"fmt"
	"mjm/app/models"

	"github.com/gobuffalo/buffalo"
)

func Create(c buffalo.Context) error {
	// tx:= c.Value("tx").(*pop.Connection)
	task := &models.Task{}

	if err := c.Bind(task); err != nil {
		return fmt.Errorf("error binding task: %w", err)
	}

	fmt.Println("task: ", task)
	return nil
}
