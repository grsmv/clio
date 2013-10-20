package main

import (
  . "github.com/cliohq/clio/core"
  "./app/routes"
)

// Mandatory function. Entering point of your application.
func main () {

    // Registering application's routes package
    routes.Register ()

    // Running application. Note that one and only parameter
    // for this function is a map of application's settings.
    Run (
        map[string]interface {} {

            // http port on which application will be available
            "port": 4567,

            // Managing assets flag. Switch it to false if assets will
            // be managed by web server, for example NGINX
            "manage-assets": true,

            // This JavaScript framework will be used with Clio.
            // Please find more at 'Generators' section of Clio's README.
            "ui-framework": "backbone",

            // Directories to watch with `clio watch`
            "directories-to-watch": []string {
                "./app/controllers/",
                "./app/routes/",
                "./app/helpers",
                "./app/views",
            },
        },
    )
}
