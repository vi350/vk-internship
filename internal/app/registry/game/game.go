package game

import (
	"database/sql"
	"errors"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/models"
	gameStorage "github.com/vi350/vk-internship/internal/app/storage/game"
	"log"
	"sync"
	"time"
)

const (
	findByUserError = "find by user error"
)

type GameRegistry struct {
	sync.RWMutex
	games       map[int]*models.Game
	ownerMap    map[int64]int
	opponentMap map[int64]int
	gameStorage *gameStorage.GameStorage
}

func New(gameStorage *gameStorage.GameStorage) *GameRegistry {
	return &GameRegistry{
		games:       make(map[int]*models.Game),
		ownerMap:    make(map[int64]int),
		opponentMap: make(map[int64]int),
		gameStorage: gameStorage,
	}
}

func (gr *GameRegistry) RemoveInactive(minutes int) {
	var start time.Time
	var actualStart time.Time
	defer func() {
		log.Printf("removed inactive games; aquiring lock took %v, removing took %v", actualStart.Sub(start), time.Since(actualStart))
	}()
	log.Println("removing inactive games")
	start = time.Now()

	gr.Lock()
	defer gr.Unlock()

	actualStart = time.Now()

	for id, g := range gr.games {
		if time.Since(g.LastAccessTime).Minutes() > float64(minutes) {
			delete(gr.ownerMap, g.OwnerID)
			delete(gr.opponentMap, g.OpponentID)
			delete(gr.games, id)
		}
	}
}

func (gr *GameRegistry) Sync() {
	var start time.Time
	var actualStart time.Time
	defer func() {
		log.Printf("synced games; aquiring lock took %v, syncing took %v", actualStart.Sub(start), time.Since(actualStart))
	}()
	log.Println("syncing games...")
	start = time.Now()

	gr.RLock()
	defer gr.RUnlock()

	actualStart = time.Now()
	gamesForStorage := make(map[int]*models.Game)
	for _, g := range gr.games {
		gamesForStorage[g.ID] = g
	}

	if err := gr.gameStorage.UpdateWithMap(gamesForStorage); err != nil {
		log.Printf("error updating users: %v", err)
	}
}

func (gr *GameRegistry) FindUsersActiveGame(userid int64) (game *models.Game, err error) {
	defer func() { err = e.WrapIfErr(findByUserError, err) }()

	gr.RLock()
	if gameID, readOk := gr.ownerMap[userid]; readOk {
		game = gr.games[gameID]
		gr.games[gameID].LastAccessTime = time.Now()
		gr.RUnlock()
		return
	}
	if gameID, readOk := gr.opponentMap[userid]; readOk {
		game = gr.games[gameID]
		gr.games[gameID].LastAccessTime = time.Now()
		gr.RUnlock()
		return
	}
	gr.RUnlock()

	game, err = gr.gameStorage.FindUsersActiveGame(userid)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
		return
	} else if err == nil {
		gr.Lock()
		defer gr.Unlock()

		gr.games[game.ID] = game
		gr.ownerMap[game.OwnerID] = game.ID
		gr.opponentMap[game.OpponentID] = game.ID
	}
	return
}

func (gr *GameRegistry) FindUsersGames(userid int64) (games []*models.Game, err error) {
	// as all games will be asked not so frequently, we can interact with storage without caching
	games, err = gr.gameStorage.FindUsersGames(userid)
	return
}
