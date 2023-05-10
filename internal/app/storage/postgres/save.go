package postgres

import (
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/storage"
)

func (ps *Storage) SaveUser(user *storage.User) (err error) {
	defer func() { err = e.WrapIfErr("error performing user insert: ", err) }()
	_, err = ps.DB.Model(&user).Insert()
	return err
}

func (ps *Storage) SetStateUser(ID int64, state int) (err error) {
	defer func() { err = e.WrapIfErr("error performing user state update: ", err) }()
	_, err = ps.DB.Model(&storage.User{ID: ID, State: state}).Column("state").WherePK().Update()
	return err
}

func (ps *Storage) SaveGame(game *storage.Game) (err error) {
	defer func() { err = e.WrapIfErr("error performing game insert: ", err) }()
	_, err = ps.DB.Model(&game).Insert()
	return err
}
