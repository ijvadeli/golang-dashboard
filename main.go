package main

import (
	"log"
	"net/http"
)

func main() {
	dir := http.Dir("./static")
	handler := http.FileServer(dir)
	http.Handle("/", handler)
	http.Handle("/files", http.FileServer(http.Dir("../../")))

	// Run on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
