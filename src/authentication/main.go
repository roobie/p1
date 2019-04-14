package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
  "github.com/roobie/p1core"
)

func main() {
	port := ":10001"
	log.Printf("Starting server at %s", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %q\r\n", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "Config: %+v", p1core.GetConfig())
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
