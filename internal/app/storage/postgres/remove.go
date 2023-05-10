package postgres

import (
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/storage"
)

func (ps *Storage) RemoveGame(game *storage.Game) (err error) {
	defer func() { err = e.WrapIfErr("error performing game removal: ", err) }()

	if _, err = ps.DB.Model(&game).Where("id = ?", game.ID).Delete(); err != nil {
		return err
	}

	return nil
}
