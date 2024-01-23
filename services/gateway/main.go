package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/services/gateway/api"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	brk := broker.NewMessageBroker()

	handler := api.NewAPIHandler(app, brk)
	handler.SetupRoutes()

	log.Fatal(app.Listen(":5175"))
}
