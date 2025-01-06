package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Structure représentant les informations à afficher
type whoami struct {
	Name  string
	Title string
	State string
}

func main() {
	// Appel à la fonction pour démarrer le serveur
	request1()
}

// Fonction qui répond à l'endpoint /whoami
func whoAmI(response http.ResponseWriter, r *http.Request) {
	// Remplace les informations existantes par le nom de ton équipe
	who := []whoami{
		whoami{
			Name:  "Team Leo et Pierre", // Remplace "Ton Équipe" par le nom réel de ton équipe
			Title: "DevOps et Déploiement Continu",
			State: "FR",
		},
	}

	// Encodage en JSON et renvoi de la réponse
	json.NewEncoder(response).Encode(who)

	// Affichage dans la console (pour débogage)
	fmt.Println("Endpoint Hit: /whoami", who)
}

// Fonction qui répond à l'endpoint racine "/"
func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Bienvenue dans l'API Web!")
	fmt.Println("Endpoint Hit: homePage")
}

// Fonction qui répond à l'endpoint "/aboutme"
func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Ton Équipe" // Remplace par le nom de ton équipe

	fmt.Fprintf(response, "Un peu à propos de nous...")
	fmt.Println("Endpoint Hit: ", who)
}

// Fonction pour définir les handlers HTTP
func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)

	// Démarre le serveur HTTP sur le port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
