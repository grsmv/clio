package main

import (
    "./store"
    /* "github.com/davecgh/go-spew/spew" */
)

func main () {
    s, _ := store.Open ("names.csv")

    /* s.Remove ("last_name", "Zibert") */

    // s.Add (map[string]string {
    //   "first_letter": "z",
    //   "first_name": "Zahria",
    //   "last_name": "Johnes" })

    s.Update("_id", "as9", map[string]string {
      "_id": "as9",
      "first_letter": "a",
      "first_name": "Adam",
      "last_name": "Zakaz" })

    /* results := s.WhereLike("last_name", "joh") */

    /* spew.Dump(results) */
}
