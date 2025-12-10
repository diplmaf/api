package handlers

import (
	"net/http"
	"travel-api/internal/models"
	"travel-api/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type TravelHandler struct {
	storage *storage.Storage
}

func NewTravelHandler(s *storage.Storage) *TravelHandler {
	return &TravelHandler{storage: s}
}

func (h *TravelHandler) GetTrips(c *fiber.Ctx) error {
	trips := h.storage.GetAllTrips()
	return c.Status(http.StatusOK).JSON(trips)
}

func (h *TravelHandler) GetTrip(c *fiber.Ctx) error {
	id := c.Params("id")
	trip, found := h.storage.GetTripByID(id)
	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Путешествие не найдено"})
	}
	return c.Status(http.StatusOK).JSON(trip)
}

func (h *TravelHandler) CreateTrip(c *fiber.Ctx) error {
	var trip models.Trip
	if err := c.BodyParser(&trip); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат JSON"})
	}
	if trip.Destination == "" || trip.Budget < 0 || trip.StartDate.After(trip.EndDate) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректные данные: проверьте даты и бюджет"})
	}
	h.storage.CreateTrip(trip)
	return c.Status(fiber.StatusCreated).JSON(trip)
}

func (h *TravelHandler) UpdateTrip(c *fiber.Ctx) error {
	id := c.Params("id")
	var trip models.Trip
	if err := c.BodyParser(&trip); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат JSON"})
	}
	if !h.storage.UpdateTrip(id, trip) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Путешествие не найдено"})
	}
	return c.Status(http.StatusOK).JSON(trip)
}

func (h *TravelHandler) DeleteTrip(c *fiber.Ctx) error {
	id := c.Params("id")
	if !h.storage.DeleteTrip(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Путешествие не найдено"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
