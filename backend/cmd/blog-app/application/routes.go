package application

import (
	"net/http"

	"github.com/Sairam-04/blog-app/backend/api/handler/blog"
	"github.com/Sairam-04/blog-app/backend/api/handler/user"
	"github.com/Sairam-04/blog-app/backend/api/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes(userHandler *user.UserHandler, blogHandler *blog.BlogHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/users", func(r chi.Router) {
		loadUserRoutes(r, userHandler)
	})

	router.Route("/blogs", func(r chi.Router) {
		loadBlogRoutes(r, blogHandler)
	})
	return router
}

func loadUserRoutes(router chi.Router, userHandler *user.UserHandler) {
	router.Post("/register", userHandler.Register)
	router.Post("/login", userHandler.Login)
	router.With(middlewares.AuthMiddleware).Get("/{id}", userHandler.GetUser)
}

func loadBlogRoutes(router chi.Router, blogHandler *blog.BlogHandler) {
	router.With(middlewares.AuthMiddleware).Post("/create", blogHandler.PostBlog)
	router.Get("/all", blogHandler.GetAllBlogs)
	router.With(middlewares.AuthMiddleware).Get("/user", blogHandler.GetUserBlogs)
	router.With(middlewares.AuthMiddleware).Patch("/update/{id}", blogHandler.UpdateBlog)
}

// Handlers - Services
//  Queries - repository
