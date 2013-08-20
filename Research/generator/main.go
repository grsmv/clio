package main

import (
  "./a"
  "./b"
  "./util"
)

func main () {

  generators := []func() {
    a.Register ,
    b.Register ,
  }

  Generate(generators)
}

func Generate (functions []func()) {
  util.Generate (functions)
}
