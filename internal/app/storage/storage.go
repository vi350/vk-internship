package storage

import "image"

const (
	Empty = iota
	WhitePawn
	WitheKing
	BlackPawn
	BlackKing
)

const (
	MainMenu = iota
	Settings
	Gaming
)

// TODO: one more layer for storing in cache?
// TODO: split user and game storage?

type Storage interface {
	SaveUser(*User) error
	ReadUser(id int64) (*User, error)
	IsExistUser(id int64) (bool, error)
	SetStateUser(id int64, state int) error
	SaveGame(*Game) error
	ReadGame(id int) (*Game, error)
	IsExistGame(id int) (bool, error)
	FindGameByUserID(id int64) (*Game, error)
	RemoveGame(game *Game) error
}

type User struct {
	ID        int64  `db:"id"`
	FirstName string `db:"first_name"`
	Username  string `db:"username"`
	StartDate int64  `db:"start_date"`
	Language  string `db:"language"`
	State     int    `db:"state"`
	Refer     string `db:"refer"`
}

type Piece struct {
	Type     int
	Position image.Point
}

type Game struct {
	ID          int     `db:"id"`
	Owner       User    `db:"owner"`
	Opponent    User    `db:"opponent"`
	WhitePieces []Piece `db:"white_pieces"`
	BlackPieces []Piece `db:"black_pieces"`
	Notation    string  `db:"notation"`
}
