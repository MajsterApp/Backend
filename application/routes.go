package application

import (
    "github.com/MajsterApp/Backend/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

)


func loadRoutes() *chi.Mux {
    router := chi.NewRouter();
    router.Use(middleware.Logger);
    router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK);
    })
    router.Route("/api/v1", loadHandlerRoutes)

    return router;

}

func loadHandlerRoutes(router chi.Router) {
    Handler := &handler.Order{}
    router.Post("/login", Handler.Login);
    router.Post("/register", Handler.Register);
    router.Get("/userData", Handler.UserData);

}

