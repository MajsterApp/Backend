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
        AllowedOrigins:   []string{"https://majsterapp.netlify.app", "https://localhost:5173"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodHead {
        w.WriteHeader(http.StatusOK)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to MajsterApp API"))
})

    router.Route("/api/v1", loadHandlerRoutes)

    return router;

}

func loadHandlerRoutes(router chi.Router) {
    Handler := &handler.Order{}
    router.Post("/login", Handler.Login);
    router.Post("/register", Handler.Register);
    router.Get("/userData", Handler.UserData);
    router.Get("/verification", Handler.Verification)
    router.Post("/change-password", Handler.PasswordChange)
}

