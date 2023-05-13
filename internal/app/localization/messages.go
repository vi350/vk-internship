package localization

type localizedMessages struct {
	unknownCommandMessage    string
	goToMenuButton           string
	menuMessage              string
	goToHelpButton           string
	helpMessage              string
	goToSettingsButton       string
	settingsMessage          string
	goToChooseLanguageButton string
	chooseLanguageMessage    string
	goToAboutButton          string
	aboutMessage             string
}

const (
	baseImagePath           = "./media/images"
	menuImagePath           = "menu.png"
	helpImagePath           = "help.png"
	settingsImagePath       = "settings.png"
	chooseLanguageImagePath = "chlang.png"
	aboutImagePath          = "about.png"
)

var enMessages = localizedMessages{
	unknownCommandMessage:    "Unknown command. Enter /help to learn more",
	goToMenuButton:           "Menu",
	menuMessage:              "Hello! I'm a bot that will help you ...",
	goToHelpButton:           "Help",
	helpMessage:              "I can help you ...",
	goToSettingsButton:       "Settings",
	settingsMessage:          "Here you can find settings",
	goToChooseLanguageButton: "Choose language",
	chooseLanguageMessage:    "Choose language",
	goToAboutButton:          "About",
	aboutMessage:             "About bot",
}

var ruMessages = localizedMessages{
	unknownCommandMessage:    "Неизвестная команда. Введите /help, чтобы узнать больше",
	goToMenuButton:           "Меню",
	menuMessage:              "Привет! Я бот, который поможет тебе ...",
	goToHelpButton:           "Помощь",
	helpMessage:              "Я могу помочь тебе ...",
	goToSettingsButton:       "Настройки",
	settingsMessage:          "Здесь ты можешь найти настройки",
	goToChooseLanguageButton: "Выбрать язык",
	chooseLanguageMessage:    "Выбери язык",
	aboutMessage:             "О боте",
}
