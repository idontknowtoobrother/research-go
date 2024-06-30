package stores

import (
	"fmt"

	models "github.com/idontknowtoobrother/monolith-api-crud/models"
)

type PlayerStore struct {
	players []models.Player
}

func (store *PlayerStore) GetPlayers() []models.Player {
	return store.players
}

func (store *PlayerStore) GetPlayerByUuid(uuid string) (*models.Player, error) {
	for _, player := range store.players {
		if player.Uuid == uuid {
			return &player, nil
		}
	}
	return nil, fmt.Errorf("player with uuid=%s not found", uuid)
}

func (store *PlayerStore) CreatePlayer(player models.Player) (*models.Player, error) {
	for _, p := range store.players {
		if p.Uuid == player.Uuid {
			return nil, fmt.Errorf("player with uuid=%s already exists", player.Uuid)
		}
		if p.Name == player.Name {
			return nil, fmt.Errorf("player with name=%s already exists", player.Name)
		}
	}
	store.players = append(store.players, player)
	return &player, nil
}

func (store *PlayerStore) UpdatePlayer(player models.Player) (*models.Player, error) {
	for i, p := range store.players {
		if p.Uuid == player.Uuid {
			store.players[i] = player
			return &(store.players[i]), nil
		}
	}

	return nil, fmt.Errorf("player with uuid=%s not found", player.Uuid)
}

func (store *PlayerStore) DeletePlayerByUuid(uuid string) error {
	for i, p := range store.players {
		if p.Uuid == uuid {
			store.players = append(store.players[:i], store.players[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("player with uuid=%s not found", uuid)
}

func NewPlayerStore() *PlayerStore {
	store := &PlayerStore{
		players: getPlayersMock(),
	}
	return store
}

func getPlayersMock() []models.Player {
	var mockPlayers = []models.Player{}

	player1 := models.Player{
		Uuid:       "1",
		Name:       "Hex GO Mode",
		Experience: 0,
		Inventory: []models.Item{
			{
				Name:     "Sword",
				Quantity: 1,
			},
		},
	}

	player2 := models.Player{
		Uuid:       "2",
		Name:       "SinceLaguna Sleepy",
		Experience: 0,
		Inventory: []models.Item{
			{
				Name:     "Shield",
				Quantity: 1,
			},
		},
	}

	mockPlayers = append(mockPlayers, player1)
	mockPlayers = append(mockPlayers, player2)
	return mockPlayers
}
