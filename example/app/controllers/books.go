package controllers

import (
    . "github.com/grsmv/clio/core"
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

type Message struct {
  Name string
  Body string
}


func Index () string {
    return Render("index")
}


func Books () string {
    data := Author { 
        Name: "Ernest",
        Surname: "Hemingway",
        Description: "(July 21, 1899 &mdash; July 2, 1961) American author and journalist",
        Works: []Work {
            Work {"The Sun Also Rises (1926)"},
            Work {"A Farewell to Arms (1929)"},
            Work {"For Whom the Bell Tolls (1940)"},
            Work {"The Old Man and the Sea (1951)"} }}

    rend := Render("books/index", data /* , Settings{ Layout: "none" } */)
    return rend
}


func Book () string {
    /* SetHeader("Content-Type", "text/plain") */
    return "Book id #" + Splat()[0] + "<br />" +
           "url: "     + Context().Request.URL.String() + "<br />" +
           "params: "  + Params()["a"]
}


func BookJ () string {
  return Json ([]Message {
    Message { "Alice", "Hello" },
    Message { "Alex",  "Bye" }})
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
