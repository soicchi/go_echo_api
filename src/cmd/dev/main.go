package main

import (
	"log"
	"os"

	"go_echo_api/controllers"
	"go_echo_api/database"
	"go_echo_api/routes"
)

func main() {
	// Database connection
	dbConfig := database.NewDBConfig(
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	dsn := dbConfig.CreateDSN()
	db, err := database.DBConnect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")

	// Database migration
	if err := database.DBMigrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println("Database migration completed")

	h := controllers.NewHandler(db)
	e := routes.SetupRoutes(h)
	e.Logger.Fatal(e.Start(":8000"))
}
