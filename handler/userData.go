package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MajsterApp/Backend/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
)

func getClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func UserData(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	var tokenString string
	_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
	if err != nil {
		http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
		return
	}

	claims, err := getClaims(tokenString)
	if err != nil {
		http.Error(w, "Problem with getting claims from the token", http.StatusUnauthorized)
		return
	}

	email, ok := claims["email"].(string)
	if !ok {
		http.Error(w, "Invalid token: email claim missing", http.StatusUnauthorized)
		return
	}

	type sqlBody struct {
		Name    string   `json:"name"`
		Surname string   `json:"surname"`
		Region  string   `json:"region"`
		Jobs    []string `json:"jobs"`
	}

	var sq sqlBody
	conn := db.DB

	err = conn.QueryRow("SELECT name, surname, region, jobs FROM users WHERE email = $1", email).
		Scan(&sq.Name, &sq.Surname, &sq.Region, pq.Array(&sq.Jobs))
	if err != nil {
		http.Error(w, "Couldn't retrieve data from the database", http.StatusInternalServerError)
		return
	}

	role, _ := claims["role"].(string)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"email":   email,
		"name":    sq.Name,
		"surname": sq.Surname,
		"region":  sq.Region,
		"jobs":    sq.Jobs,
		"role":    role,
	})
}
