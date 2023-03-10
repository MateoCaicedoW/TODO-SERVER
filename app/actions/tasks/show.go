package tasks

import (
	"fmt"
	"mjm/internal/response"
	"mjm/internal/tasks"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func Show(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	res := response.Response{}
	task, err := tasks.TaskForID(tx, uuid.FromStringOrNil(c.Param("id")))
	if err != nil {
		return fmt.Errorf("error finding task: %w", err)
	}

	res.Data = task
	res.Status = http.StatusOK
	return c.Render(http.StatusOK, r.JSON(res))

}
