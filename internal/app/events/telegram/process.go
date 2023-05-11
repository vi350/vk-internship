package telegram

import (
	"fmt"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/events"
)

func (ep *EventProcessor) Process(event events.Event) (err error) {
	switch event.Type {
	case events.Message:
		return ep.processMessage(event)
	case events.InlineQuery:
		return ep.processInlineQuery(event)
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
		// TODO: before processing as a command check entities
		if err = ep.doCommand(event.Text, meta.From); err != nil {
			return err
		}
	}
	return nil
}

func (ep *EventProcessor) processInlineQuery(event events.Event) (err error) {
	meta, err := metaInlineQuery(event)
	if err != nil {
		return e.WrapIfErr("unknown meta type", err)
	}

	// TODO: implement processing inline queries
	fmt.Println(meta.FromID)
	fmt.Println(event.Text)
	return nil
}

func metaMessage(event events.Event) (meta MetaMessage, err error) {
	res, ok := event.Meta.(MetaMessage)
	if !ok {
		return MetaMessage{}, e.WrapIfErr("unknown meta type", err)
	}
	return res, nil
}

func metaInlineQuery(event events.Event) (meta MetaInlineQuery, err error) {
	res, ok := event.Meta.(MetaInlineQuery)
	if !ok {
		return MetaInlineQuery{}, e.WrapIfErr("unknown meta type", err)
	}
	return res, nil
}
