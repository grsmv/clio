package main

import (
  . "github.com/pallada/clio/core"
  "./app/routes"
)

func AppSettings () map[string]interface {} {
  return map[string]interface {} {
    "port": 4567 }
}

func main () {
  routes.Root ()
  Run (AppSettings()["port"].(int))
}
