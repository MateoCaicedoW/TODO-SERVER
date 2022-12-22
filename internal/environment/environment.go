package environment

import (
	"os"
	"strings"

	"github.com/gobuffalo/envy"
)

// These is the place we put constants that will be used across the app
// to start these 3 constants come very handy to know which environment
// are you in.
const (
	Development = "development"
	Test        = "test"
	Production  = "production"

	// Other constants in your app
	ApplicationName = "todo"
	SessionName     = "_todo_session"
)

// Returns current environment from GO_ENV
// if empty returns Development.
func Current() string {
	if env := os.Getenv("GO_ENV"); env != "" {
		return env
	}

	return Development
}

// BaseURL loads the BASE_URL environment variable and prepends
// https in case it does not have the scheme.
func BaseURL() string {
	baseHost := envy.Get("BASE_URL", "http://localhost:3000")
	if strings.HasPrefix(baseHost, "http") {
		return baseHost
	}

	return "https://" + baseHost
}
