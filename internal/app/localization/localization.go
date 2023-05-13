package localization

import (
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	imageRegistry "github.com/vi350/vk-internship/internal/app/registry/image"
	"path"
)

type MessageType int

const (
	UnknownCommandMessage MessageType = iota
	GoToMenuButton
	MenuMessage
	GoToHelpButton
	HelpMessage
	GoToSettingsButton
	SettingsMessage
	GoToChooseLanguageButton
	ChooseLanguageMessage
	GoToAboutButton
	AboutMessage
)

func GetLocalizedText(mType MessageType, language string, variables ...interface{}) (answer string) {
	// idea: move to db/config?
	// idea: link localizedMessages fields to constants?

	switch mType {
	case UnknownCommandMessage:
		switch language {
		case "en":
			answer = enMessages.unknownCommandMessage
		case "ru":
			answer = ruMessages.unknownCommandMessage
		}
	case GoToMenuButton:
		switch language {
		case "en":
			answer = enMessages.goToMenuButton
		case "ru":
			answer = ruMessages.goToMenuButton

		}
	case MenuMessage:
		switch language {
		case "en":
			answer = enMessages.menuMessage
		case "ru":
			answer = ruMessages.menuMessage
		}
	case GoToHelpButton:
		switch language {
		case "en":
			answer = enMessages.goToHelpButton
		case "ru":
			answer = ruMessages.goToHelpButton
		}
	case HelpMessage:
		switch language {
		case "en":
			answer = enMessages.helpMessage
		case "ru":
			answer = ruMessages.helpMessage
		}
	case GoToSettingsButton:
		switch language {
		case "en":
			answer = enMessages.goToSettingsButton
		case "ru":
			answer = ruMessages.goToSettingsButton
		}
	case SettingsMessage:
		switch language {
		case "en":
			answer = enMessages.settingsMessage
		case "ru":
			answer = ruMessages.settingsMessage
		}
	case GoToChooseLanguageButton:
		switch language {
		case "en":
			answer = enMessages.goToChooseLanguageButton
		case "ru":
			answer = ruMessages.goToChooseLanguageButton
		}
	case ChooseLanguageMessage:
		switch language {
		case "en":
			answer = enMessages.chooseLanguageMessage
		case "ru":
			answer = ruMessages.chooseLanguageMessage
		}
	case GoToAboutButton:
		switch language {
		case "en":
			answer = enMessages.goToAboutButton
		case "ru":
			answer = ruMessages.goToAboutButton
		}
	case AboutMessage:
		switch language {
		case "en":
			answer = enMessages.aboutMessage
		case "ru":
			answer = ruMessages.aboutMessage
		}
	default:
		answer = "Error: Unknown message type"
	}

	return
}

func GetLocalizedInlineKeyboardMarkup(mType MessageType, language string, variables ...interface{}) tgClient.ReplyMarkup {
	switch mType {
	case UnknownCommandMessage:
		return &tgClient.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgClient.InlineKeyboardButton{
				{tgClient.InlineKeyboardButton{
					Text:         GetLocalizedText(GoToMenuButton, language),
					CallbackData: "main_menu", // TODO: move to callback constants
				}},
			},
		}
	case MenuMessage:
		return &tgClient.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgClient.InlineKeyboardButton{
				{tgClient.InlineKeyboardButton{
					Text:         GetLocalizedText(HelpMessage, language),
					CallbackData: "help",
				}},
			},
		}
	}

	return nil
}

func GetLocalizedImagePath(mType MessageType, language string, imageRegistry *imageRegistry.ImageRegistry, variables ...interface{}) (image string) {
	switch mType {
	case MenuMessage:
		switch language {
		case "en":
			image = enMessages.menuMessage
		case "ru":
			image = ruMessages.menuMessage
		}
	case HelpMessage:
		switch language {
		case "en":
			image = enMessages.helpMessage
		case "ru":
			image = ruMessages.helpMessage
		}
	case SettingsMessage:
		switch language {
		case "en":
			image = enMessages.settingsMessage
		case "ru":
			image = ruMessages.settingsMessage
		}
	case ChooseLanguageMessage:
		switch language {
		case "en":
			image = enMessages.chooseLanguageMessage
		case "ru":
			image = ruMessages.chooseLanguageMessage
		}
	case AboutMessage:
		switch language {
		case "en":
			image = enMessages.aboutMessage
		case "ru":
			image = ruMessages.aboutMessage
		}
	}
	image = imageRegistry.GetByPath(path.Join(baseImagePath, language, image))

	return
}
