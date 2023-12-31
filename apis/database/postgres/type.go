package database

import (
	"database/sql"
)

type VoteItemInfoProps struct {
	Name        string `json:"itemName" db:"item_name"`
	Description string `json:"itemDescription" db:"item_description"`
	VoteCount   int    `json:"voteCount" db:"vote_count"`
}

type VoteItemProps struct {
	VID  string            `json:"vid"`
	Info VoteItemInfoProps `json:"info"`
}

type CreateVoteItemPayload struct {
	Name        string `json:"itemName"`
	Description string `json:"itemDescription"`
}

type EditVoteItemPayload struct {
	ID          string `json:"id"`
	Name        string `json:"itemName"`
	Description string `json:"itemDescription"`
}

type VoteItemPayload struct {
	ID          string `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"itemName"`
	Description string `json:"itemDescription"`
	VoteCount   int    `json:"voteCount"`
}

type VotingPayload struct {
	ID   string `json:"id"`
	IsUp bool   `json:"isUp"`
}

type VoteItemIDPayload struct {
	VID string `json:"id"`
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

	// Vote item
	CreateVoteItem(uid string, payload *CreateVoteItemPayload) ([]uint8, error)
	GetVoteItemByID(uid string, vid string) (*VoteItemProps, error)
	UpdateVoteItemByID(uid string, item *EditVoteItemPayload) error
	UpVote(uid string, item *VotingPayload) error
	DownVote(uid string, item *VotingPayload) error
	DeleteVoteItemByID(uid string, payload *VoteItemIDPayload) error

	// Vote list
	GetVoteList() ([]VoteItemPayload, error)
	DeleteVoteList() error
}
