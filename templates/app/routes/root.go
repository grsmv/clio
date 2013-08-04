package routes

import (
  . "github.com/pallada/clio/core"
  "../../app/controllers"
)

func Root () {
  Get ("/", controllers.Root)
}
