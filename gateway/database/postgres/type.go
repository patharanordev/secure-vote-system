package database

import (
	"database/sql"
)

type AccountInfoProps struct {
	Username string
	Password string
	IsAdmin  bool
}

type AccountProps struct {
	UID  []uint8
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
	SetDB(db *sql.DB)
	CreateAccount(usr string, pwd string, isAdmin bool) ([]uint8, error)
	GetAccount(usr string, pwd string) (*AccountProps, error)
	GetAccountByID(uid string) (*AccountProps, error)
	UpdateAccount(uid string, usr string, isAdmin bool) error
	DeleteAccountByID(uid string) error
}
