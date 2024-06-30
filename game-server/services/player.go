package services

import (
	"github.com/idontknowtoobrother/monolith-api-crud/models"
	"github.com/idontknowtoobrother/monolith-api-crud/stores"
)

type PlayerService struct {
	playerStore *stores.PlayerStore
}

// recieve database store
func NewPlayerService(playerStore *stores.PlayerStore) *PlayerService {
	service := &PlayerService{
		playerStore: playerStore,
	}

	return service
}

func (service *PlayerService) GetPlayers() []models.Player {
	players := service.playerStore.GetPlayers()
	return players
}

func (service *PlayerService) GetPlayerByUuid(uuid string) (*models.Player, error) {
	player, err := service.playerStore.GetPlayerByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (service *PlayerService) CreatePlayer(player models.Player) (*models.Player, error) {
	createdPlayer, err := service.playerStore.CreatePlayer(player)
	if err != nil {
		return nil, err
	}
	return createdPlayer, nil
}

func (service *PlayerService) UpdatePlayer(player models.Player) (*models.Player, error) {
	updatedPlayer, err := service.playerStore.UpdatePlayer(player)
	if err != nil {
		return nil, err
	}
	return updatedPlayer, nil
}

func (service *PlayerService) DeletePlayerByUuid(uuid string) error {

	err := service.playerStore.DeletePlayerByUuid(uuid)
	if err != nil {
		return err
	}

	return nil
}
