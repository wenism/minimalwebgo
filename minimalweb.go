package main

import (
	"fmt"
  "net/http"
)

// This is injected at build time
var AppVersion = "undefined"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Minimal Web in Go</h1><p>Version: %s </p>\n", AppVersion)
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", AppVersion)
}

func handleLiveness(w http.ResponseWriter, r *http.Request) {
  // running?
	fmt.Fprintf(w, "OK\n")
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	// ready to serve?
  fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/", handleIndex)
  http.HandleFunc("/version", handleVersion)
  http.HandleFunc("/liveness", handleLiveness)
  http.HandleFunc("/readiness", handleReadiness)
	http.ListenAndServe(":80", nil)
}