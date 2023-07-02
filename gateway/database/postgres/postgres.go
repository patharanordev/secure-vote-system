package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Initial(conn PGConnProps) IDatabase {
	// Initial
	p := &PGProps{}
	p.conn.DB_HOST = conn.DB_HOST
	p.conn.DB_PORT = conn.DB_PORT
	p.conn.DB_NAME = conn.DB_NAME
	p.conn.DB_PASSWORD = conn.DB_PASSWORD
	p.conn.DB_USER = conn.DB_USER

	return p
}

func (p *PGProps) Close() {
	defer p.db.Close()
}

func (p *PGProps) Connect() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.conn.DB_HOST,
		p.conn.DB_PORT,
		p.conn.DB_USER,
		p.conn.DB_PASSWORD,
		p.conn.DB_NAME,
	)
	db, err := sql.Open("postgres", dbinfo)
	p.db = db

	return db, err
}

func (p *PGProps) CreateAccount(usr string, pwd string) ([]uint8, error) {
	fmt.Println("Creating account...")

	var lastInsertId []uint8
	err := p.db.QueryRow(
		"INSERT INTO user_info(username, password) VALUES( $1, crypt($2, gen_salt('bf')) ) returning uid;",
		usr, pwd,
	).Scan(&lastInsertId)

	return lastInsertId, err
}

func (p *PGProps) GetAccount(usr string) AccountProps {

	// TODO: Adding select data
	// rows, err := db.Query("SELECT * FROM user_info")
	// if err != nil

	return AccountProps{}
}

func (p *PGProps) UpdateAccount(usr string, pwd string, props AccountInfoProps) error {

	// TODO: Update data
	// stmt, err := db.Prepare("update user_info set username=$1 where uid=$2")
	// if err != nil

	return nil
}

func (p *PGProps) DeleteAccount(usr string, pwd string) error {
	// TODO: Delete data

	return nil
}
