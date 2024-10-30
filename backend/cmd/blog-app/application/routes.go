package application

import (
	"net/http"

	"github.com/Sairam-04/blog-app/backend/api/handler/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/users", loadUserRoutes)
	return router
}

func loadUserRoutes(router chi.Router) {
	userHandler := &user.User{}

	router.Post("/register", userHandler.Register)
	router.Get("/login", userHandler.Login)
}

// Handlers - Services
//  Queries - repository
