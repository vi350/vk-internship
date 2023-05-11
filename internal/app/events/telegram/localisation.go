package telegram

import tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"

type messageType int

const (
	unknownCommandMessage messageType = iota
	startMessage
	helpMessage
	settingsMessage
	ChooseLanguageMessage
)

func GetLocalizedText(mType messageType, language string, variables ...interface{}) string {
	// idea: move to db/config?
	// TODO: link localizedMessages fields to constants?

	switch mType {
	case unknownCommandMessage:
		switch language {
		case "en":
			return enMessages.unknownCommandMessage
		case "ru":
			return ruMessages.unknownCommandMessage
		}
	case startMessage:
		switch language {
		case "en":
			return enMessages.startMessage
		case "ru":
			return ruMessages.startMessage
		}
	case helpMessage:
		switch language {
		case "en":
			return enMessages.helpMessage
		case "ru":
			return ruMessages.helpMessage
		}
	case settingsMessage:
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
	}

	return ""
}
