package routes

import (
	"travel-api/internal/handlers"
	"travel-api/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes() *fiber.App {
	app := fiber.New()

	store := storage.NewJSONStorage("data/trips.json")
	travelHandler := handlers.NewTravelHandler(store)

	api := app.Group("/api/v1")
	api.Get("/trips", travelHandler.GetTrips)
	api.Get("/trips/:id", travelHandler.GetTrip)
	api.Post("/trips", travelHandler.CreateTrip)
	api.Put("/trips/:id", travelHandler.UpdateTrip)
	api.Delete("/trips/:id", travelHandler.DeleteTrip)

	return app
}
