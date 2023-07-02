package database

import (
	"database/sql"
)

type AccountInfoProps struct {
	Password string
	IsAdmin  bool
}

type AccountProps struct {
	User string
	Info AccountInfoProps
}

type PGConnProps struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

type PGProps struct {
	db   *sql.DB
	conn PGConnProps
}

type IDatabase interface {
	Connect() (*sql.DB, error)
	Close()
	CreateAccount(usr string, pwd string) ([]uint8, error)
	GetAccount(usr string) AccountProps
	UpdateAccount(usr string, pwd string, props AccountInfoProps) error
	DeleteAccount(usr string, pwd string) error
}
