package telegram

import (
	"github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/storage/game_storage"
	"github.com/vi350/vk-internship/internal/registry/user"
)

type EventProcessor struct {
	tgcli        *telegram.Client
	offset       int64
	userRegistry *user.UserRegistry
	gameStorage  *game_storage.GameStorage
}

type MetaMessage struct {
	From     telegram.User
	ChatID   int64
	Entities []telegram.Entity
}

type MetaInlineQuery struct {
	FromID int64
}

func New(tgcli *telegram.Client, userRegistry *user.UserRegistry, gameStorage *game_storage.GameStorage) *EventProcessor {
	return &EventProcessor{
		tgcli:        tgcli,
		userRegistry: userRegistry,
		gameStorage:  gameStorage,
	}
}
