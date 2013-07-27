# Clio


Clio is a simple DSL for creating web applications in Go programming language with minimal effort, built to use with [carcass](https://github.com/grsmv/carcass). It includes routes for URL-matching, simple template system with support of application-wide and custom layouts and templates for each action, and bunch of useful utilities.


##### Command-line tools


### Creating new application

To create new Clio-based application's skeleton, you need just to type few words in terminal:

``` bash
$> clio create APPLICATION_NAME
```

It will creates whole application's tree with some configuration assumptions, which you can change anyway. Also this will create basic structure of __app__ folder, which you can modify during your work on application.


### Routes

As an example of complex route usage here you can see the whole stack of REST-routes for curtain purposes (say, controller):

``` go
func ControllerRoutes () {
    Get    ("/books",   Books      )
    Get    ("/books/*", Book       )
    Post   ("/books",   BooksCreate)
    Put    ("/books/*", BookUpdate )
    Delete ("/books",   BooksRemove)
    Delete ("/books/*", BookRemove )
}
```

You can see, that Clio supports next HTTP methods: GET, POST, PUT, DELETE. 
Every route method took to arguments - route pattern and function to call if route matches. You can place closure as second argument, for example:

``` go
Get ("/", func () string {
    return "Here's index"
})
```

One and only requirement for functions or closures that calls when pattern match - they should return string. 


### Splats

As you noted, route pattern can be given in a form of wild card:

``` go
Get ("/books/*", Books)
```

So it can match, for example such URL as `/books/12`. This irregular parts of pattern is accessible as content of slice, returned by `Splat()` function.
Pattern also can hold few irregular parts:

``` go
Get ("/books/*/download/*", BooksDownload)

func  BooksDownload () string {
    return "Book with id " + Splat()[0] + "should be downloaded as: " + Splat()[1]
}
```


### Views

Clio has application-wide layout system, also you can define custom layout for specific routes. Also you can give away content without any layouts at all. Let's take a look at few examples:

To define view for specific view, you need to call `Render()` method with path to a template as a first argument:

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

Clio reuses standard Go [text/template](http://golang.org/pkg/text/template/) package, but also defines `partial` method, that can be pretty usable right in `.template` files. Lets imagine that you want to include some template right in other template. To do so you need just to call `render` method and use partial's file name as first function's argument:

``` template
{{ partial "specific_header" }}
```

### Headers

For example, you want to give away specific data not as html, but as plain text. To do so you need just to call `SetHeader ()` method. take a look at this example:

``` go
func BookPlain () string {
    SetHeader("Content-Type", "text/plain")
    return "(...)"
}
```

### Example application

To look and feel how real Clio-based application works, please take a look at __example__ folder in the root.

##### Request details
##### Cookies and sessions?

---

##### Big things to do

- Testing with [gocheck](http://labix.org/gocheck)
- Documentation
- More sophisticated example application
