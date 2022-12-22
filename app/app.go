package app

import (
	"fmt"
	"mjm/internal/environment"
	"strconv"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop/v6"
)

// App creates a new application with default settings and reading
// GO_ENV. It calls setRoutes to setup the routes for the app that's being
// created before returning it

func New() (*buffalo.App, error) {
	app := buffalo.New(buffalo.Options{
		Env:           environment.Current(),
		SessionName:   environment.SessionName,
		WorkerOff:     true,
		CompressFiles: true,
	})

	err := configure(app)
	if err != nil {
		return nil, err
	}

	setRoutes(app)

	return app, nil
}

func configure(app *buffalo.App) error {
	binding.RegisterTimeFormats("01-02-2006")
	binding.RegisterTimeFormats("2006-01-02T03:04:05")
	binding.RegisterCustomDecoder(func(vals []string) (interface{}, error) {
		// don't try to parse empty time values, it will raise an error
		if len(vals) == 0 || vals[0] == "" {
			return 0.0, nil
		}

		if val, err := strconv.ParseFloat(vals[0], 64); err == nil {
			return val, nil
		}

		return 0.0, nil
	}, []interface{}{0.0}, []interface{}{})

	binding.RegisterCustomDecoder(func(vals []string) (interface{}, error) {
		if len(vals) == 0 || vals[0] == "" {
			return 0, nil

		}
		if val, err := strconv.Atoi(vals[0]); err == nil {
			return val, nil
		}

		return 0, nil

	}, []interface{}{int(0)}, nil)

	//custom decoder for time.time written as string
	binding.RegisterCustomDecoder(func(vals []string) (interface{}, error) {
		if len(vals) == 0 || vals[0] == "" {
			return time.Time{}, nil
		}
		fmt.Println("vals[0]: ", vals[0])

		if val, err := time.Parse("2006-01-02T15:04:05", vals[0]); err == nil {
			return val, nil
		}

		return time.Time{}, nil
	}, []interface{}{time.Time{}}, nil)

	pop.PaginatorPerPageDefault = 20

	return nil
}
