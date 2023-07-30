package Handlers

import (
	"html/template"
	"net/http"
)

func CustomNotFound404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	p := map[string]interface{}{"Erreur": "Page not found", "NumError": http.StatusNotFound, "Message": "The page you’re looking for doesn’t exist."}
	var templates, _ = template.ParseFiles("Templates/error.html")
	templates.Execute(w, p)
}

func CustomNotFound500(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	p := map[string]interface{}{"Erreur": "Internal server error", "NumError": http.StatusInternalServerError, "Message": ""}
	var templates, _ = template.ParseFiles("Templates/error.html")
	templates.Execute(w, p)
}
func CustomSearchNotFound(w http.ResponseWriter, r *http.Request, search string) {
	w.WriteHeader(http.StatusNotFound)
	p := map[string]interface{}{"Erreur": "Search not found", "NumError": search, "Message": "No results found for your search."}
	var templates, _ = template.ParseFiles("Templates/error.html")
	templates.Execute(w, p)
}

func CustomNotFound400NoPostRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	p := map[string]interface{}{"Erreur": "No access", "NumError": http.StatusMethodNotAllowed, "Message": "You cannot access this endpoint."}
	var templates, _ = template.ParseFiles("Templates/error.html")
	templates.Execute(w, p)
}
