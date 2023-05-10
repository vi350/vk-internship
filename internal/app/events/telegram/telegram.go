package telegram

import (
	"github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/storage"
)

type EventProcessor struct {
	tgcli   *telegram.Client
	offset  int64
	storage storage.Storage
}

type MetaMessage struct {
	From     telegram.User
	ChatID   int64
	Entities []telegram.Entity
}

type MetaInlineQuery struct {
	FromID int64
}

func New(tgcli *telegram.Client, storage storage.Storage) *EventProcessor {
	return &EventProcessor{
		tgcli:   tgcli,
		storage: storage,
	}
}
