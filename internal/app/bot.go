package app

import (
	"github.com/joho/godotenv"
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/events"
	"github.com/vi350/vk-internship/internal/app/events/telegram"
	"github.com/vi350/vk-internship/internal/app/storage"
	"github.com/vi350/vk-internship/internal/app/storage/postgres"
	"github.com/vi350/vk-internship/internal/consumer/event_consumer"
	"os"
)

const batchSize = 100

type Bot struct {
	tgcli            *tgClient.Client
	storage          storage.Storage
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
	if b.storage, err = postgres.New(); err != nil {
		return nil, err
	}
	b.tgEventProcessor = telegram.New(b.tgcli, b.storage)
	b.consumer = event_consumer.New(b.tgEventProcessor, b.tgEventProcessor, batchSize)

	return b, nil
}

func (b *Bot) Run() (err error) {
	if err = b.consumer.Start(); err != nil {
		return err
	}
	return nil
}
