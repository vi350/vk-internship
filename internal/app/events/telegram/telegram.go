package telegram

import (
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/registry/game"
	"github.com/vi350/vk-internship/internal/app/registry/user"
)

type EventProcessor struct {
	tgClient     *tgClient.Client
	offset       int64
	userRegistry *user.UserRegistry
	gameRegistry *game.GameRegistry
}

type MetaMessage struct {
	From     tgClient.User
	ChatID   int64
	Entities []tgClient.MessageEntity
}

type MetaCallbackQuery struct {
	ID      string
	From    tgClient.User
	Message tgClient.Message
}

func New(tgcli *tgClient.Client, userRegistry *user.UserRegistry, gameRegistry *game.GameRegistry) *EventProcessor {
	return &EventProcessor{
		tgClient:     tgcli,
		userRegistry: userRegistry,
		gameRegistry: gameRegistry,
	}
}
