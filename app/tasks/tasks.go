package grifts

import (
	"log"
	"mjm/app"

	"github.com/gobuffalo/buffalo"
)

// Init the tasks with some common tasks that come from
// grift
func init() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	buffalo.Grifts(app)
}
