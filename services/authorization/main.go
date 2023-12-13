package main

import (
	"os"

	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/logging"
	"github.com/nickonos/Spotify/services/authorization/api"
)

func main() {
	logger := logging.NewLogger("Authorization")

	logger.Print(os.Getenv("NATS_URL"))
	logger.Print("Hello")
	logger.Print(os.Getenv("SPOTIFY_CLIENT_ID"))
	brk := broker.NewMessageBroker()
	handler := api.NewAPIHandler(*brk)
	handler.Subscribe()

	select {}
}
