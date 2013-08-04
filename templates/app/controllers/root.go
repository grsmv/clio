package controllers

import (
  . "github.com/pallada/clio/core"
)

func Root () string {
  return Render ("root")
}
