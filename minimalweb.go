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
  w.Write([]byte("Alive!\n"))
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Ready!\n"))
}

func main() {
	http.HandleFunc("/", handleIndex)
  http.HandleFunc("/version", handleVersion)
  http.HandleFunc("/liveness", handleLiveness)
  http.HandleFunc("/readiness", handleReadiness)
	http.ListenAndServe(":2888", nil)
}