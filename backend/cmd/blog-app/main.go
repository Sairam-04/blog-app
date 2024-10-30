package main

import (
	"context"
	"log"

	"github.com/Sairam-04/blog-app/backend/cmd/blog-app/application"
	"github.com/Sairam-04/blog-app/backend/internal/config"
)

func main() {
	// Loading configurations
	cfg := config.MustLoad()

	// creating app instance
	app := application.New()

	// starting a server
	err := app.Start(context.TODO(), cfg)
	if err != nil {
		log.Fatal("failed to start app:", err)
	}

}
