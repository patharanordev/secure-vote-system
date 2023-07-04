package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

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
	p.SetDB(db)

	return db, err
}

func (p *PGProps) SetDB(db *sql.DB) {
	p.db = db
}

func (p *PGProps) CreateAccount(usr string, pwd string, isAdmin bool) ([]uint8, error) {
	fmt.Println("Creating account...")

	var lastInsertId []uint8
	err := p.db.QueryRow(`
		INSERT INTO user_info(username, password, is_admin) 
		VALUES( $1, crypt($2, gen_salt('bf')), $3 ) 
		returning uid;
		`, usr, pwd, isAdmin,
	).Scan(&lastInsertId)

	return lastInsertId, err
}

func (p *PGProps) GetAccount(usr string, pwd string) (*AccountProps, error) {

	var account AccountProps

	qStr := fmt.Sprintf(`
		SELECT uid, username, password, is_admin 
		FROM user_info 
		WHERE username = '%s' 
		AND password = crypt('%s', password)
		`, usr, pwd,
	)

	rows, err := p.db.Query(qStr)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	accounts := []AccountProps{}
	for rows.Next() {
		var aProps AccountProps
		if errScan := rows.Scan(
			&aProps.UID,
			&aProps.Info.Username,
			&aProps.Info.Password,
			&aProps.Info.IsAdmin,
		); errScan != nil {
			return nil, errScan
		}

		accounts = append(accounts, aProps)
	}

	if len(accounts) <= 0 {
		return nil, errors.New("Unauthorized")
	}

	account = accounts[0]
	return &account, nil
}

func (p *PGProps) GetAccountByID(uid string) (*AccountProps, error) {

	qStr := fmt.Sprintf(`
		SELECT uid, username, password, is_admin 
		FROM user_info 
		WHERE uid = '%s'`, uid)

	rows, err := p.db.Query(qStr)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	accounts := []AccountProps{}
	for rows.Next() {
		var account AccountProps
		if errScan := rows.Scan(
			&account.UID,
			&account.Info.Username,
			&account.Info.Password,
			&account.Info.IsAdmin,
		); errScan != nil {
			return nil, errScan
		}

		accounts = append(accounts, account)
	}

	if len(accounts) <= 0 {
		errNotFound := errors.New("Account not found.")
		return nil, errNotFound
	}

	account := accounts[0]
	return &account, nil
}

func (p *PGProps) UpdateAccount(uid string, usr string, isAdmin bool) error {

	result, err := p.db.Exec(`
	UPDATE user_info 
	SET username=$1, is_admin=$2, updated_at=$3 
	WHERE uid=$4
	`, usr, isAdmin, time.Now(), uid)
	if err != nil {
		return err
	}

	fmt.Printf("update result: %v\n", result)

	return nil
}

func (p *PGProps) DeleteAccountByID(uid string) error {

	result, err := p.db.Exec(`
	DELETE FROM user_info 
	WHERE uid=$1
	`, uid)
	if err != nil {
		return err
	}

	fmt.Printf("delete result: %v\n", result)

	return nil
}
