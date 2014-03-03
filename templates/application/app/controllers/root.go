package controllers

import (
	. "github.com/grsmv/clio/core"
)

func Root() string {
	return Render("root")
}
