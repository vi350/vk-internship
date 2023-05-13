package user_storage

const (
	MainMenu = iota
	Settings
	ChooseLanguage
	HisTurn
	WaitingTurn
)

type User struct {
	ID        int64  `db:"id"`
	FirstName string `db:"first_name"`
	Username  string `db:"username"`
	StartDate int64  `db:"start_date"`
	Language  string `db:"language"`
	State     int    `db:"state"`
	Refer     string `db:"refer"`
}
