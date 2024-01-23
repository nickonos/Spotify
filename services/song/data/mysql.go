package data

import (
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nickonos/Spotify/packages/routes"
)

type DB interface {
	AddSong(ctx context.Context, id int64, name string, artist string, cover_url string) error
	GetSong(ctx context.Context, name string) (routes.Song, error)
	GetAllSongs(ctx context.Context) ([]routes.Song, error)
}

type MysqlDatabase struct {
	db *sqlx.DB
}

func NewMysqlDatabase() (DB, error) {
	db_string := os.Getenv("DB_STRING")
	db, err := sqlx.Open("mysql", db_string)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = init_db(db)
	if err != nil {
		return nil, err
	}

	return &MysqlDatabase{
		db,
	}, nil
}

func init_db(db *sqlx.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS songs (
		id bigint NOT NULL,
		name varchar(255),
		PRIMARY KEY (id)
	);`)
	if err != nil {
		return err
	}

	return nil
}

func (p *MysqlDatabase) AddSong(ctx context.Context, id int64, name string, artist string, cover_url string) error {
	_, err := p.db.ExecContext(ctx, `INSERT INTO songs (id, name, artist, cover_url) VALUES (?, ?, ? , ?);`, id, name, artist, cover_url)
	return err
}

func (p *MysqlDatabase) GetSong(ctx context.Context, name string) (routes.Song, error) {
	var song routes.Song
	err := p.db.QueryRowxContext(ctx, `SELECT id, name, artist, cover_url FROM songs WHERE name = ?`, name).StructScan(&song)
	if err != nil {
		return routes.Song{}, err
	}
	return song, nil
}

func (p *MysqlDatabase) GetAllSongs(ctx context.Context) ([]routes.Song, error) {
	var songs []routes.Song
	rows, err := p.db.QueryxContext(ctx, "SELECT id, name FROM songs")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var song routes.Song
		err = rows.StructScan(&song)
		if err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	return songs, nil
}
