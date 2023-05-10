package event_consumer

import (
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/events"
	"log"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("failed to fetch events: %v\n", err)
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
		}

		if err = c.handleEvents(gotEvents); err != nil {
			log.Printf("failed to handle events: %v\n", err)
			continue
		}
	}
}

// TODO: add wait group
func (c Consumer) handleEvents(handlingEvents []events.Event) error {
	for _, event := range handlingEvents {
		log.Printf("handling event: %v\n", event)
		if err := c.processor.Process(event); err != nil {
			return e.WrapIfErr("failed to process event: ", err)
			// continue (if other "if" statements used)
		}
	}
	return nil
}
