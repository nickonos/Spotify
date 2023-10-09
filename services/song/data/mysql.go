package data

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nickonos/Spotify/packages/routes"
)

type DB interface {
	AddSong(id int64, name string) error
	GetSong(name string) (routes.Song, error)
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
		id bigint,
		name varchar(255)
	);`)
	if err != nil {
		return err
	}

	return nil
}

func (p *MysqlDatabase) AddSong(id int64, name string) error {
	_, err := p.db.Exec(`INSERT INTO songs (id, name) VALUES (?, ?);`, id, name)
	return err
}

func (p *MysqlDatabase) GetSong(name string) (routes.Song, error) {
	var song routes.Song
	err := p.db.QueryRowx(`SELECT id, name FROM songs WHERE name = ?`, name).StructScan(&song)
	if err != nil {
		return routes.Song{}, err
	}
	return song, nil
}
