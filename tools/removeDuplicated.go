package tools

import (
	"main/model"
	"strings"
)

// RemoveDuplicatesArtists supprime les doublons dans la liste des artistes en utilisant le nom des artistes comme clé
func RemoveDuplicatesArtists(artistsList []model.AllArtists) []model.AllArtists {
	// Crée une carte pour suivre les noms d'artistes déjà rencontrés
	uniqueMap := make(map[string]bool)
	var result []model.AllArtists

	// Parcours de chaque artiste dans la liste
	for _, artist := range artistsList {
		lowercaseName := strings.ToLower(artist.Name) // Ignorer la casse pour le filtrage

		// Vérifie si le nom de l'artiste n'a pas déjà été rencontré
		if !uniqueMap[lowercaseName] {
			uniqueMap[lowercaseName] = true // Marque le nom de l'artiste comme rencontré
			result = append(result, artist) // Ajoute l'artiste à la liste de résultats
		}
	}
	return result
}

// IsStringInArray vérifie si une chaîne de caractères existe dans un tableau de chaînes
func IsStringInArray(arr []string, str string) bool {
	for _, value := range arr {
		// Compare en ignorant la casse pour déterminer si la chaîne existe dans le tableau
		if strings.ToLower(value) == str {
			return true
		}
	}
	return false
}

// IsStringInMember vérifie si une chaîne de caractères existe dans un tableau sans comparaison de casse
func IsStringInMember(arr []string, str string) bool {
	for _, value := range arr {
		// Compare directement pour déterminer si la chaîne existe dans le tableau
		if value == str {
			return true
		}
	}
	return false
}
