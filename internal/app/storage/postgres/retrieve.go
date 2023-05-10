package postgres

import (
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/storage"
)

// todo: move to pure sql or fix models+orm

func (ps *Storage) ReadUser(ID int64) (user *storage.User, err error) {
	defer func() { err = e.WrapIfErr("error performing user read: ", err) }()

	if err = ps.DB.Model(&user).Where("id = ?", ID).Limit(1).Select(); err != nil {
		return nil, err
	}

	return user, nil
}

func (ps *Storage) IsExistUser(id int64) (answer bool, err error) {
	defer func() { err = e.WrapIfErr("error performing user isexist: ", err) }()
	user, err := ps.ReadUser(id)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return true, nil
}

func (ps *Storage) ReadGame(id int) (game *storage.Game, err error) {
	defer func() { err = e.WrapIfErr("error performing game read: ", err) }()

	if err = ps.DB.Model(&game).Where("id = ?", id).Limit(1).Select(); err != nil {
		return nil, err
	}

	return game, nil
}

func (ps *Storage) IsExistGame(id int) (answer bool, err error) {
	defer func() { err = e.WrapIfErr("error performing game isexist: ", err) }()
	game, err := ps.ReadGame(id)
	if err != nil {
		return false, err
	}
	if game == nil {
		return false, nil
	}
	return true, nil
}

func (ps *Storage) FindGameByUserID(id int64) (game *storage.Game, err error) {
	defer func() { err = e.WrapIfErr("error while finding game: ", err) }()

	if err = ps.DB.Model(&game).Where("owner_id = ?", id).Limit(1).Select(); err != nil {
		return nil, err
	}
	if err = ps.DB.Model(&game).Where("opponent_id = ?", id).Limit(1).Select(); err != nil {
		return nil, err
	}

	return game, nil
}
