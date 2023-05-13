package user_storage

import (
	"database/sql"
	"github.com/vi350/vk-internship/internal/app/clients"
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/storage"
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

func (us *UserStorage) Insert(user *User) (err error) {
	defer func() { err = e.WrapIfErr(storage.InsertError, err) }()

	query := `INSERT INTO users VALUES (?, ?, ?, ?, ?, ?, ?)`
	// TODO: move to ExecContext
	_, err = us.DBClient.DB().
		Exec(query, user.ID, user.FirstName, user.Username, user.StartDate, user.Language, user.State, user.Refer)

	return
}

func (us *UserStorage) InsertUsingTgClientUser(userFromMessage *tgClient.User, text string) (userFromStore *User, err error) {
	defer func() { err = e.WrapIfErr(storage.InsertError, err) }()

	if userFromMessage.LanguageCode == "" {
		userFromMessage.LanguageCode = "en"
	}
	if len(text) > 6 {
		text = text[7:]
	} else {
		text = ""
	}
	userFromStore = &User{
		ID:        userFromMessage.ID,
		FirstName: userFromMessage.FirstName,
		Username:  userFromMessage.Username,
		StartDate: time.Now().Unix(),
		Language:  userFromMessage.LanguageCode,
		State:     ChooseLanguage,
		Refer:     text,
	}
	err = us.Insert(userFromStore)

	return
}

func (us *UserStorage) Read(id int64) (user *User, err error) {
	defer func() { err = e.WrapIfErr(storage.ReadError, err) }()

	var u *User
	query := `SELECT * FROM users WHERE id = ?`
	err = us.DBClient.DB().
		QueryRow(query, id).
		Scan(&u.ID, &u.FirstName, &u.Username, &u.StartDate, &u.Language, &u.State, &u.Refer)

	return
}

func (us *UserStorage) UpdateWithMap(usersToUpdate map[int64]*User) (err error) {
	defer func() { err = e.WrapIfErr(storage.UpdateWithMapError, err) }()

	var tx *sql.Tx
	if tx, err = us.DBClient.DB().Begin(); err != nil {
		return e.WrapIfErr(storage.BeginError, err)
	}

	defer func() {
		if err != nil {
			if err = tx.Rollback(); err != nil {
				err = e.WrapIfErr(storage.RollbackError, err)
			}
			return
		}
		err = tx.Commit()
		if err != nil {
			err = e.WrapIfErr(storage.CommitError, err)
		}
	}()

	query := `UPDATE users SET first_name = ?, username = ?, start_date = ?, language = ?, state = ?, refer = ? WHERE id = ?`
	statement, err := tx.Prepare(query)
	if err != nil {
		return e.WrapIfErr(storage.PrepareError, err)
	}
	defer func() {
		if err = statement.Close(); err != nil {
			err = e.WrapIfErr(storage.CloseStatementError, err)
		}
	}()

	for id, user := range usersToUpdate {
		_, err = statement.Exec(user.FirstName, user.Username, user.StartDate, user.Language, user.State, user.Refer, id)
		if err != nil {
			return e.WrapIfErr(storage.ExecError, err)
		}
	}

	return
}
