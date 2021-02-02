package grifts

import (
	"github.com/EdaHikaru/coke/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
