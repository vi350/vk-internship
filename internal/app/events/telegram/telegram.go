package telegram

import (
	"github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/storage/game_storage"
	"github.com/vi350/vk-internship/internal/app/storage/user_storage"
)

type EventProcessor struct {
	tgcli       *telegram.Client
	offset      int64
	userStorage *user_storage.UserStorage
	gameStorage *game_storage.GameStorage
}

type MetaMessage struct {
	From     telegram.User
	ChatID   int64
	Entities []telegram.Entity
}

type MetaInlineQuery struct {
	FromID int64
}

func New(tgcli *telegram.Client, userStorage *user_storage.UserStorage, gameStorage *game_storage.GameStorage) *EventProcessor {
	return &EventProcessor{
		tgcli:       tgcli,
		userStorage: userStorage,
		gameStorage: gameStorage,
	}
}
