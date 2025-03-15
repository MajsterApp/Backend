package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MajsterApp/Backend/db"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("majsterApp-key")

func generateJWT(email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	conn := db.DB
	if conn == nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	var storedHash, role string
	err = conn.QueryRow("SELECT password, role FROM users WHERE email = $1", req.Email).Scan(&storedHash, &role)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(req.Email, role)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"role":  role})
}
