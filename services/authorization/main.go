package main

import (
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/logging"
	"github.com/nickonos/Spotify/services/authorization/api"
)

func main() {
	logger := logging.NewLogger("Authorization")

	logger.Print("Connecting with message broker")
	brk := broker.NewMessageBroker()
	logger.Print("Connected with message broker")
	handler := api.NewAPIHandler(*brk)
	handler.Subscribe()

	logger.Print("Started service")

	select {}
}
