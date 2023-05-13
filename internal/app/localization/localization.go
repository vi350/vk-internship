package localization

import tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"

type MessageType int

const (
	UnknownCommandMessage MessageType = iota
	StartMessage
	HelpMessage
	SettingsMessage
	ChooseLanguageMessage
	GoToMainMenuButton
)

func GetLocalizedText(mType MessageType, language string, variables ...interface{}) string {
	// idea: move to db/config?
	// TODO: link localizedMessages fields to constants?

	switch mType {
	case UnknownCommandMessage:
		switch language {
		case "en":
			return enMessages.unknownCommandMessage
		case "ru":
			return ruMessages.unknownCommandMessage
		}
	case StartMessage:
		switch language {
		case "en":
			return enMessages.startMessage
		case "ru":
			return ruMessages.startMessage
		}
	case HelpMessage:
		switch language {
		case "en":
			return enMessages.helpMessage
		case "ru":
			return ruMessages.helpMessage
		}
	case SettingsMessage:
		switch language {
		case "en":
			return enMessages.settingsMessage
		case "ru":
			return ruMessages.settingsMessage
		}
	case ChooseLanguageMessage:
		switch language {
		case "en":
			return enMessages.ChooseLanguageMessage
		case "ru":
			return ruMessages.ChooseLanguageMessage
		}
	case GoToMainMenuButton:
		switch language {
		case "en":
			return enMessages.GoToMainMenuButton
		case "ru":
			return ruMessages.GoToMainMenuButton

		}
	}

	return "."
}

func GetLocalizedInlineKeyboardMarkup(mType MessageType, language string, variables ...interface{}) tgClient.ReplyMarkup {
	switch mType {
	case UnknownCommandMessage:
		return &tgClient.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgClient.InlineKeyboardButton{
				[]tgClient.InlineKeyboardButton{tgClient.InlineKeyboardButton{
					Text:         GetLocalizedText(GoToMainMenuButton, language),
					CallbackData: "main_menu", // TODO: move to callback constants
				}},
			},
		}
	}

	return nil
}
