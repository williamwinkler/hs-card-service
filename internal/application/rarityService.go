package application

import "github.com/williamwinkler/hs-card-service/internal/application/interfaces"

type RarityService struct {
	rarityRepo  interfaces.RarityRepository
	hsClient interfaces.HsClient
}

func NewRarityService(rarityRepo interfaces.RarityRepository, hsClient interfaces.HsClient) *RarityService {
	return &RarityService{
		rarityRepo:  rarityRepo,
		hsClient: hsClient,
	}
}

func (c *RarityService) Update() error {
	rarities, err := c.hsClient.GetRarities()
	if err != nil {
		return err
	}

	err = c.rarityRepo.DeleteAll()
	if err != nil {
		return err
	}

	return c.rarityRepo.InsertMany(rarities)
}
