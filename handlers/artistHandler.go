package handlers

import (
	"html/template"
	"main/model"
	"main/tools"
	"net/http"
	"strconv"
)



func ArtistDataHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère l'ID de l'artiste à partir de l'URL
	artistID := r.URL.Path[len("/artist/"):]
	var id int
	var err error
	if tools.IsNumeric(artistID) {
		id, err = strconv.Atoi(artistID)
		if err != nil {
			http.Error(w, "Invalid artist ID", http.StatusBadRequest)
			return
		}
	}else{
		http.Error(w, "bindal chiffre sans signe", http.StatusNotFound)
		return
	}

	// Récupère les informations de l'artiste à partir de la liste des artistes
	// AllInfo := ReceiveDataFromJson()
	var artistInfo model.AllArtists
	for _, artist := range AllInfo {
		if artist.ID == id {
			artistInfo = artist
			break
		}
	}

	if artistInfo.ID < 1 || artistInfo.ID > 52 {
		http.Error(w, "artist not found", http.StatusNotFound)
		return
	}

	// Charge le modèle "artist.html" et affiche les informations de l'artiste
	tmpl, err := template.ParseFiles("public/artist.html")
	if err != nil {
		http.Error(w, "Error while parsing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artistInfo)
	if err != nil {
		http.Error(w, "Error while executing", http.StatusInternalServerError)
	}
}
