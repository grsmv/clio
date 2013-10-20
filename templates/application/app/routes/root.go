package routes

import (
    . "github.com/cliohq/clio/core"
    "../../app/controllers"
)

func init () {
    Get ("/", controllers.Root)
}
