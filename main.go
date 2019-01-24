package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandler)
	addr := "0.0.0.0:8910"
	fmt.Printf("Listening on http://%s\n", addr)
	http.ListenAndServe(addr, nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, "got: %q\n", html.EscapeString(r.URL.Path))
}
