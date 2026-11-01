package main

import (
	"log"
	"net/http"
)

func main() {
	dir := http.Dir("./static")

	handler := http.FileServer(dir)

	http.Handle("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
