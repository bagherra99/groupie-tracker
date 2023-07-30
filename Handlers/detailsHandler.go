package Handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

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
