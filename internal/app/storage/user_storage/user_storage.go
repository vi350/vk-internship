package user_storage

import (
	"github.com/vi350/vk-internship/internal/app/clients"
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/e"
	"time"
)

type UserStorage struct {
	DBClient clients.DBClient
}

func New(dbClient clients.DBClient) *UserStorage {
	return &UserStorage{
		DBClient: dbClient,
	}
}

const (
	saveError     = "error performing user insert"
	readError     = "error performing user read"
	setStateError = "error performing user setstate"
)

func (us *UserStorage) Save(user *User) (err error) {
	defer func() { err = e.WrapIfErr(saveError, err) }()

	query := `INSERT INTO users VALUES (?, ?, ?, ?, ?, ?, ?)`
	// TODO: move to ExecContext
	_, err = us.DBClient.DB().
		Exec(query, user.ID, user.FirstName, user.Username, user.StartDate, user.Language, user.State, user.Refer)
	return err
}

func (us *UserStorage) SaveFromTg(user *tgClient.User, text string) (err error) {
	defer func() { err = e.WrapIfErr(saveError, err) }()

	if user.LanguageCode == "" {
		user.LanguageCode = "en"
	}
	if len(text) > 6 {
		text = text[7:]
	} else {
		text = ""
	}
	u := &User{
		ID:        user.ID,
		FirstName: user.FirstName,
		Username:  user.Username,
		StartDate: time.Now().Unix(),
		Language:  user.LanguageCode,
		State:     MainMenu,
		Refer:     text,
	}

	return us.Save(u)
}

func (us *UserStorage) Read(id int64) (user *User, err error) {
	defer func() { err = e.WrapIfErr(readError, err) }()

	var u User
	query := `SELECT * FROM users WHERE id = ?`
	err = us.DBClient.DB().
		QueryRow(query, id).
		Scan(&u.ID, &u.FirstName, &u.Username, &u.StartDate, &u.Language, &u.State, &u.Refer)
	return &u, err
}

func (us *UserStorage) SetState(id int64, state int) (err error) {
	defer func() { err = e.WrapIfErr(setStateError, err) }()

	query := `UPDATE users SET state = ? WHERE id = ?`
	_, err = us.DBClient.DB().
		Exec(query, state, id)
	return err
}
