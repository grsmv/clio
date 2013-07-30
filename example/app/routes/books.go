package routes

import (
    . "github.com/grsmv/clio/core"
    "../../app/controllers"
)

func BooksRoutes () {
    Get("/",              controllers.Index)       // root
    Post ("/books",       controllers.BooksCreate)
    Get ("/books",        controllers.Books)
    Get ("/books/*",      controllers.Book)
    Get ("/books/*/json", controllers.BookJ)
    Put ("/books/*",      controllers.BookUpdate)
    Delete ("/books",     controllers.BooksRemove)
    Delete ("/books/*",   controllers.BookRemove)
}
