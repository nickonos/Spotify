package data

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgressDatabase struct {
	db *sqlx.DB
}

func NewPostgresDatabase() (*PostgressDatabase, error) {
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
	_, err := db.Exec(`SELECT 'CREATE DATABASE song' 
	WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'song')`)
	if err != nil {
		return err
	}

	return nil
}
