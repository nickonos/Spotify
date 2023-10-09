package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/services/gateway/api"
)

func main() {
	app := fiber.New()
	brk := broker.NewMessageBroker()

	handler := api.NewAPIHandler(app, brk)
	handler.SetupRoutes()

	log.Fatal(app.Listen(":5175"))
}
