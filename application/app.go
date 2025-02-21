package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type App struct {
	router http.Handler
}

func New() *App {
	router := loadRoutes()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	app := &App{
		router: corsHandler,
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	fmt.Println("Server is running on localhost:3000")
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to Start server: %w", err)
	}
	return nil
}
