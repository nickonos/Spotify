package data

import (
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	AddUserRole(ctx context.Context, email string) error
	GetUserRole(ctx context.Context, email string) (string, error)
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
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS user_roles (
		email varchar(255) NOT NULL,
		role varchar(10) NOT NULL DEFAULT 'user',
		PRIMARY KEY (email)
	  );`)
	if err != nil {
		return err
	}

	return nil
}

func (p *MysqlDatabase) AddUserRole(ctx context.Context, email string) error {
	_, err := p.db.ExecContext(ctx, `INSERT INTO user_roles (email) VALUES (?);`, email)
	return err
}

type UserRole struct {
	role string `db:"role"`
}

func (p *MysqlDatabase) GetUserRole(ctx context.Context, email string) (string, error) {
	var user_role UserRole
	err := p.db.QueryRowxContext(ctx, `SELECT role FROM user_roles WHERE email = ?`, email).StructScan(&user_role)
	if err != nil {
		return "", err
	}
	return user_role.role, nil
}
