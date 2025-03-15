package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MajsterApp/Backend/db"
)

func Verification(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, `{"error": "Missing Authorization header"}`, http.StatusUnauthorized)
		return
	}

	var tokenString string
	_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
	if err != nil {
		http.Error(w, `{"error": "Invalid Authorization header format"}`, http.StatusUnauthorized)
		return
	}

	claims, err := getClaims(tokenString)
	if err != nil {
		http.Error(w, `{"error": "Problem with getting claims from the token"}`, http.StatusUnauthorized)
		return
	}

	email, ok := claims["email"].(string)
	if !ok {
		http.Error(w, `{"error": "Invalid token: email claim missing"}`, http.StatusUnauthorized)
		return
	}

	conn := db.DB
	_, err = conn.Exec("UPDATE users SET status = TRUE WHERE email = $1", email)
	if err != nil {
		http.Error(w, `{"error": "Couldn't change status"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Konto zweryfikowane"})
}
