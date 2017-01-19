package main

import (
    "os"
    "net/http"
    "strings"
)

// This is injected at build time
var AppVersion = "undefined"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Minimal Web in Go</h1><p>Version: %s </p>", AppVersion)
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", AppVersion)
}

func main() {
	http.HandleFunc("/", handleIndex)
  http.HandleFunc("/version", handleVersion)
	http.ListenAndServe(":80", nil)
}