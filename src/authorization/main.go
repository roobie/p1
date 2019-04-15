package main

import (
	"fmt"
	"github.com/roobie/p1core"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	port := ":10002"
	log.Printf("Starting server at %s", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %q\r\n\r\n",
			html.EscapeString(r.URL.Path))
		config := p1core.GetConfig()
		auth := config.Mappings.Authentication
		fmt.Fprintf(w, "Config: %+v\r\n\r\n", config)
		// targetUrl := fmt.Sprintf("http://%s:%s/log_in",
		// 	auth.Hostname, auth.Port)
		targetUrl := fmt.Sprintf("http://%s/log_in",
			auth.Hostname)

		fmt.Fprintf(w, "Url: %s\r\n\r\n", targetUrl)
		resp, err := http.PostForm(
			targetUrl,
			url.Values{"key": {"Value"}, "id": {"123"}})
		if err != nil {
			fmt.Fprintf(w, "Err: %+v\r\n\r\n", err)
		} else {
			fmt.Fprintf(w, "Res: %+v\r\n\r\n", resp)
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(w, "Err: %+v\r\n\r\n", err)
			} else {
				fmt.Fprintf(w, "Body: %+v\r\n\r\n", string(body))
			}
		}
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
