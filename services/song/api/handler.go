package api

import (
	"context"

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
	broker.Subscribe(api.broker, func(ctx context.Context, message routes.CreateSongRequest) (routes.CreateSongResponse, error) {
		song, err := api.service.CreateSong(ctx, message.Name)
		if err != nil {
			return routes.CreateSongResponse{}, err
		}

		return routes.CreateSongResponse{
			Song: song,
		}, nil
	})

	broker.Subscribe(api.broker, func(ctx context.Context, message routes.GetSongRequest) (routes.GetSongResponse, error) {
		song, err := api.service.GetSong(ctx, message.Name)
		if err != nil {
			return routes.GetSongResponse{}, err
		}

		return routes.GetSongResponse{
			Song: song,
		}, nil
	})

	broker.Subscribe(api.broker, func(ctx context.Context, message routes.GetSongsRequest) (routes.GetSongsResponse, error) {
		songs, err := api.service.GetAllSongs(ctx)
		if err != nil {
			return routes.GetSongsResponse{}, err
		}

		return routes.GetSongsResponse{
			Songs: songs,
		}, nil
	})
}
