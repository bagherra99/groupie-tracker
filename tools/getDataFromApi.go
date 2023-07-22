package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const API = "https://groupietrackers.herokuapp.com/api"


func GetDataFromApi() []string {
	resp, err := http.Get(API)
	if err != nil {
		fmt.Println("No response: ", err)
		os.Exit(0)
	}
	defer resp.Body.Close()

	// var target map[string]string
	var target = make(map[string]string)
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(0)
	}

	var tab []string
	var ArtistsApi, LocationsApi, RelationsApi, DatesApi string
	ArtistsApi = target["artists"]
	tab = append(tab, ArtistsApi)
	LocationsApi = target["locations"]
	tab = append(tab, LocationsApi)
	RelationsApi = target["dates"]
	tab = append(tab, RelationsApi)
	DatesApi = target["relation"]
	tab = append(tab, DatesApi)

	return tab
	// fmt.Println(tab)
}
