package telegram

import (
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/localization"
	"log"
	"strings"
)

func (ep *EventProcessor) doCommand(text string, userFromMessage tgClient.User) (err error) {
	text = strings.TrimSpace(text)

	log.Printf("user %s:\n sent: %s", userFromMessage.Username, text)

	userFromRegistry, isNew, err := ep.userRegistry.GetByTgUser(&userFromMessage, text)

	switch {
	case strings.HasPrefix(text, "/start"):
		if isNew {
			err = ep.tgcli.SendTextMessage(userFromMessage.ID,
				localization.GetLocalizedText(localization.StartMessage, userFromMessage.LanguageCode),
				localization.GetLocalizedInlineKeyboardMarkup(localization.StartMessage, userFromMessage.LanguageCode))
		} else {
			err = ep.tgcli.SendTextMessage(userFromMessage.ID,
				localization.GetLocalizedText(localization.StartMessage, userFromRegistry.Language),
				localization.GetLocalizedInlineKeyboardMarkup(localization.StartMessage, userFromRegistry.Language))
		}

	case strings.HasPrefix(text, "/help"):
		err = ep.tgcli.SendTextMessage(userFromMessage.ID,
			localization.GetLocalizedText(localization.HelpMessage, userFromRegistry.Language),
			localization.GetLocalizedInlineKeyboardMarkup(localization.HelpMessage, userFromRegistry.Language))

	default:
		err = ep.tgcli.SendTextMessage(userFromMessage.ID,
			localization.GetLocalizedText(localization.UnknownCommandMessage, userFromRegistry.Language),
			localization.GetLocalizedInlineKeyboardMarkup(localization.UnknownCommandMessage, userFromRegistry.Language))
	}

	return
}
