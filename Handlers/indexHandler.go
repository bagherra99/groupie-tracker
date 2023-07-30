package Handlers

import (
	ResponseData "groupie/ResponseData"
	"groupie/Tools"
	"html/template"
	"net/http"
)

func IndexDataHandler(w http.ResponseWriter, r *http.Request) {
	var Tab []string
	var TabMember []string
	for _, item := range AllInfo {
		for _, v := range item.Locations {
			if !Tools.IsStringInArray(Tab, v) {
				Tab = append(Tab, v)
			}
		}
		for _, l := range item.Members {
			if !Tools.IsStringInMember(TabMember, l) {
				TabMember = append(TabMember, l)
			}
		}
	}

	data := struct {
		Allartist []ResponseData.AllArtists
		Locat     []string
		Memb      []string
	}{
		Allartist: AllInfo,
		Locat:     Tab,
		Memb:      TabMember,
	}

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		http.Error(w, "error while parsing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "error while executing", http.StatusInternalServerError)
	}
}
