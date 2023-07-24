package handlers

import (
	"html/template"
	"main/utils"
	"net/http"
	tools "main/tools"
	model "main/model"
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
	input := r.URL.Query().Get("search")
	var Tab []string
	var TabMember []string
	for _, item := range AllInfo {
		for _, v := range item.Locations {
			if !tools.IsStringInArray(Tab, v) {
				Tab = append(Tab, v)
			}
		}
		for _, l := range item.Members {
			if !tools.IsStringInMember(TabMember, l) {
				TabMember = append(TabMember, l)
			}
		}
	}
	searched := SearchArtist(input)

	data := struct {
		Allartist []model.AllArtists
		// Input     string
		Locat     []string
		Memb      []string
		Searched    []model.AllArtists
	}{
		Allartist: AllInfo,
		// Input:     input,
		Locat:     Tab,
		Memb:      TabMember,
		Searched:    searched,
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "error while parsing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
}
