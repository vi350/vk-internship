package localization

import (
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
	WriteToTheCreatorButton
	StartGameButton
	LanguageButton
	LanguageSet
)

const (
	EnglishLanguage = "en"
	RussianLanguage = "ru"
)

func GetLocalizedText(mType MessageType, language string, variables ...string) (answer string) {
	// idea: move to db/config?localization.GetLocalizedText(localization.Language, "en")
	// idea: link localizedMessages fields to constants?

	switch mType {
	case UnknownCommandMessage:
		switch language {
		case EnglishLanguage:
			answer = enMessages.unknownCommandMessage
		case RussianLanguage:
			answer = ruMessages.unknownCommandMessage
		}
	case GoToMenuButton:
		switch language {
		case EnglishLanguage:
			answer = enMessages.goToMenuButton
		case RussianLanguage:
			answer = ruMessages.goToMenuButton

		}
	case MenuMessage:
		switch language {
		case EnglishLanguage:
			answer = enMessages.menuMessage
		case RussianLanguage:
			answer = ruMessages.menuMessage
		}
	case GoToHelpButton:
		switch language {
		case EnglishLanguage:
			answer = enMessages.goToHelpButton
		case RussianLanguage:
			answer = ruMessages.goToHelpButton
		}
	case HelpMessage:
		switch language {
		case EnglishLanguage:
			answer = enMessages.helpMessage
		case RussianLanguage:
			answer = ruMessages.helpMessage
		}
	case GoToSettingsButton:
		switch language {
		case EnglishLanguage:
			answer = enMessages.goToSettingsButton
		case RussianLanguage:
			answer = ruMessages.goToSettingsButton
		}
	case SettingsMessage:
		switch language {
		case EnglishLanguage:
			answer = enMessages.settingsMessage
		case RussianLanguage:
			answer = ruMessages.settingsMessage
		}
	case GoToChooseLanguageButton:
		switch language {
		case EnglishLanguage:
			answer = enMessages.goToChooseLanguageButton
		case RussianLanguage:
			answer = ruMessages.goToChooseLanguageButton
		}
	case ChooseLanguageMessage:
		switch language {
		case EnglishLanguage:
			answer = enMessages.chooseLanguageMessage
		case RussianLanguage:
			answer = ruMessages.chooseLanguageMessage
		}
	case GoToAboutButton:
		switch language {
		case EnglishLanguage:
			answer = enMessages.goToAboutButton
		case RussianLanguage:
			answer = ruMessages.goToAboutButton
		}
	case AboutMessage:
		switch language {
		case EnglishLanguage:
			answer = enMessages.aboutMessage
		case RussianLanguage:
			answer = ruMessages.aboutMessage
		}
	case WriteToTheCreatorButton:
		switch language {
		case EnglishLanguage:
			answer = enMessages.writeToTheCreatorButton
		case RussianLanguage:
			answer = ruMessages.writeToTheCreatorButton
		}
	case StartGameButton:
		switch language {
		case EnglishLanguage:
			answer = enMessages.startGameButton
		case RussianLanguage:
			answer = ruMessages.startGameButton
		}
	case LanguageButton:
		switch language {
		case EnglishLanguage:
			answer = enMessages.languageButton
		case RussianLanguage:
			answer = ruMessages.languageButton
		}
	case LanguageSet:
		switch language {
		case EnglishLanguage:
			answer = enMessages.languageSet
		case RussianLanguage:
			answer = ruMessages.languageSet
		}
	default:
		answer = "Error: Unknown message type"
	}

	return
}

func GetLocalizedImagePath(mType MessageType, language string, imageRegistry *imageRegistry.ImageRegistry, variables ...string) (image string) {
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
