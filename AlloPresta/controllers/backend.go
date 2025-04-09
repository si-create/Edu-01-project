package controllers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"golang.org/x/crypto/bcrypt"
	"database/sql" // Ajoutez cette ligne
	"html/template"
	
)
type User struct {
    ID          int
    Name        string
    PhoneNumber string
    Password    string
}


// Regex pour numéros ivoiriens
var phoneRegex = regexp.MustCompile(`^(01|05|07|25|27|21|22|23|24|30|31|32|33|34|35|36|37|38|39)\d{8}$`)

// Vérification du mot de passe (minimum 6 caractères, une majuscule, un chiffre, et des lettres)
func isValidPassword(password string) bool {
	if len(password) < 6 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
		if char >= 'a' && char <= 'z' {
			hasLower = true
		}
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}

	PhoneNumber := r.Form.Get("Number")
	Name := r.Form.Get("Name")
	Password := r.Form.Get("Password")
	Prenom := r.Form.Get("Prenom")

	// Initialiser la base de données
	Init()
	defer Close() // Fermer la base de données après l'utilisation

	// Vérifications
	if Name == "" || len(Name) < 3 {
		http.Error(w, "Le nom doit contenir au moins 3 caractères", http.StatusBadRequest)
		return
	}
// Vérifications
if Prenom == "" || len(Prenom) < 3 {
	http.Error(w, "Le nom doit contenir au moins 3 caractères", http.StatusBadRequest)
	return
}
	if Password == "" || !isValidPassword(Password) {
		http.Error(w, "Le mot de passe doit contenir au moins 6 caractères, une majuscule, un chiffre et des lettres", http.StatusBadRequest)
		return
	}

	if PhoneNumber == "" {
		http.Error(w, "Numéro de téléphone requis", http.StatusBadRequest)
		return
	}

	if len(PhoneNumber) > 10 && PhoneNumber[:4] == "+225" {
		PhoneNumber = PhoneNumber[4:]
	}

	if len(PhoneNumber) != 10 {
		http.Error(w, "Le numéro doit contenir exactement 10 chiffres", http.StatusBadRequest)
		return
	}

	if !regexp.MustCompile(`^\d+$`).MatchString(PhoneNumber) {
		http.Error(w, "Le numéro ne doit contenir que des chiffres", http.StatusBadRequest)
		return
	}

	if !phoneRegex.MatchString(PhoneNumber) {
		http.Error(w, "Numéro de téléphone invalide", http.StatusBadRequest)
		return
	}

	// Vérifier si le numéro est déjà utilisé
	if checkPhoneNumberExists(PhoneNumber) {
		http.Error(w, "Ce numéro est déjà utilisé", http.StatusConflict)
		return
	}

	// Hash du mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erreur lors du hachage du mot de passe", http.StatusInternalServerError)
		return
	}

	fmt.Println(Name,PhoneNumber,hashedPassword)
	// Insérer les données dans la table utilisateurs
	query := `INSERT INTO utilisateurs (nom,prenom, phone_number, password_hash )VALUES (?, ?,?, ?)`
	result, err := DB.Exec(query, Name,Prenom, PhoneNumber, hashedPassword)
	if err != nil {
		log.Printf("Erreur lors de l'insertion des données: %v", err)
    http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
		return
	}

// Récupérer l'ID de l'utilisateur inséré (nous en avons besoin pour la session)
userID, err := result.LastInsertId()
if err != nil {
	log.Printf("Erreur lors de l'obtention de l'ID de l'utilisateur: %v", err)
	http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
	return
}

    // Créer une session pour l'utilisateur après l'inscription
    sessionID := CreateSession(w, int(userID))
	fmt.Println("Session ID : ", sessionID)

	
     // Charger les fichiers HTML : page principale, header, et footer
     tmpl, err := template.ParseFiles("frontend/templates/accueil.html", "frontend/templates/Header.html", "frontend/templates/footer.html")
     if err != nil {
         log.Println("Erreur lors du chargement des templates :", err)
         http.Error(w, "Erreur interne", http.StatusInternalServerError)
         return
     }

    // Passer le numéro de téléphone au template
    tmpl.Execute(w, map[string]interface{}{
        "PhoneNumber": PhoneNumber,
    })
    

}

func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
        return
    }

    PhoneNumber := r.Form.Get("Number")
    Password := r.Form.Get("passwordSignin")

    // Vérifier si les champs sont vides
    if PhoneNumber == "" || Password == "" {
        http.Error(w, "Numéro de téléphone et mot de passe requis", http.StatusBadRequest)
        return
    }

	if PhoneNumber == "" {
		http.Error(w, "Numéro de téléphone requis", http.StatusBadRequest)
		return
	}

	if len(PhoneNumber) > 10 && PhoneNumber[:4] == "+225" {
		PhoneNumber = PhoneNumber[4:]
	}

	if len(PhoneNumber) != 10 {
		http.Error(w, "Le numéro doit contenir exactement 10 chiffres", http.StatusBadRequest)
		return
	}

	if !regexp.MustCompile(`^\d+$`).MatchString(PhoneNumber) {
		http.Error(w, "Le numéro ne doit contenir que des chiffres", http.StatusBadRequest)
		return
	}

	if !phoneRegex.MatchString(PhoneNumber) {
		http.Error(w, "Numéro de téléphone invalide", http.StatusBadRequest)
		return
	}

	

    // Initialiser la base de données
    Init()
    defer Close() // Fermer la base de données après l'utilisation

    // Vérifier si le numéro de téléphone existe dans la base de données
    user, err := getUserByPhoneNumber(PhoneNumber)
    if err != nil {
        http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
        return
    }

    // Vérifier si le mot de passe correspond
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
    if err != nil {
        http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
        return
    }

    // Créer une session pour l'utilisateur après l'inscription
    sessionID := CreateSession(w, int(user.ID))
	fmt.Println("Session ID : ", sessionID)

	
     // Charger les fichiers HTML : page principale, header, et footer
     tmpl, err := template.ParseFiles("frontend/templates/Accueil.html", "frontend/templates/Header.html", "frontend/templates/footer.html")
     if err != nil {
         log.Println("Erreur lors du chargement des templates :", err)
         http.Error(w, "Erreur interne", http.StatusInternalServerError)
         return
     }

    // Passer le numéro de téléphone au template
    tmpl.Execute(w, map[string]interface{}{
        "PhoneNumber": PhoneNumber,
    })
     
}

func getUserByPhoneNumber(phoneNumber string) (*User, error) {
    fmt.Println("Recherche de l'utilisateur avec le numéro de téléphone :", phoneNumber)
    query := `SELECT id, nom, phone_number, password_hash FROM utilisateurs WHERE phone_number = ?`
    row := DB.QueryRow(query, phoneNumber)

    var user User
    err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("Utilisateur non trouvé pour le numéro:", phoneNumber)
            return nil, fmt.Errorf("Utilisateur non trouvé")
        }
        fmt.Println("Erreur lors de la récupération de l'utilisateur:", err)
        return nil, fmt.Errorf("Erreur lors de la récupération de l'utilisateur : %v", err)
    }

    fmt.Println("Utilisateur trouvé :", user)
    return &user, nil
}

func checkPhoneNumberExists(phone string) bool {
	// Vérifier si le numéro existe déjà dans la base de données
	query := "SELECT COUNT(*) FROM utilisateurs WHERE phone_number = ?"
	var count int
	err := DB.QueryRow(query, phone).Scan(&count)
	if err != nil {
		log.Fatal("Erreur lors de la vérification du numéro de téléphone: ", err)
		return false
	}

	return count > 0
}




