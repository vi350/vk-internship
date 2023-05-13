package app

import (
	"github.com/joho/godotenv"
	pgClient "github.com/vi350/vk-internship/internal/app/clients/postgres"
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/consumer/event_consumer"
	"github.com/vi350/vk-internship/internal/app/events"
	"github.com/vi350/vk-internship/internal/app/events/telegram"
	gameRegistry "github.com/vi350/vk-internship/internal/app/registry/game"
	imageRegistry "github.com/vi350/vk-internship/internal/app/registry/image"
	userRegistry "github.com/vi350/vk-internship/internal/app/registry/user"
	gameStorage "github.com/vi350/vk-internship/internal/app/storage/game"
	userStorage "github.com/vi350/vk-internship/internal/app/storage/user"
	"os"
	"time"
)

const batchSize = 100

type Bot struct {
	tgcli            *tgClient.Client
	userRegistry     *userRegistry.UserRegistry
	gameRegistry     *gameRegistry.GameRegistry
	tgEventProcessor events.EventProcessor
	consumer         event_consumer.Consumer
}

func New() (*Bot, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	b := &Bot{}
	if b.tgcli, err = tgClient.New(
		os.Getenv("TELEGRAM_HOST"),
		os.Getenv("TELEGRAM_TOKEN"),
		imageRegistry.New()); err != nil {
		return nil, err
	}
	b.userRegistry = userRegistry.New(userStorage.New(pgClient.New()))
	b.gameRegistry = gameRegistry.New(gameStorage.New(pgClient.New()))

	b.tgEventProcessor = telegram.New(b.tgcli, b.userRegistry, b.gameRegistry)
	b.consumer = event_consumer.New(b.tgEventProcessor, b.tgEventProcessor, batchSize)

	return b, nil
}

func (b *Bot) Run() (err error) {

	go func() {
		for {
			b.userRegistry.RemoveInactive(10)
			b.gameRegistry.RemoveInactive(10)
			time.Sleep(time.Minute * 2)
		}
	}()

	go func() {
		for {
			b.userRegistry.Sync()
			b.gameRegistry.Sync()
			time.Sleep(time.Minute)
		}
	}()

	if err = b.consumer.Start(); err != nil {
		return err
	}
	return nil
}
