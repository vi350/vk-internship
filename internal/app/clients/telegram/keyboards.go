package telegram

import (
	"github.com/vi350/vk-internship/internal/app/localization"
	"os"
	"strconv"
)

func GetLocalizedInlineKeyboardMarkup(mType localization.MessageType, language string, variables ...string) (ikm ReplyMarkup) {
	switch mType {
	case localization.MenuMessage:
		ikm = &InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.StartGameButton, language),
						CallbackData: strconv.Itoa(int(localization.StartGameButton)),
					},
				},
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.GoToHelpButton, language),
						CallbackData: strconv.Itoa(int(localization.GoToHelpButton)),
					},
				},
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.GoToSettingsButton, language),
						CallbackData: strconv.Itoa(int(localization.GoToSettingsButton)),
					},
				},
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.GoToAboutButton, language),
						CallbackData: strconv.Itoa(int(localization.GoToAboutButton)),
					},
				},
			},
		}
	case localization.HelpMessage:
		ikm = &InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.GoToMenuButton, language),
						CallbackData: strconv.Itoa(int(localization.GoToMenuButton)),
					},
				},
			},
		}
	case localization.SettingsMessage:
		ikm = &InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.GoToChooseLanguageButton, language),
						CallbackData: strconv.Itoa(int(localization.GoToChooseLanguageButton)),
					},
				},
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.GoToMenuButton, language),
						CallbackData: strconv.Itoa(int(localization.GoToMenuButton)),
					},
				},
			},
		}
	case localization.ChooseLanguageMessage:
		ikm = &InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.LanguageButton, "en"),
						CallbackData: localization.GetLocalizedText(localization.Language, "en"),
					},
				},
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.LanguageButton, "ru"),
						CallbackData: localization.GetLocalizedText(localization.Language, "ru"),
					},
				},
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.GoToMenuButton, language),
						CallbackData: strconv.Itoa(int(localization.GoToMenuButton)),
					},
				},
			},
		}
	case localization.AboutMessage:
		ikm = &InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					InlineKeyboardButton{
						Text: localization.GetLocalizedText(localization.WriteToTheCreatorButton, language),
						Url:  "tg://user?id=" + os.Getenv("ADMIN_ID"),
					},
				},
				{
					InlineKeyboardButton{
						Text:         localization.GetLocalizedText(localization.GoToMenuButton, language),
						CallbackData: strconv.Itoa(int(localization.GoToMenuButton)),
					},
				},
			},
		}
	}

	return
}
