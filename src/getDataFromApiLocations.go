package src

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/tools"
	"net/http"
	"os"
)

func GetDataFromApiLocations() model.Locations {
	ApiLocations := tools.GetDataFromApi()[1]
	resp, err := http.Get(ApiLocations)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	var place model.Locations
	err = json.NewDecoder(resp.Body).Decode(&place)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(1)
	}

	return place
}
