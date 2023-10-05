package routes

import "github.com/nickonos/Spotify/packages/broker"

type RequestCreateSong struct {
	name string
}

type CreateSongResponseData struct {
	name string
	id   int64
}

type CreateSongResponse broker.Response[CreateSongResponseData]

type CreateSongRequest struct {
	name string
}

type CreateSong struct {
	response struct {
		err  string
		data CreateSongResponseData
	}
	request CreateSongRequest
}
