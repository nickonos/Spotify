package main

import (
	"log"

	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/identity"
	"github.com/nickonos/Spotify/services/song/api"
	"github.com/nickonos/Spotify/services/song/data"
	"github.com/nickonos/Spotify/services/song/service"
)

func main() {
	brk := broker.NewMessageBroker()
	db, err := data.NewMysqlDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	id := identity.NewIdentityHelper(brk)
	srv := service.NewSongService(db, id)
	handler := api.NewAPIHandler(srv, brk)

	handler.Subscribe()

	select {}
}
