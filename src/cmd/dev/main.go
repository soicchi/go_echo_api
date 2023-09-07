package main

import (
	"log"
	"os"

	"go_echo_api/database"
	"go_echo_api/routes"
)

func main() {
	log.Println("Starting Lambda")

	// Database connection
	dbConfig := database.NewDBConfig(
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	dsn := dbConfig.CreateDSN()
	_, err := database.DBConnect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")

	e := routes.SetupRoutes()
	e.Logger.Fatal(e.Start(":8000"))
}