package localization

type localizedMessages struct {
	unknownCommandMessage string
	startMessage          string
	helpMessage           string
	settingsMessage       string
	ChooseLanguageMessage string
	GoToMainMenuButton    string
}

var enMessages = localizedMessages{
	unknownCommandMessage: "Unknown command. Enter /help to see all commands.",
	startMessage:          "Hello! I'm a bot that will help you ...",
	helpMessage:           "I can help you ...",
	settingsMessage:       "Settings",
	ChooseLanguageMessage: "Choose language",
	GoToMainMenuButton:    "Go to main menu",
}

var ruMessages = localizedMessages{
	unknownCommandMessage: "Неизвестная команда. Введите /help, чтобы увидеть все команды.",
	startMessage:          "Привет! Я бот, который поможет тебе ...",
	helpMessage:           "Я могу помочь тебе ...",
	settingsMessage:       "Settings",
	ChooseLanguageMessage: "Выберите язык",
	GoToMainMenuButton:    "Вернуться в главное меню",
}