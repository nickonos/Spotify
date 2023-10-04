package api

import (
	"github.com/nickonos/Spotify/packages/broker"
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

}
