package main

import (
	"log"
	"net/http"
)

func main() {
	dir := http.Dir("./static")
	handler := http.FileServer(dir)
	http.Handle("/", handler)

	filesDir := http.Dir("../../samba/myfiles")
	mediaDir := http.Dir("../../samba/myfiles/media")

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(filesDir))))
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir(mediaDir))))

	// Run on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
