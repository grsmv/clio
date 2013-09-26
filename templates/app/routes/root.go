package routes

import (
  . "github.com/cliohq/clio/core"
  "../../app/controllers"
)

func Root () {
  Get ("/", controllers.Root)
}
