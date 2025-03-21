package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "postgresql://neondb_owner:npg_aYWmNVB8F1bv@ep-snowy-snow-a8ndujox-pooler.eastus2.azure.neon.tech/neondb?sslmode=require"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connection successful")
}
