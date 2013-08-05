package main

import (
  . "github.com/pallada/clio/core"
  "./app/routes"
)

func AppSettings () map[string]interface {} {
  return map[string]interface {} {
    "port":           4567,
    "foreign-port":   4568,
    "manage-assets":  true,
    "ui-framework":   "backbone"
  }
}

func main () {
  routes.Root ()
  Run (AppSettings()["port"].(int))
}
