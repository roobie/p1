package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
  port := ":10001"
  log.Printf("Starting server at %s", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, ": %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/log_in", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      http.Error(w, "Invalid invocation", http.StatusMethodNotAllowed)
      return
    }
    fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
