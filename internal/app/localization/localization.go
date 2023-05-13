package localization

import (
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	imageRegistry "github.com/vi350/vk-internship/internal/app/registry/image"
	"os"
	"path"
	"strconv"
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
	WriteToAdminButton
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
	case WriteToAdminButton:
		switch language {
		case "en":
			answer = enMessages.writeToAdminButton
		case "ru":
			answer = ruMessages.writeToAdminButton
		}
	default:
		answer = "Error: Unknown message type"
	}

	return
}

func GetLocalizedInlineKeyboardMarkup(mType MessageType, language string, variables ...interface{}) (ikm tgClient.ReplyMarkup) {
	switch mType {
	case MenuMessage:
		ikm = &tgClient.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgClient.InlineKeyboardButton{
				//{
				//	tgClient.InlineKeyboardButton{
				//		Text:         GetLocalizedText(StartGameButton, language),
				//		CallbackData: "startGame",
				//	},
				//},
				{
					tgClient.InlineKeyboardButton{
						Text:         GetLocalizedText(GoToHelpButton, language),
						CallbackData: strconv.Itoa(int(GoToHelpButton)),
					},
				},
				{
					tgClient.InlineKeyboardButton{
						Text:         GetLocalizedText(GoToSettingsButton, language),
						CallbackData: strconv.Itoa(int(GoToSettingsButton)),
					},
				},
				{
					tgClient.InlineKeyboardButton{
						Text:         GetLocalizedText(GoToAboutButton, language),
						CallbackData: strconv.Itoa(int(GoToAboutButton)),
					},
				},
			},
		}
	case HelpMessage:
		ikm = &tgClient.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgClient.InlineKeyboardButton{
				{
					tgClient.InlineKeyboardButton{
						Text:         GetLocalizedText(GoToMenuButton, language),
						CallbackData: strconv.Itoa(int(GoToMenuButton)),
					},
				},
			},
		}
	case SettingsMessage:
		ikm = &tgClient.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgClient.InlineKeyboardButton{
				{
					tgClient.InlineKeyboardButton{
						Text:         GetLocalizedText(GoToChooseLanguageButton, language),
						CallbackData: strconv.Itoa(int(GoToChooseLanguageButton)),
					},
				},
				{
					tgClient.InlineKeyboardButton{
						Text:         GetLocalizedText(GoToMenuButton, language),
						CallbackData: strconv.Itoa(int(GoToMenuButton)),
					},
				},
			},
		}
	case ChooseLanguageMessage:
		ikm = &tgClient.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgClient.InlineKeyboardButton{
				{
					tgClient.InlineKeyboardButton{
						Text:         "🇺🇸English🇺🇸",
						CallbackData: "en",
					},
				},
				{
					tgClient.InlineKeyboardButton{
						Text:         "🇷🇺Русский🇷🇺",
						CallbackData: "ru",
					},
				},
				{
					tgClient.InlineKeyboardButton{
						Text:         GetLocalizedText(GoToMenuButton, language),
						CallbackData: strconv.Itoa(int(GoToMenuButton)),
					},
				},
			},
		}
	case AboutMessage:
		ikm = &tgClient.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgClient.InlineKeyboardButton{
				{
					tgClient.InlineKeyboardButton{
						Text: GetLocalizedText(WriteToAdminButton, language),
						Url:  "tg://user?id=" + os.Getenv("ADMIN_ID"),
					},
				},
				{
					tgClient.InlineKeyboardButton{
						Text:         GetLocalizedText(GoToMenuButton, language),
						CallbackData: strconv.Itoa(int(GoToMenuButton)),
					},
				},
			},
		}
	}

	return
}

func GetLocalizedImagePath(mType MessageType, language string, imageRegistry *imageRegistry.ImageRegistry, variables ...interface{}) (image string) {
	switch mType {
	case MenuMessage:
		image = menuImagePath
	case HelpMessage:
		image = helpImagePath
	case SettingsMessage:
		image = settingsImagePath
	case ChooseLanguageMessage:
		image = chooseLanguageImagePath
	case AboutMessage:
		image = aboutImagePath
	}
	image = imageRegistry.GetByPath(path.Join(baseImagePath, language, image))

	return
}
