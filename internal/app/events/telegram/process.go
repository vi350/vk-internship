package telegram

import (
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/events"
)

func (ep *EventProcessor) Process(event events.Event) (err error) {
	switch event.Type {
	case events.Message:
		return ep.processMessage(event)
	case events.CallbackQuery:
		return ep.processCallbackQuery(event)
	default:
		return e.WrapIfErr("unknown event type", err)
	}
}

func (ep *EventProcessor) processMessage(event events.Event) (err error) {
	meta, err := metaMessage(event)
	if err != nil {
		return e.WrapIfErr("failed to get meta", err)
	}

	if meta.From.ID == meta.ChatID {
		var isCommand bool
		for _, entity := range meta.Entities {
			if entity.Type == "bot_command" {
				isCommand = true
			}
		}
		if isCommand {
			if err = ep.doCommand(event.Text, &meta.From); err != nil {
				return
			}
		}
	}
	return
}

func (ep *EventProcessor) processCallbackQuery(event events.Event) (err error) {
	meta, err := metaCallbackQuery(event)
	if err != nil {
		return e.WrapIfErr("unknown meta type", err)
	}

	if err = ep.doCallbackQuery(meta.ID, event.Text, &meta.From, &meta.Message); err != nil {
		return
	}
	return
}

func metaMessage(event events.Event) (meta MetaMessage, err error) {
	res, ok := event.Meta.(MetaMessage)
	if !ok {
		return MetaMessage{}, e.WrapIfErr("unknown meta type", err)
	}
	return res, nil
}

func metaCallbackQuery(event events.Event) (meta MetaCallbackQuery, err error) {
	res, ok := event.Meta.(MetaCallbackQuery)
	if !ok {
		return MetaCallbackQuery{}, e.WrapIfErr("unknown meta type", err)
	}
	return res, nil
}
