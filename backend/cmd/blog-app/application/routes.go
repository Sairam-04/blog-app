package application

import (
	"net/http"

	"github.com/Sairam-04/blog-app/backend/api/handler/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes(userHandler *user.UserHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/users", func(r chi.Router) {
		loadUserRoutes(r, userHandler)
	})
	return router
}

func loadUserRoutes(router chi.Router, userHandler *user.UserHandler) {
	router.Post("/register", userHandler.Register)
	router.Get("/login", userHandler.Login)
}

// Handlers - Services
//  Queries - repository
