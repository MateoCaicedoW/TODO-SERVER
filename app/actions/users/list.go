package users

import (
	"mjm/app/models"

	"github.com/gobuffalo/buffalo"
)

func List(c buffalo.Context) error {

	users := models.User{
		FirstName:    "John",
		LastName:     "Doe",
		DNI:          "123",
		EmailAddress: "a",
	}

	return c.Render(200, r.JSON(users))
}
