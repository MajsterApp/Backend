package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MajsterApp/Backend/db"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const sqlStatement = `
    INSERT INTO users (email, name, surname, password, region, jobs, role)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
`

func RegisterFunc(w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		Email    string   `json:"email"`
		Name     string   `json:"name"`
		Surname  string   `json:"surname"`
		Password string   `json:"password"`
		Region   string   `json:"region"`
		Jobs     []string `json:"jobs"`
		Role     string   `json:"role"`
	}
	conn := db.DB

	var rq requestBody
	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Can't hash the password: %v", err)
	}

	_, err = conn.Exec(sqlStatement, rq.Email, rq.Name, rq.Surname, string(hash), rq.Region, pq.Array(rq.Jobs), rq.Role)
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Failed to insert data into the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}
