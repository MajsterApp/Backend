package main

import (
	"context"
	"log"

	"github.com/MajsterApp/Backend/application"
	"github.com/MajsterApp/Backend/db"
)

func main() {
	app := application.New()

    db.InitDB()
	defer db.DB.Close()
	if err := app.Start(context.TODO()); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}


	log.Println("Server running on localhost:3000")
}

