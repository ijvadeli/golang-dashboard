package main

import (
	"log"
	"net/http"
)

func main() {
	dir := http.Dir("./static")
	handler := http.FileServer(dir)
	http.Handle("/", handler)
	getSamba()

	// Run on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getSamba() {
	dir := http.Dir("../../../")
	handler := http.FileServer(dir)
	http.Handle("/files", handler)
}
