package telegram

import (
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/events"
)

func (ep *EventProcessor) Fetch(limit int) (result []events.Event, err error) {
	defer func() { err = e.WrapIfErr("error in telegram.Processor.Fetch", err) }()

	updates, err := ep.tgClient.Updates(ep.offset, limit)
	if err != nil {
		return nil, err
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))
	for _, update := range updates {
		res = append(res, event(update))
	}

	ep.offset = updates[len(updates)-1].UpdateID + 1

	return res, nil
}

func event(update tgClient.Update) events.Event {
	res := events.Event{
		Type: fetchType(update),
		Text: fetchText(update),
	}
	if update.Message != nil {
		res.Meta = MetaMessage{
			From:     update.Message.From,
			ChatID:   update.Message.Chat.ID,
			Entities: update.Message.Entity,
		}
	} else if update.СallbackQuery != nil {
		res.Meta = MetaCallbackQuery{
			ID:      update.СallbackQuery.ID,
			From:    update.СallbackQuery.From,
			Message: update.СallbackQuery.Message,
		}
	}
	return res
}

func fetchType(update tgClient.Update) events.EventType {
	if update.Message != nil {
		return events.Message
	} else if update.СallbackQuery != nil {
		return events.CallbackQuery
	}
	return events.Unknown
}

func fetchText(update tgClient.Update) string {
	if update.Message != nil {
		return update.Message.Text
	} else if update.СallbackQuery != nil {
		return update.СallbackQuery.Data
	}
	return ""
}
