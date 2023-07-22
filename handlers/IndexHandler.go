package handlers

import (
	"html/template"
	"main/utils"
	"net/http"
)

var AllInfo = utils.ReceiveDataFromJson()


func IndexDataHandler(w http.ResponseWriter, r *http.Request) {
	// AllInfo := ReceiveDataFromJson()
	// Génère une liste de noms d'artistes pour les suggestions
	// Créer une liste d'options pour la barre de recherche (noms d'artistes et membres)
	// var searchOptions []string
	// for _, artist := range AllInfo {
	//     searchOptions = append(searchOptions, artist.Name)
	//     searchOptions = append(searchOptions, artist.Members...)
	// }
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "error while parsing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, AllInfo)
}
