package handlers

import (
  "net/http"
  "../config"
)

func PreHandleRequest(w http.ResponseWriter, r *http.Request) {
   if r.Host != config.AllowedOrigin || r.URL.Query()["key"][0] != config.SecretKey {
      return;
   }
   Router[r.URL.Path](w, r);
}
