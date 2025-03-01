package application

import (
    "context"
    "fmt"
    "net/http"
    "os"
)

type App struct {
    router http.Handler
}

func New() *App {
    return &App{
        router: loadRoutes(),
    }
}

func (a *App) Start(ctx context.Context) error {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    server := &http.Server{
        Addr:    ":" + port,
        Handler: a.router,
    }

    err := server.ListenAndServe()
    if err != nil {
        return fmt.Errorf("failed to start server: %w", err)
    }
    return nil
}

