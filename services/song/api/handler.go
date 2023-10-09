package api

import (
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
	broker.Subscribe(api.broker, func(message routes.CreateSongRequest) (routes.CreateSongResponse, error) {
		song, err := api.service.CreateSong(message.Name)
		if err != nil {
			return routes.CreateSongResponse{}, err
		}

		return routes.CreateSongResponse{
			Song: song,
		}, nil
	})

	broker.Subscribe(api.broker, func(message routes.GetSongRequest) (routes.GetSongResponse, error) {
		song, err := api.service.GetSong(message.Name)

		if err != nil {
			return routes.GetSongResponse{}, err
		}

		return routes.GetSongResponse{
			Song: song,
		}, nil

	})
}
