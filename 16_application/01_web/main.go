package main

import (
	"log"
	"net/http"
)

func main() {
	// Register handlers to DefaultServeMux
	http.HandleFunc("/", frontPageHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// It initializes the server at port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
