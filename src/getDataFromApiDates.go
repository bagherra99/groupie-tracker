package src

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/tools"
	"net/http"
	"os"
)

func GetDataFromApiDates() model.Dates {
	ApiDates := tools.GetDataFromApi()[2]
	resp, err := http.Get(ApiDates)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	var dates model.Dates
	err = json.NewDecoder(resp.Body).Decode(&dates)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(1)
	}

	return dates
}
