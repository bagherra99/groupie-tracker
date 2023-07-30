package Handlers

/*
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère le terme de recherche saisi par l'utilisateur
	input := r.URL.Query().Get("search")
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
	searched, err := SearchArtist(input)
	if err != nil {
		CustomSearchNotFound(w, r, input)
		return
	}

	data := struct {
		Allartist []ResponseData.AllArtists
		Input     string
		Locat     []string
		Memb      []string
		searched    []ResponseData.AllArtists
	}{
		Allartist: AllInfo,
		Input:     input,
		Locat:     Tab,
		Memb:      TabMember,
		searched:    searched,
	}

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		CustomNotFound500(w)
		return
	}

	err = tmpl.Execute(w, AllInfo)
}

var filteredArtists []ResponseData.AllArtists

func SearchArtist(s string) ([]ResponseData.AllArtists) {
	tabInput := strings.Split(s, "==>")
	fmt.Println(tabInput)
	if len(tabInput) == 2 {
		query := tabInput[0]
		category := tabInput[1]
		fmt.Println(query)
		fmt.Println(category)
		var searchedArtist []ResponseData.AllArtists
		// // Si un terme de recherche est saisi, filtre les artistes correspondants
		if s != "" && strings.ToLower(category) == "artist/band" {
			for _, artist := range AllInfo {
				if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
					searchedArtist = append(searchedArtist, artist)
				}
			}
			filteredArtists = searchedArtist
		} else if s != "" && strings.ToLower(category) == "member" {
			for _, artist := range AllInfo {
				for _, member := range artist.Members {
					if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
						searchedArtist = append(searchedArtist, artist)
					}
				}
			}
			filteredArtists = searchedArtist
		}
	}/*  else {
		Erroor := errors.New("Search not found")
		return nil, Erroor
	}
	fmt.Println(filteredArtists)
	return filteredArtists
}

/////////////////////////////////////////////////////////////////////
/*
func FiltersHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the search query parameters from the URL
	name := r.FormValue("name")
	creationDateFromStr := r.FormValue("creationDateFrom")
	creationDateToStr := r.FormValue("creationDateTo")
	firstAlbumDateFromStr := r.FormValue("firstAlbumDateFrom")
	firstAlbumDateToStr := r.FormValue("firstAlbumDateTo")
	numMembersFromStr := r.FormValue("numMembersFrom")
	numMembersToStr := r.FormValue("numMembersTo")
	concertLocations := r.FormValue("concertLocations")

	// Convert the date and number values to integers (you can add more error handling here)
	creationDateFrom, _ := strconv.Atoi(creationDateFromStr)
	creationDateTo, _ := strconv.Atoi(creationDateToStr)
	firstAlbumDateFrom, _ := strconv.Atoi(firstAlbumDateFromStr)
	firstAlbumDateTo, _ := strconv.Atoi(firstAlbumDateToStr)
	numMembersFrom, _ := strconv.Atoi(numMembersFromStr)
	numMembersTo, _ := strconv.Atoi(numMembersToStr)
	// Perform the search based on the name
	searchResults := FiltersArtist(name, creationDateFrom, creationDateTo, firstAlbumDateFrom, firstAlbumDateTo, numMembersFrom, numMembersTo, concertLocations)

	var templates, err = template.ParseFiles("Templates/search.html")
	if err != nil {
		CustomNotFound500(w)
		return
	}

	templates.Execute(w, searchResults)
}

// Sample function to search artists based on name
func FiltersArtist(name string, creationDateFrom int, creationDateTo int, firstAlbumDateFrom int, firstAlbumDateTo int, numMembersFrom int, numMembersTo int, concertLocations string) []ResponseData.AllArtists {
	var results []ResponseData.AllArtists

	for _, artist := range AllInfo {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(name)) &&
			(artist.CreationDate >= creationDateFrom && artist.CreationDate <= creationDateTo) &&
			(artist.FirstAlbum >= firstAlbumDateFrom && artist.FirstAlbum <= firstAlbumDateTo) &&
			(len(artist.Members) >= numMembersFrom && len(artist.Members) <= numMembersTo) &&
			strings.Contains(strings.ToLower(artist.ConcertLocations), strings.ToLower(concertLocations)) {
			results = append(results, artist)
		}
	}

	return results
} */
