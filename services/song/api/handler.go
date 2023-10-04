package api

import (
	"github.com/nats-io/nats.go"
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/routes"
	"github.com/nickonos/Spotify/services/song/service"
)

type APIHandler struct {
	service *service.SongService
	broker  *broker.Broker
}

func NewAPIHandler(s *service.SongService, b *broker.Broker) APIHandler {
	return APIHandler{
		service: s,
		broker:  b,
	}
}

func (api *APIHandler) Subscribe() {
	broker.Subscribe[routes.CreateSong](api.broker, func(message routes.CreateSongRequest, raw *nats.Msg) (routes.CreateSongResponse, error) {
		return routes.CreateSongResponse{}, nil
	})
}
