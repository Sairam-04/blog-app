package application

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Sairam-04/blog-app/backend/internal/config"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context, cfg *config.Config) error {
	server := &http.Server{
		Addr:    cfg.Port,
		Handler: a.router,
	}
	slog.Info("server running at", slog.String("port", cfg.Port))
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server %s", err)
	}
	return nil
}
