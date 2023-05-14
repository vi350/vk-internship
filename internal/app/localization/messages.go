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
	writeToTheCreatorButton  string
	startGameButton          string
	languageButton           string
	languageSet              string
}

const (
	baseImagePath           = "./res/images"
	menuImagePath           = "menu.png"
	helpImagePath           = "help.png"
	settingsImagePath       = "settings.png"
	chooseLanguageImagePath = "chlang.png"
	aboutImagePath          = "about.png"
)

var enMessages = localizedMessages{
	unknownCommandMessage:    "Unknown command. Enter /start to get to menu",
	goToMenuButton:           "🏁Menu",
	menuMessage:              "Hi! This is a bot for checkers game.\nDeveloping is in progress.\nStay tuned!",
	goToHelpButton:           "📕Help",
	helpMessage:              "Checkers, is a group of strategy board games for two players which involve diagonal moves of uniform game pieces and mandatory captures by jumping over opponent pieces. The term \"checkers\" derives from the checkered board which the game is played on.",
	goToSettingsButton:       "🛠Settings",
	settingsMessage:          "Here you can find settings",
	goToChooseLanguageButton: "🌍Choose language",
	chooseLanguageMessage:    "Choose language",
	goToAboutButton:          "🔎About",
	aboutMessage:             "This bot was created during VK team selection",
	writeToTheCreatorButton:  "Write to the creator",
	startGameButton:          "♟️Start game",
	languageButton:           "🇺🇸English",
	languageSet:              "English language set",
}

var ruMessages = localizedMessages{
	unknownCommandMessage:    "Неизвестная команда. Введите /start чтобы перейти в меню",
	goToMenuButton:           "🏁Меню",
	menuMessage:              "Привет! Это бот для игры в шашки.\nРазработка в процессе.\nСледите за обновлениями!",
	goToHelpButton:           "📕Помощь",
	helpMessage:              "Шашки - это группа стратегических настольных игр для двух игроков, которые предполагают диагональные ходы одинаковых игровых фигур и обязательные захваты путем перепрыгивания через фигуры противника. Термин \"шашки\" происходит от клетчатой доски, на которой ведется игра.",
	goToSettingsButton:       "🛠Настройки",
	settingsMessage:          "Здесь ты можешь найти настройки",
	goToChooseLanguageButton: "🌍Выбрать язык",
	chooseLanguageMessage:    "Выбери язык",
	goToAboutButton:          "🔎О боте",
	aboutMessage:             "Это бот, разработанный в рамках отбора в команду ВК",
	writeToTheCreatorButton:  "Написать создателю",
	startGameButton:          "♟️Начать игру",
	languageButton:           "🇷🇺Русский",
	languageSet:              "Установлен русский язык",
}
