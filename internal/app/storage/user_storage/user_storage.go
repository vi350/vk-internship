package user_storage

import (
	"fmt"
	"github.com/vi350/vk-internship/internal/app/clients/postgres"
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/e"
	"time"
)

type UserStorage struct {
	DBClient postgres.Client
}

func New(dbClient *postgres.Client) *UserStorage {
	return &UserStorage{
		DBClient: *dbClient,
	}
}

const (
	saveError     = "error performing user insert"
	readError     = "error performing user read"
	isExistError  = "error performing user isexist"
	setStateError = "error performing user setstate"
)

func (us *UserStorage) Save(user *User) (err error) {
	defer func() { err = e.WrapIfErr(saveError, err) }()

	query := fmt.Sprintf(
		"INSERT INTO users VALUES ('%d', '%s', '%s', '%d', '%s', '%d', '%s') RETURNING id;",
		user.ID, user.FirstName, user.Username, user.StartDate, user.Language, user.State, user.Refer,
	)
	err = us.DBClient.DB.QueryRow(query).Scan(&user.ID)
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

	query := fmt.Sprintf(
		"INSERT INTO users VALUES ('%d', '%s', '%s', '%d', '%s', '%d', '%s') RETURNING id;",
		u.ID, u.FirstName, u.Username, u.StartDate, u.Language, u.State, u.Refer,
	)
	err = us.DBClient.DB.QueryRow(query).Scan(&user.ID)
	return err
}

func (us *UserStorage) Read(id int64) (user *User, err error) {
	defer func() { err = e.WrapIfErr(readError, err) }()

	var u User
	query := fmt.Sprintf(
		"SELECT * FROM users WHERE id = '%d';",
		id,
	)
	err = us.DBClient.DB.QueryRow(query).Scan(&u.ID, &u.FirstName, &u.Username, &u.StartDate, &u.Language, &u.State, &u.Refer)
	return &u, err
}

func (us *UserStorage) SetState(id int64, state int) (err error) {
	defer func() { err = e.WrapIfErr(setStateError, err) }()

	query := fmt.Sprintf(
		"UPDATE users SET state = '%d' WHERE id = '%d';",
		state, id,
	)
	_, err = us.DBClient.DB.Exec(query)
	return err
}
