# Clio [![Build Status](https://travis-ci.org/grsmv/clio.png?branch=master)](https://travis-ci.org/grsmv/clio)


Clio is a simple DSL for creating web applications in Go programming language with minimal effort. It includes routes for URL-matching, simple template system with support of application-wide and custom layouts and templates for each action, and bunch of useful utilities.


##### Command-line tools

Clio comes with command line interface, which makes developing of Clio-based applications much easier. This is a bench of tools. You can find all needed information about them typing

``` bash
$> clio help
```
in your terminal.

## Creating new application

To create new Clio-based application skeleton, you need just to type few words in terminal:

``` bash
$> clio create APPLICATION_NAME
```

It will create whole application's tree with some configuration assumptions, which you can change anyway. Also this will create basic structure of __app__ folder, which you can modify during your work on application.


## Routes

As an example of complex route usage here you can see the whole stack of REST-routes for curtain purposes (say, controller):

``` go
func ControllerRoutes () {
    Get    ("/books",     Books      )
    Get    ("/books/:id", Book       )
    Post   ("/books",     BooksCreate)
    Put    ("/books/:id", BookUpdate )
    Delete ("/books",     BooksRemove)
    Delete ("/books/:id", BookRemove )
}
```

You can see that Clio supports next HTTP methods: GET, POST, PUT, DELETE.
Every route method takes two arguments - route pattern and function to call if route matches. You can place closure as second argument, for example:

``` go
Get ("/", func () string {
    return "Here's index"
})
```

One and only requirement for functions or closures that are called when pattern match - they should return string.


### Splats

As you noted, route pattern can be given in a form of wild card:

``` go
Get ("/books/:id", Books)
```

So it can match, for example URLs like `/books/12`. This irregular part of pattern is accessible as content of key-value map, returned by `Splat()` function.
Pattern can also hold few irregular parts:

``` go
Get ("/books/:id/download/:format", BooksDownload)

func  BooksDownload () string {
    return "Book with id " + Splat()["id"] + "should be downloaded as: " + Splat()["format"]
}
```

### Query

Somethimes, splats is not enough. Also you want to remember your young years and code in PHP-like style :) It's not a problem with `Query()`. For example, you want to handle all stuff after question mark in this URI: `http://site.com/articles?number=20&offset=100`. To get values for `number` and `offset` you can type just:

```go
func Articles () string {
    return Query()["number"] + " " + Query()["offset"]
}
```


## Views

Clio has application-wide layout system, although you can define custom layout for specific routes. You can give away content without any layouts at all. Let's take a look at few examples:

To define view for specific action, you need to call `Render()` method with path to a template as a first argument:

``` go
func Index () string {
    return Render("index") // this will render `app/views/index.template` file
}
```

To send some data to template, you need to use second [optional] argument for `Render()` method:

``` go
data := Author { Name: "Ernest", Surname: "Hemingway" }
view := Render ("author", data)
```

To specify custom layout for specific route you need to use third [optional] argument for `Render()` method:

```go
view := Render ("books/index", data, Settings { Layout: "hemingway" })
```

##### Partials

Clio reuses standard Go [text/template](http://golang.org/pkg/text/template/) package, but also defines `partial` method, that can be pretty usable right in `.template` files. Let's imagine that you want to include some template right in other template. To do so you need just to call `render` method and use partial's file name as first function's argument:

``` template
{{ partial "specific_header" }}
```

##### Json

Clio provides simple helper to represent your data as JSON without any exceed actions. To output JSON you just need to use `Json ()` function in your controller:

``` go
func JsonBook () string {
  return Json (Book { "Robert Graves", "Good-Bye to All That" })
}
```

it will automatically convert your data to JSON and set appropriate headers.

## WebSockets

You can easily write websocket server-side with Clio. In order to use it, add a new route with `Ws` function:

```go
Ws ("/books/updates")
```

That's all. Now you have channel to connect and to perform all this JavaScript magic.
Also you can send messages to channel right from backend anytime you want using:

```go
WsSend("/books/updates", "51th Shade of Gray is born")
```

You can create as many separated channels as you want and use this feature, for example for role-based updates.


## Headers

For example, you want to give away specific data not as html, but as plain text. To do so you need just to call `SetHeader ()` method. Take a look at this example:

``` go
func BookPlain () string {
    SetHeader("Content-Type", "text/plain")
    return "(...)"
}
```

## Example application

To see how real Clio-based application works, please take a look at https://github.com/grsmv/clio-example.

##### Request details
##### Cookies and sessions?
##### Workers

---

##### Big things to do

- Testing with [gocheck](http://labix.org/gocheck)
- Documentation
- More sophisticated example application
