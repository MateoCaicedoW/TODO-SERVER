package app

import (
	"net/http"

	"mjm/app/actions/tasks"
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
	usrs.PUT("/{id}", users.Update)
	usrs.GET("/{id}", users.Show)
	usrs.DELETE("/{id}", users.Destroy)

	tsks := root.Group("/tasks")
	tsks.POST("/", tasks.Create)
	tsks.GET("/", tasks.List)
	tsks.GET("/{id}", tasks.Show)
	tsks.PUT("/{id}", tasks.Update)
	tsks.DELETE("/{id}", tasks.Destroy)
	tsks.PUT("/{id}/complete", tasks.Complete)

	root.ServeFiles("/", http.FS(public.FS()))
}
