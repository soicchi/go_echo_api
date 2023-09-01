package main

import (
	"go_echo_api/routes"
)

func main() {
	e := routes.SetupRoutes()
	e.Logger.Fatal(e.Start(":8000"))
}