package data

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB interface {
	AddSong(id int64, name string) error
}

type PostgressDatabase struct {
	db *sqlx.DB
}

func NewPostgresDatabase() (DB, error) {
	db_string := os.Getenv("DB_STRING")
	db, err := sqlx.Open("postgres", db_string)
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

	return &PostgressDatabase{
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

func (p *PostgressDatabase) AddSong(id int64, name string) error {
	_, err := p.db.Exec(`INSERT INTO songs (id, name) VALUES (? , ?)`, id, name)
	return err
}
