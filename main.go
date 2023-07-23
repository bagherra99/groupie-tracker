package main

import (
	"fmt"
	Handlers "groupie/Handlers"
	"net/http"
)

var fs = http.FileServer(http.Dir("static"))

func main() {

	http.HandleFunc("/", Handlers.IndexDataHandler)
	http.HandleFunc("/details/", Handlers.DetailsHandler)
	http.HandleFunc("/search/", Handlers.SearchHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("http://localhost:8080\nServeur démarré sur le port :8080")
	http.ListenAndServe(":8080", nil)

}
