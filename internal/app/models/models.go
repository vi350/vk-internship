package models

import "time"

type Model interface {
	Model()
}

const (
	// TODO: actually use states or remove them
	MainMenu = iota
	Settings
	ChooseLanguage
	HisTurn
	WaitingTurn
)

type User struct {
	ID             int64  `db:"id"`
	FirstName      string `db:"first_name"`
	Username       string `db:"username"`
	StartDate      int64  `db:"start_date"`
	Language       string `db:"language"`
	State          int    `db:"state"`
	Refer          string `db:"refer"`
	LastAccessTime time.Time
}

const (
	Empty = iota
	WhitePawn
	WitheKing
	BlackPawn
	BlackKing
	GameStateInProgress
	GameStateFinished
)

type Piece struct {
	Type int `json:"type"`
	X    int `json:"x"`
	Y    int `json:"y"`
}

type Game struct {
	ID             int     `db:"id"`
	OwnerID        int64   `db:"owner_id"`
	OpponentID     int64   `db:"opponent_id"`
	WhitePieces    []Piece `db:"white_pieces"`
	BlackPieces    []Piece `db:"black_pieces"`
	Notation       string  `db:"notation"`
	State          int     `db:"state"`
	LastAccessTime time.Time
}
