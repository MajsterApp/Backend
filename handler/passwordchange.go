package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MajsterApp/Backend/db"
	"golang.org/x/crypto/bcrypt"
)

func PasswordChange(w http.ResponseWriter, r *http.Request) {

	var rq struct {
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		http.Error(w, "wrong request", http.StatusUnauthorized)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Can't hash the password: %v", err)
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, `{"error": "Missing Authorization header"}`, http.StatusUnauthorized)
		return
	}

	var tokenString string
	_, err = fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
	if err != nil {
		http.Error(w, `{"error": "Invalid Authorization header format"}`, http.StatusUnauthorized)
		return
	}

	claims, err := getClaims(tokenString)
	if err != nil {
		http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
		return
	}

	email, ok := claims["email"].(string)
	if !ok {
		http.Error(w, `{"error": "Invalid token: email claim missing"}`, http.StatusUnauthorized)
		return
	}
	conn := db.DB
	_, err = conn.Exec("UPDATE users SET password = $1 WHERE email = $2", hash, email)
	if err != nil {
		log.Printf("Database update error: %v", err)
		http.Error(w, `{"error": "Couldn't update password"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Hasło zmienione pomyślnie"})
}
