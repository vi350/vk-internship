package telegram

import (
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/localization"
	"log"
	"strconv"
)

func (ep *EventProcessor) doCallbackQuery(queryId string, query string, userFromCallback *tgClient.User, message *tgClient.Message) (err error) {
	log.Printf("user %s:\n callback: %s", userFromCallback.Username, query)
	notification := ""
	defer func() { err = ep.answerCallbackQuery(queryId, notification) }()

	userFromRegistry, _, err := ep.userRegistry.GetByTgUser(userFromCallback, "")
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
		case localization.StartGameButton:
			err = ep.tgClient.SendTextMessage(userFromRegistry.ID, "🚧🚧🚧🚧🚧", nil)
		}

	} else {
		switch query {
		case localization.EnglishLanguage:
			userFromRegistry.Language = localization.EnglishLanguage
			notification = localization.GetLocalizedText(localization.LanguageSet, userFromRegistry.Language)
			err = ep.tgClient.EditMediaMessageByUser(userFromRegistry, message, localization.MenuMessage)
		case localization.RussianLanguage:
			userFromRegistry.Language = localization.RussianLanguage
			notification = localization.GetLocalizedText(localization.LanguageSet, userFromRegistry.Language)
			err = ep.tgClient.EditMediaMessageByUser(userFromRegistry, message, localization.MenuMessage)
		}
	}

	return
}

func (ep *EventProcessor) answerCallbackQuery(queryID string, notification string) (err error) {
	return ep.tgClient.AnswerCallbackQuery(queryID, notification)
}
