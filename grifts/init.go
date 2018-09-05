package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rxx/buffalo_rest_api/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
