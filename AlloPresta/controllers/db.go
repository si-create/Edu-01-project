package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3" // Assurez-vous que le driver SQLite est importé
)

// Déclarez une variable pour la base de données
var DB *sql.DB

// Fonction d'initialisation de la base de données
func Init() {
	var err error
	// Ouvrir la base de données SQLite
	DB, err = sql.Open("sqlite3", "./AlloPresta.db") // Remplacez le chemin par celui de votre fichier SQLite
	if err != nil {
		log.Fatal("Erreur lors de la connexion à la base de données: ", err)
	}

	// Vérifiez la connexion
	err = DB.Ping()
	if err != nil {
		log.Fatal("Impossible de se connecter à la base de données: ", err)
	}

	// Exécuter le script d'initialisation de la base de données
	err = initDatabase()
	if err != nil {
		log.Fatal("Erreur lors de l'initialisation de la base de données: ", err)
	}

	log.Println("Connexion à la base de données réussie et base de données initialisée!")
}

func initDatabase() error {
	// Lire le fichier SQL contenant les requêtes de création des tables
	sqlFile, err := ioutil.ReadFile("./controllers/init_db.sql")

	if err != nil {
		return fmt.Errorf("Erreur lors de la lecture du fichier SQL: %v", err)
	}

	// Exécuter le script SQL pour créer les tables
	_, err = DB.Exec(string(sqlFile))
	if err != nil {
		return fmt.Errorf("Erreur lors de l'exécution du script SQL: %v", err)
	}

	return nil
}


// Fonction pour fermer la connexion à la base de données
func Close() {
	if err := DB.Close(); err != nil {
		log.Fatal("Erreur lors de la fermeture de la base de données: ", err)
	}
}
