package main

import (
	"fmt"

	"github.com/Sairam-04/blog-app/backend/internal/config"
)

func main() {
	// Loading configurations
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
