package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID                int                 `json:"id"`
		DatesandLocations map[string][]string `json:"datesLocations"`
	}
}

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

// func ReceiveInformaionFromJsonFile() {
	
// }

func main() {
	//getDataFromApi(API, INFO_API)
	// var target map[string]interface{}
	// fmt.Println(getDataFromApi(API))
	// fmt.Println(getDataFromApiArtist())
	// artists := getDataFromApiArtist()

	// jsonData, err := json.MarshalIndent(artists, "", " ")
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// 	os.Exit(1)
	// }

	// dates := getDataFromApiDates()

	// jsonData1, err := json.MarshalIndent(dates, "", " ")
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// 	os.Exit(1)
	// }

	// place := getDataFromApiLocations()

	// jsonData1, err := json.MarshalIndent(place, "", " ")
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// 	os.Exit(1)
	// }

	info := ReceiveDataFromJson()

	jsonData1, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		os.Exit(1)
	}
	
	fmt.Println(string(jsonData1))
	//===========================================//
}
