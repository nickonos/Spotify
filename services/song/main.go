package main

import (
	"log"

	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/services/song/api"
	"github.com/nickonos/Spotify/services/song/data"
	"github.com/nickonos/Spotify/services/song/service"
)

func main() {
	brk := broker.NewMessageBroker()
	db, err := data.NewPostgresDatabase()
	if err != nil {
		log.Fatal(err)
	}

	srv := service.NewSongService(db)
	handler := api.NewAPIHandler(srv, brk)

	handler.Subscribe()

	select {}
}
