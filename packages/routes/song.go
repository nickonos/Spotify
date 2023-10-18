package routes

type Song struct {
	Name string `db:"name"`
	Id   int64  `db:"id"`
}

type CreateSongResponse struct {
	Song Song
}

type CreateSongRequest struct {
	Name string
}

func (CreateSongRequest) Subject() string {
	return "create_song"
}

type GetSongResponse struct {
	Song Song
}

type GetSongRequest struct {
	Name string
}

func (GetSongRequest) Subject() string {
	return "get_song"
}

type GetSongsRequest struct {
}

type GetSongsResponse struct {
	Songs []Song
}

func (GetSongsRequest) Subject() string {
	return "get_songs"
}
