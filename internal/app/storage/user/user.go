package user

import (
	"database/sql"
	"github.com/vi350/vk-internship/internal/app/clients"
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/models"
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

func (us *UserStorage) Storage() {}

func (us *UserStorage) Insert(user *models.User) (err error) {
	defer func() { err = e.WrapIfErr(storage.InsertError, err) }()

	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7);`
	// TODO: move to ExecContext
	_, err = us.DBClient.DB().
		Exec(query, user.ID, user.FirstName, user.Username, user.StartDate, user.Language, user.State, user.Refer)

	return
}

func (us *UserStorage) InsertUsingTgClientUser(userFromMessage *tgClient.User, text string) (userFromStore *models.User, err error) {
	defer func() { err = e.WrapIfErr(storage.InsertError, err) }()

	if userFromMessage.LanguageCode == "" {
		userFromMessage.LanguageCode = "en"
	}
	if len(text) > 6 {
		text = text[7:]
	} else {
		text = ""
	}
	userFromStore = &models.User{
		ID:        userFromMessage.ID,
		FirstName: userFromMessage.FirstName,
		Username:  userFromMessage.Username,
		StartDate: time.Now().Unix(),
		Language:  userFromMessage.LanguageCode,
		State:     models.ChooseLanguage,
		Refer:     text,
	}
	err = us.Insert(userFromStore)

	return
}

func (us *UserStorage) Read(id int64) (user *models.User, err error) {
	defer func() { err = e.WrapIfErr(storage.ReadError, err) }()

	var u models.User
	query := `SELECT * FROM users WHERE id = $1`
	err = us.DBClient.DB().
		QueryRow(query, id).
		Scan(&u.ID, &u.FirstName, &u.Username, &u.StartDate, &u.Language, &u.State, &u.Refer)
	user = &u

	return
}

func (us *UserStorage) UpdateWithMap(usersToUpdate map[int64]*models.User) (err error) {
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

	query := `UPDATE users SET first_name = $1, username = $2, start_date = $3, language = $4, state = $5, refer = $6 WHERE id = $7;`
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
