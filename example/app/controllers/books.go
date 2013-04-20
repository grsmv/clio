package controllers

import (
    "../../../../clio"
)

func Index () string {
    return "<a href=\"books\">Books index</a> <br />" +
           "<a href=\"books/12\">Certain book</a>"
}

func Books () string {
    clio.SetHeader("Content-Type", "text/plain")
    return "Books list"
}

func Book () string {
    return "Book id #" + clio.Splat()[0] + "<br />" +
           "url: "     + clio.Context().Request.URL.String() + "<br />" +
           "params: "  + clio.Params()["a"]
}

func BooksCreate () string {
    return "Create new book"
}

func BookUpdate () string {
    return "Ok, let's update book with id #" + clio.Splat()[0]
}

func BooksRemove () string {
    return "Remove all books"
}

func BookRemove () string {
    return "Remove book with id #" + clio.Splat()[0]
}
