package telegram

import (
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/localization"
	"github.com/vi350/vk-internship/internal/app/models"
	"log"
	"strings"
)

func (ep *EventProcessor) doCommand(text string, userFromMessage *tgClient.User) (err error) {
	text = strings.TrimSpace(text)

	log.Printf("user %s:\n sent: %s", userFromMessage.Username, text)

	var userFromRegistry *models.User
	var isNew bool
	if userFromRegistry, isNew, err = ep.userRegistry.GetByTgUser(userFromMessage, text); err != nil {
		return err
	}

	switch {
	case strings.HasPrefix(text, "/start"):
		if isNew {
			err = ep.tgClient.SendImageByUser(userFromRegistry, localization.ChooseLanguageMessage)
		} else {
			err = ep.tgClient.SendImageByUser(userFromRegistry, localization.MenuMessage)
		}

	default:
		err = ep.tgClient.SendTextMessageByUser(userFromRegistry, localization.UnknownCommandMessage)
	}

	return
}
