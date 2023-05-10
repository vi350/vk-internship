package app

import (
	"github.com/joho/godotenv"
	pgClient "github.com/vi350/vk-internship/internal/app/clients/postgres"
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/consumer/event_consumer"
	"github.com/vi350/vk-internship/internal/app/events"
	"github.com/vi350/vk-internship/internal/app/events/telegram"
	"github.com/vi350/vk-internship/internal/app/storage/game_storage"
	"github.com/vi350/vk-internship/internal/app/storage/user_storage"
	"os"
)

const batchSize = 100

type Bot struct {
	tgcli            *tgClient.Client
	userStorage      *user_storage.UserStorage
	gameStorage      *game_storage.GameStorage
	tgEventProcessor events.EventProcessor
	consumer         event_consumer.Consumer
}

func New() (*Bot, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	b := &Bot{}
	if b.tgcli, err = tgClient.New(os.Getenv("TELEGRAM_HOST"), os.Getenv("TELEGRAM_TOKEN")); err != nil {
		return nil, err
	}
	b.userStorage = user_storage.New(pgClient.New())
	b.gameStorage = game_storage.New(pgClient.New())

	b.tgEventProcessor = telegram.New(b.tgcli, b.userStorage, b.gameStorage)
	b.consumer = event_consumer.New(b.tgEventProcessor, b.tgEventProcessor, batchSize)

	return b, nil
}

func (b *Bot) Run() (err error) {
	if err = b.consumer.Start(); err != nil {
		return err
	}
	return nil
}
