package service

import (
	"reflect"

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

	var temp routes.RequestCreateSong

	logger.Trace(reflect.TypeOf(temp).Name())

	return &SongService{
		logger,
		db,
		id,
	}
}

func (s *SongService) CreateSong(name string) error {
	id, err := s.id.NewID()
	if err != nil {
		return err
	}

	return s.db.AddSong(id, name)
}
