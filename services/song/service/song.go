package service

import (
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

func (s *SongService) CreateSong(name string) (routes.Song, error) {
	id, err := s.id.NewID()
	if err != nil {
		return routes.Song{}, err
	}

	err = s.db.AddSong(id, name)
	if err != nil {
		return routes.Song{}, err
	}

	return routes.Song{
		Id:   id,
		Name: name,
	}, nil
}

func (s *SongService) GetSong(name string) (routes.Song, error) {
	return s.db.GetSong(name)
}
