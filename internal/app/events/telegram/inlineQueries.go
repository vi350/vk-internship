package telegram

import (
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/localization"
	"log"
	"strconv"
)

func (ep *EventProcessor) doCallbackQuery(query string, userFromCallback tgClient.User, message tgClient.Message) (err error) {
	log.Printf("user %s:\n callback: %s", userFromCallback.Username, query)

	userFromRegistry, _, err := ep.userRegistry.GetByTgUser(&userFromCallback, "")
	q, err := strconv.Atoi(query)
	if err == nil {
		mtq := localization.MessageType(q)

		switch mtq {
		case localization.GoToMenuButton:
			err = ep.tgClient.EditMediaMessageByUser(userFromRegistry, message, localization.MenuMessage)
		case localization.GoToHelpButton:
			err = ep.tgClient.EditMediaMessageByUser(userFromRegistry, message, localization.HelpMessage)
		case localization.GoToSettingsButton:
			err = ep.tgClient.EditMediaMessageByUser(userFromRegistry, message, localization.SettingsMessage)
		case localization.GoToChooseLanguageButton:
			err = ep.tgClient.EditMediaMessageByUser(userFromRegistry, message, localization.ChooseLanguageMessage)
		case localization.GoToAboutButton:
			err = ep.tgClient.EditMediaMessageByUser(userFromRegistry, message, localization.AboutMessage)
		}

	} else {
		//
	}

	return
}
