package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nickonos/Spotify/packages/broker"
)

type API struct {
	app    *fiber.App
	broker *broker.Broker
}

func NewAPIHandler(a *fiber.App, brk *broker.Broker) *API {
	return &API{
		app:    a,
		broker: brk,
	}
}

func (api *API) SetupRoutes() {
	router := api.app.Group("/api")

	// Song Routes
	router.Get("/song", api.GetSong)
	router.Get("/songs", api.GetAllSongs)
	router.Post("/song", api.CreateSong)
	router.Get("/login", api.Login)
}
