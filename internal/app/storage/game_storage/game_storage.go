package game_storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/vi350/vk-internship/internal/app/clients/postgres"
	"github.com/vi350/vk-internship/internal/app/e"
)

type GameStorage struct {
	DBClient postgres.Client
}

func New(dbClient *postgres.Client) *GameStorage {
	return &GameStorage{
		DBClient: *dbClient,
	}
}

const (
	saveError         = "error performing game insert"
	readError         = "error performing game read"
	isExistError      = "error performing game isexist"
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

	query := fmt.Sprintf(
		"INSERT INTO games VALUES ('%d', '%d', '%d', '%s', '%s', '%s') RETURNING id;",
		game.ID, game.Owner.ID, game.Opponent.ID, string(whp), string(blp), game.Notation,
	)
	err = gs.DBClient.DB.QueryRow(query).Scan(&game.ID)

	return err
}

func (gs *GameStorage) Read(id int) (game *Game, err error) {
	defer func() { err = e.WrapIfErr(readError, err) }()

	var g Game
	query := fmt.Sprintf(
		"SELECT * FROM games WHERE id = '%d';",
		id,
	)
	err = gs.DBClient.DB.QueryRow(query).Scan(&g.ID, &g.Owner.ID, &g.Opponent.ID, &g.WhitePieces, &g.BlackPieces, &g.Notation)

	return &g, err
}

func (gs *GameStorage) IsExist(id int) (answer bool, err error) {
	defer func() { err = e.WrapIfErr(isExistError, err) }()

	// TODO: reuse read code properly (faced error comparison error)
	// _, err = gs.Read(id)

	var g Game
	query := fmt.Sprintf(
		"SELECT * FROM games WHERE id = '%d';",
		id,
	)
	err = gs.DBClient.DB.QueryRow(query).Scan(&g.ID, &g.Owner.ID, &g.Opponent.ID, &g.WhitePieces, &g.BlackPieces, &g.Notation)

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (gs *GameStorage) FindByUserID(id int64) (game *Game, err error) {
	defer func() { err = e.WrapIfErr(findByUserIDError, err) }()

	query := fmt.Sprintf(
		"SELECT * FROM games WHERE owner = '%d' OR opponent = '%d';",
		id, id,
	)
	err = gs.DBClient.DB.QueryRow(query).Scan(&game.ID, &game.Owner.ID, &game.Opponent.ID, &game.WhitePieces, &game.BlackPieces, &game.Notation)

	return game, err
}

func (gs *GameStorage) Remove(id int) (err error) {
	defer func() { err = e.WrapIfErr(removeError, err) }()

	query := fmt.Sprintf(
		"DELETE FROM games WHERE id = '%d';",
		id,
	)
	_, err = gs.DBClient.DB.Exec(query)

	return err
}
