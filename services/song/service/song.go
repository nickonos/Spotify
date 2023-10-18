package service

import (
	"context"

	"github.com/nickonos/Spotify/packages/identity"
	"github.com/nickonos/Spotify/packages/logging"
	"github.com/nickonos/Spotify/packages/routes"
	"github.com/nickonos/Spotify/services/song/data"
)

type SongService struct {
	logger *logging.Logger
	db     data.DB
	id     identity.IdentityHelper
}

func NewSongService(db data.DB, id identity.IdentityHelper) *SongService {
	logger := logging.NewLogger("song")

	logger.Trace("Song Service Started")

	return &SongService{
		logger,
		db,
		id,
	}
}

func (s *SongService) CreateSong(ctx context.Context, name string) (routes.Song, error) {
	id, err := s.id.NewID()
	if err != nil {
		return routes.Song{}, err
	}

	err = s.db.AddSong(ctx, id, name)
	if err != nil {
		return routes.Song{}, err
	}

	return routes.Song{
		Id:   id,
		Name: name,
	}, nil
}

func (s *SongService) GetSong(ctx context.Context, name string) (routes.Song, error) {
	return s.db.GetSong(ctx, name)
}

func (s *SongService) GetAllSongs(ctx context.Context) ([]routes.Song, error) {
	s.logger.Print("Get all songs")
	songs, err := s.db.GetAllSongs(ctx)

	s.logger.Print("%v", songs)
	return songs, err
}
