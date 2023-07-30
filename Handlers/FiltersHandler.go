package Handlers

/* 
import (
	ResponseData "groupie/ResponseData"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func FiltersHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the search query parameters from the URL
	name := r.FormValue("search")
	creationDateFromStr := r.FormValue("creationDateFrom")
	creationDateToStr := r.FormValue("creationDateTo")
	firstAlbumDateStr := r.FormValue("firstAlbumDate")
	numMembersStr := r.FormValue("numMembers")
	concertLocations := r.FormValue("concertLocations")

	// Convert the date and number values to integers (you can add more error handling here)
	creationDateFrom, _ := strconv.Atoi(creationDateFromStr)
	creationDateTo, _ := strconv.Atoi(creationDateToStr)
	firstAlbumDate, _ := strconv.Atoi(firstAlbumDateStr)
	numMembers, _ := strconv.Atoi(numMembersStr)

	// Perform the search based on the filter values
	searchResults := FiltersArtist(name, creationDateFrom, creationDateTo, firstAlbumDate, numMembers, concertLocations)

	data := struct {
		Allartist     []ResponseData.AllArtists
		Locat         []string
		Memb          []string
		SearchResults []ResponseData.AllArtists
	}{
		Allartist:     AllInfo,
		Locat:         getUniqueLocations(),
		Memb:          getUniqueMembers(),
		SearchResults: searchResults,
	}

	templates, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		http.Error(w, "error while parsing", http.StatusInternalServerError)
		return
	}

	err = templates.Execute(w, data)
	if err != nil {
		http.Error(w, "error while executing", http.StatusInternalServerError)
	}
}

// Helper function to get unique locations from the AllInfo data
func getUniqueLocations() []string {
	var locations []string
	for _, artist := range AllInfo {
		for _, location := range artist.Locations {
			if !containsLocation(locations, location) {
				locations = append(locations, location)
			}
		}
	}
	return locations
}

// Helper function to check if a location is already in the list
func containsLocation(locations []string, location string) bool {
	for _, loc := range locations {
		if loc == location {
			return true
		}
	}
	return false
}

// Helper function to get unique members from the AllInfo data
func getUniqueMembers() []string {
	var members []string
	for _, artist := range AllInfo {
		for _, member := range artist.Members {
			if !containsMember(members, member) {
				members = append(members, member)
			}
		}
	}
	return members
}

// Helper function to check if a member is already in the list
func containsMember(members []string, member string) bool {
	for _, mem := range members {
		if mem == member {
			return true
		}
	}
	return false
}

// Sample function to search artists based on filter values
func FiltersArtist(name string, creationDateFrom int, creationDateTo int, firstAlbumDate int, numMembers int, concertLocations string) []ResponseData.AllArtists {
	var results []ResponseData.AllArtists

	for _, artist := range AllInfo {
		if (name == "" || strings.Contains(strings.ToLower(artist.Name), strings.ToLower(name))) &&
			(creationDateFrom == 0 || artist.CreationDate >= creationDateFrom) &&
			(creationDateTo == 0 || artist.CreationDate <= creationDateTo) &&
			(firstAlbumDate == 0 || artist.FirstAlbum == firstAlbumDate) &&
			(numMembers == 0 || len(artist.Members) == numMembers) &&
			(concertLocations == "" || containsLocation(artist.Locations, concertLocations)) {
			results = append(results, artist)
		}
	}

	return results
}
 */