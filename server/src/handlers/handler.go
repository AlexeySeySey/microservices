package handlers

import (
  "fmt"
  "net/http"
)

func Foo(w http.ResponseWriter, r *http.Request) {

  fmt.Fprintf(w, "100")
}
