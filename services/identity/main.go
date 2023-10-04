package main

import (
	"log"

	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/services/identity/api"
	"github.com/nickonos/Spotify/services/identity/service"
)

func main() {
	brk := broker.NewMessageBroker()

	svc, err := service.NewIdentiyService(brk)
	if err != nil {
		log.Fatal(err)
	}

	handler := api.NewAPIHandler(svc, brk)

	handler.Subscribe()

	// Wait forever
	select {}
}
