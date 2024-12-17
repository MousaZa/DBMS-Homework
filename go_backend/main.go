package main

import (
	// "net"
	"os"

	"github.com/MousaZa/DBMS-Homework/go_backend/storage"
	// "github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
)

func main() {
	log := hclog.Default()
	err := godotenv.Load(".env")
	log.Info("Starting borrows server" + os.Getenv("DB_HOST"))

	if err != nil {
		log.Error("Unable to get env", "error", err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}

	log.Info("Connected to database", "db", db)

}
