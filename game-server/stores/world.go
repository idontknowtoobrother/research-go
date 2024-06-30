package stores

import "github.com/idontknowtoobrother/monolith-api-crud/models"

type WorldStore struct {
	Worlds map[string]models.World
}

func NewWorldStore() *WorldStore {
	store := &WorldStore{
		Worlds: getWorldsMock(),
	}
	return store
}

func getWorldsMock() map[string]models.World {
	world1 := models.World{
		Uuid: "1",
		Players: []models.Player{
			{
				Uuid:       "1",
				Name:       "Hex GO Mode",
				Experience: 0,
				Inventory: []models.Item{
					{
						Name:     "Sword",
						Quantity: 1,
					},
				},
			},
		},
	}

	world2 := models.World{
		Uuid: "2",
		Players: []models.Player{
			{
				Uuid:       "1",
				Name:       "SinceLaguna Sleepy",
				Experience: 0,
				Inventory: []models.Item{
					{
						Name:     "Shield",
						Quantity: 1,
					},
				},
			},
		},
	}

	worlds := map[string]models.World{
		world1.Uuid: world1,
		world2.Uuid: world2,
	}

	return worlds
}
