package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func generateJWTEmail(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
 func CreateToken(w http.ResponseWriter, r *http.Request) {

    var rq struct {
        Email string `json:"email"`
    }

    err := json.NewDecoder(r.Body).Decode(&rq)
    if err != nil {
        http.Error(w, "wrong request ", http.StatusUnauthorized)
        return
    }

	token, err := generateJWTEmail(rq.Email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
    })


}
