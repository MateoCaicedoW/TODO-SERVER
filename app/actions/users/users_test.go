package users_test

import (
	"mjm/app"
	"testing"

	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	todo, err := app.New()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	as := &ActionSuite{suite.NewAction(todo)}
	suite.Run(t, as)
}
