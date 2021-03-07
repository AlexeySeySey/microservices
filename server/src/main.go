package main

import (
    "log"
    "net/http"
    "./handlers"
)

func init(){
  handlers.InitRouter();
}

func main() {

    var routes []string;

    for route := range handlers.Router {
        routes = append(routes, route);
    }

    for _, route := range routes {
        http.HandleFunc(route, handlers.PreHandleRequest);
    }

    log.Fatal(http.ListenAndServe(":8081", nil));
}
