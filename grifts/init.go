package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.ibm.com/bluebox/todo2/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
