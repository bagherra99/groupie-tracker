package ResponseData

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const API = "https://groupietrackers.herokuapp.com/api"

func getDataFromApi(url string) []string {
	resp, err := http.Get(API)
	if err != nil {
		fmt.Println("No response: ", err)
	}
	defer resp.Body.Close()

	var target map[string]string
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(1)
	}
	//var target map[string]interface{}
	// target, ok := (target.(*map[string]interface{}))
	// if !ok {
	// 	fmt.Println("Invalid target type")
	// 	os.Exit(1)
	// }
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

func getDataFromApiArtist() []Artist {
	ApiArtist := getDataFromApi(API)[0]
	resp, err := http.Get(ApiArtist)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	// create a slice for our JSON data to be unmarshalled into.
	Artist := []Artist{}
	err = json.NewDecoder(resp.Body).Decode(&Artist)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(1)
	}

	return Artist
}

func getDataFromApiDates() Dates {
	ApiDates := getDataFromApi(API)[2]
	resp, err := http.Get(ApiDates)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	var dates Dates
	err = json.NewDecoder(resp.Body).Decode(&dates)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(1)
	}

	return dates
}

func getDataFromApiLocations() Locations {
	ApiLocations := getDataFromApi(API)[1]
	resp, err := http.Get(ApiLocations)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	var place Locations
	err = json.NewDecoder(resp.Body).Decode(&place)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(1)
	}

	return place
}

func getDataFromApiRelation() Relation {
	ApiRelation := getDataFromApi(API)[3]
	resp, err := http.Get(ApiRelation)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	var Info Relation

	err = json.NewDecoder(resp.Body).Decode(&Info)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(1)
	}

	return Info
}
