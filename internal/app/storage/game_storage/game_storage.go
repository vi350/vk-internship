package game_storage

import (
	"database/sql"
	"encoding/json"
	"github.com/vi350/vk-internship/internal/app/clients"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/storage"
)

type GameStorage struct {
	DBClient clients.DBClient
}

func New(dbClient clients.DBClient) *GameStorage {
	return &GameStorage{
		DBClient: dbClient,
	}
}

func (gs *GameStorage) Insert(game *Game) (err error) {
	defer func() { err = e.WrapIfErr(storage.InsertError, err) }()

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
		Exec(query, game.ID, game.OwnerID, game.OpponentID, string(whp), string(blp), game.Notation)

	return
}

func (gs *GameStorage) Read(id int) (game *Game, err error) {
	defer func() { err = e.WrapIfErr(storage.ReadError, err) }()

	query := `SELECT * FROM games WHERE id = ?`
	err = gs.DBClient.DB().
		QueryRow(query, id).
		Scan(&game.ID, &game.OwnerID, &game.OpponentID, &game.WhitePieces, &game.BlackPieces, &game.Notation)

	return
}

func (gs *GameStorage) FindUsersActiveGame(userid int64) (game *Game, err error) {
	defer func() { err = e.WrapIfErr(storage.FindError, err) }()

	query := `SELECT * FROM games WHERE (owner_id = ? OR opponent_id = ?) AND state = ?`
	err = gs.DBClient.DB().
		QueryRow(query, userid, userid, GameStateInProgress).
		Scan(&game.ID, &game.OwnerID, &game.OpponentID, &game.WhitePieces, &game.BlackPieces, &game.Notation)

	return
}

func (gs *GameStorage) FindUsersGames(userid int64) (games []*Game, err error) {
	defer func() { err = e.WrapIfErr(storage.FindError, err) }()

	query := `SELECT * FROM games WHERE owner_id = ? OR opponent_id = ?`
	rows, err := gs.DBClient.DB().
		Query(query, userid, userid)
	for rows.Next() {
		var game *Game
		err = rows.Scan(&game.ID, &game.OwnerID, &game.OpponentID, &game.WhitePieces, &game.BlackPieces, &game.Notation, &game.State)
		games = append(games, game)
	}

	return
}

func (gs *GameStorage) UpdateWithMap(gamesToUpdate map[int]*Game) (err error) {
	defer func() { err = e.WrapIfErr(storage.UpdateWithMapError, err) }()

	var tx *sql.Tx
	if tx, err = gs.DBClient.DB().Begin(); err != nil {
		return e.WrapIfErr(storage.BeginError, err)
	}

	defer func() {
		if err != nil {
			if err = tx.Rollback(); err != nil {
				err = e.WrapIfErr(storage.RollbackError, err)
			}
			return
		}
		err = tx.Commit()
		if err != nil {
			err = e.WrapIfErr(storage.CommitError, err)
		}
	}()

	query := `UPDATE games SET owner_id = ?, opponent_id = ?, white_pieces = ?, black_pieces = ?, notation = ?, state = ? WHERE id = ?`
	statement, err := tx.Prepare(query)
	if err != nil {
		return e.WrapIfErr(storage.PrepareError, err)
	}
	defer func() {
		if err = statement.Close(); err != nil {
			err = e.WrapIfErr(storage.CloseStatementError, err)
		}
	}()

	for id, game := range gamesToUpdate {
		_, err = statement.Exec(game.OwnerID, game.OpponentID, game.WhitePieces, game.BlackPieces, game.Notation, game.State, id)
		if err != nil {
			return e.WrapIfErr(storage.ExecError, err)
		}
	}

	return
}
