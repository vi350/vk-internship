package game_storage

import (
	"github.com/vi350/vk-internship/internal/app/storage/user_storage"
)

const (
	Empty = iota
	WhitePawn
	WitheKing
	BlackPawn
	BlackKing
)

type Piece struct {
	Type int `json:"type"`
	X    int `json:"x"`
	Y    int `json:"y"`
}

type Game struct {
	ID          int               `db:"id"`
	Owner       user_storage.User `db:"owner"`
	Opponent    user_storage.User `db:"opponent"`
	WhitePieces []Piece           `db:"white_pieces"`
	BlackPieces []Piece           `db:"black_pieces"`
	Notation    string            `db:"notation"`
}
