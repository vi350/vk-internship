package game

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
	ID          int     `db:"id"`
	OwnerID     int64   `db:"owner_id"`
	OpponentID  int64   `db:"opponent_id"`
	WhitePieces []Piece `db:"white_pieces"`
	BlackPieces []Piece `db:"black_pieces"`
	Notation    string  `db:"notation"`
	State       int     `db:"state"`
}
