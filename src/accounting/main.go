package main

import (
	"fmt"
	"github.com/roobie/p1core"
	"html"
	"log"
	"net/http"
)

func main() {
	port := ":10003"
	log.Printf("Starting server at %s", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %q\r\n", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "Config: %+v", p1core.GetConfig())
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
