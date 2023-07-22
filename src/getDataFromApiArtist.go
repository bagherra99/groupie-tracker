package src

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/tools"
	"net/http"
	"os"
)

func GetDataFromApiArtist() []model.Artist {
	ApiArtist := tools.GetDataFromApi()[0]
	resp, err := http.Get(ApiArtist)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	// create a slice for our JSON data to be unmarshalled into.
	Artist := []model.Artist{}
	err = json.NewDecoder(resp.Body).Decode(&Artist)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(1)
	}

	return Artist
}
