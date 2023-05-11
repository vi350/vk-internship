package telegram

import (
	"github.com/vi350/vk-internship/internal/app/e"
)

type messageType int

const (
	unknownCommandMessage messageType = iota
	startMessage
	helpMessage
	settingsMessage
	ChooseLanguageMessage
)

func GetLocalizedText(mType messageType, language string, variables ...interface{}) (answer string, err error) {
	defer func() { err = e.WrapIfErr("error localising", err) }()

	// idea: move to db/config?
	// TODO: link localizedMessages fields to constants?

	switch mType {
	case unknownCommandMessage:
		switch language {
		case "en":
			return enMessages.unknownCommandMessage, nil
		case "ru":
			return ruMessages.unknownCommandMessage, nil
		}
	case startMessage:
		switch language {
		case "en":
			return enMessages.startMessage, nil
		case "ru":
			return ruMessages.startMessage, nil
		}
	case helpMessage:
		switch language {
		case "en":
			return enMessages.helpMessage, nil
		case "ru":
			return ruMessages.helpMessage, nil
		}
	case settingsMessage:
		switch language {
		case "en":
			return enMessages.settingsMessage, nil
		case "ru":
			return ruMessages.settingsMessage, nil
		}
	case ChooseLanguageMessage:
		switch language {
		case "en":
			return enMessages.ChooseLanguageMessage, nil
		case "ru":
			return ruMessages.ChooseLanguageMessage, nil
		}
	}

	return "", err
}
