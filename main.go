package main

import (
	"context"
	"log"

	"github.com/MajsterApp/Backend/application"
	"github.com/MajsterApp/Backend/db"
	"github.com/MajsterApp/Backend/handler"
)

func main() {
    if err := handler.LoadPlacesFromFile("assets/cities.json"); err != nil {
		log.Fatal("Error loading cities:", err)
	}
	app := application.New()

    db.InitDB()
	defer db.DB.Close()
	if err := app.Start(context.TODO()); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}


}
