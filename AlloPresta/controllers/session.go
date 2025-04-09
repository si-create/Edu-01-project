package controllers

import (
	"fmt"
	"net/http"
	"github.com/google/uuid"
	"time"
)

// Structure pour la gestion des sessions
type Session struct {
	ID        string
	UserID    int
	ExpiresAt time.Time
}

// Map pour stocker les sessions actives
var sessions = make(map[string]Session)

// Fonction pour créer une nouvelle session
func CreateSession(w http.ResponseWriter, userID int) string {
	// Générer un identifiant unique pour la session
	sessionID := uuid.New().String()

	// Définir une durée d'expiration pour la session (par exemple 1 heure)
	expiration := time.Now().Add(1 * time.Hour)

	// Créer une nouvelle session
	session := Session{
		ID:        sessionID,
		UserID:    userID,
		ExpiresAt: expiration,
	}

	// Stocker la session dans la map (vous pourriez utiliser une base de données ici)
	sessions[sessionID] = session

	// Créer un cookie de session
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionID,
		Expires: expiration,
		Path:    "/",
	})

	return sessionID
}

// Fonction pour récupérer la session active depuis un cookie
func GetSession(r *http.Request) (*Session, error) {
	// Récupérer le cookie de session
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, fmt.Errorf("cookie non trouvé")
	}

	// Vérifier si la session existe dans la map
	session, exists := sessions[cookie.Value]
	if !exists || session.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("session expirée ou non trouvée")
	}

	return &session, nil
}

// Fonction pour supprimer une session
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return
	}

	// Supprimer la session de la map
	delete(sessions, cookie.Value)

	// Effacer le cookie du navigateur
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

