package src

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/tools"
	"net/http"
	"os"
)

func GetDataFromApiRelation() model.Relation {
	ApiRelation := tools.GetDataFromApi()[3]
	resp, err := http.Get(ApiRelation)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	var Info model.Relation

	err = json.NewDecoder(resp.Body).Decode(&Info)
	if err != nil {
		fmt.Println("Error decoding API response:", err)
		os.Exit(1)
	}

	return Info
}
