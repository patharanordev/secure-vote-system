package database

import (
	"context"
	"database/sql"
	"errors"
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
	p.SetDB(db)

	return db, err
}

func (p *PGProps) SetDB(db *sql.DB) {
	p.db = db
}

// -------------------- item --------------------

func (p *PGProps) CreateVoteItem(uid string, payload *CreateVoteItemPayload) ([]uint8, error) {
	fmt.Println("Creating vote item...")

	var lastInsertId []uint8
	err := p.db.QueryRow(`
		INSERT INTO vote(uid, item_name, item_description) 
		VALUES( $1, $2, $3 ) 
		returning vid;
		`, uid, payload.Name, payload.Description,
	).Scan(&lastInsertId)

	return lastInsertId, err
}

func (p *PGProps) GetVoteItemByID(uid string, vid string) (*VoteItemProps, error) {

	qStr := fmt.Sprintf(`
		SELECT vid, item_name, item_description, vote_count 
		FROM vote 
		WHERE vid = '%s'`, vid)

	rows, err := p.db.Query(qStr)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	voteItems := []VoteItemProps{}
	for rows.Next() {
		var voteItem VoteItemProps
		if errScan := rows.Scan(
			&voteItem.VID,
			&voteItem.Info.Name,
			&voteItem.Info.Description,
			&voteItem.Info.VoteCount,
		); errScan != nil {
			return nil, errScan
		}

		voteItems = append(voteItems, voteItem)
	}

	voteItem := voteItems[0]

	return &voteItem, nil
}

func (p *PGProps) UpdateVoteItemByID(uid string, item *EditVoteItemPayload) error {

	result, err := p.db.Exec(`
	UPDATE vote 
	SET item_name=$1, item_description=$2, updated_at=NOW() 
	WHERE uid=$3 
	AND vid=$4;
	`, item.Name,
		item.Description,
		uid,
		item.ID,
	)

	if err != nil {
		return err
	}

	fmt.Printf("update result: %v\n", result)

	return nil
}

func (p *PGProps) UpVote(uid string, item *VotingPayload) error {

	ctx := context.Background()

	tx, errTx := p.db.BeginTx(ctx, nil)
	if errTx != nil {
		return errTx
	}

	// Check ever vote or not
	countRow := 0
	rows, _ := tx.QueryContext(ctx, `
		SELECT * 
		FROM vote_history 
		WHERE vid=$1 AND uid=$2
	`, item.ID, uid)
	defer rows.Close()
	for rows.Next() {
		countRow++
	}
	if countRow > 0 {
		return errors.New("You already vote this item.")
	}

	// try to update
	_, errUpdate := tx.ExecContext(ctx, `
		UPDATE vote 
		SET vote_count= vote_count+1, updated_at=NOW() 
		WHERE vid=$1;
	`, item.ID)

	if errUpdate != nil {
		tx.Rollback()
		return errUpdate
	}

	// add history
	_, errHistory := tx.ExecContext(ctx, `
		INSERT INTO vote_history(vid, uid) 
		VALUES( $1, $2);
	`, item.ID, uid)

	if errHistory != nil {
		tx.Rollback()
		return errHistory
	}

	// commit all commands
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *PGProps) DownVote(uid string, item *VotingPayload) error {

	ctx := context.Background()

	tx, errTx := p.db.BeginTx(ctx, nil)
	if errTx != nil {
		return errTx
	}

	// Check ever vote or not
	countRow := 0
	rows, _ := tx.QueryContext(ctx, `
		SELECT * 
		FROM vote_history 
		WHERE vid=$1 AND uid=$2
	`, item.ID, uid)
	defer rows.Close()
	for rows.Next() {
		countRow++
	}
	if countRow == 0 {
		return errors.New("You never vote this item.")
	}

	// try to update
	_, errUpdate := tx.ExecContext(ctx, `
		UPDATE vote 
		SET vote_count= vote_count-1, updated_at=NOW() 
		WHERE vid=$1;
	`, item.ID)

	if errUpdate != nil {
		tx.Rollback()
		return errUpdate
	}

	// add history
	_, errHistory := tx.ExecContext(ctx, `
		DELETE FROM vote_history 
		WHERE vid=$1 AND uid=$2;
	`, item.ID, uid)

	if errHistory != nil {
		tx.Rollback()
		return errHistory
	}

	// commit all commands
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *PGProps) DeleteVoteItemByID(uid string, payload *VoteItemIDPayload) error {

	result, err := p.db.Exec(`
	DELETE FROM vote 
	WHERE vid=$1
	`, payload.VID)
	if err != nil {
		return err
	}

	fmt.Printf("delete result: %v\n", result)

	return nil
}

// -------------------- list --------------------

func (p *PGProps) GetVoteList() ([]VoteItemPayload, error) {

	qStr := fmt.Sprintf(`
		SELECT vid, uid, item_name, item_description, vote_count 
		FROM vote 
		ORDER BY vote_count DESC
	`)

	rows, err := p.db.Query(qStr)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	voteItems := []VoteItemPayload{}
	for rows.Next() {
		var voteItem VoteItemPayload
		if errScan := rows.Scan(
			&voteItem.ID,
			&voteItem.UserID,
			&voteItem.Name,
			&voteItem.Description,
			&voteItem.VoteCount,
		); errScan != nil {
			return nil, errScan
		}

		voteItems = append(voteItems, voteItem)
	}

	return voteItems, nil
}

func (p *PGProps) DeleteVoteList() error {

	result, err := p.db.Exec(`DELETE FROM vote`)
	if err != nil {
		return err
	}

	fmt.Printf("delete result: %v\n", result)

	return nil
}
