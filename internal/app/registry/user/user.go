package user

import (
	"database/sql"
	"errors"
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/e"
	userStorage "github.com/vi350/vk-internship/internal/app/storage/user"
	"log"
	"sync"
	"time"
)

const (
	getByTgUserError = "can't get user by tg user"
)

type UserRegistry struct {
	sync.RWMutex
	users       map[int64]*User
	userStorage *userStorage.UserStorage
}

type User struct {
	userFromStorage *userStorage.User
	lastAccessTime  time.Time
}

func New(userStorage *userStorage.UserStorage) *UserRegistry {
	return &UserRegistry{
		users:       make(map[int64]*User),
		userStorage: userStorage,
	}
}

func (ur *UserRegistry) RemoveInactive(minutes int) {
	var start time.Time
	var actualStart time.Time
	defer func() {
		log.Printf("removed inactive users; aquiring lock took %v, removing took %v", actualStart.Sub(start), time.Since(actualStart))
	}()
	log.Println("removing inactive users")
	start = time.Now()

	ur.Lock()
	defer ur.Unlock()

	actualStart = time.Now()
	for id, u := range ur.users {
		if time.Since(u.lastAccessTime).Minutes() > float64(minutes) {
			delete(ur.users, id)
		}
	}
}

func (ur *UserRegistry) Sync() {
	var start time.Time
	var actualStart time.Time
	defer func() {
		log.Printf("synced users; aquiring lock took %v, syncing took %v", actualStart.Sub(start), time.Since(actualStart))
	}()
	log.Println("syncing users...")
	start = time.Now()

	ur.RLock()
	defer ur.RUnlock()

	actualStart = time.Now()
	usersForStorage := make(map[int64]*userStorage.User)
	for id, u := range ur.users {
		usersForStorage[id] = u.userFromStorage
	}

	if err := ur.userStorage.UpdateWithMap(usersForStorage); err != nil {
		log.Printf("error updating users: %v", err)
	}
}

func (ur *UserRegistry) GetByTgUser(userFromMessage *tgClient.User, text string) (userFromStorage *userStorage.User, isNew bool, err error) {
	defer func() { err = e.WrapIfErr(getByTgUserError, err) }()
	isNew = false

	ur.RLock()
	if u, readOk := ur.users[userFromMessage.ID]; readOk {
		userFromStorage = u.userFromStorage
		// TODO: check if it's ok to update lastAccessTime without locking
		u.lastAccessTime = time.Now()
		ur.RUnlock()
		return
	}
	ur.RUnlock()

	userFromStorage, err = ur.userStorage.Read(userFromMessage.ID)
	if errors.Is(err, sql.ErrNoRows) {
		if userFromStorage, err = ur.userStorage.InsertUsingTgClientUser(userFromMessage, text); err != nil {
			return
		}
		isNew = true
	}

	ur.Lock()
	defer ur.Unlock()
	ur.users[userFromMessage.ID].userFromStorage = userFromStorage
	ur.users[userFromMessage.ID].lastAccessTime = time.Now()

	return
}
