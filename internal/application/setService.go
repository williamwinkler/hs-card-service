package application

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/application/interfaces"
)

type SetService struct {
	setRepo  interfaces.SetRepository
	hsClient interfaces.HsClient
}

func NewSetService(setRepo interfaces.SetRepository, hsClient interfaces.HsClient) *SetService {
	return &SetService{
		setRepo:  setRepo,
		hsClient: hsClient,
	}
}

func (c *SetService) Update(ctx context.Context) error {
	sets, err := c.hsClient.GetSets(ctx)
	if err != nil {
		return err
	}

	err = c.setRepo.DeleteAll()
	if err != nil {
		return err
	}

	return c.setRepo.InsertMany(sets)
}
