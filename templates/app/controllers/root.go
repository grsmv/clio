package controllers

import (
  . "github.com/cliohq/clio/core"
)

func Root () string {
  return Render ("root")
}
