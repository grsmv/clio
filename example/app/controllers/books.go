package controllers

import (
    . "../../../../clio"
)

type Work struct {
    Name string
}

type Author struct {
    Name        string
    Surname     string
    Description string
    Works       []Work
}


func Index () string {
    return "<a href=\"books\">Books index</a> <br />" +
           "<a href=\"books/12\">Certain book</a>"
}


func Books () string {
    data := Author { 
        Name: "Ernest",
        Surname: "Hemingway",
        Description: "Ernest Miller Hemingway (July 21, 1899 &mdash; July 2, 1961) was an American author" +
                     " and journalist. His economical and understated style had a strong influence" +
                     " on 20th-century fiction, while his life of adventure and his public image" +
                     " influenced later generations. ",
        Works: []Work {
            Work {"The Sun Also Rises (1926)"},
            Work {"A Farewell to Arms (1929)"},
            Work {"For Whom the Bell Tolls (1940)"},
            Work {"The Old Man and the Sea (1951)"} }}

    // SetHeader("Content-Type", "text/plain")
    rend := Render("books/index", data)
    return rend
}


func Book () string {
    return "Book id #" + Splat()[0] + "<br />" +
           "url: "     + Context().Request.URL.String() + "<br />" +
           "params: "  + Params()["a"]
}


func BooksCreate () string {
    return "Create new book"
}


func BookUpdate () string {
    return "Ok, let's update book with id #" + Splat()[0]
}


func BooksRemove () string {
    return "Remove all books"
}


func BookRemove () string {
    return "Remove book with id #" + Splat()[0]
}