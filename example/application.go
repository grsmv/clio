package main

import (
    "../../clio"
    "./app/routes"
)

func main() {
    routes.BooksRoutes()
    clio.Run (4567)
}
