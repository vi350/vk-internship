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
	Language
)

func GetLocalizedText(mType MessageType, language string, variables ...string) (answer string) {
	// idea: move to db/config?localization.GetLocalizedText(localization.Language, "en")
	// idea: link localizedMessages fields to constants?

	switch mType {
	case UnknownCommandMessage:
		switch language {
		case enMessages.language:
			answer = enMessages.unknownCommandMessage
		case ruMessages.language:
			answer = ruMessages.unknownCommandMessage
		}
	case GoToMenuButton:
		switch language {
		case enMessages.language:
			answer = enMessages.goToMenuButton
		case ruMessages.language:
			answer = ruMessages.goToMenuButton

		}
	case MenuMessage:
		switch language {
		case enMessages.language:
			answer = enMessages.menuMessage
		case ruMessages.language:
			answer = ruMessages.menuMessage
		}
	case GoToHelpButton:
		switch language {
		case enMessages.language:
			answer = enMessages.goToHelpButton
		case ruMessages.language:
			answer = ruMessages.goToHelpButton
		}
	case HelpMessage:
		switch language {
		case enMessages.language:
			answer = enMessages.helpMessage
		case ruMessages.language:
			answer = ruMessages.helpMessage
		}
	case GoToSettingsButton:
		switch language {
		case enMessages.language:
			answer = enMessages.goToSettingsButton
		case ruMessages.language:
			answer = ruMessages.goToSettingsButton
		}
	case SettingsMessage:
		switch language {
		case enMessages.language:
			answer = enMessages.settingsMessage
		case ruMessages.language:
			answer = ruMessages.settingsMessage
		}
	case GoToChooseLanguageButton:
		switch language {
		case enMessages.language:
			answer = enMessages.goToChooseLanguageButton
		case ruMessages.language:
			answer = ruMessages.goToChooseLanguageButton
		}
	case ChooseLanguageMessage:
		switch language {
		case enMessages.language:
			answer = enMessages.chooseLanguageMessage
		case ruMessages.language:
			answer = ruMessages.chooseLanguageMessage
		}
	case GoToAboutButton:
		switch language {
		case enMessages.language:
			answer = enMessages.goToAboutButton
		case ruMessages.language:
			answer = ruMessages.goToAboutButton
		}
	case AboutMessage:
		switch language {
		case enMessages.language:
			answer = enMessages.aboutMessage
		case ruMessages.language:
			answer = ruMessages.aboutMessage
		}
	case WriteToTheCreatorButton:
		switch language {
		case enMessages.language:
			answer = enMessages.writeToTheCreatorButton
		case ruMessages.language:
			answer = ruMessages.writeToTheCreatorButton
		}
	case StartGameButton:
		switch language {
		case enMessages.language:
			answer = enMessages.startGameButton
		case ruMessages.language:
			answer = ruMessages.startGameButton
		}
	case LanguageButton:
		switch language {
		case enMessages.language:
			answer = enMessages.languageButton
		case ruMessages.language:
			answer = ruMessages.languageButton
		}
	case Language:
		switch language {
		case enMessages.language:
			answer = enMessages.language
		case ruMessages.language:
			answer = ruMessages.language
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
