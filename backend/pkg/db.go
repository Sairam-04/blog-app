package pkg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Sairam-04/blog-app/backend/internal/config"
)

func NewDBConnection(cfg *config.Config) *sql.DB {
	return connectDB(cfg)
}

func connectDB(cfg *config.Config) *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("unable to connect to db", err)
	}
	pingDatabase(db)
	return db

}

func pingDatabase(db *sql.DB) {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("ping error", pingErr)
	}
}
