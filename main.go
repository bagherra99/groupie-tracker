package main

import (
	"fmt"
	"main/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.IndexDataHandler)
	http.HandleFunc("/artist/", handlers.ArtistDataHandler)
	http.HandleFunc("/search/", handlers.SearchHandler)
	// http.Handle("/", http.FileServer(http.Dir("public"))) // Remplace "public" par le répertoire contenant ton fichier HTML

	fmt.Println("Serveur démarré sur http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Verifier votre connexion")
	}
}
