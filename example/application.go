package main

import (
    . "github.com/grsmv/clio/core"
    "./app/routes"
)

func main() {
    routes.BooksRoutes()
    Run (4567)
}
