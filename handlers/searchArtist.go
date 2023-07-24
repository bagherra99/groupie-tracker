package handlers

import (
	"fmt"
	"main/model"
	"main/tools"

	// "main/handlers"
	"strings"
)

var filteredArtists []model.AllArtists

func SearchArtist(s string) []model.AllArtists {

	tabInput := strings.Split(s, "==>")
	// fmt.Println(tabInput)
	if len(tabInput) == 2 {
		query := tabInput[0]
		category := tabInput[1]
		fmt.Println(query)
		fmt.Println(category)

		// Si un terme de recherche est saisi, filtre les artistes correspondants
		if s != "" {
			switch strings.ToLower(category) {
			case "artist/band":
				// Filtrer les artistes dont le nom correspond à la recherche
				var searchedArtist []model.AllArtists
				for _, artist := range AllInfo {
					if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
						searchedArtist = append(searchedArtist, artist)
						// fmt.Println("================================")
						// fmt.Println(searchedArtist)
					}
				}
				// Utilise les artistes filtrés pour l'affichage
				filteredArtists = searchedArtist
				// fmt.Println("============================")
				// fmt.Println(filteredArtists)
			case "member":
				// Filtrer les artistes dont le nom correspond à la recherche
				var searchedArtist []model.AllArtists
				for _, artist := range AllInfo {
					for _, member := range artist.Members {
						if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
							searchedArtist = append(searchedArtist, artist)
							// fmt.Println("================================")
							// fmt.Println(searchedArtist)
						}
					}
				}
				// Utilise les artistes filtrés pour l'affichage
				filteredArtists = searchedArtist
				// fmt.Println("================================")
				// fmt.Println(filteredArtists)
			case "location":
				// Filtrer les artistes dont le nom correspond à la recherche
				var searchedArtist []model.AllArtists
				for _, artist := range AllInfo {
					for _, location := range artist.Locations {
						if strings.Contains(strings.ToLower(location), strings.ToLower(query)) {
							searchedArtist = append(searchedArtist, artist)
							// fmt.Println("================================")
							// fmt.Println(searchedArtist)
						}
					}
				}
				// Utilise les artistes filtrés pour l'affichage
				filteredArtists = searchedArtist
				// fmt.Println("================================")
				// fmt.Println(filteredArtists)
			}
		}
	}
	filteredArtists = tools.RemoveDuplicatesArtists(filteredArtists)
	return filteredArtists
}
