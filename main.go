package main

import (
	"log"
	"staffsmart-lite/internal/database"
	"staffsmart-lite/internal/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	databaseURL := "postgres://postgres:postgres@localhost:5432/staffsmart?sslmode=disable"
	db, err := database.Connect(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := database.NewOrderRepository(db)
	app := fiber.New()
	app.Post("/orders", func(c *fiber.Ctx) error {
		var input models.CreateOrderInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid request body",
			})

		}
		if input.SlotID <= 0 || strings.TrimSpace(input.CustomerName) == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid input",
			})
		}
		order, err := repo.CreateOrder(input)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(201).JSON(order)

	})

	app.Get("/slots", func(c *fiber.Ctx) error {
		slots, err := repo.GetAllSlots()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(slots)
	})
	log.Fatal(app.Listen(":8080"))
}
