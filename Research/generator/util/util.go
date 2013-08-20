package util

func Generate (functions []func()) {
  for _, function := range functions {
    function ()
  }
}
