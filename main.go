package main

import (
	"log"
	"net/http"
)

func main() {
	// Website directory
	dir := http.Dir("./static")

	// Handling the website directory as fileserver
	handler := http.FileServer(dir)

	// Set home as default
	http.Handle("/", handler)

	// Run on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
