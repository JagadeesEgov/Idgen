package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/yourorg/idgen/internal"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get Postgres DSN from env
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("POSTGRES_DSN not set in environment")
	}

	// TODO: Initialize DB connection and pass to handlers
	db, err := internal.InitDB(dsn)
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.POST("/template", internal.RegisterTemplateHandler(db))
	r.POST("/generate", internal.GenerateIdHandler(db))

	log.Println("Starting server on :8080")
	r.Run(":8080")
} 