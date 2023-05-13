package telegram

import (
	"github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/registry/game"
	"github.com/vi350/vk-internship/internal/registry/user"
)

type EventProcessor struct {
	tgcli        *telegram.Client
	offset       int64
	userRegistry *user.UserRegistry
	gameRegistry *game.GameRegistry
}

type MetaMessage struct {
	From     telegram.User
	ChatID   int64
	Entities []telegram.Entity
}

type MetaInlineQuery struct {
	FromID int64
	Query  string
}

func New(tgcli *telegram.Client, userRegistry *user.UserRegistry, gameRegistry *game.GameRegistry) *EventProcessor {
	return &EventProcessor{
		tgcli:        tgcli,
		userRegistry: userRegistry,
		gameRegistry: gameRegistry,
	}
}
