package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
    router http.Handler
}

func New() *App {
    app := &App{
        router: loadRoutes(),
    }

    return app;
}

func (a *App) Start(ctx context.Context) error {
    server := &http.Server{
        Addr: ":3000",
        Handler: a.router,
    }


    fmt.Println("Server is running on localhost:3000")
    err := server.ListenAndServe();
    if err != nil {
        return fmt.Errorf("Failed to Start server: %w", err)
    }
    return nil;
}

