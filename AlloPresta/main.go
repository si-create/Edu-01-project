package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"AlloPresta/controllers" 
	// Utilisation correcte du chemin relatif dans le module
)

func main() {
	fmt.Println("Démarrage du serveur...")

	// Initialiser la base de données
	controllers.Init()
	defer controllers.Close()
	// Chemin relatif pour accéder au dossier 'frontend/static' depuis 'backend'

	staticDir := filepath.Join("frontend", "static")
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Définir la route pour la page d'accueil
	http.HandleFunc("/", controllers.Home)              // Utilisation correcte du package controllers
	http.HandleFunc("/Inscription", controllers.Signup) // Utilisation correcte du package controllers

	http.HandleFunc("/Login", controllers.Login) // Utilisation correcte du package controllers

	http.HandleFunc("/Loginpage", controllers.LoginPage) // Utilisation correcte du package controllers
	// Lancement du serveur

	
	fmt.Println("Serveur démarré sur le port: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erreur lors du démarrage du serveur:", err)
	}
}
