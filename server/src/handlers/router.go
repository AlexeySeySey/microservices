package handlers

import (
  "net/http"
)

var Router map[string]func(w http.ResponseWriter, r *http.Request);

func InitRouter(){
   Router = map[string]func(w http.ResponseWriter, r *http.Request) {
       		  "/links-with-word/": LinksWithWord,
   }
}
