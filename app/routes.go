package app

import (
	"net/http"

	"mjm/app/actions/users"
	"mjm/app/middleware"
	"mjm/public"

	"github.com/gobuffalo/buffalo"
)

func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	// root.Use(middleware.CSRF)

	usrs := root.Group("/users")
	usrs.POST("/", users.Create)
	usrs.GET("/", users.List)

	root.ServeFiles("/", http.FS(public.FS()))
}
