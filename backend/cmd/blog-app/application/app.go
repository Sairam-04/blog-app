package application

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Sairam-04/blog-app/backend/internal/config"
)

type App struct {
	router http.Handler
	server *http.Server // Add the server as a field
}

func New(cfg *config.Config) *App {
	return &App{
		router: loadRoutes(),
		server: &http.Server{
			Addr:    cfg.Port,
			Handler: loadRoutes(),
		},
	}
}

func (a *App) Start(ctx context.Context, cfg *config.Config) error {
	slog.Info("Server running at", slog.String("port", cfg.Port))
	err := a.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %s", err)
	}
	return nil
}

// gracefully shuts down the server
func (a *App) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %s", err)
	}

	slog.Info("Server shut down gracefully")
	return nil
}
