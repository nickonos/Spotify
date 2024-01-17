package main

import (
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/logging"
	"github.com/nickonos/Spotify/services/authorization/api"
)

func main() {
	logger := logging.NewLogger("Authorization")

	brk := broker.NewMessageBroker()
	handler := api.NewAPIHandler(*brk)
	handler.Subscribe()

	logger.Print("Started service")

	select {}
}
