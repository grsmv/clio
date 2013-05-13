package main

import (
    . "../../clio"
    "./app/routes"
)

func main() {
    routes.BooksRoutes()
    Run (4567)
}
