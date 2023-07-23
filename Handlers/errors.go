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
