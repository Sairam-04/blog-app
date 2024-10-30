package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sairam-04/blog-app/backend/cmd/blog-app/application"
	"github.com/Sairam-04/blog-app/backend/internal/config"
	"github.com/Sairam-04/blog-app/backend/pkg"
	_ "github.com/lib/pq"
)

func main() {
	// Loading configurations
	cfg := config.MustLoad()

	// db connection
	db := pkg.NewDBConnection(cfg)
	fmt.Println("db connection successful", db)
	// creating app instance
	app := application.New(cfg)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Start the application
		if err := app.Start(context.Background(), cfg); err != nil {
			log.Fatal("Failed to start app:", err)
		}
	}()

	// Wait for interrupt signal
	<-done
	slog.Info("Shutting down the server...")

	// Create a timeout context for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call the shutdown method of the app
	if err := app.Shutdown(ctx); err != nil {
		log.Fatal("Error during shutdown:", err)
	}
}
