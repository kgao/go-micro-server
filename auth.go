package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {

      auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)

      if len(auth) != 2 || auth[0] != "Basic" {
          http.Error(w, "bad syntax", http.StatusBadRequest)
          return
      }

      payload, _ := base64.StdEncoding.DecodeString(auth[1])
      pair := strings.SplitN(string(payload), ":", 2)

      if len(pair) != 2 || !Validate(pair[0], pair[1]) {
          http.Error(w, "authorization failed", http.StatusUnauthorized)
          return
      }

      handler(w, r)
  }
}


func Validate(username, password string) bool{
  //TODO: read from encoded db file.
  return username == "username" && password == "password"
}
