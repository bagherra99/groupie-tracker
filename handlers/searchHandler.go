package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// var AllInfo = ReceiveDataFromJson()

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère le terme de recherche saisi par l'utilisateur
	input := r.URL.Query().Get("search")
	fmt.Println(input)
	searched := SearchArtist(input)
	fmt.Println(searched)

	tmpl, err := template.ParseFiles("public/search.html")
	if err != nil {
		http.Error(w, "error while parsing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, searched)
}
