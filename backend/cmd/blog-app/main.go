package main

import (
	"context"
	"fmt"
	"log"

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
	app := application.New()

	// starting a server
	err := app.Start(context.TODO(), cfg)
	if err != nil {
		log.Fatal("failed to start app:", err)
	}

}
