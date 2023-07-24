package tools

import (
	"main/model"
	"strings"
)

// func RemoveDuplicates(input []string) []string {
//     var result []string
//     seen := make(map[string]bool)

//     for _, item := range input {
//         lowercaseItem := strings.ToLower(item) // Ignore the case for filtering
//         if !seen[lowercaseItem] {
//             seen[lowercaseItem] = true
//             result = append(result, item)
//         }
//     }

//     return result
// }

// removeDuplicatesArtists supprime les doublons dans la liste des artistes en utilisant le nom des artistes comme clé
func RemoveDuplicatesArtists(artistsList []model.AllArtists) []model.AllArtists {
	uniqueMap := make(map[string]bool)
	var result []model.AllArtists
	for _, artist := range artistsList {
		lowercaseName := strings.ToLower(artist.Name) // Ignorer la casse pour le filtrage
		if !uniqueMap[lowercaseName] {
			uniqueMap[lowercaseName] = true
			result = append(result, artist)
		}
	}
	return result
}

func IsStringInArray(arr []string, str string) bool {
	for _, value := range arr {
		if strings.ToLower(value) == str {
			return true
		}
	}
	return false
}

func IsStringInMember(arr []string, str string) bool {
	for _, value := range arr {
		if value == str {
			return true
		}
	}
	return false
}
