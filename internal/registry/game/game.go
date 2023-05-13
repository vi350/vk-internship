package game

import (
	"database/sql"
	"errors"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/storage/game_storage"
	"log"
	"sync"
	"time"
)

const (
	findByUserError = "find by user error"
)

type GameRegistry struct {
	sync.RWMutex
	games       map[int]*Game
	ownerMap    map[int64]int
	opponentMap map[int64]int
	gameStorage *game_storage.GameStorage
}

type Game struct {
	gameFromStorage *game_storage.Game
	lastAccessTime  time.Time
}

func New(gameStorage *game_storage.GameStorage) *GameRegistry {
	return &GameRegistry{
		games:       make(map[int]*Game),
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
		if time.Since(g.lastAccessTime).Minutes() > float64(minutes) {
			delete(gr.ownerMap, g.gameFromStorage.OwnerID)
			delete(gr.opponentMap, g.gameFromStorage.OpponentID)
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
	gamesForStorage := make(map[int]*game_storage.Game)
	for _, g := range gr.games {
		gamesForStorage[g.gameFromStorage.ID] = g.gameFromStorage
	}

	if err := gr.gameStorage.UpdateWithMap(gamesForStorage); err != nil {
		log.Printf("error updating users: %v", err)
	}
}

func (gr *GameRegistry) FindUsersActiveGame(userid int64) (game *game_storage.Game, err error) {
	defer func() { err = e.WrapIfErr(findByUserError, err) }()

	gr.RLock()
	if gameID, readOk := gr.ownerMap[userid]; readOk {
		game = gr.games[gameID].gameFromStorage
		gr.games[gameID].lastAccessTime = time.Now()
		gr.RUnlock()
		return
	}
	if gameID, readOk := gr.opponentMap[userid]; readOk {
		game = gr.games[gameID].gameFromStorage
		gr.games[gameID].lastAccessTime = time.Now()
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

		gr.games[game.ID] = &Game{
			gameFromStorage: game,
			lastAccessTime:  time.Now(),
		}
		gr.ownerMap[game.OwnerID] = game.ID
		gr.opponentMap[game.OpponentID] = game.ID
	}
	return
}

func (gr *GameRegistry) FindUsersGames(userid int64) (games []*game_storage.Game, err error) {
	// as all games will be asked not so frequently, we can interact with storage without caching
	games, err = gr.gameStorage.FindUsersGames(userid)
	return
}
