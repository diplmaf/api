package main

import (
	"log"
	"travel-api/internal/routes"
)

func main() {
	app := routes.SetupRoutes()
	log.Println("Сервер запущен на http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}