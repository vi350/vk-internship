package game_storage

import (
	"encoding/json"
	"github.com/vi350/vk-internship/internal/app/clients"
	"github.com/vi350/vk-internship/internal/app/e"
)

type GameStorage struct {
	DBClient clients.DBClient
}

func New(dbClient clients.DBClient) *GameStorage {
	return &GameStorage{
		DBClient: dbClient,
	}
}

const (
	saveError         = "error performing game insert"
	readError         = "error performing game read"
	findByUserIDError = "error performing game findbyuserid"
	removeError       = "error performing game remove"
)

func (gs *GameStorage) Save(game *Game) (err error) {
	defer func() { err = e.WrapIfErr(saveError, err) }()

	whp, err := json.Marshal(game.WhitePieces)
	if err != nil {
		return err
	}
	blp, err := json.Marshal(game.BlackPieces)
	if err != nil {
		return err
	}

	query := `INSERT INTO games VALUES (?, ?, ?, ?, ?, ?)`
	_, err = gs.DBClient.DB().
		Exec(query, game.ID, game.Owner.ID, game.Opponent.ID, string(whp), string(blp), game.Notation)

	return err
}

func (gs *GameStorage) Read(id int) (game *Game, err error) {
	defer func() { err = e.WrapIfErr(readError, err) }()

	query := `SELECT * FROM games WHERE id = ?`
	err = gs.DBClient.DB().
		QueryRow(query, id).
		Scan(&game.ID, &game.Owner.ID, &game.Opponent.ID, &game.WhitePieces, &game.BlackPieces, &game.Notation)

	return game, err
}

func (gs *GameStorage) FindByUserID(id int64) (game *Game, err error) {
	defer func() { err = e.WrapIfErr(findByUserIDError, err) }()

	query := `SELECT * FROM games WHERE owner = ? OR opponent = ?`
	err = gs.DBClient.DB().
		QueryRow(query, id, id).
		Scan(&game.ID, &game.Owner.ID, &game.Opponent.ID, &game.WhitePieces, &game.BlackPieces, &game.Notation)

	return game, err
}

func (gs *GameStorage) Remove(id int) (err error) {
	defer func() { err = e.WrapIfErr(removeError, err) }()

	query := `DELETE FROM games WHERE id = ?`
	_, err = gs.DBClient.DB().
		Exec(query, id)

	return err
}
