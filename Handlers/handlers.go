package Handlers

import (
	ResponseData "groupie/ResponseData"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var AllInfo = ResponseData.ReceiveDataFromJson()

func IndexDataHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		http.Error(w, "error while parsing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, AllInfo)
	if err != nil {
		http.Error(w, "error while executing", http.StatusInternalServerError)
	}
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	tab := strings.Split(r.URL.Path, "/")
	id := tab[len(tab)-1]
	val, errs := strconv.Atoi(id)

	data := AllInfo[val-1]
	if len(tab) != 3 || id == "" || tab[len(tab)-2] != "details" || errs != nil {
		CustomNotFound404(w)
	} else {
		if val < 1 || val > 52 || id == "" {
			CustomNotFound404(w)
			return
		}
		var templates, err = template.ParseFiles("Templates/details.html")
		if err != nil {
			CustomNotFound500(w)
			return
		}
		templates.Execute(w, data)
	}

}
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the search query parameters from the URL
	name := r.FormValue("name")
	// Perform the search based on the name
	searchResults := SearchArtist(name)

	var templates, err = template.ParseFiles("Templates/search.html")
	if err != nil {
		CustomNotFound500(w)
		return
	}

	templates.Execute(w, searchResults)
}

// Sample function to search artists based on name
func SearchArtist(name string) []ResponseData.AllArtists {
	var results []ResponseData.AllArtists

	for _, artist := range AllInfo {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(name)) {
			results = append(results, artist)
		}
	}

	return results
}
