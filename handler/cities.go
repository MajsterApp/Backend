package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

var cities []string

func LoadPlacesFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data struct {
		Cities []string `json:"cities"`
	}
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	cities = data.Cities


	return nil
}

func GetCities(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}


	var filtered []string
	for _, city := range cities {
		if strings.HasPrefix(strings.ToLower(city), strings.ToLower(query)) {
			filtered = append(filtered, city)
		}
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]string{"cities": filtered})
}

