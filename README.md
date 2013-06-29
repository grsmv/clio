# Clio


Clio is a simple DSL for creating web applications in Go programming language with minimal effort, built to use with [carcass](https://github.com/grsmv/carcass). It includes routes for URL-matching, simple template system with support of application-wide and custom layouts and templates for each action, and bunch of useful utilities. 


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

As you noted, route pattern can be given in a form of wildcard:

``` go
Get ("/books/*", Books)
```

So it can match, for example such url as `/books/12`. This irregular parts of pattern is accessible as content of slice, returned by `Splat()` function. 
Pattern also can hold few irregular parts:

``` go
Get ("/books/*/download/*", BooksDownload)

func  BooksDownload () string {
    return "Book with id " + Splat()[0] + "should be downloaded as: " + Splat()[1]
}
```


##### Views

##### Headers

##### Request details

##### Cookies and sessions?

##### Example application

---

##### Big things to do

- Testing with [gocheck](http://labix.org/gocheck)
- Documentation
- More sophisticated example application
