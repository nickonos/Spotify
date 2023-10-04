package service

import (
	"github.com/nickonos/Spotify/packages/logging"
	"github.com/nickonos/Spotify/services/song/data"
)

type SongService struct {
	logger *logging.Logger
	db     *data.PostgressDatabase
}

func NewSongService(db *data.PostgressDatabase) *SongService {
	logger := logging.NewLogger("song")

	logger.Trace("Song Service Started")

	return &SongService{
		logger,
		db,
	}
}
