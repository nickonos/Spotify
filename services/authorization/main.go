package main

import (
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/logging"
	"github.com/nickonos/Spotify/services/authorization/api"
)

func main() {
	logger := logging.NewLogger("Authorization")

	brk := broker.NewMessageBroker()
	handler, err := api.NewAPIHandler(*brk)
	if err != nil {
		// fail if we could not start service 
		logger.Fatal(err)
	}

	handler.Subscribe()

	logger.Print("Started service")

	select {}
}
