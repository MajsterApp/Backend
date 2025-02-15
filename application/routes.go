package application

import (
    "github.com/MajsterApp/Backend/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func loadRoutes() *chi.Mux {
    router := chi.NewRouter();
    router.Use(middleware.Logger);
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK);
    })
    router.Route("/api/v1", loadHandlerRoutes)

    return router;

}

func loadHandlerRoutes(router chi.Router) {
    Handler := &handler.Order{}
   // handler routes right here
    router.Post("/login", Handler.Login);
    router.Post("/register", Handler.Register);

}

