package Tools

import (
	"groupie/ResponseData"
	"strings"
)

func RemoveDuplicatesArtists(artistsList []ResponseData.AllArtists) []ResponseData.AllArtists {
	uniqueMap := make(map[string]bool)
	var result []ResponseData.AllArtists
	for _, artist := range artistsList {
		lowercaseName := strings.ToLower(artist.Name) // Ignorer la casse pour le filtrage
		if !uniqueMap[lowercaseName] {
			uniqueMap[lowercaseName] = true
			result = append(result, artist)
		}
	}
	return result
}

func IsStringInString(arr []string, str string) bool {
	for _, value := range arr {
		search := strings.ToLower(value)
		if strings.Contains(search, str) {
			return true
		}
	}
	return false
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
